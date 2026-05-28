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

## Install

    go get github.com/nscaledev/nscale-sdk-go@latest

Pre-1.0: the package layout and types may change before `v1.0.0`.

## Versioning

Released automatically by [release-please](https://github.com/googleapis/release-please) from Conventional Commits on `main`. See `CHANGELOG.md` (created on the first release).

## Contributing

**Do not hand-edit the `*.gen.go` files.** They are produced by [`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen) from the vendored `openapi.yaml` spec in each service directory.

To refresh a client, update its `openapi.yaml` from the upstream service repo and run:

    go generate ./...

Bug reports for API behaviour belong in the relevant upstream service repo (`nscaledev/uni-*` or `nscaledev/reservation`). PRs against this repo should be limited to packaging, the codegen config, or the vendored specs.

## Licence

[Apache-2.0](./LICENSE).
