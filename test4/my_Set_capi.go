package main

/*
#cgo CXXFLAGS: -std=c++11
#include "my_Set_capi.h"
*/
import "C"
import "unsafe"

type cgo_myset C.myset_T

func newmyset() *cgo_myset {
	p := C.newmyset()
	return (*cgo_myset)(p)
}
func cgo_insert(p *cgo_myset, i interface{}) {
	C.insertMyset((*C.myset_T)(p), unsafe.Pointer(&i))
}
func cgo_size(p *cgo_myset) int {
	return int(C.sizeMyset((*C.myset_T)(p)))
}
func cgo_find(p *cgo_myset, i interface{}) bool {
	if bool(C.findMyset((*C.myset_T)(p), unsafe.Pointer(&i))) {
		return true
	}
	return false
}
func cgo_erase(p *cgo_myset, i interface{}) {
	C.eraseMyset((*C.myset_T)(p), unsafe.Pointer(&i))
}
