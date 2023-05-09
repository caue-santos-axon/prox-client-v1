package mutex

import (
	"syscall"
	"unsafe"
)

var (
	kernel32        = syscall.NewLazyDLL("kernel32.dll")
	procCreateMutex = kernel32.NewProc("CreateMutexW")
)

// Create a mutex to assure that only one instance of Prox Client will open
func CreateMutex(name string) (uintptr, error) {
	p, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return 0, err
	}
	ret, _, err := procCreateMutex.Call(
		0,
		0,
		uintptr(unsafe.Pointer(p)),
	)
	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}
