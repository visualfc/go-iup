// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iupimglib

/*
#cgo CFLAGS : -I../../libs/iup/include
#cgo LDFLAGS: -L../../libs/iup -liup -liupimglib

#include <stdlib.h>
#include <iup.h>
*/
import "C"
import "github.com/visualfc/go-iup/iup"

func Open() *iup.Error {
	C.IupImageLibOpen()
	return nil
}
