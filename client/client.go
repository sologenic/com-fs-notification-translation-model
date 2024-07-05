package client

import (
	grpcdef "github.com/sologenic/com-fs-notification-translation-model"
	grpcclient "github.com/sologenic/fs-utils-lib/go/grpc-client"
)

const endpoint = "NOTIFICATION_PARSER_LIB"

var (
	client     *grpcdef.TranslationServiceClient
	grpcClient *grpcclient.GRPCClient
)

func initClient() {
	grpcClient = grpcclient.InitClient(endpoint)
	cl := grpcdef.NewTranslationServiceClient(grpcClient.Conn)
	client = &cl
}

func Client() *grpcdef.TranslationServiceClient {
	if client == nil {
		initClient()
	}
	return client
}
