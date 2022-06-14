package utils

// #cgo LDFLAGS:-L. -lclosessl
//#include<closessl.h>
import "C"
import (
	"unsafe"
)

func Encode(msg []byte) []byte {
	goUnsafePointer := C.CBytes(msg)
	cPointer := (*C.uint8_t)(goUnsafePointer)
	C.closessl_encrypt(cPointer, (C.ulong)(len(msg)))
	response := C.GoBytes(unsafe.Pointer(cPointer), (C.int)(len(msg)))
	return response
}

func Decode(msg []byte) []byte {
	goUnsafePointer := C.CBytes(msg)
	cPointer := (*C.uint8_t)(goUnsafePointer)
	C.closessl_decrypt(cPointer, (C.ulong)(len(msg)))
	response := C.GoBytes(unsafe.Pointer(cPointer), (C.int)(len(msg)))
	return response
}
