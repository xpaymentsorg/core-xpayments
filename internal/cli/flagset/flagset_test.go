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

package flagset

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFlagsetBool(t *testing.T) {
	f := NewFlagSet("")

	value := false
	f.BoolFlag(&BoolFlag{
		Name:  "flag",
		Value: &value,
	})

	assert.NoError(t, f.Parse([]string{"--flag", "true"}))
	assert.Equal(t, value, true)
}

func TestFlagsetSliceString(t *testing.T) {
	f := NewFlagSet("")

	value := []string{}
	f.SliceStringFlag(&SliceStringFlag{
		Name:  "flag",
		Value: &value,
	})

	assert.NoError(t, f.Parse([]string{"--flag", "a,b", "--flag", "c"}))
	assert.Equal(t, value, []string{"a", "b", "c"})
}

func TestFlagsetDuration(t *testing.T) {
	f := NewFlagSet("")

	value := time.Duration(0)
	f.DurationFlag(&DurationFlag{
		Name:  "flag",
		Value: &value,
	})

	assert.NoError(t, f.Parse([]string{"--flag", "1m"}))
	assert.Equal(t, value, 1*time.Minute)
}

func TestFlagsetMapString(t *testing.T) {
	f := NewFlagSet("")

	value := map[string]string{}
	f.MapStringFlag(&MapStringFlag{
		Name:  "flag",
		Value: &value,
	})

	assert.NoError(t, f.Parse([]string{"--flag", "a=b,c=d"}))
	assert.Equal(t, value, map[string]string{"a": "b", "c": "d"})
}
