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
