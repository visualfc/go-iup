// Copyright (C) 2011 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iuppplot

/*
#include <stdlib.h>
#include <iup.h>
#include <iup_pplot.h>
*/
import "C"
import "unsafe"
import "vfc/iup"

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
	C.IupPlotBegin(h.CHandle(), C.int(strXdata))
}

func (h *IupPPlot) End() {
	C.IupPlotEnd(h.CHandle())
}

func (h *IupPPlot) Add(x, y float32) {
	C.IupPlotAdd(h.CHandle(), C.float(x), C.float(y))
}

func (h *IupPPlot) AddStr(x string, y float32) {
	cx := iup.NewCS(x)
	iup.FreeCS(cx)
	C.IupPlotAddStr(h.CHandle(), (*C.char)(cx), C.float(y))
}

func (h *IupPPlot) Insert(index, sample_index int, x, y float32) {
	C.IupPlotInsert(h.CHandle(), C.int(index), C.int(sample_index), C.float(x), C.float(y))
}

func (h *IupPPlot) InsertStr(index, sample_index int, x string, y float32) {
	cx := iup.NewCS(x)
	iup.FreeCS(cx)
	C.IupPlotInsertStr(h.CHandle(), C.int(index), C.int(sample_index), (*C.char)(cx), C.float(y))
}

func (h *IupPPlot) InsertPoints(index, sample_index int, x, y []float32) {
	count := len(x)
	if len(x) > len(y) {
		count = len(y)
	}
	C.IupPlotInsertPoints(h.CHandle(), C.int(index), C.int(sample_index), (*C.float)(unsafe.Pointer(&x[0])), (*C.float)(unsafe.Pointer(&y[0])), C.int(count))
}

func (h *IupPPlot) InsertStrPoints(index, sample_index int, x []string, y []float32) {
	count := len(x)
	if len(x) > len(y) {
		count = len(y)
	}
	cx := iup.NewCSA(x)
	defer iup.FreeCSA(cx)
	C.IupPlotInsertStrPoints(h.CHandle(), C.int(index), C.int(sample_index), (**C.char)(unsafe.Pointer(&cx[0])), (*C.float)(unsafe.Pointer(&y[0])), C.int(count))
}

func (h *IupPPlot) InsertAddPoints(index int, x, y []float32) {
	count := len(x)
	if len(x) > len(y) {
		count = len(y)
	}
	C.IupPlotAddPoints(h.CHandle(), C.int(index), (*C.float)(unsafe.Pointer(&x[0])), (*C.float)(unsafe.Pointer(&y[0])), C.int(count))
}

func (h *IupPPlot) AddStrPoints(index int, x []string, y []float32) {
	count := len(x)
	if len(x) > len(y) {
		count = len(y)
	}
	cx := iup.NewCSA(x)
	defer iup.FreeCSA(cx)
	C.IupPlotAddStrPoints(h.CHandle(), C.int(index), (**C.char)(unsafe.Pointer(&cx[0])), (*C.float)(unsafe.Pointer(&y[0])), C.int(count))
}

func (h *IupPPlot) Transform(x, y float32, ix, iy *int) {
	C.IupPlotTransform(h.CHandle(), C.float(x), C.float(y), (*C.int)(unsafe.Pointer(ix)), (*C.int)(unsafe.Pointer(iy)))
}

func (h *IupPPlot) PaintTo(cnv uintptr) {
	C.IupPlotPaintTo(h.CHandle(), unsafe.Pointer(cnv))
}
