package exchanges

import "hamgit.ir/novin-backend/trader-bot/internal/core/port"

type okx struct {
}

func NewOkx() port.Exchange {
	return &okx{}
}
