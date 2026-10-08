// Package reservation is the Go client for the Nscale Reservation service.
// The vendored openapi.yaml is copied verbatim from nscaledev/openapi
// (reservation/main/openapi.yaml). Do not hand-edit it or reservation.gen.go;
// refresh both with `just update reservation` from the repo root.
//
// Reservation releases are not reaching reservation/latest, so this client
// tracks the source service's main branch rather than a tagged version: unlike
// most packages here, its API surface can change without a version bump.
package reservation

//go:generate go tool oapi-codegen -config config.yaml openapi.yaml
