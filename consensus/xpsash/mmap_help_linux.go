// Copyright 2022 The go-xpayments Authors
// This file is part of the go-xpayments library.
//
// The go-xpayments library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-xpayments library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-xpayments library. If not, see <http://www.gnu.org/licenses/>.

//go:build linux
// +build linux

package xpsash

import (
	"os"

	"golang.org/x/sys/unix"
)

// ensureSize expands the file to the given size. This is to prevent runtime
// errors later on, if the underlying file expands beyond the disk capacity,
// even though it ostensibly is already expanded, but due to being sparse
// does not actually occupy the full declared size on disk.
func ensureSize(f *os.File, size int64) error {
	// Docs: https://www.man7.org/linux/man-pages/man2/fallocate.2.html
	return unix.Fallocate(int(f.Fd()), 0, 0, size)
}
