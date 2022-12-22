package main

import (
	"context"
	"log"
	"strconv"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
)

var Snowflake = NewNode(0)

func main() {

	s, err := daprd.NewService(":50001")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	if err := s.AddServiceInvocationHandler("gen", genHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

}

func genHandler(_ context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	log.Printf("gen - ContentType:%s, Verb:%s, QueryString:%s, %+v", in.ContentType, in.Verb, in.QueryString, string(in.Data))

	i := Snowflake.Generate()

	out = &common.Content{
		Data:        i.UintBytes(),
		ContentType: "application/octet-stream",
		DataTypeURL: "qianjunakasumi/fur/gen/" + strconv.FormatUint(i.Uint64(), 10),
	}
	return
}
