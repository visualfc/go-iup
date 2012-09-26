// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iup

/*
#include <iup.h>

#define LDESTROY_CB "LDESTROY_CB"
extern void goLDestroyCb(void*);
static void iupSetLDestroyCb(Ihandle* ih)
{
	IupSetCallback(ih,LDESTROY_CB,(Icallback)&goLDestroyCb);
}
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

func parser(a []interface{}) (children []IHandle, attrs []IAttribute, sattrs []string, funcs []interface{}) {
	for i := 0; i < len(a); i++ {
		switch v := a[i].(type) {
		case IHandle:
			children = append(children, v)
		case IAttribute:
			attrs = append(attrs, v)
		case []IAttribute:
			attrs = append(attrs, v[:]...)
		case string:
			sattrs = append(sattrs, v)
		default:
			funcs = append(funcs, v)
		}
	}
	return
}

func _ParserAttrs(a string) (attrs []IAttribute) {
	var n1, n2 int
	var s1, s2 string
	b1, b2 := true, false
	for n, v := range a {
		if b1 && v == '=' {
			s1 = a[n1:n]
			b2 = true
			n2 = n + 1
		} else if b2 && v == ',' {
			s2 = a[n2:n]
			b1 = true
			n1 = n + 1
			attrs = append(attrs, &Attribute{s1, s2})
		}
	}
	if b1 && b2 {
		attrs = append(attrs, &Attribute{s1, s2})
	}
	return
}

func NewClassInfo(classname string, fn func(*Handle, interface{}) bool) *ClassInfo {
	return &ClassInfo{classname, fn}
}

func (i ClassInfo) ClassName() string {
	return i.className
}

func (i *ClassInfo) SetCallback(h *Handle, fn interface{}) bool {
	return i.setCallback(h, fn)
}

func (i *ClassInfo) GetClassAttributes() []string {
	return GetClassAttributes(i.className)
}

func (i *ClassInfo) GetClassCallbacks() []string {
	return GetClassCallbacks(i.className)
}

func (i *ClassInfo) SetClassDefaultAttribute(name, value string) {
	SetClassDefaultAttribute(i.className, name, value)
}

func GetClassAttributes(classname string) []string {
	c := NewCS(classname)
	defer FreeCS(c)
	total_count := C.IupGetClassAttributes(c, nil, 0)
	if total_count == 0 {
		return nil
	}
	parray := make([]*C.char, total_count)
	count := C.IupGetClassAttributes(c, &parray[0], total_count)
	return FromCSA(parray[0:count])
}

func GetClassCallbacks(classname string) []string {
	c := NewCS(classname)
	defer FreeCS(c)
	total_count := C.IupGetClassCallbacks(c, nil, 0)
	if total_count == 0 {
		return nil
	}
	parray := make([]*C.char, total_count)
	count := C.IupGetClassCallbacks(c, &parray[0], total_count)
	return FromCSA(parray[0:count])
}

func SetClassDefaultAttribute(classname, name, value string) {
	c := NewCS(classname)
	defer FreeCS(c)
	n := NewCS(name)
	defer FreeCS(n)
	v := NewCS(value)
	defer FreeCS(v)
	C.IupSetClassDefaultAttribute(c, n, v)
}

var classMap = make(map[string]*ClassInfo)

func RegisterClass(name string, info *ClassInfo) {
	classMap[name] = info
}

type IHandle interface {
	Native() uintptr
}

var handleMap = make(map[uintptr]IHandle)

func toHandle(native *C.Ihandle) IHandle {
	return handleMap[uintptr(unsafe.Pointer(native))]
}

func ptoHandle(native unsafe.Pointer) IHandle {
	return handleMap[uintptr(native)]
}

func toNative(h IHandle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(h.Native()))
}

//export goLDestroyCb
func goLDestroyCb(p unsafe.Pointer) {
	delete(handleMap, uintptr(p))
}

// iup.Attr(key,value), iup.Attrs(k1,v1,v2,v2,...) ,
// string , IHandle , func(arg *Action)
func New(classname string, a ...interface{}) *Handle {
	cname := NewCS(classname)
	defer FreeCS(cname)
	ptr := C.IupCreate(cname)
	if ptr == nil {
		return nil
	}
	h := NewHandle(ptr)
	if len(a) == 0 {
		return h
	}
	children, attrs, sattrs, funcs := parser(a)
	for i := 0; i < len(children); i++ {
		h.Append(children[i])
	}
	for i := 0; i < len(attrs); i++ {
		h.SetAttribute(attrs[i].Name(), attrs[i].Value())
	}
	for i := 0; i < len(sattrs); i++ {
		h.SetAttributes(sattrs[i])
	}
	for i := 0; i < len(funcs); i++ {
		h.SetCallback(funcs[i])
	}
	return h
}

type ClassInfo struct {
	className   string
	setCallback func(*Handle, interface{}) bool
}

func SetHandle(name string, h IHandle) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupSetHandle(cname, toNative(h))
}

func GetHandle(name string) IHandle {
	cname := NewCS(name)
	defer FreeCS(cname)
	return toHandle(C.IupGetHandle(cname))
}

func GetName(h IHandle) string {
	return FromCS(C.IupGetName(toNative(h)))
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

type Handle struct {
	p     *C.Ihandle
	funcs map[string]interface{}
}

func NewHandle(p *C.Ihandle) *Handle {
	h := new(Handle)
	h.p = p
	h.funcs = make(map[string]interface{})
	handleMap[uintptr(unsafe.Pointer(p))] = h
	C.iupSetLDestroyCb(p)
	return h
}

func (h *Handle) Native() uintptr {
	return (uintptr)(unsafe.Pointer(h.p))
}

func (h *Handle) GetClassType() string {
	return FromCS(C.IupGetClassType(h.p))
}

func (h *Handle) GetClassName() string {
	return FromCS(C.IupGetClassName(h.p))
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
	return FromCS(C.IupGetName(h.p))
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

func (h *Handle) Detach() {
	C.IupDetach(h.p)
}

func (h *Handle) Update() {
	C.IupUpdate(h.p)
}

func (h *Handle) UpdateChildren() {
	C.IupUpdateChildren(h.p)
}

func (h *Handle) Redraw(children int) {
	C.IupRedraw(h.p, C.int(children))
}

func (h *Handle) Refresh() {
	C.IupRefresh(h.p)
}

func (h *Handle) RefreshChildren() {
	C.IupRefreshChildren(h.p)
}

func (h *Handle) ConvertXYToPos(x, y int) int {
	return int(C.IupConvertXYToPos(h.p, C.int(x), C.int(y)))
}

func (h *Handle) Destroy() {
	C.IupDestroy(h.p)
}

func (h *Handle) Append(child IHandle) error {
	if C.IupAppend(h.p, toNative(child)) == nil {
		return &Error{"IupAppend"}
	}
	return nil
}

func (h *Handle) Insert(ref_child, child IHandle) error {
	if C.IupInsert(h.p, toNative(ref_child), toNative(child)) == nil {
		return &Error{"IupInsert"}
	}
	return nil
}

func (h *Handle) GetChild(pos int) IHandle {
	return toHandle(C.IupGetChild(h.p, C.int(pos)))
}

func (h *Handle) GetChildPos(child IHandle) int {
	return int(C.IupGetChildPos(h.p, toNative(child)))
}

func (h *Handle) GetChildCount() int {
	return int(C.IupGetChildCount(h.p))
}

func (h *Handle) GetNextChild(child IHandle) IHandle {
	return toHandle(C.IupGetNextChild(h.p, toNative(child)))
}

func (h *Handle) GetBrother() IHandle {
	return toHandle(C.IupGetBrother(h.p))
}

func (h *Handle) GetParent() IHandle {
	return toHandle(C.IupGetParent(h.p))
}

func (h *Handle) GetDialog() IHandle {
	return toHandle(C.IupGetDialog(h.p))
}

func (h *Handle) GetDialogChild(name string) IHandle {
	cname := NewCS(name)
	defer FreeCS(cname)
	return toHandle(C.IupGetDialogChild(h.p, cname))
}

func (h *Handle) Reparent(new_parent, ref_child IHandle) int {
	return int(C.IupReparent(h.p, toNative(new_parent), toNative(ref_child)))
}

func (h *Handle) Popup(x, y int) int {
	return int(C.IupPopup(h.p, C.int(x), C.int(y)))
}

func (h *Handle) PopupA() int {
	return int(C.IupPopup(h.p, C.IUP_CURRENT, C.IUP_CURRENT))
}

func (h *Handle) PopupC() int {
	return int(C.IupPopup(h.p, C.IUP_CENTER, C.IUP_CENTER))
}

func (h *Handle) Show() int {
	return int(C.IupShow(h.p))
}

func (h *Handle) ShowXY(x, y int) int {
	return int(C.IupShowXY(h.p, C.int(x), C.int(y)))
}

func (h *Handle) Hide() int {
	return int(C.IupHide(h.p))
}

func (h *Handle) Map() int {
	return int(C.IupMap(h.p))
}

func (h *Handle) Unmap() {
	C.IupUnmap(h.p)
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
	C.IupStoreAttribute(h.p, cname, cvalue)
}

func (h *Handle) GetAttribute(name string) string {
	cname := NewCS(name)
	defer FreeCS(cname)
	return FromCS(C.IupGetAttribute(h.p, cname))
}

func (h *Handle) SetAttributeId(name string, id int, value string) {
	cname := NewCS(name)
	defer FreeCS(cname)
	cvalue := NewCS(value)
	defer FreeCS(cvalue)
	C.IupStoreAttributeId(h.p, cname, C.int(id), cvalue)
}

func (h *Handle) GetAttributeId(name string, id int) string {
	cname := NewCS(name)
	defer FreeCS(cname)
	return FromCS(C.IupGetAttributeId(h.p, cname, C.int(id)))
}

func (h *Handle) SetAttributes(values string) {
	cvalues := NewCS(values)
	defer FreeCS(cvalues)
	C.IupSetAttributes(h.p, cvalues)
}

func (h *Handle) GetAttributes() string {
	return FromCS(C.IupGetAttributes(h.p))
}

func (h *Handle) SetAttributeData(name string, value uintptr) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupSetAttribute(h.p, cname, (*C.char)(unsafe.Pointer(value)))
}

func (h *Handle) GetAttributeData(name string) uintptr {
	cname := NewCS(name)
	defer FreeCS(cname)
	return uintptr(unsafe.Pointer(C.IupGetAttribute(h.p, cname)))
}

func (h *Handle) SetAttributeDataId(name string, id int, value uintptr) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupSetAttributeId(h.p, cname, C.int(id), (*C.char)(unsafe.Pointer(value)))
}

func (h *Handle) GetAttributeDataId(name string, id int) uintptr {
	cname := NewCS(name)
	defer FreeCS(cname)
	return uintptr(unsafe.Pointer(C.IupGetAttributeId(h.p, cname, C.int(id))))
}

func (h *Handle) SetAttributeHandle(name string, h_named IHandle) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupSetAttributeHandle(h.p, cname, toNative(h_named))
}

func (h *Handle) GetAttributeHandle(name string) IHandle {
	cname := NewCS(name)
	defer FreeCS(cname)
	return toHandle(C.IupGetAttributeHandle(h.p, cname))
}

func (h *Handle) ResetAttribute(name string) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupResetAttribute(h.p, cname)
}

func (h *Handle) GetInt(name string) int {
	cname := NewCS(name)
	defer FreeCS(cname)
	return int(C.IupGetInt(h.p, cname))
}

func (h *Handle) GetIntId(name string, id int) int {
	cname := NewCS(name)
	defer FreeCS(cname)
	return int(C.IupGetIntId(h.p, cname, C.int(id)))
}

func (h *Handle) GetIntInt(name string) (count int, n1 int, n2 int) {
	cname := NewCS(name)
	defer FreeCS(cname)
	var v1, v2 C.int
	r := C.IupGetIntInt(h.p, cname, &v1, &v2)
	return int(r), int(v1), int(v2)
}

func (h *Handle) GetFloat(name string) float32 {
	cname := NewCS(name)
	defer FreeCS(cname)
	return float32(C.IupGetFloat(h.p, cname))
}

func (h *Handle) GetFloatId(name string, id int) float32 {
	cname := NewCS(name)
	defer FreeCS(cname)
	return float32(C.IupGetFloatId(h.p, cname, C.int(id)))
}

func (h *Handle) NextField() IHandle {
	return toHandle(C.IupNextField(h.p))
}

func (h *Handle) PreviousField() IHandle {
	return toHandle(C.IupPreviousField(h.p))
}

func GetFocus() IHandle {
	return toHandle(C.IupGetFocus())
}

func SetFocus(h IHandle) IHandle {
	return toHandle(C.IupSetFocus(toNative(h)))
}

func (h *Handle) SaveClassAttributes() {
	C.IupSaveClassAttributes(h.p)
}

func (h *Handle) CopyClassAttributes(src IHandle) {
	C.IupCopyClassAttributes(h.p, toNative(src))
}

func (h *Handle) SetCallbackProc(name string, proc uintptr) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupSetCallback(h.p, cname, (C.Icallback)(unsafe.Pointer(proc)))
}

//func GetAttribute(h IHandle, name string) string {
//	return h.GetHandle().GetAttribute(name)
//}

//func GetAttributeData(h IHandle, name string) uintptr {
//	return h.GetHandle().GetAttributeData(name)
//}

//func SetAttribute(h IHandle, name string, value string) {
//	h.GetHandle().SetAttribute(name, value)
//}

//func SetAttributeData(h IHandle, name string, value uintptr) {
//	h.GetHandle().SetAttributeData(name, value)
//}

func init() {
	RegisterAllClass()
}
