/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package capability

// AuditEvent is a single recorded capability-sandbox event.
//
// Kind is one of "grant", "deny", or "use". At is a virtual monotonic counter
// assigned by the AuditLog (never time.Now), guaranteeing deterministic ordering
// across runs. Result reports whether the action succeeded.
type AuditEvent struct {
	// Kind is "grant", "deny", or "use".
	Kind string

	// PluginID labels the plugin the event concerns.
	PluginID string

	// Capability is the capability bit involved in the event.
	Capability Capability

	// At is the virtual sequence number assigned by the AuditLog.
	At int64

	// Result reports whether the action succeeded.
	Result bool
}

// IsGrant reports whether the event is a grant event.
func (e AuditEvent) IsGrant() bool { return e.Kind == "grant" }

// IsDeny reports whether the event is a deny event.
func (e AuditEvent) IsDeny() bool { return e.Kind == "deny" }

// IsUse reports whether the event is a use event.
func (e AuditEvent) IsUse() bool { return e.Kind == "use" }
