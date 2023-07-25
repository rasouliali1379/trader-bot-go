package domain

import (
	"context"
	"go.uber.org/zap"
)

type observer func(context.Context) error

type Observer struct {
	list []observer
}

func (o *Observer) NotifyAll(ctx context.Context) {
	for i := range o.list {
		go func(index int) {
			if err := o.list[index](ctx); err != nil {
				zap.L().Error("error while informing an observer", zap.Error(err))
			}
		}(i)
	}
}

func (o *Observer) Register(f observer) {
	o.list = append(o.list, f)
}
