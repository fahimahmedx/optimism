package timeint

import "github.com/ethereum/go-ethereum/common/hexutil"

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
