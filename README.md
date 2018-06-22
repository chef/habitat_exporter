# Habitat prometheus exporter

Queries a [habitat](https://habitat.sh) supervisor and exports useful metrics
from it.

## Usage

```
./habitat_exporter
```

There are two optional flags:

* `-listen-address` - Change the address and port the exporter listens on
* `-habitat-address` - Change the address of the habitat supervisor that is
  queried.

## Metrics

* `habitat_service_health` - The health of each habitat service that is
  active. Because prometheus can only output numeric values, the metric values
  match the habitat health_check hook (and also nagios check) exit codes: 0
  for OK, 1 for Warning, 2 for Critical, 3 for Unknown.
* `habitat_service_package_release` - The package release number (e.g.
  20180622123456) of the habitat package for each habitat service.

## Building

```
go build
```
