# sish-gitlab-keys

[![Build](https://github.com/trancong12102/sish-gitlab-keys/actions/workflows/build.yml/badge.svg)](https://github.com/trancong12102/sish-gitlab-keys/actions/workflows/build.yml/badge.svg)
[![codecov](https://codecov.io/gh/trancong12102/sish-gitlab-keys/graph/badge.svg?token=YWN7WWNFH9)](https://codecov.io/gh/trancong12102/sish-gitlab-keys)
[![Maintainability](https://api.codeclimate.com/v1/badges/eb2146e7afe5633a0023/maintainability)](https://codeclimate.com/github/trancong12102/sish-gitlab-keys/maintainability)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is an authentication server for [sish](https://github.com/antoniomika/sish) that uses GitLab SSH keys for
authentication.

## Features

- [x] Authenticate **sish** users with GitLab SSH keys
- [x] Health check endpoint at `/health`
- [x] **sish** key authentication endpoint at `/auth`

## Usage

See environment variables: [environments.md](./docs/environments.md)

### Binary

Download the latest release from the [releases page](https://github.com/trancong12102/sish-gitlab-keys/releases).

```shell
./sish-gitlab-keys
```

### Docker

Use the Docker image from [Docker Hub](https://hub.docker.com/r/trancong12102/sish-gitlab-keys).

```shell
docker run -d -p 8080:8080 -e GITLAB_URL=https://gitlab.com -e GITLAB_TOKEN=your_token trancong12102/sish-gitlab-keys
```

Example docker deployment with sish, let's encrypt in [deploy](./deploy) directory.

## Development

Environment variables: [environments.md](./docs/environments.md)

```shell
go download
go run ./cmd/server
```
