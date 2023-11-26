package storage

import (
	"bytes"
	"encoding/binary"
)

type SlotInfo struct {
	Name   string
	Type   string
	Slot   int64
	Size   int64
	Offset int64
}

func (si *SlotInfo) Bytes() []byte {
	buffer := new(bytes.Buffer)

	// Write the Name as bytes
	buffer.Write([]byte(si.Name))

	// Write the Type as bytes
	buffer.Write([]byte(si.Type))

	// Convert Slot, Size, and Offset to bytes and write
	binary.Write(buffer, binary.BigEndian, si.Slot)
	binary.Write(buffer, binary.BigEndian, si.Size)
	binary.Write(buffer, binary.BigEndian, si.Offset)

	return buffer.Bytes()
}
