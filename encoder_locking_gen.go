package msgpack

// Auto-generated by internal/cmd/genencoder/genencoder.go. DO NOT EDIT!

import (
	"time"
)

func (d *encoder) Encode(v interface{}) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.Encode(v)
}

func (d *encoder) EncodeArray(v interface{}) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeArray(v)
}

func (d *encoder) EncodeArrayHeader(v int) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeArrayHeader(v)
}

func (d *encoder) EncodeBool(v bool) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeBool(v)
}

func (d *encoder) EncodeBytes(v []byte) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeBytes(v)
}

func (d *encoder) EncodeExt(v EncodeMsgpacker) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeExt(v)
}

func (d *encoder) EncodeExtHeader(v int) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeExtHeader(v)
}

func (d *encoder) EncodeExtType(v EncodeMsgpacker) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeExtType(v)
}

func (d *encoder) EncodeMap(v interface{}) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeMap(v)
}

func (d *encoder) EncodeNegativeFixNum(v int8) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeNegativeFixNum(v)
}

func (d *encoder) EncodeNil() error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeNil()
}

func (d *encoder) EncodePositiveFixNum(v uint8) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodePositiveFixNum(v)
}

func (d *encoder) EncodeString(v string) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeString(v)
}

func (d *encoder) EncodeStruct(v interface{}) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeStruct(v)
}

func (d *encoder) EncodeTime(v time.Time) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeTime(v)
}

func (d *encoder) Writer() Writer {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.Writer()
}

func (d *encoder) EncodeInt(v int) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeInt(v)
}

func (d *encoder) EncodeInt8(v int8) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeInt8(v)
}

func (d *encoder) EncodeInt16(v int16) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeInt16(v)
}

func (d *encoder) EncodeInt32(v int32) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeInt32(v)
}

func (d *encoder) EncodeInt64(v int64) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeInt64(v)
}

func (d *encoder) EncodeUint(v uint) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeUint(v)
}

func (d *encoder) EncodeUint8(v uint8) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeUint8(v)
}

func (d *encoder) EncodeUint16(v uint16) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeUint16(v)
}

func (d *encoder) EncodeUint32(v uint32) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeUint32(v)
}

func (d *encoder) EncodeUint64(v uint64) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeUint64(v)
}

func (d *encoder) EncodeFloat32(v float32) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeFloat32(v)
}

func (d *encoder) EncodeFloat64(v float64) error {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.nl.EncodeFloat64(v)
}
