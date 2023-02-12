# Changelog

All notable changes to this project will be documented in this file.

## [2.0.0](https://github.com/atanaspam/splunkacs-api-go/compare/v1.3.0...v2.0.0) (2023-02-12)


### ⚠ BREAKING CHANGES

* update API operations results to include SplunkACSResponse (#7)
* Update API operations results to include SplunkACSResponse instead of the raw http.Response in preparation for decoupling

### Features

* update API operations results to include SplunkACSResponse ([#7](https://github.com/atanaspam/splunkacs-api-go/issues/7)) ([d309384](https://github.com/atanaspam/splunkacs-api-go/commit/d309384bdae3f8d67757e114a943f89012f301de))
* Update API operations results to include SplunkACSResponse instead of the raw http.Response in preparation for decoupling ([2b529e9](https://github.com/atanaspam/splunkacs-api-go/commit/2b529e9d9a9fd2ac6956e9e289745a4c28b31319))

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
