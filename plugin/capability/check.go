/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package capability

// Check reports whether have grants want. It is a pure gate with no side effects
// and returns ErrCapabilityDenied when the capability is missing.
func Check(have Capabilities, want Capability) error {
	if have.Has(want) {
		return nil
	}
	return ErrCapabilityDenied
}

// CheckMany reports the first missing capability in wants, or nil if all are
// granted. It short-circuits on the first denial.
func CheckMany(have Capabilities, wants []Capability) error {
	for _, w := range wants {
		if err := Check(have, w); err != nil {
			return err
		}
	}
	return nil
}

// CheckAudit is the audited variant of Check. When log is non-nil it appends a
// "use" event on success or a "deny" event on denial. PluginID labels the event
// for forensic attribution.
func CheckAudit(have Capabilities, want Capability, log *AuditLog, pluginID string) error {
	if have.Has(want) {
		if log != nil {
			log.Append(AuditEvent{Kind: "use", PluginID: pluginID, Capability: want, Result: true})
		}
		return nil
	}
	if log != nil {
		log.Append(AuditEvent{Kind: "deny", PluginID: pluginID, Capability: want, Result: false})
	}
	return ErrCapabilityDenied
}

// CheckManyAudit is the audited variant of CheckMany, appending one event per
// checked capability until the first denial.
func CheckManyAudit(have Capabilities, wants []Capability, log *AuditLog, pluginID string) error {
	for _, w := range wants {
		if err := CheckAudit(have, w, log, pluginID); err != nil {
			return err
		}
	}
	return nil
}
