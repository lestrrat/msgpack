package msgpack

// Auto-generated by internal/cmd/gendecoder/gendecoder.go. DO NOT EDIT!

import (
	"math"

	"github.com/pkg/errors"
)

func (dnl *decoderNL) DecodeInt(v *int) error {
	code, err := dnl.src.ReadByte()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code for Int64`)
	}
	if IsFixNumFamily(Code(code)) {
		*v = int(code)
		return nil
	}
	switch Code(code) {
	case Int64:
		x, err := dnl.src.ReadUint64()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int`)
		}
		*v = int(x)
		return nil
	case Int32:
		x, err := dnl.src.ReadUint32()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int`)
		}
		*v = int(x)
		return nil
	case Int16:
		x, err := dnl.src.ReadUint16()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int`)
		}
		*v = int(x)
		return nil
	case Int8:
		x, err := dnl.src.ReadUint8()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int`)
		}
		*v = int(x)
		return nil
	}
	return errors.Errorf(`msgpack: invalid numeric type %s for int`, Code(code))
}

func (dnl *decoderNL) DecodeInt8(v *int8) error {
	code, err := dnl.src.ReadByte()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code for Int8`)
	}
	if IsFixNumFamily(Code(code)) {
		*v = int8(code)
		return nil
	}
	switch Code(code) {
	case Int8:
		x, err := dnl.src.ReadUint8()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int8`)
		}
		*v = int8(x)
		return nil
	}
	return errors.Errorf(`msgpack: invalid numeric type %s for int8`, Code(code))
}

func (dnl *decoderNL) DecodeInt16(v *int16) error {
	code, err := dnl.src.ReadByte()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code for Int16`)
	}
	if IsFixNumFamily(Code(code)) {
		*v = int16(code)
		return nil
	}
	switch Code(code) {
	case Int16:
		x, err := dnl.src.ReadUint16()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int16`)
		}
		*v = int16(x)
		return nil
	case Int8:
		x, err := dnl.src.ReadUint8()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int16`)
		}
		*v = int16(x)
		return nil
	}
	return errors.Errorf(`msgpack: invalid numeric type %s for int16`, Code(code))
}

func (dnl *decoderNL) DecodeInt32(v *int32) error {
	code, err := dnl.src.ReadByte()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code for Int32`)
	}
	if IsFixNumFamily(Code(code)) {
		*v = int32(code)
		return nil
	}
	switch Code(code) {
	case Int32:
		x, err := dnl.src.ReadUint32()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int32`)
		}
		*v = int32(x)
		return nil
	case Int16:
		x, err := dnl.src.ReadUint16()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int32`)
		}
		*v = int32(x)
		return nil
	case Int8:
		x, err := dnl.src.ReadUint8()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int32`)
		}
		*v = int32(x)
		return nil
	}
	return errors.Errorf(`msgpack: invalid numeric type %s for int32`, Code(code))
}

func (dnl *decoderNL) DecodeInt64(v *int64) error {
	code, err := dnl.src.ReadByte()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code for Int64`)
	}
	if IsFixNumFamily(Code(code)) {
		*v = int64(code)
		return nil
	}
	switch Code(code) {
	case Int64:
		x, err := dnl.src.ReadUint64()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int64`)
		}
		*v = int64(x)
		return nil
	case Int32:
		x, err := dnl.src.ReadUint32()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int64`)
		}
		*v = int64(x)
		return nil
	case Int16:
		x, err := dnl.src.ReadUint16()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int64`)
		}
		*v = int64(x)
		return nil
	case Int8:
		x, err := dnl.src.ReadUint8()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for int64`)
		}
		*v = int64(x)
		return nil
	}
	return errors.Errorf(`msgpack: invalid numeric type %s for int64`, Code(code))
}

func (dnl *decoderNL) DecodeUint(v *uint) error {
	code, err := dnl.src.ReadByte()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code for Uint64`)
	}
	if IsFixNumFamily(Code(code)) {
		*v = uint(code)
		return nil
	}
	switch Code(code) {
	case Uint64:
		x, err := dnl.src.ReadUint64()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint`)
		}
		*v = uint(x)
		return nil
	case Uint32:
		x, err := dnl.src.ReadUint32()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint`)
		}
		*v = uint(x)
		return nil
	case Uint16:
		x, err := dnl.src.ReadUint16()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint`)
		}
		*v = uint(x)
		return nil
	case Uint8:
		x, err := dnl.src.ReadUint8()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint`)
		}
		*v = uint(x)
		return nil
	}
	return errors.Errorf(`msgpack: invalid numeric type %s for uint`, Code(code))
}

