// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iupim

/*
#cgo CFLAGS : -I../../libs/iup/include
#cgo LDFLAGS : -L../../libs/iup -liup -liupim
#cgo linux CFLAGS : -I../../libs/im/include
#cgo linux LDFLAGS : -L../../libs/im -lim
#include <iup.h>
#include <iupim.h>
*/
import "C"
import "github.com/visualfc/go-iup/iup"

func LoadImage(filename string) *iup.Handle {
	cname := iup.NewCS(filename)
	defer iup.FreeCS(cname)
	return iup.NewHandle(C.IupLoadImage((*C.char)(cname)))
}
