package domain

import (
	"fmt"
)

type Market struct {
	Give     string
	Take     string
	Exchange *Exchange
	Price    *Price
}

func (m *Market) String() string {
	return fmt.Sprintf("%s-%s", m.Give, m.Take)
}
