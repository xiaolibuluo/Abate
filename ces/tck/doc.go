/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

// Package tck is the CES Technology Compatibility Kit.
//
// A third-party CES implementation becomes "CES compliant" by passing every
// case in this kit. The kit is the neutral arbiter of the open CES
// specification: every case is data (YAML under cases/) and each case cites
// the specification clause it exercises, so a case's expected outcome can be
// derived from the spec text rather than from Abate's private behavior.
//
// The core of the kit (this package, excluding abate.go) depends only on
// CES-SCHEMA and the standard library. abate.go ships an AbateDriver implementing
// the Driver interface using Abate's own parse/normalize/validate pipeline; it
// is the first implementation that must pass the kit 100% (eating our own dog
// food). A minimal independent reference driver (see tck_test.go) proves the
// kit does not depend on Abate-specific behavior.
//
// Layout:
//
//	tck.go     Runner and the Run entry point.
//	driver.go  Driver interface and the neutral result shapes.
//	case.go    Case model, loading and embedding of the case library.
//	report.go  Compliance Report and the Compliant judgment.
//	abate.go   AbateDriver: the bundled reference implementation.
//	cases/     Data-driven case library (parse / validate / roundtrip).
package tck
