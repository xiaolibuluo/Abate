// Abate —— 止息混沌 · 全语言构建/检查/测试命令分发
// Copyright (C) 2026 小黎的部落（www.qingdujian.com）
// AGPLv3; 商业用途须事先获得小黎的部落书面授权

package test

// BuildCommand 返回项目类型对应的构建命令。
func BuildCommand(pt ProjectType) []string {
	switch pt {
	case ProjectGo:
		return []string{"go", "build", "./..."}
	case ProjectGradle:
		return []string{"gradle", "compileDebugKotlin", "--no-daemon"}
	case ProjectMaven:
		return []string{"mvn", "compile", "-q"}
	case ProjectNpm:
		return []string{"npm", "run", "build"}
	case ProjectCargo:
		return []string{"cargo", "build"}
	case ProjectPython:
		return []string{"python3", "-m", "compileall", "."}
	case ProjectSwift:
		return []string{"swift", "build"}
	case ProjectRuby:
		return []string{"gem", "build"}
	case ProjectPhp:
		return []string{"composer", "install", "--no-dev", "--quiet"}
	case ProjectDotNet:
		return []string{"dotnet", "build"}
	case ProjectCpp:
		return []string{"cmake", "--build", "."}
	case ProjectScala:
		return []string{"sbt", "compile"}
	case ProjectDart:
		return []string{"dart", "compile", "exe"}
	default:
		return nil
	}
}

// VetCommand 返回项目类型对应的静态检查命令。
func VetCommand(pt ProjectType) []string {
	switch pt {
	case ProjectGo:
		return []string{"go", "vet", "./..."}
	case ProjectGradle:
		return []string{"gradle", "lint", "--no-daemon"}
	case ProjectMaven:
		return []string{"mvn", "checkstyle:check", "-q"}
	case ProjectNpm:
		return []string{"npm", "run", "lint"}
	case ProjectCargo:
		return []string{"cargo", "clippy"}
	case ProjectPython:
		return []string{"python3", "-m", "pyflakes", "."}
	case ProjectSwift:
		return []string{"swift", "lint"}
	case ProjectRuby:
		return []string{"rubocop"}
	case ProjectPhp:
		return []string{"php", "-l", "src"}
	case ProjectDotNet:
		return []string{"dotnet", "format", "--verify-no-changes"}
	case ProjectCpp:
		return []string{"cppcheck", "."}
	case ProjectScala:
		return []string{"sbt", "scalafmtCheck"}
	case ProjectDart:
		return []string{"dart", "analyze"}
	default:
		return nil
	}
}

// TestCommand 返回项目类型对应的测试命令。
func TestCommand(pt ProjectType) []string {
	switch pt {
	case ProjectGo:
		return []string{"go", "test", "./..."}
	case ProjectGradle:
		return []string{"gradle", "test", "--no-daemon"}
	case ProjectMaven:
		return []string{"mvn", "test", "-q"}
	case ProjectNpm:
		return []string{"npm", "test"}
	case ProjectCargo:
		return []string{"cargo", "test"}
	case ProjectPython:
		return []string{"python3", "-m", "pytest"}
	case ProjectSwift:
		return []string{"swift", "test"}
	case ProjectRuby:
		return []string{"bundle", "exec", "rspec"}
	case ProjectPhp:
		return []string{"vendor/bin/phpunit"}
	case ProjectDotNet:
		return []string{"dotnet", "test"}
	case ProjectCpp:
		return []string{"ctest", "--output-on-failure"}
	case ProjectScala:
		return []string{"sbt", "test"}
	case ProjectDart:
		return []string{"dart", "test"}
	default:
		return nil
	}
}
