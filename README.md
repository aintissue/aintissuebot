# Free Telegram Chat Bot (Without A Database)

Aint Chat Bot is a free Telegram user and customer support bot that helps you organize user support for your startup, SaaS, Open Source project or a company in a convenient way. It's written in Go (Golang).

## Install

To install the bot as a self-hosted solution, copy the config file, edit it and set it up and then run docker compose command.

```
cp data/config.sample.yaml data/config.yaml
sudo docker compose up -d
```