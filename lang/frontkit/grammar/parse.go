/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package grammar

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"
)

// ErrMalformedGrammar reports an upstream document that cannot be decoded.
var ErrMalformedGrammar = errors.New("grammar: malformed upstream document")

// nodeTypesEntry is one entry of the upstream node-types.json document.
//
// The upstream file mixes three shapes in one array: plain named nodes with
// fields, anonymous nodes with children, and abstract supertypes with
// subtypes. The children slot is a boolean in some entries and an object in
// others, so it stays raw and only the field names are decoded.
type nodeTypesEntry struct {
	Type     string                     `json:"type"`
	Named    bool                       `json:"named"`
	Fields   map[string]json.RawMessage `json:"fields"`
	Children json.RawMessage            `json:"children"`
	Subtypes []nodeTypesEntry           `json:"subtypes"`
}

// grammarFile is the decoded upstream grammar.json document.
//
// Only the fields the Abate frontends consume are typed; every other section
// is kept as raw JSON so an upstream change never breaks the decoder.
type grammarFile struct {
	Name       string                     `json:"name"`
	Word       string                     `json:"word"`
	Rules      map[string]json.RawMessage `json:"rules"`
	Supertypes []string                   `json:"supertypes"`
	Extras     []json.RawMessage          `json:"extras"`
	Inline     []string                   `json:"inline"`
	Externals  []json.RawMessage          `json:"externals"`
	Conflicts  []json.RawMessage          `json:"conflicts"`
	Reserved   json.RawMessage            `json:"reserved"`
}

// production is one decoded production body of the upstream grammar.json.
//
// The value field holds a string for ALIAS, PATTERN and STRING but a number
// for the PREC family, so it stays raw until the shape is known.
type production struct {
	Type    string          `json:"type"`
	Name    string          `json:"name"`
	Value   json.RawMessage `json:"value"`
	Members []production    `json:"members"`
	Content *production     `json:"content"`
}

// Parse decodes one upstream node-types.json and grammar.json pair into a
// grammar fact set.
//
// The language tag, version and source are recorded verbatim so a report can
// state which upstream revision the inventory was taken from.
func Parse(language, version, source string,
	nodeTypes, productions []byte) (*Grammar, error) {
	if language == "" {
		return nil, fmt.Errorf("%w: missing language tag", ErrMalformedGrammar)
	}
	entries, err := decodeNodeTypes(nodeTypes)
	if err != nil {
		return nil, err
	}
	file, err := decodeGrammar(productions)
	if err != nil {
		return nil, err
	}
	return build(language, version, source, entries, file)
}

// decodeNodeTypes decodes the upstream node type array.
func decodeNodeTypes(data []byte) ([]nodeTypesEntry, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("%w: empty node-types document", ErrMalformedGrammar)
	}
	var entries []nodeTypesEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, fmt.Errorf("%w: node types: %w", ErrMalformedGrammar, err)
	}
	return entries, nil
}

// decodeGrammar decodes the upstream grammar document.
func decodeGrammar(data []byte) (grammarFile, error) {
	var file grammarFile
	if len(data) == 0 {
		return file, fmt.Errorf("%w: empty grammar document", ErrMalformedGrammar)
	}
	if err := json.Unmarshal(data, &file); err != nil {
		return grammarFile{}, fmt.Errorf("%w: grammar: %w", ErrMalformedGrammar, err)
	}
	if len(file.Rules) == 0 {
		return grammarFile{}, fmt.Errorf("%w: grammar has no rules", ErrMalformedGrammar)
	}
	return file, nil
}

// build merges the two upstream documents into one sorted fact set.
func build(language, version, source string,
	entries []nodeTypesEntry, file grammarFile) (*Grammar, error) {
	supertypes := make(map[string]struct{}, len(file.Supertypes))
	for _, name := range file.Supertypes {
		supertypes[name] = struct{}{}
	}
	nodes := make([]NodeType, 0, len(entries))
	for _, entry := range entries {
		if entry.Type == "" {
			return nil, fmt.Errorf("%w: node type without a name", ErrMalformedGrammar)
		}
		_, super := supertypes[entry.Type]
		nodes = append(nodes, NodeType{
			Name:      entry.Type,
			Named:     entry.Named,
			Supertype: super,
			Fields:    sortedKeys(entry.Fields),
			Subtypes:  subtypeNames(entry.Subtypes),
		})
	}
	slices.SortFunc(nodes, func(a, b NodeType) int {
		return strings.Compare(a.Name, b.Name)
	})
	rules := make([]Rule, 0, len(file.Rules))
	for _, name := range sortedKeys(file.Rules) {
		body, err := decodeProduction(file.Rules[name])
		if err != nil {
			return nil, fmt.Errorf("%w: rule %q: %w", ErrMalformedGrammar, name, err)
		}
		members := make([]string, 0, 4)
		collectMembers(body, &members, make(map[string]struct{}))
		rules = append(rules, Rule{Name: name, Kind: body.Type, Members: members})
	}
	slices.SortFunc(rules, func(a, b Rule) int {
		return strings.Compare(a.Name, b.Name)
	})
	supertypeNames := slices.Clone(file.Supertypes)
	slices.Sort(supertypeNames)
	return &Grammar{
		Language:   language,
		Version:    version,
		Source:     source,
		Nodes:      nodes,
		Rules:      rules,
		Supertypes: supertypeNames,
	}, nil
}

// decodeProduction decodes one production body of the upstream grammar.
func decodeProduction(raw json.RawMessage) (production, error) {
	var body production
	if err := json.Unmarshal(raw, &body); err != nil {
		return production{}, err
	}
	return body, nil
}

// collectMembers appends the production names referenced by one body.
//
// Symbols are collected in source order and kept once, so the member list is
// a faithful and deterministic rendering of the upstream rule.
func collectMembers(body production, out *[]string, seen map[string]struct{}) {
	switch body.Type {
	case "SYMBOL":
		appendMember(out, seen, body.Name)
	case "ALIAS":
		var alias string
		if err := json.Unmarshal(body.Value, &alias); err == nil {
			appendMember(out, seen, alias)
		}
	}
	for _, member := range body.Members {
		collectMembers(member, out, seen)
	}
	if body.Content != nil {
		collectMembers(*body.Content, out, seen)
	}
}

// appendMember appends one production name once.
func appendMember(out *[]string, seen map[string]struct{}, name string) {
	if name == "" {
		return
	}
	if _, ok := seen[name]; ok {
		return
	}
	seen[name] = struct{}{}
	*out = append(*out, name)
}

// sortedKeys returns the keys of one JSON object in ascending order.
func sortedKeys[V any](in map[string]V) []string {
	out := make([]string, 0, len(in))
	for key := range in {
		out = append(out, key)
	}
	slices.Sort(out)
	return out
}

// subtypeNames returns the sorted names of the subtypes of one entry.
func subtypeNames(entries []nodeTypesEntry) []string {
	out := make([]string, 0, len(entries))
	for _, entry := range entries {
		out = append(out, entry.Type)
	}
	slices.Sort(out)
	return out
}
