// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package baraddur is designed to act as a
// small and lightweight file-watcher that
// works fine in docker projects at least.
package baraddur

import (
	"time"
)

// Config is to store the data of the separate files.
type Config struct {
	Files []File
}

// File is to save the path as filepath, the
// interval as time duration for the range ofter
// time ticker and execute to store the function.
type File struct {
	Path     string
	Interval time.Duration
	Fault    func(err error)
	Execute  func()
}
