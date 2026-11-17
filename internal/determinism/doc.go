/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package determinism provides the deterministic primitives shared by the
// simulator and the native plugin host: canonical hashing and a same-seed
// double-run assertion for reproducibility, plus the cooperative fuel meter
// (FuelMeter), virtual clock (VirtualClock) and pure-data execution config
// (Config) that were migrated into the host layer as part of the WAS-HOST
// decision. Their behavior is frozen by the T01 golden observation under
// docs/refactor/golden/.
package determinism
