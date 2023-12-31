# CamCheckBot

CamCheckBot is a Discord bot designed to check if everyone in a voice channel has opened their camera.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
  - [Configuration](#configuration)
  - [Running Locally](#running-locally)
  - [Docker Compose](#docker-compose)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

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

2. **Create a `.env` file in the root directory with the following content:**

    ```env
    # Discord Bot Token
    DISCORD_TOKEN=your_discord_bot_token

    # Redis Configuration
    REDIS_ADDR=localhost:6379
    REDIS_PASSWORD=

    # MongoDB Configuration
    MONGO_URI=mongodb://localhost:27017
    MONGO_DB=camcheckbot
    ```

### Running Locally

```bash
go run main.go
