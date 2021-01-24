# DeletedBot Backend

This repo contains the code for DeletedBot's backend, including:

- The API protocol to be used by the frontend
- The endpoint for Telegram's webhook

It uses Go Gin as web server and Redis as database. 
This project is released under GNU GPL v3.0, read the LICENSE file to know more. Contributions are always welcome!

> The backend is split apart from the frontend for everyone who wants to follow the **Jamstack** way.
> If you don't care about that, and you want an all-in-one solution which is **easier to deploy,** check out the [allinone](https://github.com/deleted-bot/allinone) repo.



## How to self-host

Notice: The backend, by itself, is pretty useless. So, once you're done self hosting the backend, make sure to check out how you can deploy the frontend [here](https://github.com/deleted-bot/frontend).

### The zoomers' way

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

### The good'ol way

1. Download a pre-built binary from the release page or clone the repo and run `go build . -o deletedbot-backend` in the same directory

2. Get yourself a Redis server. If you're on a VPS or dedicated server, check out [this](https://redis.io/download) link; if you're on Heroku or similar services, take a look [here](https://redislabs.com/try-free/) to get a free managed Redis database, or get an add-on for the specified service.

3. Set the following environment variables:

   | Key            | Example                  | Explaination                                                 |
   | -------------- | ------------------------ | ------------------------------------------------------------ |
   | HOST           | deleted.bot.nu           | The host to which your website will be served. This link will appear in the messages sent by the bot as well. |
   | PORT           | 80                       | The port the web server will listen on.                      |
   | REDIS_ADDR     | localhost                | The hostname of your Redis database                          |
   | REDIS_PASSWORD | ExtremelySafePassword    | The password of your Redis database. Leave empty if no password is set. |

4. Run your instance: if you're on a server, use `sudo nohup ./deletedbot-backend`, if you're using anything else use your common deploying technique.

5. That's it. Here's a cat to fill some space:

![Maurizio the cat](https://i.imgur.com/58ZHjlC.gif)

