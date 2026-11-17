/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package net provides a deterministic in-memory virtual network for the Abate
// simulator.
//
// Nodes, links and routes are modeled in memory; packet propagation, queuing,
// loss and latency are all resolved through the SIM clock and SIM-RNG so that a
// given seed and traffic pattern always yields a byte-identical event stream.
//
// Bandwidth queueing is a standard-library rewrite of the token-bucket and
// packet-slicing algorithms from Shopify/toxiproxy. The improvement is a
// deterministic rewrite: virtual-time byte tokens are integers instead of
// wall-clock sleeps, large payloads are sliced before their transmit duration
// is summed, and reservations sharing a virtual instant are totally ordered by
// packet id. The result keeps the hundred-run event-stream hash unique.
package net
