// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// flag文件夹下有flag_test包，是因为该文件夹下包含核心代码flag.go和测试代码*_test.go
// flag_test包存在的意义是将测试代码从核心代码区分出来。
package flag_test

import (
	"flag"
	"fmt"
	"net/url"
)

type URLValue struct {
	URL *url.URL
}

// String方法：调用url.URL实现的String方法
func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

// Set方法： 调用url.Parse来解析url字符串
func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u // v.URL是*url.URL类型，也就是url.URL指针类型。u也是同类型。这里是想将值赋值（安全），而不是地址
	}
	return nil
}

var u = &url.URL{}

func ExampleValue() {
	fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	fs.Var(&URLValue{u}, "url", "URL to parse")

	fs.Parse([]string{"-url", "https://golang.org/pkg/flag/"})
	fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)

	// Output:
	// {scheme: "https", host: "golang.org", path: "/pkg/flag/"}
}
