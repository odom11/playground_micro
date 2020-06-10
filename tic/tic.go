package tic

import (
	"context"
	"github.com/odom11/playground_micro/api"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

type Tic struct {
	toc api.TocService;
	counter int
}

func New(toc api.TocService) *Tic {
	return &Tic{toc: toc}
}

func (t *Tic) Shout(ctx context.Context, req *api.Bounce, rsp *api.Bounce) error  {
	_, err := t.toc.Shout(context.Background(), &api.Bounce{Message: "foo"})
	if err != nil {
		fmt.Println("successfully shouted", 1)
		fmt.Println("got message: ", req.Message)
	}
	return nil
}
