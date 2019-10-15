package conn

import (
	"io"
	"reflect"

	"github.com/rbrick/mcserver/util"
)

var CodecMap map[reflect.Type]Codec = map[reflect.Type]Codec{}

func init() {
	CodecMap[reflect.TypeOf("")] = &StringCodec{}
	CodecMap[reflect.TypeOf(int64(0))] = &Int64Codec{}
	CodecMap[reflect.TypeOf(uint16(46))] = &Uint16Codec{}
	CodecMap[reflect.TypeOf(int16(-1))] = &Int16Codec{}
	CodecMap[reflect.TypeOf(int(0))] = &IntCodec{}
	CodecMap[reflect.TypeOf(uint(0))] = &UintCodec{}
	CodecMap[reflect.TypeOf(&util.Varint{})] = &VarintCodec{}
	CodecMap[reflect.TypeOf(&util.Uvarint{})] = &UvarintCodec{}
}

type Codec interface {
	Decode(io.Reader) interface{}
	Encode(io.Writer, interface{})
}