func (dnl *decoderNL) DecodeUint8(v *uint8) error {
	code, err := dnl.src.ReadByte()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code for Uint8`)
	}
	if IsFixNumFamily(Code(code)) {
		*v = uint8(code)
		return nil
	}
	switch Code(code) {
	case Uint8:
		x, err := dnl.src.ReadUint8()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint8`)
		}
		*v = uint8(x)
		return nil
	}
	return errors.Errorf(`msgpack: invalid numeric type %s for uint8`, Code(code))
}

func (dnl *decoderNL) DecodeUint16(v *uint16) error {
	code, err := dnl.src.ReadByte()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code for Uint16`)
	}
	if IsFixNumFamily(Code(code)) {
		*v = uint16(code)
		return nil
	}
	switch Code(code) {
	case Uint16:
		x, err := dnl.src.ReadUint16()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint16`)
		}
		*v = uint16(x)
		return nil
	case Uint8:
		x, err := dnl.src.ReadUint8()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint16`)
		}
		*v = uint16(x)
		return nil
	}
	return errors.Errorf(`msgpack: invalid numeric type %s for uint16`, Code(code))
}

func (dnl *decoderNL) DecodeUint32(v *uint32) error {
	code, err := dnl.src.ReadByte()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code for Uint32`)
	}
	if IsFixNumFamily(Code(code)) {
		*v = uint32(code)
		return nil
	}
	switch Code(code) {
	case Uint32:
		x, err := dnl.src.ReadUint32()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint32`)
		}
		*v = uint32(x)
		return nil
	case Uint16:
		x, err := dnl.src.ReadUint16()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint32`)
		}
		*v = uint32(x)
		return nil
	case Uint8:
		x, err := dnl.src.ReadUint8()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint32`)
		}
		*v = uint32(x)
		return nil
	}
	return errors.Errorf(`msgpack: invalid numeric type %s for uint32`, Code(code))
}

func (dnl *decoderNL) DecodeUint64(v *uint64) error {
	code, err := dnl.src.ReadByte()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read code for Uint64`)
	}
	if IsFixNumFamily(Code(code)) {
		*v = uint64(code)
		return nil
	}
	switch Code(code) {
	case Uint64:
		x, err := dnl.src.ReadUint64()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint64`)
		}
		*v = uint64(x)
		return nil
	case Uint32:
		x, err := dnl.src.ReadUint32()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint64`)
		}
		*v = uint64(x)
		return nil
	case Uint16:
		x, err := dnl.src.ReadUint16()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint64`)
		}
		*v = uint64(x)
		return nil
	case Uint8:
		x, err := dnl.src.ReadUint8()
		if err != nil {
			return errors.Wrap(err, `msgpack: failed to read payload for uint64`)
		}
		*v = uint64(x)
		return nil
	}
	return errors.Errorf(`msgpack: invalid numeric type %s for uint64`, Code(code))
}

func (dnl *decoderNL) DecodeFloat32(v *float32) error {
	code, x, err := dnl.src.ReadByteUint32()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read float32`)
	}

	if code != Float.Byte() {
		return errors.Errorf(`msgpack: expected Float, got %s`, Code(code))
	}

	*v = math.Float32frombits(x)
	return nil
}

func (dnl *decoderNL) DecodeFloat64(v *float64) error {
	code, x, err := dnl.src.ReadByteUint64()
	if err != nil {
		return errors.Wrap(err, `msgpack: failed to read float64`)
	}

	if code != Double.Byte() {
		return errors.Errorf(`msgpack: expected Double, got %s`, Code(code))
	}

	*v = math.Float64frombits(x)
	return nil
}
