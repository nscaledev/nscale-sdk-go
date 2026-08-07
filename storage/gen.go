// Package storage is the Go client for the Nscale Storage service.
// The vendored openapi.yaml is copied verbatim from nscaledev/openapi
// (storage/latest/openapi.yaml). Do not hand-edit it or storage.gen.go.
package storage

//go:generate go tool oapi-codegen -config config.yaml openapi.yaml
