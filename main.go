/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package main is the entry point of the Abate service.
package main

import (
	"os"

	"abate/ctl/cli"
)

func main() {
	os.Exit(cli.Execute(os.Args[1:]))
}
