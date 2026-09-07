# nscale-sdk-go

Generated Go clients for the [Nscale](https://nscale.com) public APIs.

Each subdirectory is one service. Every client is generated from the canonical spec published in [`nscaledev/openapi`](https://github.com/nscaledev/openapi):

| Package     | Import path                                      | Spec source                |
| ----------- | ------------------------------------------------ | -------------------------- |
| Identity    | `github.com/nscaledev/nscale-sdk-go/identity`    | `identity/latest`          |
| Region      | `github.com/nscaledev/nscale-sdk-go/region`      | `region/latest`            |
| Storage     | `github.com/nscaledev/nscale-sdk-go/storage`     | `storage/latest`           |
| Compute     | `github.com/nscaledev/nscale-sdk-go/compute`     | `compute/latest`           |
| Kubernetes  | `github.com/nscaledev/nscale-sdk-go/kubernetes`  | `nks-core/main`            |
| Reservation | `github.com/nscaledev/nscale-sdk-go/reservation` | `reservation/latest`       |

Published specs are bundled — every `$ref` is dereferenced before publication — so each package is self-contained and shares no types with the others. There is no `common` package.

**⚠️ Alpha:** The Reservation organization topology API is in alpha. Its endpoints and types may change without notice, so expect bugs and breaking changes.

**⚠️ Unstable:** The Kubernetes client tracks NKS's `main` branch, because NKS has not cut a stable release yet. Every other package tracks a tagged release. Expect the Kubernetes surface to change without a version bump.

## Install

    go get github.com/nscaledev/nscale-sdk-go@latest

Pre-1.0: the package layout and types may change before `v1.0.0`.

## Versioning

Semver, released by [release-please](https://github.com/googleapis/release-please). Land [conventional commits](https://www.conventionalcommits.org/) on `main` and release-please maintains a rolling `chore: release X.Y.Z` PR; merging it writes `CHANGELOG.md`, tags `vX.Y.Z`, and publishes the GitHub release. Nothing is tagged by hand.

Pre-1.0, a breaking change (`feat!:` / `BREAKING CHANGE:`) bumps the minor, not the major — so `v0.3.0` → `v0.4.0`, and a spec refresh that removes a field is a minor bump.

## Contributing

**Do not hand-edit the `*.gen.go` files.** They are produced by [`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen) from the vendored `openapi.yaml` spec in each service directory.

**Do not hand-edit the vendored `openapi.yaml` files either.** Each is a verbatim copy of a spec published in [`nscaledev/openapi`](https://github.com/nscaledev/openapi). Specs are vendored rather than fetched at generate time so `go generate` stays hermetic and works offline, and so an upstream spec change lands as a reviewable diff rather than shifting silently under CI.

### Refreshing a client

Everything runs through [`just`](https://github.com/casey/just). The Justfile's `specs` variable is the single source of truth for which upstream spec each package tracks:

    just sources             # show each package and the spec it tracks
    just drift               # which vendored specs have fallen behind upstream
    just update kubernetes   # fetch that spec from nscaledev/openapi, regenerate, build
    just update              # same, for every package
    just check               # what CI runs: build, test, and confirm codegen is current

`just sync` is the only recipe that touches the network. `OPENAPI_REF=<sha> just sync` pins the fetch to a specific commit of `nscaledev/openapi` instead of its `main`.

`<service>/latest/` always mirrors that service's newest stable release, so a refresh picks up new releases without editing the Justfile. Note the `info.version` field inside a published spec is not maintained upstream and does not track the release — the version folder name does.

Commit the regenerated `openapi.yaml` and `*.gen.go` together, and describe the API change rather than the mechanics, so the changelog is useful: `feat(kubernetes)!: refresh client from nks-core`.

Bug reports for API behaviour belong in the relevant upstream service repo (`nscaledev/uni-*` or `nscaledev/reservation`), not here and not in `nscaledev/openapi`, where specs are published by CI and never hand-edited. PRs against this repo should be limited to packaging, the codegen config, or refreshing the vendored specs.

## Licence

[Apache-2.0](./LICENSE).
