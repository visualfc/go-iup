// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iup

/*
#include <stdlib.h>
#include <windows.h>
char *Utf8toAnsi(char *buf, int len)
{
	int ulen = MultiByteToWideChar(CP_UTF8,0,buf,len,0,0);
	unsigned short *utf = malloc(ulen*2);
	ulen = MultiByteToWideChar(CP_UTF8,0,buf,len,utf,ulen*2);
	int clen = WideCharToMultiByte(CP_ACP,0,utf,ulen,0,0,0,0);
	char *ansi = malloc(clen+1);
	clen = WideCharToMultiByte(CP_ACP,0,utf,ulen,ansi,clen,0,0);
	ansi[clen] = '\0';
	free(utf);
	return ansi;
}

char *Utf8toAnsiN(char *buf, int len, int max)
{
	int ulen = MultiByteToWideChar(CP_UTF8,0,buf,len,0,0);
	unsigned short *utf = malloc(ulen*2);
	ulen = MultiByteToWideChar(CP_UTF8,0,buf,len,utf,ulen*2);
	int clen = WideCharToMultiByte(CP_ACP,0,utf,ulen,0,0,0,0);
	char *ansi = malloc(max);
	memset(ansi,max,0);
	clen = WideCharToMultiByte(CP_ACP,0,utf,ulen,ansi,clen,0,0);
	ansi[clen] = '\0';
	free(utf);
	return ansi;
}

char *AnsiToUtf8(char *buf, int *size)
{
	int len = strlen(buf);
	int ulen = MultiByteToWideChar(CP_ACP,0,buf,len,0,0);
	unsigned short *utf = malloc(ulen*2);
	ulen = MultiByteToWideChar(CP_ACP,0,buf,len,utf,ulen*2);
	len = WideCharToMultiByte(CP_UTF8,0,utf,ulen,0,0,0,0);
	char *utf8 = malloc(len+1);
	len = WideCharToMultiByte(CP_UTF8,0,utf,ulen,utf8,len,0,0);
	utf8[len] = '\0';
	free(utf);	
	*size = len;
	return utf8;
}
*/
import "C"
import "unsafe"
import "reflect"

func NewCS(s string) *C.char {
	if len(s) == 0 {
		return C.CString(s)
	}
	head := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return C.Utf8toAnsi((*C.char)(unsafe.Pointer(head.Data)), C.int(len(s)))
}

func NewCSN(s string, max int) *C.char {
	head := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return C.Utf8toAnsiN((*C.char)(unsafe.Pointer(head.Data)), C.int(len(s)), C.int(max))
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
	if cs == nil {
		return ""
	}
	var size C.int
	buf := C.AnsiToUtf8(cs, &size)
	defer C.free(unsafe.Pointer(buf))
	return string((*[1 << 30]byte)(unsafe.Pointer(buf))[0:size])
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
