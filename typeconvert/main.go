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

created by lvxin for project golang-test at 19-2-20 下午6:49
*/
package main

import "fmt"

func convertToInt(i interface{}) int{
	 r,_:=i.(int)
	 return r
}
func main() {
	p:=2.3
	fmt.Println("aaaa",convertToInt(p))
}