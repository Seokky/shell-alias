# shell-alias

It is my handy helper to work with shell aliases.
I've created this project to get some practice with Go, so it is not some serious program, but still can be usefull for you.

## Getting help info

`shell-alias help`

Print a little help information to standard out

## Print list of aliases you have

`shell-alias list`

Print a list of aliases you have in name=command format.

## Add new alias

`shell-alias add --name="" --command=""`

Add new alias with name and command provided.

## Delete existing alias

`shell-alias delete --name=""`

Delete existing alias from.

## About .shellrc file auto detection

Program will auto detect .shellrc file in your home directory: first it will try to find .zshrc (for Mac users), then it goes to find .bashrc.

If one of them here, working with it.

If no one here, program will ask you for name of .shellrc file you want to create, and then create it for you.

**Also you can force path of your .shellrc file by providing it via --path parameter to each command (list, add, delete).**

## How to install

#### 1. Install Go on your machine

#### 2. Clone repository

- ssh: `git clone git@github.com:Seokky/shell-alias.git`
- https: `git clone https://github.com/Seokky/shell-alias.git`

#### 3. Go to repo directory

#### 4. Build and install program

```
go build
go install
```

#### 5. Add do PATH

Discover the Go install path, where the go installed the package

`go list -f '{{.Target}}'`

It will print something like _/Users/\<username\>/go/bin/shell-alias_.

**Copy this output.**

Then open your shell rc file. Probably is:

- MacOs: `nano ~/.zshrc`
- Linux: `nano ~/.bashrc`

Add this line to the end of file:

```
PATH="$PATH:*/Users/\<username\>/go/bin/shell-alias*"
```

Press ctrl+X, then Y, and Enter to save changes.

Then run in shell:

- MacOs: `source ~/.zshrc`
- Linux: `source ~/.bashrc`

Now you can use `shell-alias`!
