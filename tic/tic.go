package tic

import "github.com/ob-vss-20ss/blatt2-headcover_of_eternal_torment/api"

type Tic struct {
	toc api.TocService;
}

func New(toc api.TocService) *Tic {
	return &Tic{toc: toc}
}

func (t *Tic) Shout()
