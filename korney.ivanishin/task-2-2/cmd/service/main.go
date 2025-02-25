package main

import (
	"fmt"
	"github.com/cafe/internal/process"
)

func main() {
	dishHeap, err := process.NewDishHeap()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = process.ScanDishHeap(dishHeap)
	if err != nil {
		fmt.Println(err)
		return
	}

	designPos, err := process.ScanDesignPos()
	if err != nil {
		fmt.Println(err)
		return
	}

	designScore, err := process.GetDesignScore(dishHeap, designPos)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(designScore)
}
