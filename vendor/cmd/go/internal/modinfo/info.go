// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modinfo

import "time"

// Note that these structs are publicly visible (part of go list's API)
// and the fields are documented in the help text in ../list/list.go

type ModulePublic struct {
	Path    string        `json:",omitempty"` // module path
	Version string        `json:",omitempty"` // module version
	Replace *ModulePublic `json:",omitempty"` // replaced by this module
	Time    *time.Time    `json:",omitempty"` // time version was created
	Update  *ModulePublic `json:",omitempty"` // available update (with -u)
	Main    bool          `json:",omitempty"` // is this the main module?
	Dir     string        `json:",omitempty"` // directory holding local copy of files, if any
	Error   *ModuleError  `json:",omitempty"` // error loading module
}

type ModuleError struct {
	Err string // error text
}

func (m *ModulePublic) String() string {
	s := m.Path
	if m.Version != "" {
		s += " " + m.Version
		if m.Update != nil {
			s += " [" + m.Update.Version + "]"
		}
	}
	if m.Replace != nil {
		s += " => " + m.Replace.Path
		if m.Replace.Version != "" {
			s += " " + m.Replace.Version
			if m.Replace.Update != nil {
				s += " [" + m.Replace.Update.Version + "]"
			}
		}
	}
	return s
}
