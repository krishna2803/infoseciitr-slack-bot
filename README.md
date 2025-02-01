# InfoSecIITR Slack-Bot

Custom Slack Bot for InfoSecIITR

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Commands](#commands)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Introduction

A custom Slack bot designed for InfoSecIITR. It provides various functionalities to interact with Slack channels, manage keys, and more.

## Features

- Database in `postgresql`
- Logging with `slog`.
- Uses `slack-io/slacker` API in socket mode.
- Fully Dockerized.
- Key ownership.
- **WIP**: User management
- **WIP**: CTF Event Fetching

## Deployment

### Prerequisites

- Docker
- Docker Compose

### Steps
1. Clone the repository:

    ```sh
    git clone https://github.com/krishna2803/infoseciitr-slack-bot.git
    cd infoseciitr-slack-bot
    ```

2. Create `.env`based on `.env.sample`:

    ```sh
    cp .env.sample .env
    ```
3. Create and add a slack bot to your slack community using [manifest.yaml](manifest.yaml)

4. Fill in the required environment variables in `.env` according to [configuration](#configuration).

5. Build and run the Docker containers:

    ```sh
    docker compose up
    ```

## Configuration

- `SLACK_BOT_TOKEN`: Slack bot token.
- `SLACK_APP_TOKEN`: Slack app token.
- `DB_USER`: Database username.
- `DB_PASS`: Database password.
- `DB_NAME`: Database name.
- `DB_HOST`: Database host (default is `localhost`).
- `ENV`: Environment (`dev` or `prod`).


## Commands

### Ping

- **Description**: Pings the bot.
- **Command**: *bot ping*

### Who Has The Keys

- **Description**: Fetches the current key owners.
- **Command**: *bot who has the keys*

### Transfer Keys

- **Description**: Sets `username` or `you` as the owner of the key you provide.
- **Command**: *bot {username/i} has/have the {name} keys*
- **Examples**:
  - bot segfault has the master keys
  - bot i have the master keys

## Development

### Prerequisites

- Go
- Make

### Running
```sh
make vendor
make dev # uses `air` for hot reloading
```


### Linting and Formatting

To lint and format the code, run:

```sh
make lint
make format
```