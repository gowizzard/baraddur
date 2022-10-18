// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package baraddur

import (
	"os"
	"reflect"
	"sync"
	"time"
)

// Watch is to check the mod time of the different
// files via wait groups. When a file receives a new
// modification time, the function from the configuration struct
// is executed. Also, can define the interval for each file.
func (c *Config) Watch() {

	var routines sync.WaitGroup
	routines.Add(reflect.ValueOf(c.Files).Len())

	defer routines.Wait()

	for _, value := range c.Files {

		go func(path string, interval time.Duration, fault func(err error), execute func()) {

			defer routines.Done()

			var modification time.Time
			for range time.Tick(interval) {

				stat, err := os.Stat(path)
				if err != nil {
					fault(err)
				}

				if !modification.IsZero() && modification.Before(stat.ModTime()) {
					execute()
				}
				modification = stat.ModTime()

			}

		}(value.Path, value.Interval, value.Fault, value.Execute)

	}

}
