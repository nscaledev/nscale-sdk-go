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

Semver. Releases are tagged manually by maintainers (`git tag vX.Y.Z && git push --tags && gh release create vX.Y.Z --generate-notes`). Automated release tooling will be added once a dedicated release bot is in place.

## Contributing

**Do not hand-edit the `*.gen.go` files.** They are produced by [`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen) from the vendored `openapi.yaml` spec in each service directory.

**Do not hand-edit the vendored `openapi.yaml` files either.** Each is a verbatim copy of a spec published in [`nscaledev/openapi`](https://github.com/nscaledev/openapi); the source path is recorded in that package's `gen.go`. Specs are vendored rather than fetched at generate time so `go generate` stays hermetic and works offline.

To refresh a client, re-copy its spec from `nscaledev/openapi` and regenerate:

    curl -sfL https://raw.githubusercontent.com/nscaledev/openapi/main/identity/latest/openapi.yaml \
      -o identity/openapi.yaml
    go generate ./...

`<service>/latest/` always mirrors that service's newest stable release, so a refresh picks up new releases without changing any path here. Note the `info.version` field inside a published spec is not maintained upstream and does not track the release — the version folder name does.

Bug reports for API behaviour belong in the relevant upstream service repo (`nscaledev/uni-*` or `nscaledev/reservation`), not here and not in `nscaledev/openapi`, where specs are published by CI and never hand-edited. PRs against this repo should be limited to packaging, the codegen config, or refreshing the vendored specs.

## Licence

[Apache-2.0](./LICENSE).
