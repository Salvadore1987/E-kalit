package main

import (
	"E-kalit/signer"
	"fmt"
)

func main()  {

	ekd := signer.Create()
	if signer.IsConnected(ekd) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	signer.PtrToString(ekd)
	signer.Free(ekd)
}
