package jimage

type Header struct {
	magic          uint32
	majorVersion   uint16
	minorVersion   uint16
	flags          uint32
	resourceCount  uint32
	tableLength    uint32
	attributesSize uint32
	stringsSize    uint32
}
