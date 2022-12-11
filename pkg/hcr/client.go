package hcr

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/souk4711/playcode/pkg/hcr/pb/languages"
	"github.com/souk4711/playcode/pkg/hcr/pb/runs"
)

var API *Client

type Client struct {
	Languages languages.LanguagesClient
	Runs      runs.RunsClient
}

func NewClient(uri string) (*Client, error) {
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	languagesClient := languages.NewLanguagesClient(conn)
	runsClient := runs.NewRunsClient(conn)

	return &Client{
		Languages: languagesClient,
		Runs:      runsClient,
	}, nil
}
