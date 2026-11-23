/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package csmutate

import (
	"context"

	"abate/lang/mutcore"
)

// CsSandboxPolicy returns the CopyPolicy for a native C# project. Source
// files, project manifests and the optional solution file are copied per
// mutant. NuGet packages are restored into the global package cache and
// the bin/obj directories are excluded at the session layer (skippedDirs
// already drops bin/ and obj/). No Once directory is declared: the
// dotnet build picks up restored packages from the host cache. When the
// project carries a local nuget.config or Directory.Build.props, those
// are included in Always.
// NewNative builds a native-sandbox mutation session over a real C#
// project. It validates the toolchain, creates a sandbox root (C# projects
// use the host NuGet cache so Once is empty), and follows the same
// scan-discover-measure-assemble flow as New.
func NewNative(ctx context.Context, root string) (*mutcore.Session, error) {
	cfg, err := defaultConfig()
	if err != nil {
		return nil, err
	}
	policy := CsSandboxPolicy()
	if err := policy.Validate(); err != nil {
		return nil, err
	}
	return mutcore.NewNative(ctx, driver{cfg: cfg}, root, policy)
}

func CsSandboxPolicy() mutcore.CopyPolicy {
	return mutcore.CopyPolicy{
		Always:  []string{"*.cs", "*.csproj", "*.sln", "nuget.config", "Directory.Build.props", "Directory.Build.targets"},
		Once:    nil,
		Exclude: []string{"bin", "obj", "*.user", "*.suo", "*.cache"},
	}
}
