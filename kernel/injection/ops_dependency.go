/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package injection

import "abate/kernel/algebra"

// applyDependency sets one pure in-memory dependency fault flag.
func applyDependency(w *World, op algebra.Operator) error {
	params, err := paramsOf(op)
	if err != nil {
		return err
	}
	switch params.class {
	case "dependency-error":
		w.depError = true
	case "dependency-timeout":
		w.depTimeout = true
	default:
		return ErrUnknownOperator
	}
	return nil
}

// recoverDependency clears the dependency fault flags.
func recoverDependency(w *World, op algebra.Operator) error {
	params, err := paramsOf(op)
	if err != nil {
		return err
	}
	switch params.class {
	case "dependency-error", "dependency-timeout":
		w.depError = false
		w.depTimeout = false
	default:
		return ErrUnknownOperator
	}
	return nil
}
