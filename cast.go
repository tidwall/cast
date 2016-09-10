package cast

import (
	"reflect"
	"unsafe"
)

// ToString will quickly return a string representation of a []byte without
// allocating new memory or having to copy data. Converting a []byte -> string
// will result in a mutable string. Editing the originial bytes will change the
// string too.
func ToString(b []byte) string {
	// A string is 16 bytes and []byte is 24 bytes, therefore []byte can be
	// directly assigned to a string.
	return *(*string)(unsafe.Pointer(&b))
}

// ToBytes will quickly return a []byte representation of a string without
// allocating new memory or having to copy data. Converting a string -> []byte
// will result in an immutable byte slice. Editing will cause a panic.
func ToBytes(s string) []byte {
	// A string is smaller than a []byte so we need to create a new byte
	// header and fill in the fields.
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{sh.Data, sh.Len, sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
