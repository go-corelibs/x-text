// Copyright (c) 2024  The Go-CoreLibs Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package message

import (
	"context"
	"net/http"

	"github.com/go-corelibs/x-text/language"
	"github.com/go-corelibs/x-text/message/catalog"
)

type privateContextKey string

const (
	gLanguageTag     privateContextKey = "corelibs-x-text-language-tag"
	gLanguagePrinter privateContextKey = "corelibs-x-text-language-printer"
	gLanguageDefault privateContextKey = "corelibs-x-text-language-default"
)

// NewCatalogPrinter constructs a NewPrinter with the specific catalog.Catalog
// and the parsed `lang` code. If `lang` fails to parse, language.Und is
// returned
//
// NewCatalogPrinter was added by Go-Enjin
func NewCatalogPrinter(lang string, c catalog.Catalog) (tag language.Tag, printer *Printer) {
	var err error
	if tag, err = language.Parse(lang); err != nil {
		tag = language.Und
		printer = NewPrinter(language.Und, Catalog(c))
	} else {
		printer = NewPrinter(tag, Catalog(c))
	}
	return
}

// SetTag associates a private context key with the language.Tag given
//
// SetTag was added by Go-Enjin
func SetTag(r *http.Request, tag language.Tag) (modified *http.Request) {
	modified = r.Clone(context.WithValue(r.Context(), gLanguageTag, tag))
	return
}

// SetDefaultTag associates a private context key with the language.Tag given
//
// SetDefaultTag was added by Go-Enjin
func SetDefaultTag(r *http.Request, tag language.Tag) (modified *http.Request) {
	modified = r.Clone(context.WithValue(r.Context(), gLanguageDefault, tag))
	return
}

// GetTag retrieves the language.Tag assigned to the given request. GetTag
// first checks for the SetTag private context key, and if not found, checks
// for the SetDefaultTag private context key, and if still not found returns
// language.Und
//
// GetTag was added by Go-Enjin
func GetTag(r *http.Request) (tag language.Tag) {
	if p, ok := r.Context().Value(gLanguageTag).(language.Tag); ok {
		tag = p
	} else if p, ok := r.Context().Value(gLanguageDefault).(language.Tag); ok {
		tag = p
	} else {
		tag = language.Und
	}
	return
}

// GetLanguageTag returns the language.Tag assigned to the given request with a
// prior call to SetTag
//
// GetLanguageTag was added by Go-Enjin
func GetLanguageTag(r *http.Request) (tag language.Tag, ok bool) {
	tag, ok = r.Context().Value(gLanguageTag).(language.Tag)
	return
}

// SetPrinter associates a private context key with the given Printer instance
//
// SetPrinter was added by Go-Enjin
func SetPrinter(r *http.Request, printer *Printer) (modified *http.Request) {
	modified = r.Clone(context.WithValue(r.Context(), gLanguagePrinter, printer))
	return
}

// GetPrinter looks for a Printer instance associated with the given
// http.Request and if not present, constructs a NewPrinter instance using the
// SetTag, SetDefaultTag or language.Und
//
// GetPrinter was added by Go-Enjin
func GetPrinter(r *http.Request) (printer *Printer) {
	if p, ok := r.Context().Value(gLanguagePrinter).(*Printer); ok {
		printer = p
	} else if tag, ok := r.Context().Value(gLanguageTag).(language.Tag); ok {
		printer = NewPrinter(tag)
	} else if tag, ok := r.Context().Value(gLanguageDefault).(language.Tag); ok {
		printer = NewPrinter(tag)
	} else {
		printer = NewPrinter(language.Und)
	}
	return
}
