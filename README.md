# nscale-sdk-go

Generated Go clients for the [Nscale](https://nscale.com) public APIs.

Each subdirectory is one service:

| Package     | Import path                                     |
| ----------- | ----------------------------------------------- |
| Identity    | `github.com/nscaledev/nscale-sdk-go/identity`   |
| Region      | `github.com/nscaledev/nscale-sdk-go/region`     |
| Storage     | `github.com/nscaledev/nscale-sdk-go/storage`    |
| Compute     | `github.com/nscaledev/nscale-sdk-go/compute`    |
| Kubernetes  | `github.com/nscaledev/nscale-sdk-go/kubernetes` |
| Reservation | `github.com/nscaledev/nscale-sdk-go/reservation` |

The `common/` package holds shared OpenAPI types referenced by every service client.

**⚠️ Alpha:** The Reservation organization topology API is in alpha. Its endpoints and types may change without notice, so expect bugs and breaking changes.

## Install

    go get github.com/nscaledev/nscale-sdk-go@latest

Pre-1.0: the package layout and types may change before `v1.0.0`.

## Versioning

Semver. Releases are tagged manually by maintainers (`git tag vX.Y.Z && git push --tags && gh release create vX.Y.Z --generate-notes`). Automated release tooling will be added once a dedicated release bot is in place.

## Contributing

**Do not hand-edit the `*.gen.go` files.** They are produced by [`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen) from the vendored `openapi.yaml` spec in each service directory.

To refresh a client, update its `openapi.yaml` from the upstream service repo and run:

    go generate ./...

Bug reports for API behaviour belong in the relevant upstream service repo (`nscaledev/uni-*` or `nscaledev/reservation`). PRs against this repo should be limited to packaging, the codegen config, or the vendored specs.

## Licence

[Apache-2.0](./LICENSE).
