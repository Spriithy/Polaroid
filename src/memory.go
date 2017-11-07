package yasm

import "unsafe"

func getu8(p uintptr) u8       { return *(*u8)(unsafe.Pointer(p)) }
func geti8(p uintptr) i8       { return *(*i8)(unsafe.Pointer(p)) }
func getu16(p uintptr) u16     { return *(*u16)(unsafe.Pointer(p)) }
func geti16(p uintptr) i16     { return *(*i16)(unsafe.Pointer(p)) }
func getu32(p uintptr) u32     { return *(*u32)(unsafe.Pointer(p)) }
func geti32(p uintptr) i32     { return *(*i32)(unsafe.Pointer(p)) }
func getu64(p uintptr) u64     { return *(*u64)(unsafe.Pointer(p)) }
func geti64(p uintptr) i64     { return *(*i64)(unsafe.Pointer(p)) }
func getf32(p uintptr) f32     { return *(*f32)(unsafe.Pointer(p)) }
func getf64(p uintptr) f64     { return *(*f64)(unsafe.Pointer(p)) }
func getPtr(p uintptr) uintptr { return *(*uintptr)(unsafe.Pointer(p)) }

func setu8(p uintptr, v u8)       { *(*u8)(unsafe.Pointer(p)) = v }
func seti8(p uintptr, v i8)       { *(*i8)(unsafe.Pointer(p)) = v }
func setu16(p uintptr, v u16)     { *(*u16)(unsafe.Pointer(p)) = v }
func seti16(p uintptr, v i16)     { *(*i16)(unsafe.Pointer(p)) = v }
func setu32(p uintptr, v u32)     { *(*u32)(unsafe.Pointer(p)) = v }
func seti32(p uintptr, v i32)     { *(*i32)(unsafe.Pointer(p)) = v }
func setu64(p uintptr, v u64)     { *(*u64)(unsafe.Pointer(p)) = v }
func seti64(p uintptr, v i64)     { *(*i64)(unsafe.Pointer(p)) = v }
func setf32(p uintptr, v f32)     { *(*f32)(unsafe.Pointer(p)) = v }
func setf64(p uintptr, v f64)     { *(*f64)(unsafe.Pointer(p)) = v }
func setPtr(p uintptr, v uintptr) { *(*uintptr)(unsafe.Pointer(p)) = v }

func (c *cpu) getLocal8(offset int) u8        { return getu8(c.fp + uintptr(offset)) }
func (c *cpu) getLocal16(offset int) u16      { return getu16(c.fp + uintptr(offset)) }
func (c *cpu) getLocal32(offset int) u32      { return getu32(c.fp + uintptr(offset)) }
func (c *cpu) getLocal64(offset int) u64      { return getu64(c.fp + uintptr(offset)) }
func (c *cpu) getLocalf32(offset int) f32     { return getf32(c.fp + uintptr(offset)) }
func (c *cpu) getLocalf64(offset int) f64     { return getf64(c.fp + uintptr(offset)) }
func (c *cpu) getLocalPtr(offset int) uintptr { return getPtr(c.fp + uintptr(offset)) }

func (c *cpu) setLocal8(offset int, v u8)        { setu8(c.fp+uintptr(offset), v) }
func (c *cpu) setLocal16(offset int, v u16)      { setu16(c.fp+uintptr(offset), v) }
func (c *cpu) setLocal32(offset int, v u32)      { setu32(c.fp+uintptr(offset), v) }
func (c *cpu) setLocal64(offset int, v u64)      { setu64(c.fp+uintptr(offset), v) }
func (c *cpu) setLocalf32(offset int, v f32)     { setf32(c.fp+uintptr(offset), v) }
func (c *cpu) setLocalf64(offset int, v f64)     { setf64(c.fp+uintptr(offset), v) }
func (c *cpu) setLocalPtr(offset int, v uintptr) { setPtr(c.fp+uintptr(offset), v) }