package main

import (
	"context"
	"os"
	"time"

	"github.com/pedidopago/trainingsvc-clients/protos/pb"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "server",
			EnvVars: []string{"SERVER"},
			Value:   "127.0.0.1:6000",
		},
	}

	app.Action = func(c *cli.Context) error {
		conn, err := grpc.Dial(c.String("server"), grpc.WithInsecure())
		if err != nil {
			return cli.NewExitError(err.Error(), 2)
		}
		defer conn.Close()
		cl := pb.NewClientsServiceClient(conn)
		return runTests(cl)
	}

	cli.HandleExitCoder(app.Run(os.Args))
}

func runTests(cl pb.ClientsServiceClient) error {

	ctx := context.Background()

	if _, err := cl.DeleteAllClients(ctx, &pb.DeleteAllClientsRequest{}); err != nil {
		return cli.NewExitError(err.Error(), 3)
	}

	bobidResp, err := cl.NewClient(ctx, &pb.NewClientRequest{
		Name:     "Bob",
		Birthday: time.Date(1970, 1, 10, 12, 0, 0, 0, time.UTC).UnixNano(),
		Score:    0,
	})
	if err != nil {
		return cli.NewExitError(err.Error(), 4)
	}

	aliceidResp, err := cl.NewClient(ctx, &pb.NewClientRequest{
		Name:     "Alice",
		Birthday: time.Date(1987, 3, 13, 12, 0, 0, 0, time.UTC).UnixNano(),
		Score:    0,
	})
	if err != nil {
		return cli.NewExitError(err.Error(), 5)
	}

	// adding score to Alice, so it needs to be the first result
	if _, err := cl.NewMatch(ctx, &pb.NewMatchRequest{
		ClientId: aliceidResp.Id,
		Score:    100,
	}); err != nil {
		return cli.NewExitError(err.Error(), 6)
	}

	{
		r, err := cl.QueryClients(ctx, &pb.QueryClientsRequest{})
		if err != nil {
			return cli.NewExitError("should be Alice but is err: "+err.Error(), 9)
		}
		if r.Ids[0] != aliceidResp.Id {
			return cli.NewExitError("should be Alice ID", 9)
		}
	}

	xclients, err := cl.GetClients(ctx, &pb.GetClientsRequest{
		Ids: []string{bobidResp.Id, aliceidResp.Id},
	})
	if err != nil {
		return cli.NewExitError(err.Error(), 7)
	}
	if len(xclients.Clients) != 2 {
		return cli.NewExitError("invalid length", 8)
	}
	if xclients.Clients[1].Name != "Alice" {
		return cli.NewExitError("should be Alice but is "+xclients.Clients[1].Name, 9)
	}

	// adding a negative score to Alice, so it needs to be the last result
	if _, err := cl.NewMatch(ctx, &pb.NewMatchRequest{
		ClientId: aliceidResp.Id,
		Score:    -500,
	}); err != nil {
		return cli.NewExitError(err.Error(), 10)
	}

	/// /// ///
	xclients2, err := cl.QueryClients(ctx, &pb.QueryClientsRequest{})
	if err != nil {
		return cli.NewExitError(err.Error(), 7)
	}
	if len(xclients2.Ids) != 2 {
		return cli.NewExitError("invalid length (2)", 8)
	}

	if xclients2.Ids[1] != aliceidResp.Id {
		return cli.NewExitError("should be Alice (2)", 9)
	}
	/// /// ///

	if _, err := cl.DeleteClient(ctx, &pb.DeleteClientRequest{
		Id: aliceidResp.Id,
	}); err != nil {
		return cli.NewExitError(err, 11)
	}

	// adding a score to Alice should result in an error now
	if _, err := cl.NewMatch(ctx, &pb.NewMatchRequest{
		ClientId: aliceidResp.Id,
		Score:    10,
	}); err == nil {
		return cli.NewExitError("should not add score to Alice (deleted)", 12)
	}

	xclients, err = cl.GetClients(ctx, &pb.GetClientsRequest{
		Ids: []string{bobidResp.Id},
	})
	if err != nil {
		return cli.NewExitError(err.Error(), 13)
	}
	if len(xclients.Clients) != 1 {
		return cli.NewExitError("invalid length", 14)
	}

	resp, err := cl.Sort(ctx, &pb.SortRequest{
		Items:            []string{"a", "c", "b", "b", "x", "aaa", "z1"},
		RemoveDuplicates: true,
	})

	if len(resp.Items) != 6 {
		return cli.NewExitError("invalid length", 15)
	}

	if resp.Items[0] != "a" {
		return cli.NewExitError("invalid val", 16)
	}

	if resp.Items[1] != "aaa" {
		return cli.NewExitError("invalid val", 16)
	}

	if resp.Items[5] != "z1" {
		return cli.NewExitError("invalid val", 16)
	}

	println("SUCCESS!")

	return nil
}
