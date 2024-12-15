package session

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	v1 "github.com/teachme-group/session/pkg/api/grpc/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client v1.SessionServiceClient
	conn   *grpc.ClientConn
}

func NewClient(cfg Config) (*Client, error) {
	var (
		client     = &Client{}
		connection *grpc.ClientConn
		err        error
	)

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	}

	connection, err = grpc.NewClient(
		cfg.Address,
		options...,
	)
	if err != nil {
		log.Errorf("failed to connect collector err: %v", err)
	} else {
		client.conn = connection
	}

	client.client = v1.NewSessionServiceClient(client.conn)

	return client, err
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) ClientSetSession(ctx context.Context, in *v1.ClientSetSessionRequest) (*v1.ClientSetSessionResponse, error) {
	return c.client.ClientSetSession(ctx, in)
}
