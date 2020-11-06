package yt_test

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"log"
	"testing"

	"github.com/yuantingapis/go-client/yuanting/yt/v1"
	ytpb "github.com/yuantingapis/go-genproto/yuanting/yt/v1"
)

func TestVendorClient_GetVendor(t *testing.T) {
	ctx := context.Background()
	c, err := yt.NewVendorClient(ctx,
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithInsecure()))
	if err != nil {
		t.Fatalf("NewVendorClient returned error: %v", err)
	}

	req := &ytpb.GetVendorRequest{
		Name: "Bruce Lee",
	}
	resp, err := c.GetVendor(ctx, req)
	if err != nil {
		t.Fatalf("GetVendor returned error: %v", err)
	}
	log.Printf("got vendor: %v", resp.Description)
}
