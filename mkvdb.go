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

package mkvdb

import (
	"sync"
)

// MKVDB ...
type MKVDB struct {
	// mu Mutex protects data structures below.
	mu sync.Mutex

	// store is the Set underlying store of values.
	store map[interface{}]interface{}
}

// New Database
func New() *MKVDB {
	db := &MKVDB{
		store: make(map[interface{}]interface{}),
	}

	return db
}

// Put add an item into the db.
func (db *MKVDB) Put(key interface{}, value interface{}) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.store[key] = value
}

// Get returns the value of a key from the db.
func (db *MKVDB) Get(key interface{}) (value interface{}, found bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

	value, found = db.store[key]
	return value, found
}

// Remove deletes a key-value pair from the db.
//
// If a key is not found in the map it doesn't fails, just does nothing.
func (db *MKVDB) Remove(key interface{}) {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Check key exists
	if _, found := db.store[key]; !found {
		return
	}

	// Remove the value from the store
	delete(db.store, key)
}

// Empty return if the map in empty or not.
func (db *MKVDB) Empty() bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	return len(db.store) == 0
}
