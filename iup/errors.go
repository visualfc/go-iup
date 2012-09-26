// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iup

import (
	"fmt"
)

type Error struct {
	Name string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error %s", e.Name)
}
