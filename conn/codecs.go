package conn

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/rbrick/mcserver/util"
)

type Int64Codec struct{}

func (*Int64Codec) Encode(w io.Writer, v interface{}) {
	binary.Write(w, binary.BigEndian, v.(int64))
}

func (*Int64Codec) Decode(r io.Reader) interface{} {
	var val int64
	binary.Read(r, binary.BigEndian, val)
	return val
}

type Uint16Codec struct{}

func (*Uint16Codec) Encode(w io.Writer, v interface{}) {
	binary.Write(w, binary.BigEndian, v.(uint16))
}

func (*Uint16Codec) Decode(r io.Reader) interface{} {
	var val uint16
	binary.Read(r, binary.BigEndian, val)
	return val
}

type Int16Codec struct{}

func (*Int16Codec) Encode(w io.Writer, v interface{}) {
	binary.Write(w, binary.BigEndian, v.(int16))
}

func (*Int16Codec) Decode(r io.Reader) interface{} {
	var val int16
	binary.Read(r, binary.BigEndian, val)
	return val
}

type IntCodec struct{}

func (*IntCodec) Encode(w io.Writer, v interface{}) {
	binary.Write(w, binary.BigEndian, v.(int))
}

func (*IntCodec) Decode(r io.Reader) interface{} {
	var val int
	binary.Read(r, binary.BigEndian, val)
	return val
}

type UintCodec struct{}

func (*UintCodec) Encode(w io.Writer, v interface{}) {
	binary.Write(w, binary.BigEndian, v.(uint))
}

func (*UintCodec) Decode(r io.Reader) interface{} {
	var val uint
	binary.Read(r, binary.BigEndian, val)
	return val
}

type StringCodec struct{}

func (*StringCodec) Encode(w io.Writer, v interface{}) {
	s := v.(string)
	l := &util.Varint{int32(len(s))}

	l.Encode(w)
	w.Write([]byte(s))
}

func (*StringCodec) Decode(r io.Reader) interface{} {
	l, err := util.DecodeUvarint(r)
	if err != nil {
		return nil
	}

	// fmt.Println("len:", l)

	d := make([]byte, l.Val)
	r.Read(d)

	fmt.Println(string(d))

	return string(d)
}

type VarintCodec struct{}

func (*VarintCodec) Decode(r io.Reader) interface{} {
	v, err := util.DecodeVarint(r)
	if err != nil {
		return nil
	}
	return v
}

func (*VarintCodec) Encode(w io.Writer, v interface{}) {
	v.(*util.Varint).Encode(w)
}

type UvarintCodec struct{}

func (*UvarintCodec) Decode(r io.Reader) interface{} {
	v, err := util.DecodeUvarint(r)
	if err != nil {
		return nil
	}
	return v
}

func (*UvarintCodec) Encode(w io.Writer, v interface{}) {
	v.(*util.Uvarint).Encode(w)
}
