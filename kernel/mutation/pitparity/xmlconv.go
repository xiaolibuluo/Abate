/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package pitparity

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

// pitXMLDocument is the root element of the PIT XML mutation report.
type pitXMLDocument struct {
	XMLName   xml.Name         `xml:"mutations"`
	Mutations []pitXMLMutation `xml:"mutation"`
}

// pitXMLMutation is one mutation element of the PIT XML mutation report.
type pitXMLMutation struct {
	Detected      string `xml:"detected,attr"`
	Status        string `xml:"status,attr"`
	SourceFile    string `xml:"sourceFile"`
	MutatedMethod string `xml:"mutatedMethod"`
	LineNumber    int    `xml:"lineNumber"`
	Mutator       string `xml:"mutator"`
	KillingTest   string `xml:"killingTest"`
	Description   string `xml:"description"`
}

// killingTestMarker separates the class and method parts of a killing test.
const killingTestMarker = "[method:"

// ConvertXML converts a PIT XML mutation report into the canonical JSON form.
//
// The conversion reproduces the vendored reference report byte for byte, which
// is what makes the --rerun-pit mode of scripts/pit-parity.sh comparable with
// the checked in artifact.
func ConvertXML(data []byte) ([]byte, error) {
	var doc pitXMLDocument
	if err := xml.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidReport, err)
	}
	if len(doc.Mutations) == 0 {
		return nil, fmt.Errorf("%w: XML report holds no mutation", ErrInvalidReport)
	}
	records := make([]Mutation, 0, len(doc.Mutations))
	for _, item := range doc.Mutations {
		records = append(records, Mutation{
			Status:        item.Status,
			Detected:      item.Detected,
			SourceFile:    item.SourceFile,
			MutatedMethod: item.MutatedMethod,
			LineNumber:    item.LineNumber,
			Mutator:       shortMutator(item.Mutator),
			Description:   item.Description,
			KillingTest:   shortKillingTest(item.KillingTest),
		})
	}
	out, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("pitparity: encode JSON: %w", err)
	}
	return out, nil
}

// shortMutator drops the package prefix of a fully qualified mutator name.
func shortMutator(name string) string {
	if index := strings.LastIndex(name, "."); index >= 0 {
		return name[index+1:]
	}
	return name
}

// shortKillingTest keeps the method signature tail of a killing test.
//
// The reference report stores the fully qualified killing test reduced to its
// method part, so the conversion reproduces the artifact byte for byte.
func shortKillingTest(name string) string {
	if index := strings.LastIndex(name, killingTestMarker); index >= 0 {
		return name[index+len(killingTestMarker):]
	}
	return name
}
