# Yet Another Personal Assistant (YAPA)

<img src="0.png" alt="Drawing" width="200px" />

A small utility tool to speed up and ease your day to day work.

## Install



## Setup

After setting up yapa, next thing is to setup your yapa profile. This can be done by running,

```bash
$ yapa setup
```

It will create a default `.yapa` directory, place a few files in it and then ask for for username.

Now you are good to go. Run `yapa help` to explore what yapa can do for you.

```bash
NAME:
   yapa - Yet Another Personal Assistant

USAGE:
   yapa [global options] command [command options] [arguments...]

VERSION:
   v0.0.1

COMMANDS:
   setup                        Setup a new yapa profile
   clean                        Clear all yapa settings
   count                        Count number of files/folders in directory
   key                          Get current user\'s public key
   ping [HOSTNAME]              Check if host is online. Defaults to 8.8.8.8.
   list                         List all servers listed in config.json
   bye                          Shutdown system
   uptime [USER] [IP]           Display uptime of a server
   cool
   hackernews, hn               Display Hacker News
   scan                         Scan a hostname
   toss                         Flips a coin
   dice                         Roll a dice
   help, h                      Display help
   todo                         Show list of todo\'s
      list, l                   Show list of todo\'s
         completed, c           Show completed todo\'s
         incompleted, in        Show incomplete todo\'s
      remove, r [id]            Remove a todo from list
      add, a                    Add a new todo
      complete, c [id]          Mark a todo as completed
   all-users, allusr            List all users
   investigate, inv [username]  Get detail of the user specified

GLOBAL OPTIONS:
   -h, --help     Display help
```

## Commands

### ping

Pings the google dns to check if internet connection is up.

### scan

Scans a host for open ports.

### speedtest

Test your internet speed.

### key

Prints your default(id_rsa.pub) public key.

##### TODO

- Display all keys in `.ssh/` and option to select from them.

### count

Counts the number of files/folders in current directory.

##### TODO

- List number of hidden files.

### hackernews

**alias**: `hn`

Displays top 10 hacknews articles.

### ip

Get your public ip.

### dice

Roll a dice.

### toss

Flip a coin.

### todo

Your local todo manager.

### all-users

**alias**: `allusr`

List all users on your system.

### investigate

**alias**: `inv`

Get all info for any user.

### forever

Start a never ending mode. No need of prepending yapa to all commands.

### Removing your yapa profile

Running `yapa clean` will remove all your yapa configs.

## Contribution

Feature requests and bug fixes are welcome. :smile:

## License

MIT
