## INSTALLING

#### Installing from go

First, use go get to install the latest version. This command will install the `gtodo` and its dependencies:

`go get -u github.com/mrinjamul/gtodo`

#### Installing from Binaries

Download for your platform

[Download](https://github.com/mrinjamul/gtodo/releases)

For Linux,

```sh
wget https://github.com/mrinjamul/gtodo/releases/download/v0.2.2/gtodo-linux-amd64-v0.2.2.zip
unzip gtodo-linux-amd64-v0.2.2.zip
chmod +x gtodo
sudo mv gtodo /usr/bin
```

or you can put the executable file into your env variables `$PATH`

For Termux,

You need to have `zip unzip wget`. To install simply type `pkg install zip unzip wget`

```sh
cd ~
wget https://github.com/mrinjamul/gtodo/releases/download/v0.2.2/gtodo-linux-arm-v0.2.2.zip
unzip gtodo-linux-arm-v0.2.2.zip
chmod +x gtodo
mv gtodo ../usr/bin
```

[Note: download link will be updated!]

## Usage

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

### add

    Add will create a new todo item to the list

    Usage:
    gtodo add [flags]

    Aliases:
    add, a

    Flags:
    -h, --help           help for add
    -p, --priority int   Priority:1,2,3 (default 2)

    Global Flags:
        --config string     config file (default is $HOME/.gtodo.yaml)
        --datafile string   data file to store todos (default "$HOME/.gtodo.json")

### remove

    Remove will remove a todo item from the list by Label(index)

    Usage:
    gtodo remove [id1] [id2] ...

    Aliases:
    remove, rm

    Flags:
    -d, --done   remove only done tasks
    -h, --help   help for remove

    Global Flags:
        --config string     config file (default is $HOME/.gtodo.yaml)
        --datafile string   data file to store todos (default "$HOME/.gtodo.json")

### modify

    edit a todo

    Usage:
    gtodo modify [flags]

    Aliases:
    modify, mod, ed

    Flags:
    -h, --help   help for modify

    Global Flags:
        --config string     config file (default is $HOME/.gtodo.yaml)
        --datafile string   data file to store todos (default "$HOME/.gtodo.json")

### clear

    Clear all todos

    Usage:
    gtodo clear [flags]

    Flags:
    -f, --force   Force remove all todos
    -h, --help    help for clear

    Global Flags:
        --config string     config file (default is $HOME/.gtodo.yaml)
        --datafile string   data file to store todos (default "$HOME/.gtodo.json")

### done, do

Mark Item as Done

    Usage:
    gtodo done [flags]

    Aliases:
    done, do

    Flags:
    -h, --help   help for done

    Global Flags:
        --config string     config file (default is $HOME/.gtodo.yaml)
        --datafile string   data file to store todos (default "$HOME/.gtodo.json")

### undo

    Mark Item as UnDone

    Usage:
    gtodo undone [flags]

    Aliases:
    undone, undo

    Flags:
    -h, --help   help for undone

    Global Flags:
        --config string     config file (default is $HOME/.gtodo.yaml)
        --datafile string   data file to store todos (default "$HOME/.gtodo.json")

### list

    list all todos

    Usage:
    gtodo list [flags]

    Aliases:
    list, ls

    Flags:
    -a, --all    Show all TOdos
    -d, --done   Show 'Done' Todos
    -h, --help   help for list

    Global Flags:
        --config string     config file (default is $HOME/.gtodo.yaml)
        --datafile string   data file to store todos (default "$HOME/.gtodo.json")

### version

    Prints current version of gtodo
