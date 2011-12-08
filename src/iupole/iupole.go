// Copyright (C) 2011 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iupole

/*
#include <stdlib.h>
#include <iup.h>
#include <iupole.h>
*/
import "C"
import "vfc/iup"

func Open() *iup.Error {
	r := C.IupOleControlOpen()
	if r == C.IUP_ERROR {
		return &iup.Error{"IupOleControlOpen"}
	}
	return nil
}

func OleControl(progid string) *iup.Handle {
	id := iup.NewCS(progid)
	defer iup.FreeCS(id)
	return (*iup.Handle)(C.IupOleControl((*C.char)(id)))
}
