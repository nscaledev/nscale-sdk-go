// Package kubernetes is the Go client for the Nscale Kubernetes service (NKS).
// The vendored openapi.yaml is copied verbatim from nscaledev/openapi
// (nks-core/latest/openapi.yaml). Do not hand-edit it or kubernetes.gen.go;
// refresh both with `just update kubernetes` from the repo root.
package kubernetes

//go:generate go tool oapi-codegen -config config.yaml openapi.yaml
