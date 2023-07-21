package domain

import "context"

type observer func(context.Context) error

type Observer struct {
	list []observer
}

func (o Observer) NotifyAll(ctx context.Context) error {
	for i := range o.list {
		if err := o.list[i](ctx); err != nil {
			return err
		}
	}

	return nil
}

func (o Observer) Register(f observer) {

}
