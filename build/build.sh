#!/usr/bin/env bash

#Copyright 2019 Tiger
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreedto in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.
#
#created by lvxin for project golang-test at 19-2-13 上午10:10

BasePath=$(cd `dirname $0`; pwd)
pkg_name="${BasePath#"$GOPATH/src/"}/version"
#go  build -ldflags "-X version.Version=1.0.0 -X 'version.BuildTime=`date`' -X 'version.GoVersion=`go version`'"
go build -ldflags "-X ${pkg_name}.Version=1.0.0 -X '${pkg_name}.BuildTime=`date`' -X '${pkg_name}.GoVersion=`go version`'"
./build