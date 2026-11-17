/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

// Factory instantiates a Plugin from its verified manifest. The host rejects any
// plugin whose signature fails verification before the factory is ever called, so
// a factory only ever receives a verified manifest.
type Factory func(m Manifest) (Plugin, error)
