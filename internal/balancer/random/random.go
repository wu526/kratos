package random

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/go-kratos/kratos/v2/internal/balancer"
	"github.com/go-kratos/kratos/v2/registry"
)

var _ balancer.Balancer = &Balancer{}

type Balancer struct {
}

func New() *Balancer {
	return &Balancer{}
}

func (b *Balancer) Pick(ctx context.Context, pathPattern string, nodes []*registry.ServiceInstance) (node *registry.ServiceInstance, done func(balancer.DoneInfo), err error) {
	if len(nodes) == 0 {
		return nil, nil, fmt.Errorf("no instances avaiable")
	} else if len(nodes) == 1 {
		return nodes[0], func(di balancer.DoneInfo) {}, nil
	}
	idx := rand.Intn(len(nodes))
	return nodes[idx], func(di balancer.DoneInfo) {}, nil
}