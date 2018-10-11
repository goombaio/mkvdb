// Copyright 2018, Goomba project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package mkvdb_test

import (
	"testing"

	"github.com/goombaio/mkvdb"
)

func TestMKVDB(t *testing.T) {
	m := mkvdb.New()

	if !m.Empty() {
		t.Fatalf("New MKVDB expected to be empty but it is not")
	}
}

func TestMKVDB_Put(t *testing.T) {
	m := mkvdb.New()

	m.Put("foo", "bar")

	if m.Empty() {
		t.Fatalf("New MKVDB expected not to be empty but it is")
	}
}

func TestMKVDB_Get(t *testing.T) {
	m := mkvdb.New()

	m.Put("foo", "bar")

	if m.Empty() {
		t.Fatalf("New MKVDB expected not to be empty but it is")
	}

	value, found := m.Get("foo")
	if !found {
		t.Fatalf("Key expected to be found but it isn't")
	}
	if value != "bar" {
		t.Fatalf("Value expected to be %q, but got %q", "bar", value)
	}
}

func TestMKVDB_Remove(t *testing.T) {
	m := mkvdb.New()

	m.Put("foo", "bar")

	if m.Empty() {
		t.Fatalf("New MKVDB expected not to be empty but it is")
	}

	m.Remove("bar")

	if m.Empty() {
		t.Fatalf("New MKVDB expected not to be empty but it is")
	}

	m.Remove("foo")

	if !m.Empty() {
		t.Fatalf("New MKVDB expected to be empty but it is not")
	}
}
