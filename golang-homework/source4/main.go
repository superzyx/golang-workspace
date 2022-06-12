package main

import (
	"fmt"
	"source4/pkg/rolling"
	"sync"
)

func main() {
	sm := rolling.NewStatusMap()
	wg := sync.WaitGroup{}
	for i := 0; i<15; i++ {
		test_status := rolling.NewStatusNode()
		test_status.Add(1, 2, 3, 4)
		wg.Add(1)
		go func() {
			sm.ModifyStatus(test_status)
			fmt.Println(sm.GetStatus())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(sm.GetStatus())
}

//func main() {
//	sm := rolling.NewStatusMap()
//	test_status := rolling.NewStatusNode()
//	test_status.Add(1, 2, 3,4)
//	sm.ModifyStatus(test_status)
//	for k, v := range sm.GetAllStatus() {
//		fmt.Println(k, v)
//	}
//}
