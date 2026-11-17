/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package web

import "embed"

// assets embeds the self-contained web console build output.
//
// The console is a single static page with inline CSS and vanilla JavaScript: it
// carries no runtime CDN, font or script dependency and opens fully offline. The
// dist directory is checked into the repository so the binary stays reproducible
// and the service works without any network access (PLAN-07 CTL-WEB section 6).
//
//go:embed dist
var assets embed.FS
