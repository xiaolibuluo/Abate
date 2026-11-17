/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package injection implements the KRN-INJ chaos kernel as a pure-Go,
// fully virtual rewrite of the chaosblade fault-parameter model.
//
// 参考 chaosblade 故障参数分类模型，改进点=全虚拟化（零系统调用、零真实
// socket/进程）+ 对偶补偿（Inverse + 基线快照）+ 确定性（整数计数、排序键）。
// Every fault resolves through SIM-NET, SIM-FS, SIM-VSCHED or SIM-CLOCK state
// held in memory, so one seed always yields one byte-identical event stream and
// no chaosblade Model or ExpModel type enters the public API.
package injection
