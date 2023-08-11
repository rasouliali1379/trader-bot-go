package domain

import (
	"context"
	"go.uber.org/zap"
)

type strategy func(context.Context, Exchange, *Market) error

type Strategy string

const (
	Ema Strategy = "ema"
)

type Strategies struct {
	list []strategy
}

func (o *Strategies) NotifyAll(c context.Context, e Exchange, m *Market) {
	for i := range o.list {
		go func(index int) {
			if err := o.list[index](c, e, m); err != nil {
				zap.L().Error("error while informing an strategy", zap.Error(err))
			}
		}(i)
	}
}

func (o *Strategies) Register(f strategy) {
	o.list = append(o.list, f)
}
