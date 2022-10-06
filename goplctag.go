// Mozilla Public License 2.0, and LGPL License.

// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package goplctag

/*
#cgo pkg-config: libplctag
#include "libplctag.h"
#include "callbacks.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// DecodeError function as declared in include/libplctag.h:119
func DecodeError(err int32) string {
	cerr, cerrAllocMap := (C.int)(err), cgoAllocsUnknown
	__ret := C.plc_tag_decode_error(cerr)
	runtime.KeepAlive(cerrAllocMap)
	__v := packPCharString(__ret)
	return __v
}

// SetDebugLevel function as declared in include/libplctag.h:137
func SetDebugLevel(debugLevel int32) {
	cdebugLevel, cdebugLevelAllocMap := (C.int)(debugLevel), cgoAllocsUnknown
	C.plc_tag_set_debug_level(cdebugLevel)
	runtime.KeepAlive(cdebugLevelAllocMap)
}

// CheckLibVersion function as declared in include/libplctag.h:163
func CheckLibVersion(reqMajor int32, reqMinor int32, reqPatch int32) int32 {
	creqMajor, creqMajorAllocMap := (C.int)(reqMajor), cgoAllocsUnknown
	creqMinor, creqMinorAllocMap := (C.int)(reqMinor), cgoAllocsUnknown
	creqPatch, creqPatchAllocMap := (C.int)(reqPatch), cgoAllocsUnknown
	__ret := C.plc_tag_check_lib_version(creqMajor, creqMinor, creqPatch)
	runtime.KeepAlive(creqPatchAllocMap)
	runtime.KeepAlive(creqMinorAllocMap)
	runtime.KeepAlive(creqMajorAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Create function as declared in include/libplctag.h:196
func Create(attribStr string, timeout int32) int32 {
	cattribStr, cattribStrAllocMap := unpackPCharString(attribStr)
	ctimeout, ctimeoutAllocMap := (C.int)(timeout), cgoAllocsUnknown
	__ret := C.plc_tag_create(cattribStr, ctimeout)
	runtime.KeepAlive(ctimeoutAllocMap)
	runtime.KeepAlive(cattribStrAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Shutdown function as declared in include/libplctag.h:229
func Shutdown() {
	C.plc_tag_shutdown()
}

// UnregisterCallback function as declared in include/libplctag.h:300
func UnregisterCallback(tagId int32) int32 {
	ctagId, ctagIdAllocMap := (C.int32_t)(tagId), cgoAllocsUnknown
	__ret := C.plc_tag_unregister_callback(ctagId)
	runtime.KeepAlive(ctagIdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// UnregisterLogger function as declared in include/libplctag.h:339
func UnregisterLogger() int32 {
	__ret := C.plc_tag_unregister_logger()
	__v := (int32)(__ret)
	return __v
}

// Lock function as declared in include/libplctag.h:357
func Lock(tag int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	__ret := C.plc_tag_lock(ctag)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Unlock function as declared in include/libplctag.h:368
func Unlock(tag int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	__ret := C.plc_tag_unlock(ctag)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Abort function as declared in include/libplctag.h:386
func Abort(tag int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	__ret := C.plc_tag_abort(ctag)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Destroy function as declared in include/libplctag.h:399
func Destroy(tag int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	__ret := C.plc_tag_destroy(ctag)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Read function as declared in include/libplctag.h:416
func Read(tag int32, timeout int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	ctimeout, ctimeoutAllocMap := (C.int)(timeout), cgoAllocsUnknown
	__ret := C.plc_tag_read(ctag, ctimeout)
	runtime.KeepAlive(ctimeoutAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Status function as declared in include/libplctag.h:430
func Status(tag int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	__ret := C.plc_tag_status(ctag)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Write function as declared in include/libplctag.h:448
func Write(tag int32, timeout int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	ctimeout, ctimeoutAllocMap := (C.int)(timeout), cgoAllocsUnknown
	__ret := C.plc_tag_write(ctag, ctimeout)
	runtime.KeepAlive(ctimeoutAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetIntAttribute function as declared in include/libplctag.h:458
func GetIntAttribute(tag int32, attribName string, defaultValue int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	cattribName, cattribNameAllocMap := unpackPCharString(attribName)
	cdefaultValue, cdefaultValueAllocMap := (C.int)(defaultValue), cgoAllocsUnknown
	__ret := C.plc_tag_get_int_attribute(ctag, cattribName, cdefaultValue)
	runtime.KeepAlive(cdefaultValueAllocMap)
	runtime.KeepAlive(cattribNameAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// SetIntAttribute function as declared in include/libplctag.h:459
func SetIntAttribute(tag int32, attribName string, newValue int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	cattribName, cattribNameAllocMap := unpackPCharString(attribName)
	cnewValue, cnewValueAllocMap := (C.int)(newValue), cgoAllocsUnknown
	__ret := C.plc_tag_set_int_attribute(ctag, cattribName, cnewValue)
	runtime.KeepAlive(cnewValueAllocMap)
	runtime.KeepAlive(cattribNameAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetSize function as declared in include/libplctag.h:461
func GetSize(tag int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	__ret := C.plc_tag_get_size(ctag)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// SetSize function as declared in include/libplctag.h:463
func SetSize(tag int32, newSize int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	cnewSize, cnewSizeAllocMap := (C.int)(newSize), cgoAllocsUnknown
	__ret := C.plc_tag_set_size(ctag, cnewSize)
	runtime.KeepAlive(cnewSizeAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetBit function as declared in include/libplctag.h:465
func GetBit(tag int32, offsetBit int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffsetBit, coffsetBitAllocMap := (C.int)(offsetBit), cgoAllocsUnknown
	__ret := C.plc_tag_get_bit(ctag, coffsetBit)
	runtime.KeepAlive(coffsetBitAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// SetBit function as declared in include/libplctag.h:466
func SetBit(tag int32, offsetBit int32, val int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffsetBit, coffsetBitAllocMap := (C.int)(offsetBit), cgoAllocsUnknown
	cval, cvalAllocMap := (C.int)(val), cgoAllocsUnknown
	__ret := C.plc_tag_set_bit(ctag, coffsetBit, cval)
	runtime.KeepAlive(cvalAllocMap)
	runtime.KeepAlive(coffsetBitAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetUint64 function as declared in include/libplctag.h:468
func GetUint64(tag int32, offset int32) uint32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	__ret := C.plc_tag_get_uint64(ctag, coffset)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (uint32)(__ret)
	return __v
}

// SetUint64 function as declared in include/libplctag.h:469
func SetUint64(tag int32, offset int32, val uint32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	cval, cvalAllocMap := (C.uint64_t)(val), cgoAllocsUnknown
	__ret := C.plc_tag_set_uint64(ctag, coffset, cval)
	runtime.KeepAlive(cvalAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetInt64 function as declared in include/libplctag.h:471
func GetInt64(tag int32, offset int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	__ret := C.plc_tag_get_int64(ctag, coffset)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// SetInt64 function as declared in include/libplctag.h:472
func SetInt64(arg0 int32, offset int32, val int32) int32 {
	carg0, carg0AllocMap := (C.int32_t)(arg0), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	cval, cvalAllocMap := (C.int64_t)(val), cgoAllocsUnknown
	__ret := C.plc_tag_set_int64(carg0, coffset, cval)
	runtime.KeepAlive(cvalAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(carg0AllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetUint32 function as declared in include/libplctag.h:475
func GetUint32(tag int32, offset int32) uint32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	__ret := C.plc_tag_get_uint32(ctag, coffset)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (uint32)(__ret)
	return __v
}

// SetUint32 function as declared in include/libplctag.h:476
func SetUint32(tag int32, offset int32, val uint32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	cval, cvalAllocMap := (C.uint32_t)(val), cgoAllocsUnknown
	__ret := C.plc_tag_set_uint32(ctag, coffset, cval)
	runtime.KeepAlive(cvalAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetInt32 function as declared in include/libplctag.h:478
func GetInt32(tag int32, offset int32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	__ret := C.plc_tag_get_int32(ctag, coffset)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// SetInt32 function as declared in include/libplctag.h:479
func SetInt32(arg0 int32, offset int32, val int32) int32 {
	carg0, carg0AllocMap := (C.int32_t)(arg0), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	cval, cvalAllocMap := (C.int32_t)(val), cgoAllocsUnknown
	__ret := C.plc_tag_set_int32(carg0, coffset, cval)
	runtime.KeepAlive(cvalAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(carg0AllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetUint16 function as declared in include/libplctag.h:482
func GetUint16(tag int32, offset int32) uint16 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	__ret := C.plc_tag_get_uint16(ctag, coffset)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (uint16)(__ret)
	return __v
}

// SetUint16 function as declared in include/libplctag.h:483
func SetUint16(tag int32, offset int32, val uint16) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	cval, cvalAllocMap := (C.uint16_t)(val), cgoAllocsUnknown
	__ret := C.plc_tag_set_uint16(ctag, coffset, cval)
	runtime.KeepAlive(cvalAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetInt16 function as declared in include/libplctag.h:485
func GetInt16(tag int32, offset int32) int16 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	__ret := C.plc_tag_get_int16(ctag, coffset)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int16)(__ret)
	return __v
}

// SetInt16 function as declared in include/libplctag.h:486
func SetInt16(arg0 int32, offset int32, val int16) int32 {
	carg0, carg0AllocMap := (C.int32_t)(arg0), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	cval, cvalAllocMap := (C.int16_t)(val), cgoAllocsUnknown
	__ret := C.plc_tag_set_int16(carg0, coffset, cval)
	runtime.KeepAlive(cvalAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(carg0AllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetUint8 function as declared in include/libplctag.h:489
func GetUint8(tag int32, offset int32) byte {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	__ret := C.plc_tag_get_uint8(ctag, coffset)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (byte)(__ret)
	return __v
}

// SetUint8 function as declared in include/libplctag.h:490
func SetUint8(tag int32, offset int32, val byte) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	cval, cvalAllocMap := (C.uint8_t)(val), cgoAllocsUnknown
	__ret := C.plc_tag_set_uint8(ctag, coffset, cval)
	runtime.KeepAlive(cvalAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetInt8 function as declared in include/libplctag.h:492
func GetInt8(tag int32, offset int32) int8 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	__ret := C.plc_tag_get_int8(ctag, coffset)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int8)(__ret)
	return __v
}

// SetInt8 function as declared in include/libplctag.h:493
func SetInt8(arg0 int32, offset int32, val int8) int32 {
	carg0, carg0AllocMap := (C.int32_t)(arg0), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	cval, cvalAllocMap := (C.int8_t)(val), cgoAllocsUnknown
	__ret := C.plc_tag_set_int8(carg0, coffset, cval)
	runtime.KeepAlive(cvalAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(carg0AllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetFloat64 function as declared in include/libplctag.h:496
func GetFloat64(tag int32, offset int32) float64 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	__ret := C.plc_tag_get_float64(ctag, coffset)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (float64)(__ret)
	return __v
}

// SetFloat64 function as declared in include/libplctag.h:497
func SetFloat64(tag int32, offset int32, val float64) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	cval, cvalAllocMap := (C.double)(val), cgoAllocsUnknown
	__ret := C.plc_tag_set_float64(ctag, coffset, cval)
	runtime.KeepAlive(cvalAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetFloat32 function as declared in include/libplctag.h:499
func GetFloat32(tag int32, offset int32) float32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	__ret := C.plc_tag_get_float32(ctag, coffset)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (float32)(__ret)
	return __v
}

// SetFloat32 function as declared in include/libplctag.h:500
func SetFloat32(tag int32, offset int32, val float32) int32 {
	ctag, ctagAllocMap := (C.int32_t)(tag), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	cval, cvalAllocMap := (C.float)(val), cgoAllocsUnknown
	__ret := C.plc_tag_set_float32(ctag, coffset, cval)
	runtime.KeepAlive(cvalAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(ctagAllocMap)
	__v := (int32)(__ret)
	return __v
}

// SetRawBytes function as declared in include/libplctag.h:503
func SetRawBytes(id int32, offset int32, buffer []byte, bufferLength int32) int32 {
	cid, cidAllocMap := (C.int32_t)(id), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	cbuffer, cbufferAllocMap := copyPUint8Bytes((*sliceHeader)(unsafe.Pointer(&buffer)))
	cbufferLength, cbufferLengthAllocMap := (C.int)(bufferLength), cgoAllocsUnknown
	__ret := C.plc_tag_set_raw_bytes(cid, coffset, cbuffer, cbufferLength)
	runtime.KeepAlive(cbufferLengthAllocMap)
	runtime.KeepAlive(cbufferAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(cidAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetRawBytes function as declared in include/libplctag.h:504
func GetRawBytes(id int32, offset int32, buffer []byte, bufferLength int32) int32 {
	cid, cidAllocMap := (C.int32_t)(id), cgoAllocsUnknown
	coffset, coffsetAllocMap := (C.int)(offset), cgoAllocsUnknown
	cbuffer, cbufferAllocMap := copyPUint8Bytes((*sliceHeader)(unsafe.Pointer(&buffer)))
	cbufferLength, cbufferLengthAllocMap := (C.int)(bufferLength), cgoAllocsUnknown
	__ret := C.plc_tag_get_raw_bytes(cid, coffset, cbuffer, cbufferLength)
	runtime.KeepAlive(cbufferLengthAllocMap)
	runtime.KeepAlive(cbufferAllocMap)
	runtime.KeepAlive(coffsetAllocMap)
	runtime.KeepAlive(cidAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetString function as declared in include/libplctag.h:508
func GetString(tagId int32, stringStartOffset int32, buffer []byte, bufferLength int32) int32 {
	ctagId, ctagIdAllocMap := (C.int32_t)(tagId), cgoAllocsUnknown
	cstringStartOffset, cstringStartOffsetAllocMap := (C.int)(stringStartOffset), cgoAllocsUnknown
	cbuffer, cbufferAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&buffer)))
	cbufferLength, cbufferLengthAllocMap := (C.int)(bufferLength), cgoAllocsUnknown
	__ret := C.plc_tag_get_string(ctagId, cstringStartOffset, cbuffer, cbufferLength)
	runtime.KeepAlive(cbufferLengthAllocMap)
	runtime.KeepAlive(cbufferAllocMap)
	runtime.KeepAlive(cstringStartOffsetAllocMap)
	runtime.KeepAlive(ctagIdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// SetString function as declared in include/libplctag.h:509
func SetString(tagId int32, stringStartOffset int32, stringVal string) int32 {
	ctagId, ctagIdAllocMap := (C.int32_t)(tagId), cgoAllocsUnknown
	cstringStartOffset, cstringStartOffsetAllocMap := (C.int)(stringStartOffset), cgoAllocsUnknown
	cstringVal, cstringValAllocMap := unpackPCharString(stringVal)
	__ret := C.plc_tag_set_string(ctagId, cstringStartOffset, cstringVal)
	runtime.KeepAlive(cstringValAllocMap)
	runtime.KeepAlive(cstringStartOffsetAllocMap)
	runtime.KeepAlive(ctagIdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetStringLength function as declared in include/libplctag.h:510
func GetStringLength(tagId int32, stringStartOffset int32) int32 {
	ctagId, ctagIdAllocMap := (C.int32_t)(tagId), cgoAllocsUnknown
	cstringStartOffset, cstringStartOffsetAllocMap := (C.int)(stringStartOffset), cgoAllocsUnknown
	__ret := C.plc_tag_get_string_length(ctagId, cstringStartOffset)
	runtime.KeepAlive(cstringStartOffsetAllocMap)
	runtime.KeepAlive(ctagIdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetStringCapacity function as declared in include/libplctag.h:511
func GetStringCapacity(tagId int32, stringStartOffset int32) int32 {
	ctagId, ctagIdAllocMap := (C.int32_t)(tagId), cgoAllocsUnknown
	cstringStartOffset, cstringStartOffsetAllocMap := (C.int)(stringStartOffset), cgoAllocsUnknown
	__ret := C.plc_tag_get_string_capacity(ctagId, cstringStartOffset)
	runtime.KeepAlive(cstringStartOffsetAllocMap)
	runtime.KeepAlive(ctagIdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetStringTotalLength function as declared in include/libplctag.h:512
func GetStringTotalLength(tagId int32, stringStartOffset int32) int32 {
	ctagId, ctagIdAllocMap := (C.int32_t)(tagId), cgoAllocsUnknown
	cstringStartOffset, cstringStartOffsetAllocMap := (C.int)(stringStartOffset), cgoAllocsUnknown
	__ret := C.plc_tag_get_string_total_length(ctagId, cstringStartOffset)
	runtime.KeepAlive(cstringStartOffsetAllocMap)
	runtime.KeepAlive(ctagIdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// CreateEx function as declared in goplctag/callbacks.h:33
func CreateEx(attribStr string, callbackFunction CreateExCallback, userdata unsafe.Pointer, timeout int32) int32 {
	cattribStr, cattribStrAllocMap := unpackPCharString(attribStr)
	ccallbackFunction, ccallbackFunctionAllocMap := callbackFunction.PassValue()
	cuserdata, cuserdataAllocMap := userdata, cgoAllocsUnknown
	ctimeout, ctimeoutAllocMap := (C.int)(timeout), cgoAllocsUnknown
	__ret := C.go_plc_tag_create_ex(cattribStr, ccallbackFunction, cuserdata, ctimeout)
	runtime.KeepAlive(ctimeoutAllocMap)
	runtime.KeepAlive(cuserdataAllocMap)
	runtime.KeepAlive(ccallbackFunctionAllocMap)
	runtime.KeepAlive(cattribStrAllocMap)
	__v := (int32)(__ret)
	return __v
}

// RegisterCallbackEx function as declared in goplctag/callbacks.h:37
func RegisterCallbackEx(tagId int32, callbackFunction RegisterExCallback, userdata unsafe.Pointer) int32 {
	ctagId, ctagIdAllocMap := (C.int32_t)(tagId), cgoAllocsUnknown
	ccallbackFunction, ccallbackFunctionAllocMap := callbackFunction.PassValue()
	cuserdata, cuserdataAllocMap := userdata, cgoAllocsUnknown
	__ret := C.go_plc_tag_register_callback_ex(ctagId, ccallbackFunction, cuserdata)
	runtime.KeepAlive(cuserdataAllocMap)
	runtime.KeepAlive(ccallbackFunctionAllocMap)
	runtime.KeepAlive(ctagIdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// RegisterLogger function as declared in goplctag/callbacks.h:41
func RegisterLogger(callbackFunction RegisterLoggerCallback) int32 {
	ccallbackFunction, ccallbackFunctionAllocMap := callbackFunction.PassValue()
	__ret := C.go_plc_tag_register_logger(ccallbackFunction)
	runtime.KeepAlive(ccallbackFunctionAllocMap)
	__v := (int32)(__ret)
	return __v
}
