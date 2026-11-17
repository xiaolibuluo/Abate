/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package web implements the "abate web" command: it serves the local control
// API together with a self-contained browser console from a single process.
//
// The browser console is a single static page with inline CSS and vanilla
// JavaScript. It carries no runtime CDN, font or script dependency and opens
// fully offline, so the whole surface is served same-origin: the console never
// needs CORS configuration and the API never needs to be exposed separately.
//
// Two serving modes are supported:
//
//   - Internal (default): the embedded REST server is started in-process and the
//     console talks to it over the same listener.
//   - Proxy (--rest): an external REST base URL is reverse-proxied under /api,
//     while the console is still served locally.
package web
