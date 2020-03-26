# SSH Helper

A tool to help you manage your ssh hosts, making it easily to both quickly list your current hosts, as well as quickly generate new hosts.

## Usage

```
SSH-Helper
-----------

Usage: ssh-helper COMMAND [OPTIONS]


A tool to manage your ssh configs. By default 
uses ~/.ssh/ssh-helper to manage all configs.

Commands:
  add       Add an ssh config, can generate a key automatically
  backup    Backs up the current ~/.ssh directory
  config    Print the current configs of ssh-helper
  list      List all available hosts
  remove    Remove a host from the config. This operation will not delete the key unless asked to.

```



