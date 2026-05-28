# nscale-sdk-go

Generated Go clients for the [Nscale](https://nscale.com) public APIs.

Each subdirectory is one service:

| Package      | Import path                                     |
| ------------ | ----------------------------------------------- |
| Storage      | `github.com/nscaledev/nscale-sdk-go/storage`    |
| Region       | `github.com/nscaledev/nscale-sdk-go/region`     |
| Identity     | `github.com/nscaledev/nscale-sdk-go/identity`   |
| Compute      | `github.com/nscaledev/nscale-sdk-go/compute`    |
| Reservations | `github.com/nscaledev/nscale-sdk-go/reservations` |
| Kubernetes   | `github.com/nscaledev/nscale-sdk-go/kubernetes` |

## Install

    go get github.com/nscaledev/nscale-sdk-go@latest

Pre-1.0: the package layout and types may change before `v1.0.0`.

## Versioning

Semver. Releases are tagged manually by maintainers (`git tag vX.Y.Z && git push --tags && gh release create vX.Y.Z --generate-notes`). Automated release tooling will be added once a dedicated release bot is in place.

## Contributing

**Do not hand-edit the service directories.** They are synced from upstream OpenAPI repositories by automation; edits will be overwritten on the next sync. File issues or open PRs against the relevant upstream repo.

## Licence

[Apache-2.0](./LICENSE).
