// +build darwin,!kqueue,cgo,!go1.10

package notify

/*
#include <CoreServices/CoreServices.h>
*/
import "C"

var refZero = (*C.struct___CFAllocator)(nil)
