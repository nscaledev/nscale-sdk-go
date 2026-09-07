// Package identity is the Go client for the Nscale Identity service.
// The vendored openapi.yaml is copied verbatim from nscaledev/openapi
// (identity/latest/openapi.yaml). Do not hand-edit it or identity.gen.go;
// refresh both with `just update identity` from the repo root.
package identity

//go:generate go tool oapi-codegen -config config.yaml openapi.yaml
