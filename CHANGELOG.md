# Changelog

AI-generated changelog based on git diff and commit messages.

## [v1.0.1] - 2025-11-28

### Added
- Added `LogPrettyJSON`, `PrintPrettyJSONBytes`, and `PrintPrettyJSONReader` for formatted JSON logging
- Added `internal/jsonformat` package to handle JSON indentation logic

### Changed
- Bumped Go version to `1.25.4`
- Refactored function signatures to use `any` instead of `interface{}`
- Moved tests to a separate `tests/` package for better black-box testing isolation
- Replaced standard `encoding/json` with `github.com/goccy/go-json` via module replacement
- Updated `xprint` dependency to v0.0.7
- Pinned GitHub Action versions to specific SHAs for improved security
- Changed Dependabot schedule from weekly to monthly

### Removed
- Removed shell script wrappers (`lint.sh`, `full-lint.sh`) in favor of direct workflow commands
- Removed unused `utils.go` file

## [v1.0.0] - 2025-04-05

### Added
- Bugfixes and bump to v1

## [v0.0.6] - 2025-03-25

### Added
- Added httplogger package with HTTP request/response logging support
- Added internal/colors package for color management
- Added internal/formats package 
- Added test.sh script for running tests
- Added reflect_test.go for testing reflection features
- Added println.go for enhanced printing capabilities

### Changed
- Updated project structure with new directory organization
- Enhanced error handling in errorlogs.go
- Improved documentation in README files across the project
- Updated Go module dependencies
- Refined CI/CD workflows

### Fixed
- Various code issues addressed by linting updates
- Improved error handling and reporting

## [v0.0.5] - 2025-02-20

### Added
- Added LogStruct and LogSlice functions for improved structured logging

### Changed
- Rewrote core functionality to remove dependency on standard library's log package
- Fixed incorrect directory structure

### Fixed
- Small fixes to improve overall stability

## [v0.0.4] - 2024-09-13

### Fixed
- Fixed incorrect directory structure

## [v0.0.3] - 2024-09-13

### Added
- Added more printf style functions for enhanced logging capabilities

### Changed
- Updated module name to fix import paths

## [v0.0.2] - 2024-08-27

### Fixed
- Fixed module name for proper importing

## [v0.0.1] - 2024-08-27

### Added
- Initial commit with basic logging functionality
