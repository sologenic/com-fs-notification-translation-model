package client

import (
	"context"

	grpcdef "github.com/sologenic/com-fs-notification-translation-model"
	grpcclient "github.com/sologenic/com-fs-utils-lib/go/grpc-client"
)

const endpoint = "NOTIFICATION_PARSER_LIB"

var (
	client     grpcdef.TranslationServiceClient
	grpcClient *grpcclient.GRPCClient
)

func initClient() {
	grpcClient = grpcclient.InitClient(endpoint)
	cl := grpcdef.NewTranslationServiceClient(grpcClient.Conn)
	client = cl
}

func Client() grpcdef.TranslationServiceClient {
	if client == nil {
		initClient()
	}
	return client
}

func AuthCtx(ctx context.Context) context.Context {
	if grpcClient == nil {
		initClient()
	}
	return grpcClient.AuthCtx(ctx)
}
