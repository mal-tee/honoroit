// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE-GO file.

// Modifications Copyright 2020 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libc // import "modernc.org/sqlite/internal/libc"

import (
	"fmt"
	"os"
	"sync"
	"time"
	"unsafe"
	// "golang.org/x/sys/unix"
	// "modernc.org/sqlite/internal/libc/errno"
	// "modernc.org/sqlite/internal/libc/fcntl"
)

var randState uint32
var randStateMu sync.Mutex

func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}

func nextRandom(x uintptr) {
	randStateMu.Lock()
	r := randState
	if r == 0 {
		r = reseed()
	}
	r = r*1664525 + 1013904223
	randState = r
	randStateMu.Unlock()
	copy((*RawMem)(unsafe.Pointer(x))[:6:6], fmt.Sprintf("%06d", int(1e9+r%1e9)%1e6))
}

func tempFile(s, x uintptr, flags int32) (fd int, err error) {
	panic(todo(""))

}
