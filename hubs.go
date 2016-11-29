// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"
)

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hubs struct {
}

func newHubs() *Hubs {
	return &Hubs{
	}
}

func (h * Hubs) run(){
	forever()
	select{
	}
}

func runHub(h* Hub){
	go h.run()
}

func forever() {
    for {
        time.Sleep(time.Second)
    }
}