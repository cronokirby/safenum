// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !math_big_pure_go

package safenum

// This should be feature detected, but we can't use the internal/cpu package
var support_adx = true