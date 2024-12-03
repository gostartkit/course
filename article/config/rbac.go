// Copyright 2020-2024 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// NOTE: This file should not be edited
// see https://gostartkit.com/docs/code for more information.
package config

import (
	"sort"
	"strings"
)

const (
	Read int64 = 1 << iota
	Write
)

// CreateRbacConfig return *RightCollection
func CreateRbacConfig() *RightCollection {

	cfg := &RightCollection{
		{
			Key:   rightKey(Read),
			Value: Read,
		},
		{
			Key:   rightKey(Write),
			Value: Write,
		},
	}

	sort.Sort(cfg)

	return cfg
}

// Right struct
type Right struct {
	Key   string `json:"key"`
	Value int64  `json:"value"`
}

// RightCollection struct
type RightCollection []Right

// Len return len
func (o *RightCollection) Len() int { return len(*o) }

// Swap swap i, j
func (o *RightCollection) Swap(i, j int) { (*o)[i], (*o)[j] = (*o)[j], (*o)[i] }

// Less compare i, j
func (o *RightCollection) Less(i, j int) bool { return (*o)[i].Key < (*o)[j].Key }

// Search uses binary search to find and return the smallest index Value
func (o *RightCollection) Search(key string) int64 {

	i := sort.Search(o.Len(), func(i int) bool { return (*o)[i].Key >= key })

	if i < o.Len() && (*o)[i].Key == key {
		return (*o)[i].Value
	}

	return 0
}

// Sum sum right.value
func (o *RightCollection) Sum() int64 {
	var val int64 = 0
	for i := 0; i < o.Len(); i++ {
		val += (*o)[i].Value
	}
	return val
}

// Keys get keys by userRight
func (o *RightCollection) Keys(userRight int64) string {
	var sb strings.Builder
	for i := 0; i < o.Len(); i++ {
		if (*o)[i].Value&userRight > 0 {
			if sb.Len() > 0 {
				sb.WriteByte(',')
			}
			sb.WriteString((*o)[i].Key)
		}
	}
	return sb.String()
}

func rightKey(val int64) string {
	switch val {
	case Read:
		return "read"
	case Write:
		return "write"
	default:
		return ""
	}
}
