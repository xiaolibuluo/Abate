/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package cli

import (
	"io"

	"abate/ctl/web"
)

// cmdWeb serves the local control API together with the browser console. It is a
// thin adapter over the web package so the command stays inside the shared CLI
// dispatch and exit-code contract.
func cmdWeb(args []string, stdout, stderr io.Writer) int {
	return web.Run(args, stdout, stderr)
}
