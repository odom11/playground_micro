package toc

import (
	"context"
	"github.com/odom11/playground_micro/api"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

type Toc struct {
	tic     api.TicService;
	counter int
}

func New(toc api.TicService) *Toc {
	return &Toc{tic: toc}
}

func (t *Toc) Shout(ctx context.Context, req *api.Bounce, rsp *api.Bounce) error  {
	_, err := t.tic.Shout(context.Background(), &api.Bounce{Message: "bar"})
	if err != nil {
		fmt.Println("toc successfully shouted ", t.counter, " times")
		t.counter++
		fmt.Println("got message: ", req.Message)
	}
	return nil
}
