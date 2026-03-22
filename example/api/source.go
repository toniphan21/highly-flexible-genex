package api

import "context"

type Request struct{}
type Response struct{}

type awesomeAPI struct{}

func (op *awesomeAPI) parse(ctx context.Context, request Request) (string, error) {
	return "", nil
}

func (op *awesomeAPI) invoke(ctx context.Context, input string) (int, error) {
	return 0, nil
}

func (op *awesomeAPI) response(ctx context.Context, output int) (Response, error) {
	return Response{}, nil
}
