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

created by lvxin for project golang-test at 19-2-13 上午9:58
*/
package version1

import (
	"fmt"
)
var (
	Version    string
	BuildTime string
	GoVersion string
)
type Info struct {
	Version     string
	BuildTime   string
	GoVersion   string
}

func (b *Info) String() string {
	return fmt.Sprintf(`
Version: %v
BuildTime: %v
GoVersion: %v
`,
		b.Version,
		b.BuildTime,
		b.GoVersion,
			)
}
func NewVersion() *Info {
	return &Info{
		Version:     Version,
		BuildTime:  BuildTime,
		GoVersion: GoVersion,
	}
}

var info *Info
func init() {
	info = NewVersion()
}

func GetVersion() string{
	return info.String()
}