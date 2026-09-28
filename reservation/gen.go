// Package reservation is the Go client for the Nscale Reservation service.
// The vendored openapi.yaml is copied verbatim from nscaledev/openapi
// (reservation/latest/openapi.yaml). Do not hand-edit it or reservation.gen.go;
// refresh both with `just update reservation` from the repo root.
package reservation

//go:generate go tool oapi-codegen -config config.yaml openapi.yaml
