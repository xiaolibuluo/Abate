/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package scheduler

// Handler executes one event and returns the derived events to enqueue.
//
// It must be a pure function of its argument: the scheduler guarantees that
// derived events receive fresh sequence numbers, so a handler cannot inject
// ordering decisions from outside the queue.
type Handler func(Event) ([]Event, error)
