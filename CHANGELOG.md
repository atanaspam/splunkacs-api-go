# Changelog

All notable changes to this project will be documented in this file.

## [1.6.0](https://github.com/atanaspam/splunkacs-api-go/compare/v1.5.0...v1.6.0) (2023-02-12)


### Features

* Add IP Whitelist Get, Create and Delete ([143b7d9](https://github.com/atanaspam/splunkacs-api-go/commit/143b7d948816ed41d26d0ee5a79ebb19443372f2))


### Bug Fixes

* Fix missing examples documentation ([e3f877a](https://github.com/atanaspam/splunkacs-api-go/commit/e3f877a106f3647dcd69b66b91f8298709a846ee))

## [1.5.0](https://github.com/atanaspam/splunkacs-api-go/compare/v1.4.0...v1.5.0) (2023-02-12)


### Features

* Add ability to get the status of the current stack ([032309c](https://github.com/atanaspam/splunkacs-api-go/commit/032309cf318bebef35a7c66c462683d277fd06bb))

## [1.4.0](https://github.com/atanaspam/splunkacs-api-go/compare/v1.3.0...v1.4.0) (2023-02-12)


### Features

* Update API operations results to include SplunkACSResponse instead of the raw http.Response in preparation for decoupling ([d998195](https://github.com/atanaspam/splunkacs-api-go/commit/d998195b7c9a3b716a230c7a72b87c3322789d92))

## [1.3.0](https://github.com/atanaspam/splunkacs-api-go/compare/v1.2.0...v1.3.0) (2023-02-12)


### Features

* add rate limit aware http client logic ([#6](https://github.com/atanaspam/splunkacs-api-go/issues/6)) ([e2b02cb](https://github.com/atanaspam/splunkacs-api-go/commit/e2b02cb096e9c9e610cbd02feca0c20e205fe415))
* adding http client rate-limiting aware logic to the client using ghetto exponential backoff ([9c1d72b](https://github.com/atanaspam/splunkacs-api-go/commit/9c1d72bcd4e66f722b15e573dc53752698f12fc4))
* Update client response model to use new SplunkApiResponse type ([5dfd1e3](https://github.com/atanaspam/splunkacs-api-go/commit/5dfd1e33e12df9a63631dd7fd621c15614be9ffa))

## [1.2.0](https://github.com/atanaspam/splunkacs-api-go/compare/v1.1.2...v1.2.0) (2022-12-02)


### Features

* adding support for operations on indexes ([#5](https://github.com/atanaspam/splunkacs-api-go/issues/5)) ([fb61654](https://github.com/atanaspam/splunkacs-api-go/commit/fb61654cfb0e4397b5641b0acca8df3e8bb9fa43))

### [1.1.2](https://github.com/atanaspam/splunkacs-api-go/compare/v1.1.1...v1.1.2) (2022-11-27)


### Bug Fixes

* fixed issue with expected response code for HEC token update ([#4](https://github.com/atanaspam/splunkacs-api-go/issues/4)) ([4236684](https://github.com/atanaspam/splunkacs-api-go/commit/423668479a4a1c364c22463c2952ccbce9a19c7c))

### [1.1.1](https://github.com/atanaspam/splunkacs-api-go/compare/v1.1.0...v1.1.1) (2022-11-21)


### Bug Fixes

* bad URL for HEC update operations ([#3](https://github.com/atanaspam/splunkacs-api-go/issues/3)) ([1d13312](https://github.com/atanaspam/splunkacs-api-go/commit/1d13312423765941c66e7abfce078f0f14376929))

## [1.1.0](https://github.com/atanaspam/splunkacs-api-go/compare/v1.0.0...v1.1.0) (2022-11-20)


### Features

* hec token update now uses HTTP PUT ([#2](https://github.com/atanaspam/splunkacs-api-go/issues/2)) ([d781c2a](https://github.com/atanaspam/splunkacs-api-go/commit/d781c2a44da45e70211bbde5c10a9294aeccaa61))

## 1.0.0 (2022-11-20)


### Features

* initial release ([#1](https://github.com/atanaspam/splunkacs-api-go/issues/1)) ([b8c0864](https://github.com/atanaspam/splunkacs-api-go/commit/b8c08644b6e9c6c480100edbdc4fcc59b3448f03))
