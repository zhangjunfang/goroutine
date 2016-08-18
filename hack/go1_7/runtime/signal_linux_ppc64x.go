// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux
// +build ppc64 ppc64le

package runtime

import (
	"unsafe"
)

type sigctxt struct {
	info	*siginfo
	ctxt	unsafe.Pointer
}