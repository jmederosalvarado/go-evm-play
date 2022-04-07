package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm/runtime"
)

func main() {
	ret, state, err := runtime.Execute(common.Hex2Bytes("6C63FFFFFFFF60005260046000F3600052601360006000F0"), nil, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}
