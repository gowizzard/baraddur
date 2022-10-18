// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package baraddur_test

import (
	"github.com/gowizzard/baraddur"
	"log"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

// TestWatch is to test the Watch function from the baraddur package.
func TestWatch(t *testing.T) {

	tests := []struct {
		path   string
		data   []byte
		update []byte
		perm   os.FileMode
		config baraddur.Config
		wait   time.Duration
	}{
		{
			path:   filepath.Join(os.TempDir(), ".env"),
			data:   []byte("PATH=~/usr/local/bin"),
			update: []byte("PATH=~/usr/bin"),
			perm:   0666,
			config: baraddur.Config{
				Files: []baraddur.File{
					{
						Path:     filepath.Join(os.TempDir(), ".env"),
						Interval: 1 * time.Second,
						Fault: func(err error) {
							t.Log(err)
						},
						Execute: func() {
							t.Log("The file \".env\" was updated.")
						},
					},
				},
			},
			wait: 3 * time.Second,
		},
		{
			path:   filepath.Join(os.TempDir(), "hello.txt"),
			data:   []byte("This is a tests file."),
			update: []byte("This is a updated test file."),
			perm:   0666,
			config: baraddur.Config{
				Files: []baraddur.File{
					{
						Path:     filepath.Join(os.TempDir(), "hello.txt"),
						Interval: 5 * time.Second,
						Fault: func(err error) {
							t.Log(err)
						},
						Execute: func() {
							t.Log("The file \"hello.txt\" was updated.")
						},
					},
				},
			},
			wait: 15 * time.Second,
		},
		{
			path:   filepath.Join(os.TempDir(), "error.log"),
			data:   []byte("ERROR[2022-08-26T11:51:30] This is a error message.\tFILE=log_test.go:65"),
			update: []byte("ERROR[2022-08-26T11:52:00]  This is a updated error message.\tFILE=log_test.go:65"),
			perm:   0666,
			config: baraddur.Config{
				Files: []baraddur.File{
					{
						Path:     filepath.Join(os.TempDir(), "error.log"),
						Interval: 10 * time.Second,
						Fault: func(err error) {
							t.Log(err)
						},
						Execute: func() {
							t.Log("The file \"error.log\" was updated.")
						},
					},
				},
			},
			wait: 30 * time.Second,
		},
	}

	for _, value := range tests {

		err := os.WriteFile(value.path, value.data, value.perm)
		if err != nil {
			log.Fatal(err)
		}

		var routines sync.WaitGroup
		routines.Add(2)

		go func(c baraddur.Config) {
			defer routines.Done()
			c.Watch()
		}(value.config)

		go func(wait time.Duration, path string, update []byte, perm os.FileMode) {
			defer routines.Done()
			time.Sleep(wait)
			err := os.WriteFile(path, update, perm)
			if err != nil {
				log.Fatal(err)
			}
		}(value.wait, value.path, value.update, value.perm)

		routines.Wait()

	}

}
