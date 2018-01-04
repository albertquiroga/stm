[![Build Status](https://travis-ci.org/albertquiroga/stm.svg?branch=master)](https://travis-ci.org/albertquiroga/stm)[![Codacy Badge](https://api.codacy.com/project/badge/Grade/5cdd1b3c58264f1e8e3a8973c9150d35)](https://www.codacy.com/app/albertquiroga/stm?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=albertquiroga/stm&amp;utm_campaign=Badge_Grade)

# stm
Simple Telegram message sender written in Go.

## Introduction
stm (**s**end **t**elegram **m**essage) is a binary tool to easily send Telegram messages using the [HTTP-based Telegram bot API](https://core.telegram.org/bots/api#making-requests). Using this tool you'll easily be able to send messages to Telegram users, allowing you to automate services or provide server information to your users.

## Installation
There are no packages for now, so you can either compile the source code using a simple go install (you'll need the [kingpin.v2 go library](https://github.com/alecthomas/kingpin)) or download the proper binary from the [releases page](https://github.com/albertquiroga/stm/releases) and copy it to a location in your PATH.

## Usage
stm allows you to send messages to any user or chatroom in Telegram from a bot user. In order to do that you'll need:

* To create your own bot user, from which messages will be sent.
* To identify the chat ID of the user or chatroom you want to send the messages to.
* Properly configure stm with the results of the previous steps to make it work!

### Configuring a Telegram bot
Head to the [Telegram bot documentation page](https://core.telegram.org/bots#6-botfather) and follow the "creating a new bot" steps. Creating a bot is pretty simple and you can customize it with a name and a picture for easy recognition. Once you are done you'll get a Telegram bot API key (something along the lines of 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw). That's the key that will allow you to send messages as your bot and that you'll need to configure stm.

### Obtaining your user chat ID
Everything in Telegram is a chatroom, including one-to-one conversations. That means your conversation with your mom is identified by a chat room ID and not by your mom's username, which isn't very intuitive for one-to-one chats.

You can get that identifier in several ways, but the easiest one is to use [Fredy Kardian](https://stackoverflow.com/users/6223024/fredy-kardian)'s bot (@get_id_bot). Simply open a conversation with that bot and type the /my_id command, it'll answer you with your chat ID. If you want the chat ID of a contact, send the bot the contact. And if you want the chat ID of a group chat room, add the bot to the room and type the /my_id command.

### Using the tool and flags
Now that you have all required parameters, let's configure the tool. A message in stm is defined by:

* Its source (the bot API key)
* Its destination (the chat ID of the user/group you want to send messages to)
* The actual message you want to send

The first two can either be set as CLI flags when summoning stm or be stored in a CSV file for added privacy. CLI flags have priority over the CSV file, so you can override stored settings for a quick message to a different user.

CLI flags are:

* --help: displays the usage help message.
* -b or --bot: allows you to specify a bot API key.
* -c or --chat: allows to you specify a chat ID.
* -i or --identity: path to a CSV file containing the bot API key and chat ID.

The identity file is a simple CSV file in the form of botAPIkey,chatID. If you want to always use the same file and not specify it, you can store it in your $HOME directory as ".stm" and stm will automatically use those credentials (if they are not overriden by flags).

Example file:

`110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw,1281283`

Finally, the message will be passed as an argument when summoning stm (remember to surround it in double quotes if it includes spaces).

## Credits

Credits to @alecthomas for his [Kingpin](https://github.com/alecthomas/kingpin) library.
