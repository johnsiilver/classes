package client

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/johnsiilver/classes/workflow/proto"
)

type Workflow struct {
	conn   *grpc.ClientConn
	client pb.WorkflowClient
}

func New(addr string) (*Workflow, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Workflow{
		conn:   conn,
		client: pb.NewWorkflowClient(conn),
	}, nil
}

func (w *Workflow) Submit(ctx context.Context, req *pb.WorkReq) (string, error) {
	resp, err := w.client.Submit(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}

func (w *Workflow) Exec(ctx context.Context, id string) error {
	_, err := w.client.Exec(ctx, &pb.ExecReq{Id: id})
	return err
}

func (w *Workflow) Status(ctx context.Context, id string) (*pb.StatusResp, error) {
	return w.client.Status(ctx, &pb.StatusReq{Id: id})
}
