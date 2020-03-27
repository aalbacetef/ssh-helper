# SSH Helper

A tool to help you manage your ssh hosts, making it easy to both quickly list your current hosts, as well as quickly generate new hosts.


## Installation

You can use `go get` to install the tool:

`$ go get github.com/aalbacetef/ssh-helper/cmd/ssh-helper`


Additionally, you can use the pre-built binary (linux):

[ssh-helper v0.1.1](https://github.com/aalbacetef/ssh-helper/releases/download/v0.1.1/ssh-helper) (Linux 64-bit)

To download:

`$ wget https://github.com/aalbacetef/ssh-helper/releases/download/v0.1.1/ssh-helper && chmod +x ssh-helper`

## Examples

#### Add a new host
To add a host with an existing key pair:

```
$ ssh-helper add \
  --name my-test-host \
  --hostname my.test.host \
  --user testuser \
  --identityfile /path/to/private/key
```


To add a host with the key pair being generated:
```
$ ssh-helper add \
  --name my-test-host \
  --hostname my.test.host \
  --user testuser \
  --newkey 
```

#### Remove an existing host

Suppose you already have a host with name `existing-host`, then you could:

```
ssh-helper remove --host existing-host
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

