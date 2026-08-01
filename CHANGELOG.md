## v2.0.2 - 2026-08-01

### Security

- Upgraded `gopkg.in/yaml.v3` to v3.0.1 to address CVE-2022-28948, a denial-of-service vulnerability.

## v2.0.1 - 2026-07-27

### Fixed

- Prevented chained `Find`, `FindStrict`, `FindAll`, and `FindAllStrict` calls from panicking when the root element is missing.

### Changed

- Removed a duplicate `net/url` import.

## v2.0.0 - 2026-07-27

### Changed

- Go 1.26.5 or newer is now required.
- Upgraded `golang.org/x/net` to v0.56.0 and `golang.org/x/text` to v0.39.0.
- Replaced Travis CI with GitHub Actions running tests and a verbose, pinned `govulncheck` scan.

### Security

- Updated the Go toolchain and dependencies to resolve all findings reported by `govulncheck`.

## v1.2.0

### Added

- Error enums which can be accessed using `Root.Error.(soup.Error).Type`. Refer to `examples/errors`.

## v1.1.0

### Added

- Cookies can be added to the HTTP request, either via the `Cookies` map or the `Cookie()` function
- Function `GetWithClient()` provides the ability to send the request with a custom HTTP client
- Function `FindStrict()` finds the first instance of the mentioned tag with the exact matching values of the provided attribute (previously `Find()`)
- Function `FindAllStrict()` finds all the instances of the mentioned tag with the exact matching values of the attributes (previously `FindAll()`)

## Changed

- Function `Find()` now finds the first instance of the mentioned tag with any matching values of the provided attribute.
- Function `FindAll()` now finds all the instances of the mentioned tag with any matching values of the provided attribute.

---
