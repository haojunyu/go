// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// These examples demonstrate more intricate uses of the flag package.
package flag_test

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

// Example 1: A single string flag called "species" with default value "gopher".
// 例1： 字符串命令参数species，默认值为gopher，提示为the species we are studying
var species = flag.String("species", "gopher", "the species we are studying")

// Example 2: Two flags sharing a variable, so we can have a shorthand.
// The order of initialization is undefined, so make sure both use the
// same default value. They must be set up with an init function.
// 例2：多个命令行参数表示相同的含义，特别是参数简写（-help简写成-h）
var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
}

// Example 3: A user-defined flag type, a slice of durations.
// 例3： 用户定义类型interval
type interval []time.Duration

// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
// String方法：interval类型内容字符串输出
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
// Set方法：将包含分隔符,的字符串复制到接受者interval类型中，解析失败则报异常error
func (i *interval) Set(value string) error {
	// If we wanted to allow the flag to be set multiple times,
	// accumulating values, we would delete this if statement.
	// That would permit usages such as
	//	-deltaT 10s -deltaT 15s
	// and other combinations.
	// 这段是不允许该命令行参数被重复赋值
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

// Define a flag to accumulate durations. Because it has a special type,
// we need to use the Var function and therefore create the flag during
// init.
// 定义interval类型的变量intervalFlag，因为它是自定义类型，使用flag.Var类创建命令行参数。
var intervalFlag interval

// init函数： 该源文件的初始化函数，主要是初始化命令行参数
func init() {
	// Tie the command-line flag to the intervalFlag variable and
	// set a usage message.
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func Example() {
	// All the interesting pieces are with the variables declared above, but
	// to enable the flag package to see the flags defined there, one must
	// execute, typically at the start of main (not init!):
	//	flag.Parse()
	// We don't run it here because this is not a main function and
	// the testing suite has already parsed the flags.
	// 命令行参数的解析需要在main函数执行flag.Parse()，而不是在init函数中
}
