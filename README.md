# Patient Reservation

## Getting Started

### Requirements

- Docker >= 20.10.15
- (optional) Golang >= 1.18.7
- (optional) `make` command (see [GNU Make](https://www.gnu.org/software/make/))

### Setup

1. Build docker image

`docker build -t patientreservation:latest .`

or you can run `make image`

2. Run docker container

`docker run --rm -p 4040:4040 patientreservation:latest`

or you can run `make container`

```
$ make container
docker run --rm -p 4040:4040 patientreservation:latest

 ┌───────────────────────────────────────────────────┐
 │                   Fiber v2.43.0                   │
 │               http://127.0.0.1:4040               │
 │       (bound on host 0.0.0.0 and port 4040)       │
 │                                                   │
 │ Handlers ............. 8  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ................. 1 │
 └───────────────────────────────────────────────────┘

```

Please raise an issue if you find difficulties.

### Test

To have the unit test, run `go test -v -failfast ./app`

or you can run `make unittest`

## Product Requirement Document

[Product Requirement Document](./docs/prd.md)

## API Contract

[API Contract](./docs/api_contract.md)
