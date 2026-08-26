# Spectral Calibration Lab

Spectral Calibration Lab is a small HTTP service for research teams that prepare repeatable optical calibration runs. A laboratory operator can inspect supported calibration profiles, submit a batch of measured values, and retrieve a quality summary for a profile. The service is fully in-memory so it can be built, tested, and run without network services.

## Users and workflows

- **Laboratory operator** selects a calibration profile and submits a sequence of readings.
- **Research engineer** reviews the acceptance score, spread, and recommended next action.
- **Method steward** compares a profile's reference center and tolerance with recent runs.

The normal workflows are: inspect profiles, submit an observation batch, and request a profile quality summary.

## Layout

- `cmd/calibrationd`: service entrypoint and server configuration.
- `internal/lab`: domain types, profile catalog, scoring logic, and run storage.
- `internal/httpapi`: JSON handlers and request validation.

## Run locally

```powershell
go build ./...
go test ./...
go run ./cmd/calibrationd
```

The service listens on `127.0.0.1:18084` by default. Set `CALIBRATION_ADDR` to use another address.

## HTTP endpoints

- `GET /healthz` returns the process status.
- `GET /v1/profiles` returns supported calibration profiles.
- `POST /v1/observations` accepts a `profile_id` and numeric `values`.
- `GET /v1/reports?profile_id=thermal-stability` returns a summary for one profile.

## Environment

`CALIBRATION_ADDR` is optional. When unset, the service uses `127.0.0.1:18084`.
