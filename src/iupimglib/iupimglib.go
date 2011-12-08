// Copyright (C) 2011 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iupimglib

/*
#include <stdlib.h>
#include <iup.h>
*/
import "C"
import "vfc/iup"

func Open() *iup.Error {
	C.IupImageLibOpen()
	return nil
}