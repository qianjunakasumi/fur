package main

import (
	"context"
	"os"
	"strconv"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Snowflake *Node

func getEnv() zerolog.Level {

	node, ok := os.LookupEnv("FUR_NODE")
	if !ok {
		log.Panic().Msg("cannot get \"FUR_NODE\" env")
	}

	nodeid, err := strconv.ParseInt(node, 10, 8)
	if err != nil {
		log.Panic().Msg("cannot parse nodeid")
	}

	Snowflake = NewNode(nodeid)

	level, ok := os.LookupEnv("FUR_LOGLEVEL")
	if !ok {
		level = "1"
	}

	levelid, err := strconv.ParseInt(level, 10, 8)
	if err != nil {
		log.Panic().Msg("cannot parse log level")
	}

	return zerolog.Level(levelid)
}

func main() {

	levelid := getEnv()

	zerolog.SetGlobalLevel(levelid)

	s, err := daprd.NewService(":50001")
	if err != nil {
		log.Panic().Err(err).Msg("failed to start the server")
	}

	if err := s.AddServiceInvocationHandler("gen", genHandler); err != nil {
		log.Panic().Err(err).Msg("error adding invocation handler")
	}

	log.Info().Msg(`server is starting on ":50001"`)
	if err := s.Start(); err != nil {
		log.Panic().Err(err).Msg("server error")
	}

}

func genHandler(_ context.Context, _ *common.InvocationEvent) (out *common.Content, err error) {

	log.Debug().Str("Endpoint", "gen").Msg("called")

	i := Snowflake.Generate()
	out = &common.Content{
		Data:        i.UintBytes(),
		ContentType: "application/octet-stream",
		DataTypeURL: "qianjunakasumi/fur/gen/" + strconv.FormatUint(i.Uint64(), 10),
	}

	return
}
