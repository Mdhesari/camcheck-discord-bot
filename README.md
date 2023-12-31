# CamCheckBot

CamCheckBot is a Discord bot designed to check if everyone in a voice channel has opened their camera.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
  - [Configuration](#configuration)
  - [Docker Compose](#docker-compose)
- [Usage](#usage)
- [Contributing](#contributing)

## Features

- Real-time camera presence checks in voice channels.
- Seamless integration with Discord servers.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- Docker
- Docker Compose

## Setup

### Configuration

1. **Clone this repository:**

   ```bash
   git clone https://github.com/your-username/CamCheckBot.git
   cd CamCheckBot
   ```

2. **Create a `config.yml` file in the root directory from config-sample.yml:**

   ```bash
   cp config-sample.yml config.yml
   ```

### Docker Compose

**Build and run docker compose**

    ```bash
    docker compose up -d
    ```

This will start CamCheckBot with Redis and MongoDB containers.

### Usage

1. Invite the bot to your Discord server.
2. Use the /ping command to verify bot loading.

### Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.