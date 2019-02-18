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

created by lvxin for project golang-test at 19-2-18 下午3:09
*/
package main

import "fmt"

func R_WithNoSet() (b bool)  {
	return
}


func R_WithSet() (b bool)  {
	b=true
	return
}

func main()  {
	fmt.Println("R_WithName() = ",R_WithNoSet())
	fmt.Println("R_WithName() = ",R_WithSet())

}