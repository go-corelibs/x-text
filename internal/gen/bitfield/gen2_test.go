// Code generated by github.com/go-corelibs/x-text/internal/gen/bitfield. DO NOT EDIT.

package bitfield

type myInt2 uint32

func (m myInt2) fob() uint16 {
	return uint16((m >> 12) & 0xffff)
}

func (m myInt2) baz() int8 {
	return int8((m >> 7) & 0x1f)
}

func (m myInt2) bar() myUint8 {
	return myUint8((m >> 4) & 0x7)
}

func (m myInt2) Bool() bool {
	const bit = 1 << 3
	return m&bit == bit
}

func (m myInt2) Baz() int8 {
	return int8((m >> 0) & 0x7)
}
