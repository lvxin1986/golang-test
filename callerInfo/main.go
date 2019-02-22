/*
Copyright 2019 Tiger

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreedto in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

created by lvxin for project golang-test at 19-2-22 上午11:55
*/
package main

import (
	"runtime"
	"fmt"
	"github.com/lvxin1986/golang-test/util/stack"
)

func  methodB()  {
	stack.PrintRuntimeFullStack()
}
func  methodBs()  {
	pc := make([]uintptr, 1024)
	for skip := 0; ; skip++ {
		n := runtime.Callers(skip, pc)
		if n <= 0 {
			break
		}
		fmt.Printf("skip = %v, pc = %v\n", skip, pc[:n])
	}
	// Output:
	// skip = 0, pc = [4304486 4198562 4280114 4289760]
	// skip = 1, pc = [4198562 4280114 4289760]
	// skip = 2, pc = [4280114 4289760]
	// skip = 3, pc = [4289760]
}
func methodA(){
	methodB()
}
func methodAs(){
	methodBs()
}

func main()  {
	methodA()
	methodAs()
}
