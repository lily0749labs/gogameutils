# Changelog

本项目的重要变更都会记录在此文件中。版本号遵循 [Semantic Versioning](https://semver.org/)。

## [Unreleased]

## [0.1.1] - 2026-08-26

### Changed

- 升级 `github.com/lily0749labs/goutils` 至 v0.3.0。
- `ScoreToStrCur`、`Unset`、`ValidateData` 和兼容队列改为委托正式工具库实现。
- 修复兼容队列并发关闭时的数据竞争与重复关闭问题。

## [0.1.0] - 2026-08-26

### Added

- 首次公开发布。
- 提供扑克牌算法、行为树、定时任务、TCP 封包及服务配置工具。
- 提供 MySQL、PostgreSQL 和 Redis 连接辅助。

[Unreleased]: https://github.com/lily0749labs/gogameutils/compare/v0.1.1...HEAD
[0.1.1]: https://github.com/lily0749labs/gogameutils/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/lily0749labs/gogameutils/releases/tag/v0.1.0
