// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flag

import "os"

// Additional routines compiled into the package only during testing.
// DefaultUsage函数变量：使用flag包中定义的默认函数变量Usage
var DefaultUsage = Usage

// ResetForTesting clears all flag state and sets the usage function as directed.
// After calling ResetForTesting, parse errors in flag handling will not
// exit the program.
// ResetForTesting函数是专门为测试写的函数，它是用来清除所有命令参数状态
// 以及直接设置usage函数（置空表示解析失败后不会打印usage信息）,在调用ResetForTesting函数后
// 解析失败并不会使程序终止
func ResetForTesting(usage func()) {
	CommandLine = NewFlagSet(os.Args[0], ContinueOnError)
	CommandLine.Usage = commandLineUsage // 赋值结构体FlagSet中的函数Usage
	Usage = usage                        // 赋值flag包中的函数变量Usage
}
