// Package region is the Go client for the Nscale Region service.
// The vendored openapi.yaml is copied verbatim from nscaledev/openapi
// (region/latest/openapi.yaml). Do not hand-edit it or region.gen.go;
// refresh both with `just update region` from the repo root.
package region

//go:generate go tool oapi-codegen -config config.yaml openapi.yaml
