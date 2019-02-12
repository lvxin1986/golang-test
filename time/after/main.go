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

created by lvxin for project golang-test at 19-2-12 下午4:57
This test code is used to test time.after
*/
package main

import (
	"time"
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case info := <-time.After(time.Second * 1):
		fmt.Println("timeout, and the timeout info is: ",info)
	}
}
