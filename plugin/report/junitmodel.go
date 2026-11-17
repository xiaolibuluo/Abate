/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package report

import "encoding/xml"

// junitTimeNotMeasured is a schema-valid placeholder for the JUnit time
// attribute. Timing is deliberately NOT measured: BuildJUnit is a pure function
// of its inputs and must stay byte-stable across runs, so no wall-clock or
// elapsed duration may enter the document. The value "0" is a required numeric
// placeholder, NOT a claim that any test or run took zero seconds; the emitted
// report.time property states "NOT_MEASURED" explicitly so no consumer can
// mistake the attribute for a real timing.
const junitTimeNotMeasured = "0"

// junitTimeNotMeasuredMarker is the honest semantic label written into the
// JUnit properties to disambiguate the numeric time placeholder.
const junitTimeNotMeasuredMarker = "NOT_MEASURED"

// junitDocument is the root testsuites element of the offline JUnit document.
type junitDocument struct {
	XMLName xml.Name     `xml:"testsuites"`
	Suites  []junitSuite `xml:"testsuite"`
}

// junitSuite is one testsuite element that groups the generated test cases.
type junitSuite struct {
	Name       string          `xml:"name,attr"`
	Tests      int             `xml:"tests,attr"`
	Failures   int             `xml:"failures,attr"`
	Skipped    int             `xml:"skipped,attr"`
	Errors     int             `xml:"errors,attr"`
	Time       string          `xml:"time,attr"`
	Properties junitProperties `xml:"properties"`
	Cases      []junitCase     `xml:"testcase"`
}

// junitProperties wraps the deterministic report properties element.
type junitProperties struct {
	Properties []junitProperty `xml:"property"`
}

// junitProperty is one deterministic name/value property of the run.
type junitProperty struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// junitCase is one testcase element for a mutant or a rule finding.
type junitCase struct {
	Name      string         `xml:"name,attr"`
	Classname string         `xml:"classname,attr"`
	Time      string         `xml:"time,attr"`
	Failures  []junitFailure `xml:"failure,omitempty"`
	Skipped   *junitSkipped  `xml:"skipped,omitempty"`
}

// junitFailure is one failure element with escaped character data.
type junitFailure struct {
	Message string `xml:"message,attr"`
	Type    string `xml:"type,attr"`
	Body    string `xml:",chardata"`
}

// junitSkipped is one skipped element for equivalent or timed-out mutants.
type junitSkipped struct {
	Message string `xml:"message,attr"`
}
