/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package ktmutate

import (
	"context"

	"abate/lang/mutcore"
)

// KtSandboxPolicy returns the CopyPolicy for a native Kotlin project.
// Source files and the build manifests are copied per mutant. The Gradle
// cache directory (which holds downloaded dependencies) is linked from
// the sandbox root when present so kotlinc invocations can reuse the
// pre-resolved classpath. The build output directory is excluded at the
// session layer through skippedDirs.
// NewNative builds a native-sandbox mutation session over a real Kotlin
// project. It validates the toolchain, creates a sandbox root with the
// Gradle cache directories linked once, and follows the same
// scan-discover-measure-assemble flow as New.
func NewNative(ctx context.Context, root string) (*mutcore.Session, error) {
	cfg, err := defaultConfig()
	if err != nil {
		return nil, err
	}
	policy := KtSandboxPolicy()
	if err := policy.Validate(); err != nil {
		return nil, err
	}
	return mutcore.NewNative(ctx, driver{cfg: cfg}, root, policy)
}

func KtSandboxPolicy() mutcore.CopyPolicy {
	return mutcore.CopyPolicy{
		Always:  []string{"*.kt", "*.kts", "build.gradle.kts", "build.gradle", "gradle.properties", "settings.gradle.kts", "settings.gradle"},
		Once:    []string{"gradle", ".gradle", "lib"},
		Exclude: []string{"build", ".kotlin", "*.class", "*.jar"},
	}
}
