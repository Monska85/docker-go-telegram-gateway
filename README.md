# Telegram Gateway

Telegram Gateway is a Go HTTP server which exchanges Telegram messages over the HTTP protocol.

The idea is to provide a simple way to send and receive messages from/to Telegram without the need to use the Telegram API, but instead using a simple HTTP request. Of course, you need to protect the access to this server because it has access as your Telegram bot: whoever can post messages to this server, can send messages as your bot. At the same time, this makes very simple to integrate Telegram notifications in your applications, whatever the language they are written in.
