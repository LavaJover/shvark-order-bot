package grpcapi

import (
	"context"
	"time"

	ssopb "github.com/LavaJover/shvark-sso-service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SSOClient struct {
	conn *grpc.ClientConn
	service ssopb.SSOServiceClient
}

func NewSSOClient(addr string) (*SSOClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		return nil, err
	}
	return &SSOClient{
		conn: conn,
		service: ssopb.NewSSOServiceClient(conn),
	}, nil
}

func (c *SSOClient) ValidateToken(token string) (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	response, err := c.service.ValidateToken(
		ctx,
		&ssopb.ValidateTokenRequest{
			AccessToken: token,
		},
	)
	if err != nil {
		return false, "", err
	}
	return response.Valid, response.UserId, nil
}