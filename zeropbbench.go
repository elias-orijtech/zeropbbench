package gcbench

/*

#include <stdint.h>

static uint32_t foreign_func(uint32_t *buf, uint32_t len) {
	uint32_t accum = 0;
	for (int i = 0; i < len; i++) {
		accum += buf[i];
	}
	return accum;
}
*/
import "C"
import "unsafe"

func foreign_func(args []uint32) {
	C.foreign_func((*C.uint32_t)(unsafe.Pointer(unsafe.SliceData(args))), C.uint32_t(len(args)))
}
