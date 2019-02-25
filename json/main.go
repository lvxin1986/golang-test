package main

import (
	"encoding/json"
	"fmt"
)

/*
Copyright 2019 Tiglabs

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreedto in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

created by lvxin for project golang-test at 19-2-25 上午9:03
*/



type Partition struct {
	Id       int `json:"id,omitempty"`
	SpaceId  int     `json:"space_id,omitempty"`
	Replicas []int    `json:"replicas,omitempty"` //leader in replicas
}


type Space struct {
	Id           int           `json:"id,omitempty"`
	Name         string            `json:"name,omitempty"`
	Enabled      *bool             `json:"enabled"`
	Partitions   []*Partition      `json:"partitions"` // partitionids not sorted
	Fields       map[string]uint16 `json:"fields"`
}

func  main()  {
	spaceEnabled:=true
	space:= Space{
		Id:1,
		Name:"space1",
		Enabled:&spaceEnabled,
	}
	space.Partitions = append(space.Partitions, &Partition{
		Id:1,
		SpaceId: 1,
		Replicas:[]int {1,2,3},
	})

	space.Partitions = append(space.Partitions, &Partition{
		Id:2,
		SpaceId: 1,
		Replicas:[]int {4,5,6},
	})
	space.Fields = map[string]uint16{"key1":1,"key2":2}

	if jsonBytes,err:=json.Marshal(space);err==nil {
		fmt.Println(string(jsonBytes))
	} else {
		fmt.Println("error:",err)
	}

}