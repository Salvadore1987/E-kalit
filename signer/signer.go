package signer

import (
	"syscall"
	"fmt"
	"encoding/binary"
	"runtime"
)

func IsConnected(ekd uintptr) bool {
	return isConnected(ekd);
}


func Create() uintptr {
	return create();
}

func Free(ekd uintptr)  {
	free(ekd)
}

func GetErrorCode(ekd uintptr) uint32 {
	return getErrorCode(ekd)
}

/*func GetErrorMessage(ekd uintptr) string {
	//return getErrorMessage(ekd)
}*/

func PtrToString(ptr uintptr) {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(ptr))
	fmt.Println(b)
}

/*func getErrorMessage(ekd uintptr) string {
	var mod = syscall.NewLazyDLL("libesigner.dll");
	var proc = mod.NewProc("EKDGetError");
	ret, _, _ := proc.Call(ekd);
	str := ""
	return str
}*/

func getErrorCode(ekd uintptr) uint32  {
	var mod *syscall.LazyDLL
	if runtime.GOARCH == "amd64" {
		mod = syscall.NewLazyDLL("./64/libesigner.dll")
	} else {
		mod = syscall.NewLazyDLL("./32/libesigner.dll")
	}
	var proc = mod.NewProc("EKDGetErrorCode");
	ret, _, err := proc.Call(ekd);
	if err != nil {
		fmt.Println(err);
	}
	return uint32(ret)
}

func create() uintptr  {
	var mod *syscall.LazyDLL
	if runtime.GOARCH == "amd64" {
		mod = syscall.NewLazyDLL("./64/libesigner.dll")
	} else {
		mod = syscall.NewLazyDLL("./32/libesigner.dll")
	}
	var proc = mod.NewProc("EKDNew");
	ret, _, err := proc.Call();
	if err != nil {
		fmt.Println(err);
	}
	return ret;
}

func isConnected(ekd uintptr) bool  {
	var mod *syscall.LazyDLL
	if runtime.GOARCH == "amd64" {
		mod = syscall.NewLazyDLL("./64/libesigner.dll")
	} else {
		mod = syscall.NewLazyDLL("./32/libesigner.dll")
	}
	var proc = mod.NewProc("EKDIsConnected");
	ret, _, err := proc.Call(ekd);
	if err != nil {
		fmt.Println(err);
	}
	if ret > 0 {
		return true
	}
	return false;
}

func free(ekd uintptr) {
	var mod *syscall.LazyDLL
	if runtime.GOARCH == "amd64" {
		mod = syscall.NewLazyDLL("./64/libesigner.dll")
	} else {
		mod = syscall.NewLazyDLL("./32/libesigner.dll")
	}
	var proc = mod.NewProc("EKDFree");
	_, _, _ = proc.Call(ekd);
}