# ChatGPT-bot

> Interact with ChatGPT

Go CLI to fuels a Telegram bot that lets you interact with [ChatGPT](https://openai.com/blog/chatgpt/), a large language model trained by OpenAI.

## Installation

Before the installation. You should export `https://chat.openai.com` on your browser. You should export cookies as JSON for bypassing Cloudflare. You can use `https://chrome.google.com/webstore/detail/%E3%82%AF%E3%83%83%E3%82%AD%E3%83%BCjson%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E5%87%BA%E5%8A%9B-for-puppet/nmckokihipjgplolmcmjakknndddifde` extension or also you can copy-paste from Chrome DevTools like below structure

```json

[
  {
    "name": "cookie-name",
    "value": "value",
    "domain": ".openai.com",
    "path": "/",
    "expires": 1701377088, // timestamp format
    "httpOnly": false,
    "secure": false,
    "sameSite": "Lax"
  },
  ...
]

```

Save as json file

Clone the project

```bash

git clone https://github.com/peacecwz/chatgpt-telegram

```

Move cookies json file into repository


After you clone the project, open the `env.example` file with a text editor and fill in your credentials. 
- `TELEGRAM_TOKEN`: Your Telegram Bot token
  - Follow [this guide](https://core.telegram.org/bots/tutorial#obtain-your-bot-token) to create a bot and get the token.
- `TELEGRAM_ID` (Optional): Your Telegram User ID
  - If you set this, only you will be able to interact with the bot.
  - To get your ID, message `@userinfobot` on Telegram.
  - Multiple IDs can be provided, separated by commas.
- `EDIT_WAIT_SECONDS` (Optional): Amount of seconds to wait between edits
  - This is set to `1` by default, but you can increase if you start getting a lot of `Too Many Requests` errors.
- `COOKIE_FILE` (Required): Give path of exported cookie file
- Save the file, and rename it to `.env`.
> **Note** Make sure you rename the file to _exactly_ `.env`! The program won't work otherwise.

Finally, open the terminal in your computer (if you're on windows, look for `PowerShell`), navigate to the path you extracted the above file (you can use `cd dirname` to navigate to a directory, ask ChatGPT if you need more assistance 😉) and run `./chatgpt-telegram`.

### Running with Docker

If you're trying to run this on a server with an existing Docker setup, you might want to use our Docker image instead.

```sh
docker pull ghcr.io/m1guelpf/chatgpt-telegram
```

Here's how you'd set things up with `docker-compose`:

```yaml
services:
  chatgpt-telegram:
    image: ghcr.io/m1guelpf/chatgpt-telegram
    container_name: chatgpt-telegram
    volumes:
      # your "cookies.json" will move into container
      - cookies.json/:/root/cookies.json
    environment:
      - TELEGRAM_ID=
      - TELEGRAM_TOKEN=
      - COOKIE_FILE=cookies.json
```

> **Note** The docker setup is optimized for the Browserless authentication mechanism, described below. Make sure you update the `.config/chatgpt.json` file in this repo with your session token before running.

## Authentication

If you include cookies.json file into .env. It'll be running automatically

By default, the program will launch a browser for you to sign into your account, and close it once you're signed in. If this setup doesn't work for you (there are issues with the browser starting, you want to run this in a computer with no screen, etc.), you can manually extract your session from your browser instead.

> **Note** If you have already run the program, the file should exist but be empty. If it doesn't exist yet, you can either run the program or manually create it.

## License

This repository is licensed under the [MIT License](LICENSE).
