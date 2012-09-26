// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iup

/*
#include <iup.h>
*/
import "C"

/************************************************************************/
/*                   GO-IUP Version		                                */
/************************************************************************/
const (
	Name        = "GO-IUP - Iup Golang Binding"
	CopyRight   = "Copyright (C) 2011-2012 visualfc"
	Version     = "1.1"
	VersionDate = "2012.9.26"
)

/************************************************************************/
/*                   IUP Version		                                */
/************************************************************************/
const (
	IupName          = C.IUP_NAME
	IupCopyRight     = C.IUP_COPYRIGHT
	IupDescription   = C.IUP_DESCRIPTION
	IupVersion       = C.IUP_VERSION
	IupVersionNumber = C.IUP_VERSION_NUMBER
	IupVersionDate   = C.IUP_VERSION_DATE
)

/************************************************************************/
/*                   Common Return Values                               */
/************************************************************************/
const (
	ERROR   int = C.IUP_ERROR
	NOERROR     = C.IUP_NOERROR
	OPENDE      = C.IUP_OPENED
	INVALID     = C.IUP_INVALID
)

/************************************************************************/
/*                   Callback Return Values                             */
/************************************************************************/
const (
	IGNORE   int = C.IUP_IGNORE
	DEFAULT      = C.IUP_DEFAULT
	CLOSE        = C.IUP_CLOSE
	CONTINUE     = C.IUP_CONTINUE
)

/************************************************************************/
/*           IupPopup and IupShowXY Parameter Values                    */
/************************************************************************/
const (
	CENTER       int = C.IUP_CENTER       //0xFFFF  /* 65535 */
	LEFT             = C.IUP_LEFT         // 0xFFFE  /* 65534 */
	RIGHT            = C.IUP_RIGHT        // 0xFFFD  /* 65533 */
	MOUSEPOS         = C.IUP_MOUSEPOS     // 0xFFFC  /* 65532 */
	CURIUP           = C.IUP_CURRENT      // 0xFFFB  /* 65531 */
	CENTERPARENT     = C.IUP_CENTERPARENT // 0xFFFA  /* 65530 */
)

/************************************************************************/
/*               SHOW_CB Callback Values                                */
/************************************************************************/
const (
	SHOW     = C.IUP_SHOW
	RESTORE  = C.IUP_RESTORE
	MINIMIZE = C.IUP_MINIMIZE
	MAXIMIZE = C.IUP_MAXIMIZE
	HIDE     = C.IUP_HIDE
)

/************************************************************************/
/*               SCROLL_CB Callback Values                              */
/************************************************************************/
const (
	SBUP      = C.IUP_SBUP
	SBDN      = C.IUP_SBUP
	SBPGUP    = C.IUP_SBPGUP
	SBPGDN    = C.IUP_SBPGDN
	SBPOSV    = C.IUP_SBPOSV
	SBDRAGV   = C.IUP_SBDRAGV
	SBLEFT    = C.IUP_SBLEFT
	SBRIGHT   = C.IUP_SBRIGHT
	SBPGLEFT  = C.IUP_SBPGLEFT
	SBPGRIGHT = C.IUP_SBPGRIGHT
	SBPOSH    = C.IUP_SBPOSH
	SBDRAGH   = C.IUP_SBDRAGH
)

type KeyState int

func (k KeyState) Key() int {
	return int(k % 256)
}

func (k KeyState) Xkey() int {
	return int(k % 128)
}

func (k KeyState) IsShift() bool {
	return k > 256 && k < 512
}

func (k KeyState) IsCtrl() bool {
	return k > 512 && k < 768
}

func (k KeyState) IsAlt() bool {
	return k > 768 && k < 1024
}

func (k KeyState) IsSys() bool {
	return k > 1024 && k < 1280
}

/************************************************************************/
/*               Mouse Button Values and State                         */
/************************************************************************/
const (
	BUTTON1 = '1'
	BUTTON2 = '2'
	BUTTON3 = '3'
	BUTTON4 = '4'
	BUTTON5 = '5'
)

type MouseState struct {
	S string
}

func (s *MouseState) IsShift() bool {
	return s.S[0] == 'S'
}

func (s *MouseState) IsControl() bool {
	return s.S[1] == 'C'
}

func (s *MouseState) IsButton1() bool {
	return s.S[2] == '1'
}

func (s *MouseState) IsButton2() bool {
	return s.S[3] == '2'
}

func (s *MouseState) IsButton3() bool {
	return s.S[4] == '3'
}

func (s *MouseState) IsDouble() bool {
	return s.S[5] == 'D'
}

func (s *MouseState) IsAlt() bool {
	return s.S[6] == 'A'
}

func (s *MouseState) IsSys() bool {
	return s.S[7] == 'Y'
}

func (s *MouseState) IsButton4() bool {
	return s.S[8] == '4'
}

func (s *MouseState) IsButton5() bool {
	return s.S[9] == '5'
}

/************************************************************************/
/*                      Pre-Defined Masks                               */
/************************************************************************/
const (
	MASK_FLOAT  = "[+/-]?(/d+/.?/d*|/./d+)"
	MASK_UFLOAT = "(/d+/.?/d*|/./d+)"
	MASK_EFLOAT = "[+/-]?(/d+/.?/d*|/./d+)([eE][+/-]?/d+)?"
	MASK_INT    = "[+/-]?/d+"
	MASK_UINT   = "/d+"
)

/************************************************************************/
/*                   Record Input Modes                                 */
/************************************************************************/
const (
	RECBINARY   = C.IUP_RECBINARY
	IUP_RECTEXT = C.IUP_RECTEXT
)
