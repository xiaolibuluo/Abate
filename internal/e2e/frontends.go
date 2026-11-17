/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package e2e

import (
	"errors"
	"fmt"
	"slices"

	"abate/lang/air"
	"abate/lang/cfront"
	"abate/lang/csfront"
	"abate/lang/gofront"
	"abate/lang/jsfront"
	"abate/lang/ktfront"
	"abate/lang/pyfront"
	"abate/lang/rsfront"
)

// errNoFrontend reports a language tag outside the seven registered frontends.
var errNoFrontend = errors.New("e2e: unsupported language")

// FrontendTable binds the seven canonical language tags to their frontends and
// dialects. Go is constructed through NewGoFrontend with its default
// typeCheck=true; no type-check bypass is ever applied.
type FrontendTable struct {
	byLang  map[string]air.Frontend
	dialect map[string]string
}

// NewFrontendTable constructs all seven frontends and returns an error when any
// constructor fails, so a partial table can never be used.
func NewFrontendTable() (FrontendTable, error) {
	py, err := pyfront.NewPyFrontend()
	if err != nil {
		return FrontendTable{}, err
	}
	js, err := jsfront.NewJsFrontend()
	if err != nil {
		return FrontendTable{}, err
	}
	c, err := cfront.NewCFrontend()
	if err != nil {
		return FrontendTable{}, err
	}
	cs, err := csfront.NewCsFrontend()
	if err != nil {
		return FrontendTable{}, err
	}
	rs, err := rsfront.NewRsFrontend()
	if err != nil {
		return FrontendTable{}, err
	}
	kt, err := ktfront.NewKtFrontend()
	if err != nil {
		return FrontendTable{}, err
	}
	return FrontendTable{
		byLang: map[string]air.Frontend{
			"go": gofront.NewGoFrontend(), "py": py, "js": js,
			"c": c, "cs": cs, "rs": rs, "kt": kt,
		},
		dialect: dialectTable,
	}, nil
}

// For returns one frontend and its dialect. The error is non-nil only for a
// language tag outside the canonical table.
func (t FrontendTable) For(lang string) (air.Frontend, string, error) {
	frontend, found := t.byLang[lang]
	if !found || frontend == nil {
		return nil, "", fmt.Errorf("%w: %s", errNoFrontend, lang)
	}
	return frontend, t.dialect[lang], nil
}

// Languages returns the seven tags in canonical ascending report order.
func (FrontendTable) Languages() []string {
	return slices.Clone(languageOrder)
}
