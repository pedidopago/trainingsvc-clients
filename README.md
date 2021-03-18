# Teste Microserviço Clients

* Clonar o repositório `git clone git@github.com:pedidopago/trainingsvc-clients.git`
* Implementar e corrigir métodos neste microservice de clients (marcados com TODO e FIXME);
* Executar o `./run_server.sh` e o `./run_client.sh`;
* O `./run_client.sh` deve exibir "SUCCESS";
* Os unit tests devem executar com sucesso ao rodar `go test ./...`;

Ao concluir, enviar url do repositório p/ recrutamento \<aT\> pedidopago.com.br

## Dependências

### protoc
#### MacOS
```sh
brew install protobuf
```
#### Linux
```sh
apt install -y protobuf-compiler
```

#### mariadb 10.2+ (ou mysql 5.7+)

## Setup

#### Criar Database:
```sql
CREATE DATABASE IF NOT EXISTS `ms_training` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
USE `ms_training`;




DROP TABLE IF EXISTS `client_matches`;
DROP TABLE IF EXISTS `clients`;


CREATE TABLE `clients` (
  `id` char(26) NOT NULL,
  `name` varchar(200) NOT NULL,
  `birthday` datetime DEFAULT NULL,
  `score` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`) USING BTREE,
  KEY `idx_birthday` (`birthday`) USING BTREE,
  KEY `idx_score` (`score`) USING BTREE,
  KEY `idx_created_at` (`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `client_matches` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `client_id` char(26) NOT NULL,
  `score` int(11) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `client_matches_ibfk_1` (`client_id`),
  CONSTRAINT `client_matches_ibfk_1` FOREIGN KEY (`client_id`) REFERENCES `clients` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```
### Salvar a configuração em um arquivo .env:
```
DBCS=user:password@tcp(host:port)/ms_training?parseTime=true
```
Uma alternativa é alterar o run_server.sh:
```sh
go run cmd/service/main.go --dbcs="user:password@tcp(host:port)/ms_training?parseTime=true"
```

### Implementar e corrigir partes do serviço marcados com "TODO" e "FIXME"

### Executar go generate ao atualizar arquivos .proto
`go generate ./...`

### Executar server
```sh
./run_server.sh
```

### Executar o testclient em um outro shell:
```sh
./run_client.sh
```

### Corrigir/implementar teste:
`internal/clients-service/service/service_test.go`
