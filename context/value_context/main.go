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

created by lvxin for project golang-test at 19-2-12 下午3:22
*/
package main

import (
	"time"
	"fmt"
	"context"
)

var key string="name"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//附加值
	subCtx :=context.WithValue(ctx,key,"【监控1】")
	go watch(subCtx,"子协程1")
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context,name string) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(name,": ",ctx.Value(key),"监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(name,": ",ctx.Value(key),"goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
