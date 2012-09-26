// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iuppplot

/*
#cgo CFLAGS : -I../../libs/iup/include
#cgo LDFLAGS: -L../../libs/iup -liup -liup_pplot
#cgo linux CFLAGS : -I../../libs/cd/include
#cgo linux LDFLAGS: -L../../libs/cd -liupcd -lcd
#include <stdlib.h>
#include <iup.h>
#include <iup_pplot.h>
*/
import "C"
import "unsafe"
import "github.com/visualfc/go-iup/iup"

func toNative(h iup.IHandle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(h.Native()))
}

func Open() *iup.Error {
	C.IupPPlotOpen()
	return nil
}

type IupPPlot struct {
	*iup.Handle
}

func PPlot(a ...interface{}) *IupPPlot {
	return &IupPPlot{iup.PPlot(a...)}
}

func AttachPPloat(h *iup.Handle) *IupPPlot {
	return &IupPPlot{h}
}

func (h *IupPPlot) Begin(strXdata int) {
	C.IupPPlotBegin(toNative(h), C.int(strXdata))
}

func (h *IupPPlot) End() {
	C.IupPPlotEnd(toNative(h))
}

func (h *IupPPlot) Add(x, y float32) {
	C.IupPPlotAdd(toNative(h), C.float(x), C.float(y))
}

func (h *IupPPlot) AddStr(x string, y float32) {
	cx := iup.NewCS(x)
	iup.FreeCS(cx)
	C.IupPPlotAddStr(toNative(h), (*C.char)(cx), C.float(y))
}

func (h *IupPPlot) Insert(index, sample_index int, x, y float32) {
	C.IupPPlotInsert(toNative(h), C.int(index), C.int(sample_index), C.float(x), C.float(y))
}

func (h *IupPPlot) InsertStr(index, sample_index int, x string, y float32) {
	cx := iup.NewCS(x)
	iup.FreeCS(cx)
	C.IupPPlotInsertStr(toNative(h), C.int(index), C.int(sample_index), (*C.char)(cx), C.float(y))
}

func (h *IupPPlot) InsertPoints(index, sample_index int, x, y []float32) {
	count := len(x)
	if len(x) > len(y) {
		count = len(y)
	}
	C.IupPPlotInsertPoints(toNative(h), C.int(index), C.int(sample_index), (*C.float)(unsafe.Pointer(&x[0])), (*C.float)(unsafe.Pointer(&y[0])), C.int(count))
}

func (h *IupPPlot) InsertStrPoints(index, sample_index int, x []string, y []float32) {
	count := len(x)
	if len(x) > len(y) {
		count = len(y)
	}
	cx := iup.NewCSA(x)
	defer iup.FreeCSA(cx)
	C.IupPPlotInsertStrPoints(toNative(h), C.int(index), C.int(sample_index), (**C.char)(unsafe.Pointer(&cx[0])), (*C.float)(unsafe.Pointer(&y[0])), C.int(count))
}

func (h *IupPPlot) InsertAddPoints(index int, x, y []float32) {
	count := len(x)
	if len(x) > len(y) {
		count = len(y)
	}
	C.IupPPlotAddPoints(toNative(h), C.int(index), (*C.float)(unsafe.Pointer(&x[0])), (*C.float)(unsafe.Pointer(&y[0])), C.int(count))
}

func (h *IupPPlot) AddStrPoints(index int, x []string, y []float32) {
	count := len(x)
	if len(x) > len(y) {
		count = len(y)
	}
	cx := iup.NewCSA(x)
	defer iup.FreeCSA(cx)
	C.IupPPlotAddStrPoints(toNative(h), C.int(index), (**C.char)(unsafe.Pointer(&cx[0])), (*C.float)(unsafe.Pointer(&y[0])), C.int(count))
}

func (h *IupPPlot) Transform(x, y float32, ix, iy *int) {
	C.IupPPlotTransform(toNative(h), C.float(x), C.float(y), (*C.int)(unsafe.Pointer(ix)), (*C.int)(unsafe.Pointer(iy)))
}

func (h *IupPPlot) PaintTo(cnv uintptr) {
	C.IupPPlotPaintTo(toNative(h), unsafe.Pointer(cnv))
}
