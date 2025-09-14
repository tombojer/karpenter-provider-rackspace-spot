package main

import (
	"github.com/tombojer/karpenter-provider-rackspace-spot/pkg/operator"
	coreoperator "sigs.k8s.io/karpenter/pkg/operator"
)

func main() {
	ctx, op := operator.NewOperator(coreoperator.NewOperator())
	op.Start(ctx)
}
