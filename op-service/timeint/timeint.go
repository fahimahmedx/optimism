package timeint

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Seconds uint64

func (t Seconds) MultiplyInt(n uint64) Seconds {
	return t * Seconds(n)
}

func (t Seconds) ToUint64Ptr() *uint64 {
	t_uint64 := uint64(t)
	return &t_uint64
}

func (t *Seconds) FromSecPtrToUint64Ptr() *uint64 {
	if t == nil {
		return nil
	}

	t_uint64 := uint64(*t)
	return &t_uint64
}

func (t Seconds) ToUint64Sec() uint64 {
	return uint64(t)
}

func (t Seconds) ToSecondsPtr() *Seconds {
	return &t
}

func FromUint64SecToSec(t uint64) Seconds {
	return Seconds(t)
}

func FromUint64PtrToSecPtr(t *uint64) *Seconds {
	s := Seconds(*t)
	return &s
}

func FromHexUint64SecToSec(t hexutil.Uint64) Seconds {
	return Seconds(t)
}

type Milliseconds uint64

func (t Milliseconds) MultiplyInt(n uint64) Milliseconds {
	return t * Milliseconds(n)
}

func (t Seconds) ToMilliseconds() Milliseconds {
	return Milliseconds(t * 1000)
}

func (t Milliseconds) ToSeconds() Seconds {
	return Seconds(t / 1000)
}

func FromUint64SecToMilli(t uint64) Milliseconds {
	s := Milliseconds(t * 1000)
	return s
}

func FromHexUint64SecToMilli(t hexutil.Uint64) Milliseconds {
	s := Milliseconds(t * 1000)
	return s
}

func (t Milliseconds) ToUint64Milli() uint64 {
	return uint64(t)
}
