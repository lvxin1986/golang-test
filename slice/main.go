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

created by lvxin for project golang-test at 19-2-14 下午2:16
*/
package main

import "fmt"

var Schemas []string

func main() {
	Schemas = append(Schemas,"a")
	Schemas = append(Schemas,"a")
	Schemas = append(Schemas,"a")
	Schemas = append(Schemas,"a")
	fmt.Println(Schemas)
}