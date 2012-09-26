// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iup

/*
#include <stdlib.h>
#include <memory.h>
char *CStringN(char *buf, int len, int max)
{
	char *ansi = malloc(max);
	memcpy(ansi,buf,len);
	ansi[len] = '\0';
	return ansi;
}
*/
import "C"
import "unsafe"
import "reflect"

func NewCSN(s string, max int) *C.char {
	head := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return C.CStringN((*C.char)(unsafe.Pointer(head.Data)), C.int(len(s)), C.int(max))
}

func NewCS(s string) *C.char {
	return C.CString(s)
}

func FreeCS(cs *C.char) {
	C.free(unsafe.Pointer(cs))
}

func NewCSA(sa []string) []*C.char {
	max := len(sa)
	csa := make([]*C.char, max+1)
	for i := 0; i < max; i++ {
		csa[i] = NewCS(sa[i])
	}
	csa[max] = nil
	return csa
}

func FreeCSA(csa []*C.char) {
	max := len(csa)
	for i := 0; i < max; i++ {
		if csa[i] != nil {
			C.free(unsafe.Pointer(csa[i]))
		}
	}
}

func FromCS(cs *C.char) string {
	return C.GoString(cs)
}

func FromCSA(csa []*C.char) []string {
	size := len(csa)
	if size == 0 {
		return nil
	}
	array := make([]string, size)
	for i := 0; i < size; i++ {
		array[i] = FromCS(csa[i])
	}
	return array
}
