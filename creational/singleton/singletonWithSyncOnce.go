package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type single2 struct{}

var singleInstance2 *single2

func getInstance2() *single2 {
	if singleInstance2 == nil {
		once.Do(func() {
			fmt.Println("Creating single2 instance now.")
			singleInstance2 = &single2{}
		})

	} else {
		fmt.Println("Single instance2 already created.")
	}

	return singleInstance2
}
