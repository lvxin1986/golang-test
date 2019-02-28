package main

import "fmt"

// Copyright 2019 The ChuBao Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.
//create by bjlvxin at 13:58 for project golang-test

func main() {
	names := []string{"alex", "bob", "calina"}
	fmt.Println("Ranging array...")
	fmt.Println("[1] n:= range names...")
	for n := range names {
		fmt.Println(n)
	}
	fmt.Println("[2] n,i:= range names...")
	for n, i := range names {
		fmt.Println(n, "   ", i)
	}
	mp := map[int]string{
		1: "aa",
		2: "bb",
	}

	fmt.Println("--------------------------------")
	fmt.Println("Ranging map...")
	fmt.Println("[1] n:= range mp...")
	for n := range mp {
		fmt.Println(n)
	}
	fmt.Println("[2] n,i:= range names...")
	for n, i := range mp {
		fmt.Println(n, "   ", i)
	}
}
