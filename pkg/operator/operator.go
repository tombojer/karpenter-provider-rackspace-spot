package operator

import (
	"context"
	"sync"

	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/karpenter/pkg/operator"
)

type Operator struct {
	manager.Manager
}

func NewOperator(ctx context.Context, operator *operator.Operator) (context.Context, *Operator) {
	return context.Background(), &Operator{}
}

func (o *Operator) Start(ctx context.Context) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()
	wg.Wait()
}
