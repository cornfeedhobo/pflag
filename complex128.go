// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.15

package pflag

import "strconv"

// -- complex128 Value
type complex128Value complex128

func newComplex128Value(val complex128, p *complex128) *complex128Value {
	*p = val
	return (*complex128Value)(p)
}

func (f *complex128Value) Set(s string) error {
	v, err := strconv.ParseComplex(s, 128)
	*f = complex128Value(v)
	return err
}

func (f *complex128Value) Type() string {
	return "complex128"
}

func (f *complex128Value) String() string { return strconv.FormatComplex(complex128(*f), 'g', -1, 128) }

func complex128Conv(sval string) (interface{}, error) {
	return strconv.ParseComplex(sval, 128)
}

// GetComplex128 return the complex128 value of a flag with the given name
func (f *FlagSet) GetComplex128(name string) (complex128, error) {
	val, err := f.getFlagType(name, "complex128", complex128Conv)
	if err != nil {
		return 0, err
	}
	return val.(complex128), nil
}

// MustGetComplex128 is like GetComplex128, but panics on error.
func (f *FlagSet) MustGetComplex128(name string) complex128 {
	val, err := f.GetComplex128(name)
	if err != nil {
		panic(err)
	}
	return val
}

// Complex128Var defines a complex128 flag with specified name, default value, and usage string.
// The argument p points to a complex128 variable in which to store the value of the flag.
func (f *FlagSet) Complex128Var(p *complex128, name string, value complex128, usage string) {
	f.Complex128VarP(p, name, "", value, usage)
}

// Complex128VarP is like Complex128Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Complex128VarP(p *complex128, name, shorthand string, value complex128, usage string) {
	f.VarP(newComplex128Value(value, p), name, shorthand, usage)
}

// Complex128VarS is like Complex128Var, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) Complex128VarS(p *complex128, name, shorthand string, value complex128, usage string) {
	f.VarS(newComplex128Value(value, p), name, shorthand, usage)
}

// Complex128Var defines a complex128 flag with specified name, default value, and usage string.
// The argument p points to a complex128 variable in which to store the value of the flag.
func Complex128Var(p *complex128, name string, value complex128, usage string) {
	CommandLine.Complex128Var(p, name, value, usage)
}

// Complex128VarP is like Complex128Var, but accepts a shorthand letter that can be used after a single dash.
func Complex128VarP(p *complex128, name, shorthand string, value complex128, usage string) {
	CommandLine.Complex128VarP(p, name, shorthand, value, usage)
}

// Complex128VarS is like Complex128Var, but accepts a shorthand letter that can be used after a single dash, alone.
func Complex128VarS(p *complex128, name, shorthand string, value complex128, usage string) {
	CommandLine.Complex128VarS(p, name, shorthand, value, usage)
}

// Complex128 defines a complex128 flag with specified name, default value, and usage string.
// The return value is the address of a complex128 variable that stores the value of the flag.
func (f *FlagSet) Complex128(name string, value complex128, usage string) *complex128 {
	return f.Complex128P(name, "", value, usage)
}

// Complex128P is like Complex128, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Complex128P(name, shorthand string, value complex128, usage string) *complex128 {
	p := new(complex128)
	f.Complex128VarP(p, name, shorthand, value, usage)
	return p
}

// Complex128S is like Complex128, but accepts a shorthand letter that can be used after a single dash, alone.
func (f *FlagSet) Complex128S(name, shorthand string, value complex128, usage string) *complex128 {
	p := new(complex128)
	f.Complex128VarS(p, name, shorthand, value, usage)
	return p
}

// Complex128 defines a complex128 flag with specified name, default value, and usage string.
// The return value is the address of a complex128 variable that stores the value of the flag.
func Complex128(name string, value complex128, usage string) *complex128 {
	return CommandLine.Complex128(name, value, usage)
}

// Complex128P is like Complex128, but accepts a shorthand letter that can be used after a single dash.
func Complex128P(name, shorthand string, value complex128, usage string) *complex128 {
	return CommandLine.Complex128P(name, shorthand, value, usage)
}

// Complex128S is like Complex128, but accepts a shorthand letter that can be used after a single dash, alone.
func Complex128S(name, shorthand string, value complex128, usage string) *complex128 {
	return CommandLine.Complex128S(name, shorthand, value, usage)
}
