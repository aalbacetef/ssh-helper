# SSH Helper

A tool to help you manage your ssh hosts, making it easily to both quickly list your current hosts, as well as quickly generate new hosts.

## Examples

#### Add a new host
To add a host with an existing key pair:

```
ssh-helper add \
  --name my-test-host \
  --hostname my.test.host \
  --user testuser \
  --identityfile /path/to/private/key
```


To add a host with the key pair being generated:
```
ssh-helper add \
  --name my-test-host \
  --hostname my.test.host \
  --user testuser \
  --newkey 
```

## Usage

```
SSH-Helper
-----------

Usage: ssh-helper COMMAND [OPTIONS]

A tool to manage your ssh configs. By default uses ~/.ssh/ssh-helper/ to manage all configs.

Commands:
  add       Add an ssh config, can generate a key automatically.
  backup    Backs up the current ~/.ssh directory.
  config    Print the current configs of ssh-helper.
  list      List all available hosts. Supports outputting in JSON format.
  remove    Remove a host from the config. This operation will not delete the key unless asked to.

```


## Installation

You can use `go get` to install the tool:

`go get github.com/aalbacetef/ssh-helper/cmd/ssh-helper`


