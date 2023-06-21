package main

import (
	"ecs-hook/internal/config"
	"ecs-hook/internal/db"
	"ecs-hook/internal/transport"
	"flag"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	cfgPath := flag.String("config", "./config.yaml", "Path to yaml configuration file")
	flag.Parse()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	logger.Info().Msg("service start")

	conf := config.NewConfigStruct()
	err := conf.LoadConfig(*cfgPath)
	if err != nil {
		logger.Fatal().Err(err).Msg("config loading error")
	}

	db.Init("YourDSNName")
	/*
		обращения с таблицами
	*/
	req := transport.RequestGRPC{}
	conn, err := transport.Connect(conf.AddrGrpc)
	if err != nil {
		logger.Fatal().Err(err).Msg("grpc connection error")
	}
	resp, err := req.SendResponse(conn)
	if err != nil {
		logger.Fatal().Err(err).Msg("grpc SendResponse error")
	}
	println(resp)
}
