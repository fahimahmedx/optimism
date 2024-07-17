package clock

import (
	"sync/atomic"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/timeint"
)

type SimpleClock struct {
	v atomic.Pointer[time.Time]
}

func NewSimpleClock() *SimpleClock {
	return &SimpleClock{}
}

func (c *SimpleClock) SetTime(u timeint.Seconds) {
	t := time.Unix(int64(u), 0)
	c.v.Store(&t)
}

func (c *SimpleClock) Set(v time.Time) {
	c.v.Store(&v)
}

func (c *SimpleClock) Now() time.Time {
	v := c.v.Load()
	if v == nil {
		return time.Unix(0, 0)
	}
	return *v
}
