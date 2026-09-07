// Package compute is the Go client for the Nscale Compute service.
// The vendored openapi.yaml is copied verbatim from nscaledev/openapi
// (compute/latest/openapi.yaml). Do not hand-edit it or compute.gen.go;
// refresh both with `just update compute` from the repo root.
package compute

//go:generate go tool oapi-codegen -config config.yaml openapi.yaml
