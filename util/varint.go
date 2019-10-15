package util

import (
	"bufio"
	"encoding/binary"
	"io"
)

type Uvarint struct {
	Val uint32
}

func (u *Uvarint) Decode(r io.Reader) error {
	buf := bufio.NewReader(r)

	x, err := binary.ReadUvarint(buf)

	if err != nil {
		return err
	}
	u.Val = uint32(x)
	return nil
}

func (u *Uvarint) Encode(w io.Writer) error {
	buf := make([]byte, binary.MaxVarintLen32)
	binary.PutUvarint(buf, uint64(u.Val))
	_, err := w.Write(buf)
	return err
}

type Varint struct {
	Val int32
}

func (v *Varint) Decode(r io.Reader) error {
	buf := bufio.NewReader(r)

	x, err := binary.ReadVarint(buf)

	if err != nil {
		return err
	}
	v.Val = int32(x)
	return nil
}

func (u *Varint) Encode(w io.Writer) error {
	buf := make([]byte, binary.MaxVarintLen32)
	binary.PutVarint(buf, int64(u.Val))
	_, err := w.Write(buf)
	return err
}

func DecodeVarint(r io.Reader) (*Varint, error) {
	v := &Varint{}
	err := v.Decode(r)
	return v, err
}

func DecodeUvarint(r io.Reader) (*Uvarint, error) {
	v := &Uvarint{}
	err := v.Decode(r)
	return v, err
}
