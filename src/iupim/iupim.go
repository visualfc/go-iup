// iupim project iupim.go
package iupim

/*
#include <iup.h>
#include <iupim.h>
*/
import "C"
import "vfc/iup"

func LoadImage(filename string) *iup.Handle {
	cname := iup.NewCS(filename)
	defer iup.FreeCS(cname)
	return (*iup.Handle)(C.IupLoadImage((*C.char)(cname)))
}
