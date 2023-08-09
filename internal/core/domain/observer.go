package domain

import (
	"context"
	"go.uber.org/zap"
)

type observer func(context.Context, *Market) error

type Observer struct {
	list []observer
}

func (o *Observer) NotifyAll(c context.Context, m *Market) {
	for i := range o.list {
		go func(index int) {
			if err := o.list[index](c, m); err != nil {
				zap.L().Error("error while informing an observer", zap.Error(err))
			}
		}(i)
	}
}

func (o *Observer) Register(f observer) {
	o.list = append(o.list, f)
}
