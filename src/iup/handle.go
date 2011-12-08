// Copyright (C) 2011 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iup

/*
#include <iup.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type IAttribute interface {
	Name() string
	Value() string
}

type Attribute struct {
	name  string
	value string
}

func (a *Attribute) Name() string {
	return a.name
}

func (a *Attribute) Value() string {
	return a.value
}

func Attr(name string, value string) IAttribute {
	return &Attribute{name, value}
}

func Attrs(a ...string) (attrs []IAttribute) {
	count := len(a)
	if count == 0 || count%2 != 0 {
		return nil
	}
	for i := 0; i < count; i += 2 {
		attrs = append(attrs, &Attribute{a[i], a[i+1]})
	}
	return
}

type IHandle interface {
	GetHandle() *Handle
	CHandle() *C.Ihandle
}

type Handle C.Ihandle

func SetHandle(name string, h IHandle) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupSetHandle(cname, h.CHandle())
}

func GetHandle(name string) *Handle {
	cname := NewCS(name)
	defer FreeCS(cname)
	return (*Handle)(C.IupGetHandle(cname))
}

func GetName(h IHandle) string {
	return FromCS(C.IupGetName(h.CHandle()))
}

func GetAllNames() []string {
	total_count := C.IupGetAllNames(nil, 0)
	if total_count == 0 {
		return nil
	}
	parray := make([]*C.char, total_count)
	count := C.IupGetAllNames(&parray[0], total_count)
	return FromCSA(parray[0:count])
}

func GetAllDialogs() []string {
	total_count := C.IupGetAllDialogs(nil, 0)
	if total_count == 0 {
		return nil
	}
	parray := make([]*C.char, total_count)
	count := C.IupGetAllDialogs(&parray[0], total_count)
	return FromCSA(parray[0:count])
}

func (h *Handle) GetHandle() *Handle {
	return h
}

func (h *Handle) CHandle() *C.Ihandle {
	return (*C.Ihandle)(h)
}

func (h *Handle) GetClassType() string {
	return FromCS(C.IupGetClassType(h.CHandle()))
}

func (h *Handle) GetClassName() string {
	return FromCS(C.IupGetClassName(h.CHandle()))
}

func (h *Handle) GetClassInfo() *ClassInfo {
	classname := h.GetClassName()
	if info, ok := classMap[classname]; ok {
		return info
	}
	return nil
}

func (h *Handle) SetName(name string) {
	SetHandle(name, h)
}

func (h *Handle) GetName() string {
	return FromCS(C.IupGetName(h.CHandle()))
}

func (h *Handle) SetCallback(fn interface{}) error {
	classname := h.GetClassName()
	if info, ok := classMap[classname]; ok {
		if info.SetCallback(h, fn) {
			return nil
		}
	}
	return &Error{fmt.Sprintf("ClassName %s Unsupport %v", classname, fn)}
}

func (h *Handle) Attach(ih IHandle) {
	h = ih.GetHandle()
}

func (h *Handle) Detach() {
	C.IupDetach(h.CHandle())
}

func (h *Handle) Update() {
	C.IupUpdate(h.CHandle())
}

func (h *Handle) UpdateChildren() {
	C.IupUpdateChildren(h.CHandle())
}

func (h *Handle) Redraw(children int) {
	C.IupRedraw(h.CHandle(), C.int(children))
}

func (h *Handle) Refresh() {
	C.IupRefresh(h.CHandle())
}

func (h *Handle) RefreshChildren() {
	C.IupRefreshChildren(h.CHandle())
}

func (h *Handle) ConvertXYToPos(x, y int) int {
	return int(C.IupConvertXYToPos(h.CHandle(), C.int(x), C.int(y)))
}

func (h *Handle) Destroy() {
	C.IupDestroy(h.CHandle())
}

func (h *Handle) Append(child IHandle) error {
	if C.IupAppend(h.CHandle(), child.CHandle()) == nil {
		return &Error{"IupAppend"}
	}
	return nil
}

func (h *Handle) Insert(ref_child, child IHandle) error {
	if C.IupInsert(h.CHandle(), ref_child.CHandle(), child.CHandle()) == nil {
		return &Error{"IupInsert"}
	}
	return nil
}

func (h *Handle) GetChild(pos int) *Handle {
	return (*Handle)(C.IupGetChild(h.CHandle(), C.int(pos)))
}

func (h *Handle) GetChildPos(child IHandle) int {
	return int(C.IupGetChildPos(h.CHandle(), child.CHandle()))
}

func (h *Handle) GetChildCount() int {
	return int(C.IupGetChildCount(h.CHandle()))
}

func (h *Handle) GetNextChild(child IHandle) *Handle {
	return (*Handle)(C.IupGetNextChild(h.CHandle(), child.CHandle()))
}

func (h *Handle) GetBrother() *Handle {
	return (*Handle)(C.IupGetBrother(h.CHandle()))
}

func (h *Handle) GetParent() *Handle {
	return (*Handle)(C.IupGetParent(h.CHandle()))
}

func (h *Handle) GetDialog() *Handle {
	return (*Handle)(C.IupGetDialog(h.CHandle()))
}

func (h *Handle) GetDialogChild(name string) *Handle {
	cname := NewCS(name)
	defer FreeCS(cname)
	return (*Handle)(C.IupGetDialogChild(h.CHandle(), cname))
}

func (h *Handle) Reparent(new_parent, ref_child IHandle) int {
	return int(C.IupReparent(h.CHandle(), new_parent.CHandle(), ref_child.CHandle()))
}

func (h *Handle) Popup(x, y int) int {
	return int(C.IupPopup(h.CHandle(), C.int(x), C.int(y)))
}

func (h *Handle) PopupA() int {
	return int(C.IupPopup(h.CHandle(), C.IUP_CURRENT, C.IUP_CURRENT))
}

func (h *Handle) PopupC() int {
	return int(C.IupPopup(h.CHandle(), C.IUP_CENTER, C.IUP_CENTER))
}

func (h *Handle) Show() int {
	return int(C.IupShow(h.CHandle()))
}

func (h *Handle) ShowXY(x, y int) int {
	return int(C.IupShowXY(h.CHandle(), C.int(x), C.int(y)))
}

func (h *Handle) Hide() int {
	return int(C.IupHide(h.CHandle()))
}

func (h *Handle) Map() int {
	return int(C.IupMap(h.CHandle()))
}

func (h *Handle) Unmap() {
	C.IupUnmap(h.CHandle())
}

func (h *Handle) SetAttrs(a ...string) error {
	count := len(a)
	if count%2 != 0 {
		return &Error{"key number != value number"}
	}
	for i := 0; i < count; i += 2 {
		h.SetAttribute(a[i], a[i+1])
	}
	return nil
}

func (h *Handle) SetAttribute(name string, value string) {
	cname := NewCS(name)
	defer FreeCS(cname)
	cvalue := NewCS(value)
	defer FreeCS(cvalue)
	C.IupStoreAttribute(h.CHandle(), cname, cvalue)
}

func (h *Handle) GetAttribute(name string) string {
	cname := NewCS(name)
	defer FreeCS(cname)
	return FromCS(C.IupGetAttribute(h.CHandle(), cname))
}

func (h *Handle) SetAttributeId(name string, id int, value string) {
	cname := NewCS(name)
	defer FreeCS(cname)
	cvalue := NewCS(value)
	defer FreeCS(cvalue)
	C.IupStoreAttributeId(h.CHandle(), cname, C.int(id), cvalue)
}

func (h *Handle) GetAttributeId(name string, id int) string {
	cname := NewCS(name)
	defer FreeCS(cname)
	return FromCS(C.IupGetAttributeId(h.CHandle(), cname, C.int(id)))
}

func (h *Handle) SetAttributes(values string) {
	cvalues := NewCS(values)
	defer FreeCS(cvalues)
	C.IupSetAttributes(h.CHandle(), cvalues)
}

func (h *Handle) GetAttributes() string {
	return FromCS(C.IupGetAttributes(h.CHandle()))
}

func (h *Handle) SetAttributeData(name string, value uintptr) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupSetAttribute(h.CHandle(), cname, (*C.char)(unsafe.Pointer(value)))
}

func (h *Handle) GetAttributeData(name string) uintptr {
	cname := NewCS(name)
	defer FreeCS(cname)
	return uintptr(unsafe.Pointer(C.IupGetAttribute(h.CHandle(), cname)))
}

func (h *Handle) SetAttributeDataId(name string, id int, value uintptr) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupSetAttributeId(h.CHandle(), cname, C.int(id), (*C.char)(unsafe.Pointer(value)))
}

func (h *Handle) GetAttributeDataId(name string, id int) uintptr {
	cname := NewCS(name)
	defer FreeCS(cname)
	return uintptr(unsafe.Pointer(C.IupGetAttributeId(h.CHandle(), cname, C.int(id))))
}

func (h *Handle) SetAttributeHandle(name string, h_named IHandle) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupSetAttributeHandle(h.CHandle(), cname, h_named.CHandle())
}

func (h *Handle) GetAttributeHandle(name string) *Handle {
	cname := NewCS(name)
	defer FreeCS(cname)
	return (*Handle)(C.IupGetAttributeHandle(h.CHandle(), cname))
}

func (h *Handle) ResetAttribute(name string) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupResetAttribute(h.CHandle(), cname)
}

func (h *Handle) GetInt(name string) int {
	cname := NewCS(name)
	defer FreeCS(cname)
	return int(C.IupGetInt(h.CHandle(), cname))
}

func (h *Handle) GetIntId(name string, id int) int {
	cname := NewCS(name)
	defer FreeCS(cname)
	return int(C.IupGetIntId(h.CHandle(), cname, C.int(id)))
}

func (h *Handle) GetIntInt(name string) (count int, n1 int, n2 int) {
	cname := NewCS(name)
	defer FreeCS(cname)
	var v1, v2 C.int
	r := C.IupGetIntInt(h.CHandle(), cname, &v1, &v2)
	return int(r), int(v1), int(v2)
}

func (h *Handle) GetFloat(name string) float32 {
	cname := NewCS(name)
	defer FreeCS(cname)
	return float32(C.IupGetFloat(h.CHandle(), cname))
}

func (h *Handle) GetFloatId(name string, id int) float32 {
	cname := NewCS(name)
	defer FreeCS(cname)
	return float32(C.IupGetFloatId(h.CHandle(), cname, C.int(id)))
}

func (h *Handle) NextField() IHandle {
	return (*Handle)(C.IupNextField(h.CHandle()))
}

func (h *Handle) PreviousField() IHandle {
	return (*Handle)(C.IupPreviousField(h.CHandle()))
}

func GetFocus() IHandle {
	return (*Handle)(C.IupGetFocus())
}

func SetFocus(h IHandle) IHandle {
	return (*Handle)(C.IupSetFocus(h.CHandle()))
}

func (h *Handle) SaveClassAttributes() {
	C.IupSaveClassAttributes(h.CHandle())
}

func (h *Handle) CopyClassAttributes(src IHandle) {
	C.IupCopyClassAttributes(h.CHandle(), src.CHandle())
}

func (h *Handle) SetCallbackProc(name string, proc uintptr) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupSetCallback(h.CHandle(), cname, (C.Icallback)(unsafe.Pointer(proc)))
}
