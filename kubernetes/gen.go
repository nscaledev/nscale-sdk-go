// Package kubernetes is the Go client for the Nscale Kubernetes service (NKS).
// The vendored openapi.yaml is copied verbatim from nscaledev/openapi
// (nks-core/main/openapi.yaml). Do not hand-edit it or kubernetes.gen.go.
//
// NKS has not cut a stable release yet, so this client tracks the source
// service's main branch rather than a tagged version: unlike the other
// packages here, its API surface can change without a version bump.
package kubernetes

//go:generate go tool oapi-codegen -config config.yaml openapi.yaml
