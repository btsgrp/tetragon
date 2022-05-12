// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs_darwin.go

package socket

type iovec struct {
	Base *byte
	Len  uint32
}

type msghdr struct {
	Name       *byte
	Namelen    uint32
	Iov        *iovec
	Iovlen     int32
	Control    *byte
	Controllen uint32
	Flags      int32
}

type cmsghdr struct {
	Len   uint32
	Level int32
	Type  int32
}

const (
	sizeofIovec  = 0x8
	sizeofMsghdr = 0x1c
)