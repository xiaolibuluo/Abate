/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package algebra

// Family is the algebra family an operator belongs to.
type Family string

// FamilyMutation is the code-space (mutation) operator family.
const FamilyMutation Family = "m"

// FamilyChaos is the system-space (chaos) operator family.
const FamilyChaos Family = "c"

// FamilyIdentity is the neutral family of the identity operator.
const FamilyIdentity Family = "e"

// FamilyMixed marks a concurrent product that mixes operator families.
const FamilyMixed Family = "x"
