/*
Copyright 2018 Tiger

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreedto in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

created by lvxin for project golang-test at 18-9-12 下午3:28
*/

package doer
//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks github.com/lvxin1986/golang-test/mock/doer Doer
type Doer interface {
	DoSomething(int, string) error
}