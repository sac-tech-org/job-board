# SacTech Job Board

The SacTech Job Board is a site dedicated to connecting tech employers and job-seekers in Sacramento, CA. This has been created with following goals in mind:

1. Focus on local jobs.
2. Highlight the Sacramento's tech scene's companies, startups, or anyone looking to hire for a project.
3. Provide a project for new developers, those looking to reskill, or anyone who wants to learn, to work on.

## Tech Stack

The stack consists of a [Vue] web app, backed by data in stored in [Postgres], served by a REST API written in [Go].

### Web app

Code: [web](./web/)

### API

A REST API provided by a HTTP server written in [Go]. Documentation provided by an [OpenAPI] specification.

Code: [api](./api/)

API Definition: [OpenAPI Spec](./api/openapi.yaml)

### Database

PostgreSQL version 16.1, running in [Docker].

## Running locally

### Install Tools

Install the following:

- [Docker](https://www.docker.com/get-started/)
- [Node]
- [Go](https://go.dev/learn/)
  - [Air](https://github.com/cosmtrek/air?tab=readme-ov-file#installation) - for hot reloading the backend code. Must be install _after_ Go.

#### Installing Dependencies

Run in your terminal at the root of the project:

```sh
make setup
```

This will install the Go and Node dependencies required by the API and web app.

#### Run it!

Make sure that Docker is running on your computer, and then run in your terminal at the root of the project:

```sh
make dev
```

This will concurrently start the API, listening on HTTP port `8080`, the web app, and the database service in a Docker container.

[Docker]: https://www.docker.com
[Go]: https://go.dev
[Node]: https://nodejs.org/en
[OpenAPI]: https://www.openapis.org
[Postgres]: https://www.postgresql.org
[Vue]: https://vuejs.org
