// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package baraddur_test

import (
	"github.com/gowizzard/baraddur"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"testing"
	"time"
)

// TestWatch is to test the Watch function from the baraddur package.
func TestWatch(t *testing.T) {

	tests := []struct {
		files []struct {
			path   string
			data   []byte
			update []byte
			perm   os.FileMode
			wait   time.Duration
		}
		config baraddur.Config
	}{
		{
			files: []struct {
				path   string
				data   []byte
				update []byte
				perm   os.FileMode
				wait   time.Duration
			}{
				{
					path:   filepath.Join(os.TempDir(), ".env"),
					data:   []byte("PATH=~/usr/local/bin"),
					update: []byte("PATH=~/usr/bin"),
					perm:   0666,
					wait:   3 * time.Second,
				},
			},
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
		},
		{
			files: []struct {
				path   string
				data   []byte
				update []byte
				perm   os.FileMode
				wait   time.Duration
			}{
				{
					path:   filepath.Join(os.TempDir(), "hello.txt"),
					data:   []byte("This is a tests file."),
					update: []byte("This is a updated test file."),
					perm:   0666,
					wait:   15 * time.Second,
				},
				{
					path:   filepath.Join(os.TempDir(), "bye.txt"),
					data:   []byte("This is a tests file."),
					update: []byte("This is a updated test file."),
					perm:   0666,
					wait:   6 * time.Second,
				},
			},
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
					{
						Path:     filepath.Join(os.TempDir(), "bye.txt"),
						Interval: 2 * time.Second,
						Fault: func(err error) {
							t.Log(err)
						},
						Execute: func() {
							t.Log("The file \"bye.txt\" was updated.")
						},
					},
				},
			},
		},
		{
			files: []struct {
				path   string
				data   []byte
				update []byte
				perm   os.FileMode
				wait   time.Duration
			}{
				{
					path:   filepath.Join(os.TempDir(), "error.log"),
					data:   []byte("ERROR[2022-10-10T11:15:00] This is a error message.\tFILE=log_test.go:65"),
					update: []byte("ERROR[2022-10-10T11:30:00]  This is a updated error message.\tFILE=log_test.go:65"),
					perm:   0666,
					wait:   30 * time.Second,
				},
				{
					path:   filepath.Join(os.TempDir(), "warning.log"),
					data:   []byte("WARNING[2022-10-10T10:30:00] This is a warning message.\tFILE=log_test.go:65"),
					update: []byte("WARNING[2022-10-10T10:45:00] This is a updated warning message.\tFILE=log_test.go:65"),
					perm:   0666,
					wait:   15 * time.Second,
				},
				{
					path:   filepath.Join(os.TempDir(), "information.log"),
					data:   []byte("INFO[2022-10-10T10:45:00] This is a information message.\tFILE=log_test.go:65"),
					update: []byte("INFO[2022-10-10T11:00:00] This is a updated information message.\tFILE=log_test.go:65"),
					perm:   0666,
					wait:   9 * time.Second,
				},
			},
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
					{
						Path:     filepath.Join(os.TempDir(), "warning.log"),
						Interval: 5 * time.Second,
						Fault: func(err error) {
							t.Log(err)
						},
						Execute: func() {
							t.Log("The file \"warning.log\" was updated.")
						},
					},
					{
						Path:     filepath.Join(os.TempDir(), "information.log"),
						Interval: 3 * time.Second,
						Fault: func(err error) {
							t.Log(err)
						},
						Execute: func() {
							t.Log("The file \"information.log\" was updated.")
						},
					},
				},
			},
		},
	}

	for _, value := range tests {

		for _, value := range value.files {
			err := os.WriteFile(value.path, value.data, value.perm)
			if err != nil {
				log.Fatal(err)
			}
		}

		var routines sync.WaitGroup
		routines.Add(reflect.ValueOf(value.files).Len() + 1)

		go func(c baraddur.Config) {
			defer routines.Done()
			c.Watch()
		}(value.config)

		for _, value := range value.files {
			go func(wait time.Duration, path string, update []byte, perm os.FileMode) {
				defer routines.Done()
				time.Sleep(wait)
				err := os.WriteFile(path, update, perm)
				if err != nil {
					log.Fatal(err)
				}
			}(value.wait, value.path, value.update, value.perm)
		}

		routines.Wait()

	}

}
