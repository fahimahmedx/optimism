package timeint

type Seconds uint64

// func Seconds(n uint64) Seconds {
// 	return seconds(n)
// }

// func Seconds(n hexutil.Uint64) seconds {
// 	return seconds(n)
// }

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

// can't seem to call the function above in other packages?

// type Milliseconds uint64

// func (t Milliseconds) ToUint64Sec() uint64 {
// 	return uint64(t) / 1000
// }
