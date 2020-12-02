// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.15

package pflag

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func setUpC128SFlagSet(c128sp *[]complex128) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.Complex128SliceVar(c128sp, "c128s", []complex128{}, "Command separated list!")
	return f
}

func setUpC128SFlagSetWithDefault(c128sp *[]complex128) *FlagSet {
	f := NewFlagSet("test", ContinueOnError)
	f.Complex128SliceVar(c128sp, "c128s", []complex128{0.0, 1.0}, "Command separated list!")
	return f
}

func TestEmptyC128S(t *testing.T) {
	var c128s []complex128
	f := setUpC128SFlagSet(&c128s)
	err := f.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	getC128S, err := f.GetComplex128Slice("c128s")
	if err != nil {
		t.Fatal("got an error from GetComplex128Slice():", err)
	}
	if len(getC128S) != 0 {
		t.Fatalf("got c128s %v with len=%d but expected length=0", getC128S, len(getC128S))
	}
}

func TestC128S(t *testing.T) {
	var c128s []complex128
	f := setUpC128SFlagSet(&c128s)

	vals := []string{"1.0", "2.0", "4.0", "3.0"}
	arg := fmt.Sprintf("--c128s=%s", strings.Join(vals, ","))
	err := f.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for i, v := range c128s {
		d, err := strconv.ParseComplex(vals[i], 128)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if d != v {
			t.Fatalf("expected c128s[%d] to be %s but got: %f", i, vals[i], v)
		}
	}
	getC128S, err := f.GetComplex128Slice("c128s")
	if err != nil {
		t.Fatalf("got error: %v", err)
	}
	for i, v := range getC128S {
		d, err := strconv.ParseComplex(vals[i], 128)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if d != v {
			t.Fatalf("expected c128s[%d] to be %s but got: %f from GetComplex128Slice", i, vals[i], v)
		}
	}
}

func TestC128SDefault(t *testing.T) {
	var c128s []complex128
	f := setUpC128SFlagSetWithDefault(&c128s)

	vals := []string{"0.0", "1.0"}

	err := f.Parse([]string{})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for i, v := range c128s {
		d, err := strconv.ParseComplex(vals[i], 128)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if d != v {
			t.Fatalf("expected c128s[%d] to be %f but got: %f", i, d, v)
		}
	}

	getC128S, err := f.GetComplex128Slice("c128s")
	if err != nil {
		t.Fatal("got an error from GetComplex128Slice():", err)
	}
	for i, v := range getC128S {
		d, err := strconv.ParseComplex(vals[i], 128)
		if err != nil {
			t.Fatal("got an error from GetComplex128Slice():", err)
		}
		if d != v {
			t.Fatalf("expected c128s[%d] to be %f from GetComplex128Slice but got: %f", i, d, v)
		}
	}
}

func TestC128SWithDefault(t *testing.T) {
	var c128s []complex128
	f := setUpC128SFlagSetWithDefault(&c128s)

	vals := []string{"1.0", "2.0"}
	arg := fmt.Sprintf("--c128s=%s", strings.Join(vals, ","))
	err := f.Parse([]string{arg})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for i, v := range c128s {
		d, err := strconv.ParseComplex(vals[i], 128)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if d != v {
			t.Fatalf("expected c128s[%d] to be %f but got: %f", i, d, v)
		}
	}

	getC128S, err := f.GetComplex128Slice("c128s")
	if err != nil {
		t.Fatal("got an error from GetComplex128Slice():", err)
	}
	for i, v := range getC128S {
		d, err := strconv.ParseComplex(vals[i], 128)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if d != v {
			t.Fatalf("expected c128s[%d] to be %f from GetComplex128Slice but got: %f", i, d, v)
		}
	}
}

func TestC128SAsSliceValue(t *testing.T) {
	var c128s []complex128
	f := setUpC128SFlagSet(&c128s)

	in := []string{"1.0", "2.0"}
	argfmt := "--c128s=%s"
	arg1 := fmt.Sprintf(argfmt, in[0])
	arg2 := fmt.Sprintf(argfmt, in[1])
	err := f.Parse([]string{arg1, arg2})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}

	f.VisitAll(func(f *Flag) {
		if val, ok := f.Value.(SliceValue); ok {
			_ = val.Replace([]string{"3.1"})
		}
	})
	if len(c128s) != 1 || c128s[0] != 3.1 {
		t.Fatalf("Expected ss to be overwritten with '3.1', but got: %v", c128s)
	}
}

func TestC128SCalledTwice(t *testing.T) {
	var c128s []complex128
	f := setUpC128SFlagSet(&c128s)

	in := []string{"1.0,2.0", "3.0", "0+2i", "1,2i,2.5+3.1i"}
	expected := []complex128{1.0, 2.0, 3.0, complex(0, 2), complex(1, 0), complex(0, 2), complex(2.5, 3.1)}
	argfmt := "--c128s=%s"
	arg1 := fmt.Sprintf(argfmt, in[0])
	arg2 := fmt.Sprintf(argfmt, in[1])
	err := f.Parse([]string{arg1, arg2})
	if err != nil {
		t.Fatal("expected no error; got", err)
	}
	for i, v := range c128s {
		if expected[i] != v {
			t.Fatalf("expected c128s[%d] to be %f but got: %f", i, expected[i], v)
		}
	}
}
