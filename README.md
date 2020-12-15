# pflag

***This is a fork of [spf13/pflag](https://github.com/spf13/pflag) due to poor maintenance***

[![GoDoc](https://godoc.org/github.com/cornfeedhobo/pflag?status.svg)](https://godoc.org/github.com/cornfeedhobo/pflag)
[![Go Report Card](https://goreportcard.com/badge/github.com/cornfeedhobo/pflag)](https://goreportcard.com/report/github.com/cornfeedhobo/pflag)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/cornfeedhobo/pflag?sort=semver)](https://github.com/cornfeedhobo/pflag/releases)
[![Build Status](https://travis-ci.com/cornfeedhobo/pflag.svg?branch=master)](https://travis-ci.com/cornfeedhobo/pflag)

* [Installation](#installation)
  * [Installing this fork with spf13/cobra](#installing-this-fork-with-cobra)
* [Documentation](#documentation)
  * [Shorthand flags](#shorthand-flags)
  * [Shorthand-only flags](#shorthand-only-flags)
  * [Set a custom default for flags passed without values](#set-a-custom-default-for-flags-passed-without-values)
  * [Custom flag types in usage](#custom-flag-types-in-usage)
  * [Disable printing default value](#disable-printing-default-value)
  * [Disable sorting of flags](#disable-sorting-of-flags)
  * [Deprecating a flag or its shorthand](#deprecating-a-flag-or-its-shorthand)
  * [Hidden flags](#hidden-flags)
  * [Mutating or "Normalizing" Flag names](#mutating-or-normalizing-flag-names)
  * [Unknown flags](#unknown-flags)
  * [Supporting Go flags when using pflag](#supporting-go-flags-when-using-pflag)
  * [Command line flag syntax](#command-line-flag-syntax)

## Installation

pflag is available using the standard `go get` command.

Install by running:

``` bash
go get github.com/cornfeedhobo/pflag
```

### Installing this fork with spf13/cobra

Initialize your new app as normal

``` bash
cobra init --pkg-name example.com/hello
go mod init example.com/hello
```

Override the upstream module

``` bash
go mod edit -replace github.com/spf13/pflag=github.com/cornfeedhobo/pflag
```

## Documentation

You can see the full reference documentation of the pflag package
[at godoc.org](http://godoc.org/github.com/cornfeedhobo/pflag), querying with
[`go doc`](https://golang.org/cmd/doc/), or through go's standard documentation
system by running `godoc -http=:6060` and browsing to
[http://localhost:6060/pkg/github.com/cornfeedhobo/pflag](http://localhost:6060/pkg/github.com/cornfeedhobo/pflag)
after installation.

### Shorthand flags

A flag supporting both long and short formats can be created with any of the
flag functions suffixed with `P`:

``` go
flag.BoolP("toggle", "t", false, "toggle help message")
```

### Shorthand-only flags

A shorthand-only flag can be created with any of the flag functions suffixed
with `S`:

``` go
flag.StringS("value", "l", "", "value help message")
```

This flag can be looked up using it's long name, but will only be parsed when
the short form is passed.

### Set a custom default for flags passed without values

If a flag has a NoOptDefVal and the flag is set on the command line
without an option the flag will be set to the NoOptDefVal.

**Example**:

``` go
var ip = flag.IntP("flagname", "f", 1234, "help message")
flag.Lookup("flagname").NoOptDefVal = "4321"
```

**Results**:

| Parsed Arguments | Resulting Value |
| -------------    | -------------   |
| --flagname=1357  | ip=1357         |
| --flagname       | ip=4321         |
| [nothing]        | ip=1234         |

### Custom flag types in usage

There are two methods to set a custom type to be printed in the usage.

First, it's possible to set explicitly with `UsageType`:

``` go
flag.String("character", "", "character name")
flag.Lookup("character").UsageType = "enum"
```

Output:

``` plain
  --character enum   character name (default "")
```

Alternatively, it's possbile to include backticks around a single word in the
usage string, which will be extracted and printed with the usage:

``` go
flag.String("character", "", "`character` name")
```

Output:

``` plain
  --character character   character name (default "")
```

_Note: This unquoting behavior can be disabled with `Flag.DisableUnquoteUsage`_.

### Disable printing default value

The printing of a flag's default value can be suppressed with `Flag.DisablePrintDefault`.

**Example**:

``` go
flag.Int("in", -1, "help message")
flag.Lookup("in").DisablePrintDefault = true
```

**Output**:

``` plain
  --in int   help message
```

### Disable sorting of flags

It is possible to disable sorting of flags for help and usage message.

**Example**:

``` go
flag.BoolP("verbose", "v", false, "verbose output")
flag.String("coolflag", "yeaah", "it's really cool flag")
flag.Int("usefulflag", 777, "sometimes it's very useful")
flag.SortFlags = false
flag.PrintDefaults()
```

**Output**:

``` plain
  -v, --verbose           verbose output
      --coolflag string   it's really cool flag (default "yeaah")
      --usefulflag int    sometimes it's very useful (default 777)
```

### Deprecating a flag or its shorthand

It is possible to deprecate a flag, or just its shorthand. Deprecating a
flag/shorthand hides it from help text and prints a usage message when the
deprecated flag/shorthand is used.

**Example #1**: You want to deprecate a flag named "badflag" as well as
inform the users what flag they should use instead.

``` go
// deprecate a flag by specifying its name and a usage message
flags.MarkDeprecated("badflag", "please use --good-flag instead")
```

This hides "badflag" from help text, and prints
`Flag --badflag has been deprecated, please use --good-flag instead`
when "badflag" is used.

**Example #2**: You want to keep a flag name "noshorthandflag" but deprecate
it's shortname "n".

``` go
// deprecate a flag shorthand by specifying its flag name and a usage message
flags.MarkShorthandDeprecated("noshorthandflag", "please use --noshorthandflag only")
```

This hides the shortname "n" from help text, and prints
`Flag shorthand -n has been deprecated, please use --noshorthandflag only`
when the shorthand "n" is used.

Note that usage message is essential here, and it should not be empty.

### Hidden flags

It is possible to mark a flag as hidden, meaning it will still function as
normal, however will not show up in usage/help text.

**Example**: You have a flag named "secretFlag" that you need for internal use
only and don't want it showing up in help text, or for its usage text to be available.

``` go
// hide a flag by specifying its name
flags.MarkHidden("secretFlag")
```

### Mutating or "Normalizing" Flag names

It is possible to set a custom flag name 'normalization function.' It allows
flag names to be mutated both when created in the code and when used on the
command line to some 'normalized' form. The 'normalized' form is used for
comparison. Two examples of using the custom normalization func follow.

**Example #1**: You want -, _, and . in flags to compare the same. aka --my-flag == --my_flag == --my.flag

``` go
func wordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return pflag.NormalizedName(name)
}

myFlagSet.SetNormalizeFunc(wordSepNormalizeFunc)
```

**Example #2**: You want to alias two flags. aka --old-flag-name == --new-flag-name

``` go
func aliasNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "old-flag-name":
		name = "new-flag-name"
		break
	}
	return pflag.NormalizedName(name)
}

myFlagSet.SetNormalizeFunc(aliasNormalizeFunc)
```

### Unknown flags

Normally pflag will error when an unknown flag is passed, but it's also possible
to disable that using `FlagSet.ParseErrorsAllowList.UnknownFlags`:

``` go
flags.ParseErrorsAllowlist.UnknownFlags = true
flag.Parse()
```

Additionally, these can then be examined using `FlagSet.VisitUnknowns` or
`FlagSet.GetUnknownFlags`.

### Supporting Go flags when using pflag

In order to support flags defined using Go's `flag` package, they must be added
to the `pflag` flagset. This is usually necessary to support flags defined by
third-party dependencies (e.g. `golang/glog`).

**Example**: You want to add the Go flags to the `CommandLine` flagset

``` go
import (
	goflag "flag"
	flag "github.com/cornfeedhobo/pflag"
)

var ip *int = flag.Int("flagname", 1234, "help message for flagname")

func main() {
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()
}
```

### Command line flag syntax

``` plain
--flag    // boolean flags, or flags with no option default values
--flag x  // only on flags without a default value
--flag=x
```

Unlike the flag package, a single dash before an option means something
different than a double dash. Single dashes signify a series of shorthand
letters for flags. All but the last shorthand letter must be boolean flags
or a flag with a default value

``` plain
// boolean or flags where the 'no option default value' is set
-f
-f=true
-abc
but
-b true is INVALID

// non-boolean and flags without a 'no option default value'
-n 1234
-n=1234
-n1234

// mixed
-abcs "hello"
-absd="hello"
-abcs1234
```

Slice flags can be specified multiple times, or specified with an equal sign and csv.

``` plain
--sliceVal one --sliceVal=two
--sliceVal=one,two
```

Integer flags accept 1234, 0664, 0x1234 and may be negative.
Boolean flags (in their long form) accept 1, 0, t, f, true, false,
TRUE, FALSE, True, False.
Duration flags accept any input valid for time.ParseDuration.

Flag parsing stops after the terminator "--". Unlike the flag package,
flags can be interspersed with arguments anywhere on the command line
before this terminator.
