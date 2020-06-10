package tic

import "github.com/odom11/playground_micro"

type Tic struct {
	toc api.TocService;
}

func New(toc api.TocService) *Tic {
	return &Tic{toc: toc}
}

func (t *Tic) Shout()
