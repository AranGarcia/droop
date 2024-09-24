// go:build testing
package characters

import (
	"context"

	"google.golang.org/grpc"

	characterspb "github.com/AranGarcia/droop/proto/gen/characters"
)

type mockClient struct {
	retrieveFunc func() (*characterspb.RetrieveResponse, error)
}

func (m mockClient) Create(ctx context.Context, in *characterspb.CreateRequest, opts ...grpc.CallOption) (*characterspb.CreateResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (m mockClient) Retrieve(ctx context.Context, in *characterspb.RetrieveRequest, opts ...grpc.CallOption) (*characterspb.RetrieveResponse, error) {
	return m.retrieveFunc()
}

func (m mockClient) Update(ctx context.Context, in *characterspb.UpdateRequest, opts ...grpc.CallOption) (*characterspb.UpdateResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (m mockClient) Delete(ctx context.Context, in *characterspb.DeleteRequest, opts ...grpc.CallOption) (*characterspb.DeleteResponse, error) {
	panic("not implemented") // TODO: Implement
}
