// Copyright 2022 The go-xpayments Authors
// This file is part of the go-xpayments library.
//
// Copyright 2022 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package pprof

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

// Profile generates a pprof.Profile report for the given profile name.
func Profile(profile string, debug, gc int) ([]byte, map[string]string, error) {
	p := pprof.Lookup(profile)
	if p == nil {
		return nil, nil, fmt.Errorf("profile '%s' not found", profile)
	}

	if profile == "heap" && gc > 0 {
		runtime.GC()
	}

	var buf bytes.Buffer
	if err := p.WriteTo(&buf, debug); err != nil {
		return nil, nil, err
	}

	headers := map[string]string{
		"X-Content-Type-Options": "nosniff",
	}
	if debug != 0 {
		headers["Content-Type"] = "text/plain; charset=utf-8"
	} else {
		headers["Content-Type"] = "application/octet-stream"
		headers["Content-Disposition"] = fmt.Sprintf(`attachment; filename="%s"`, profile)
	}
	return buf.Bytes(), headers, nil
}

// CPUProfile generates a CPU Profile for a given duration
func CPUProfile(ctx context.Context, sec int) ([]byte, map[string]string, error) {
	if sec <= 0 {
		sec = 1
	}

	var buf bytes.Buffer
	if err := pprof.StartCPUProfile(&buf); err != nil {
		return nil, nil, err
	}

	sleep(ctx, time.Duration(sec)*time.Second)

	pprof.StopCPUProfile()

	return buf.Bytes(),
		map[string]string{
			"X-Content-Type-Options": "nosniff",
			"Content-Type":           "application/octet-stream",
			"Content-Disposition":    `attachment; filename="profile"`,
		}, nil
}

// Trace runs a trace profile for a given duration
func Trace(ctx context.Context, sec int) ([]byte, map[string]string, error) {
	if sec <= 0 {
		sec = 1
	}

	var buf bytes.Buffer
	if err := trace.Start(&buf); err != nil {
		return nil, nil, err
	}

	sleep(ctx, time.Duration(sec)*time.Second)

	trace.Stop()

	return buf.Bytes(),
		map[string]string{
			"X-Content-Type-Options": "nosniff",
			"Content-Type":           "application/octet-stream",
			"Content-Disposition":    `attachment; filename="trace"`,
		}, nil
}

func sleep(ctx context.Context, d time.Duration) {
	// Sleep until duration is met or ctx is cancelled
	select {
	case <-time.After(d):
	case <-ctx.Done():
	}
}
