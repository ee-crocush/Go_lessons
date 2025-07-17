# go-39-ci

## Getting started

.golangci.yml
```yaml
run:
  modules-download-mode: readonly

linters:
  disable-all: true

issues:
  max-issues-per-linter: 0
  max-same-issues: 0
```
