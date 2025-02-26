// Copyright (c) 2021 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by tailscale.com/cmd/cloner; DO NOT EDIT.

package persist

import (
	"tailscale.com/tailcfg"
	"tailscale.com/types/key"
	"tailscale.com/types/structs"
)

// Clone makes a deep copy of Persist.
// The result aliases no memory with the original.
func (src *Persist) Clone() *Persist {
	if src == nil {
		return nil
	}
	dst := new(Persist)
	*dst = *src
	dst.DisallowedTKAStateIDs = append(src.DisallowedTKAStateIDs[:0:0], src.DisallowedTKAStateIDs...)
	return dst
}

// A compilation failure here means this code must be regenerated, with the command at the top of this file.
var _PersistCloneNeedsRegeneration = Persist(struct {
	_                               structs.Incomparable
	LegacyFrontendPrivateMachineKey key.MachinePrivate
	PrivateNodeKey                  key.NodePrivate
	OldPrivateNodeKey               key.NodePrivate
	Provider                        string
	LoginName                       string
	UserProfile                     tailcfg.UserProfile
	NetworkLockKey                  key.NLPrivate
	NodeID                          tailcfg.StableNodeID
	DisallowedTKAStateIDs           []string
}{})
