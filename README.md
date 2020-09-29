# gtodo : A simple todo list application

gtodo will help you get more done in less time.
It's designed to be as simple as possible to help you
accomplish your goals.

## Build Instructions

Clone repository

`git clone https://github.com/mrinjamul/gtodo`

`cd gtodo`

Install modules

`go mod tidy`

Build

`go build`

Run ,

`./gtodo`

Install

`go install gtodo`

If you want, you can install by copying into any \$PATH

## INSTALLING

#### Installing from go

First, use go get to install the latest version. This command will install the `gtodo` and its dependencies:

`go get -u github.com/mrinjamul/gtodo`

#### Installing from Binaries

Download for your platform

[Download](https://github.com/mrinjamul/gtodo/releases)

For Linux,

```sh
unzip gtodo-linux-[whatever].zip
chmod +x gtodo
sudo mv gtodo /usr/bin
```

or you can put the executable file into your env variables `$PATH`

For Termux,

You need to have `zip unzip wget`. To install simply type `pkg install zip unzip wget`

```sh
cd ~
wget https://github.com/mrinjamul/gtodo/releases/download/v0.2.1/gtodo-linux-arm-v0.2.1.zip
unzip gtodo-linux-arm-v0.2.1.zip
chmod +x gtodo
mv gtodo ../usr/bin
```

[Note: download link will be updated!]

## Usage

```sh
gtodo add -p=1 make dinner
```

The following output is automatically generated by gtodo. Nothing beyond the
command and flag definitions are needed.

    gtodo will help you get more done in less time.
    It's designed to be as simple as possible to help you
    accomplish your goals.

    Usage:
    gtodo [command]

    Available Commands:
    add         Add a new todo
    clear       Clear all todos
    done        Mark Item as Done
    help        Help about any command
    list        list all todos
    modify      edit a todo
    remove      Remove a todo
    undone      Mark Item as UnDone
    version     Print the version number of gtodo

    Flags:
        --config string     config file (default is $HOME/.gtodo.yaml)
        --datafile string   data file to store todos (default "$HOME/.gtodo.json")
    -h, --help              help for gtodo

    Use "gtodo [command] --help" for more information about a command.

## Links

[Repo Website]()

## License

- Apache-2.0
