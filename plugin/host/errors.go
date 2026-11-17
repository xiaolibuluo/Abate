/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package host

import (
	"errors"
)

// ErrEntryForbidden reports that an entry name is not present in the instance
// entry whitelist, or that the plugin type has no callable entry in this build.
var ErrEntryForbidden = errors.New("plugin: entry forbidden")

// ErrResultTooLarge reports that a call result exceeded the 1 MiB cap.
var ErrResultTooLarge = errors.New("plugin: result too large")

// ErrPluginNotFound reports that no instance is registered under the requested id.
var ErrPluginNotFound = errors.New("plugin: plugin not found")

// ErrInstanceStale reports that the instance was superseded by a hot-swap and
// must no longer be invoked. Callers holding a stale pointer observe this error
// instead of executing against a draining or unloaded instance.
var ErrInstanceStale = errors.New("plugin: instance superseded")

// errNilPlugin is returned when a factory yields a nil plugin.
var errNilPlugin = errors.New("plugin: factory returned nil plugin")

// errFactoryExists is returned when Register is called twice for one type.
var errFactoryExists = errors.New("plugin: factory already registered for type")

// MaxResultBytes is the hard ceiling on a single call result (1 MiB). Results
// larger than this are rejected with ErrResultTooLarge to bound memory use.
const MaxResultBytes = 1 << 20

// EntryApply is the operator plugin call entry.
const EntryApply = "apply"

// EntryBuild is the reporter plugin call entry. Host.Call dispatches it to
// ReporterPlugin.Build, decoding the args into a BuildArgs payload.
const EntryBuild = "build"

// EntryDetect is the adapter plugin detect entry. Host.Call dispatches it to
// AdapterPlugin.Detect, decoding the args into an environment map.
const EntryDetect = "detect"

// EntryEmit is the adapter plugin emit entry. Host.Call dispatches it to
// AdapterPlugin.Emit, passing the args verbatim as the payload.
const EntryEmit = "emit"
