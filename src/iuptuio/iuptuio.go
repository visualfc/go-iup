// Copyright (C) 2011 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iuptuio

/*
#include <stdlib.h>
#include <iup.h>
#include <iuptuio.h>
*/
import "C"
import "vfc/iup"

func Open() *iup.Error {
	r := C.IupTuioOpen()
	if r == C.IUP_ERROR {
		return &iup.Error{"IupTuiopen"}
	}
	return nil
}

func TuioClient(port int) *iup.Handle {
	return (*iup.Handle)(C.IupTuioClient(C.int(port)))
}
