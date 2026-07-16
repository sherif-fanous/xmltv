# Changelog

## [1.2.1] - 2026-07-15

### Changed

- Flatten the `internal/types` package into `xmltv`; the public API is unchanged (`xmltv.Time` and `xmltv.Bool` retain the same names, underlying types, and methods)

## [1.2.0] - 2026-07-14

### Fixed

- Parse all DTD-permitted date and time formats, including partial dates such as `2004` and `200407` and optional timezone offsets

## [1.1.0] - 2025-03-24

### Fixed

- Handle empty `date` attributes and elements

## [1.0.0] - 2025-02-27

_First release._

[1.2.1]: https://github.com/sherif-fanous/xmltv/releases/tag/v1.2.1
[1.2.0]: https://github.com/sherif-fanous/xmltv/releases/tag/v1.2.0
[1.1.0]: https://github.com/sherif-fanous/xmltv/releases/tag/v1.1.0
[1.0.0]: https://github.com/sherif-fanous/xmltv/releases/tag/v1.0.0
