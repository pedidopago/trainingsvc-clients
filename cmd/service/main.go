package main

import (
	"context"
	"net"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // registers mariadb/mysql connection driver
	"github.com/pedidopago/trainingsvc-clients/internal/clients-service/service"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "addr",
			EnvVars: []string{"LISTEN_ADDRESS", "ADDR"},
			Usage:   "host:port",
			Aliases: []string{"l", "a"},
			Value:   ":6000",
		},
		&cli.StringFlag{
			Name:    "dbcs",
			EnvVars: []string{"DBCS"},
			Usage:   "mariadb connection string: user:password@tcp(host:port)/ms_training?parseTime=true",
		},
	}

	app.Action = run

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("app run") // = log.Error() + os.Exit(1)
	}
}

func run(c *cli.Context) error {

	lis, err := net.Listen("tcp", c.String("addr"))
	if err != nil {
		log.Error().Err(err).Str("addr", c.String("addr")).Caller().Msg("listener error")
		return err
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()

	ctx, cf := context.WithCancel(context.Background())
	defer cf()

	if err := service.New(ctx, grpcServer, service.Config{
		DBCS: c.String("dbcs"),
	}); err != nil {
		log.Error().Err(err).Caller().Msg("service starter error")
		return err
	}

	lerr := make(chan error, 1)
	go func() {
		err := grpcServer.Serve(lis)
		lerr <- err
	}()
	select {
	case err := <-lerr:
		log.Error().Err(err).Caller().Msg("listen error")
		return err
	case <-time.After(time.Millisecond * 500):
		log.Debug().Str("addr", c.String("addr")).Msg("listening")
	}

	//FIXME: fazer com que o programa só finalize ao receber um sinal os.INTERRUPT (ou +) com o channel "ch" (tip: Notify)
	ch := make(chan os.Signal, 1)
	// signal...
	panic("fazer com que o programa só finalize ao receber um sinal os.INTERRUPT")
	<-ch

	grpcServer.GracefulStop()
	log.Warn().Msg("shutting down")
	return nil
}
