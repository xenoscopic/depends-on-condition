# `depends_on: condition`

This project contains an example of using the `condition` form of `depends_on`
in Compose projects. This functionality was recently revived as part of the
[Compose Spec](https://compose-spec.io) and requires Docker Compose v1.27.0+.

This project grew out of a [talk](https://havoc.io/talks/depends-on-condition)
given at the Docker Community All Hands meeting on 11 March 2021.


## Design

There are two services: `describer` and `web`. `web` is defined to have a
dependency on `describer` (using the `condition: service_healthy` form of
`depends_on`), so Compose waits for `describer` to start and become healthy
before starting `web`.

The `describer` service has a simulated startup time of 10 seconds and a health
check defined that dependent services can use to wait on. The `web` service has
a simulated hard dependency on the `describer` service, exiting if the latter
isn't available immediately when `web` starts.


## Usage

To run the project, use:

```
docker-compose up --build
```

After the `describer` service starts, you'll see the project wait for about 10
seconds while it waits for the service to become healthy. After that, the `web`
service will start and be available on
[`http://localhost:8080`](http://localhost:8080).

Refresh the page repeatedly for inspirational messages (and to ensure that the
services are communicating).

To tear down the project, use:

```
docker-compose down --rmi=local
```
