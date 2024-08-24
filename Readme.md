# ssh_box

ssh_box is a lightweight command-line tool for managing and connecting to your servers via SSH. With ssh_box, you can easily add, list, remove, and SSH into your registered servers using simple commands.

## Features

- Add Servers: Register servers with SSH keys, IP addresses, and tags.
- List Servers: View all registered servers, with detailed information if needed.
- Remove Servers: Unregister servers by their tags.
- SSH into Servers: Quickly SSH into your registered servers using a simple command.

## Installation

Build and Install using Makefile

1. Clone the repository:

```bash
git clone
cd ssh_box
```

2. Build the application:

```bash
make build
```

3. Install the binary:

```bash
sudo make install
```

This will place the ssh_box binary in /usr/local/bin.

### Running without Installation

You can also run the application without installing it globally:

```bash
make run
```

## Usage

### Add a Server

```bash
ssh_box add --key /path/to/key --server <IP_ADDRESS> --tag <SERVER_TAG> [--user <USERNAME>]
```

- --key: Path to the SSH key (required).
- --server: IP address of the server (required).
- --tag: A tag to identify the server (required).
- --user: Username for SSH (optional, defaults to root).

## List Servers

```bash
ssh_box ls
```

To list servers with detailed information:

```bash
ssh_box ls -a
```

## Remove a Server

```bash
ssh_box rm --tag <SERVER_TAG>
```

- --tag: The tag of the server to remove (required).

### SSH into a Server

````bash
 ssh_box exec --tag <SERVER_TAG>```
````

- --tag: The tag of the server to SSH into (required).

## Uninstallation

To remove the installed ssh_box binary:

```bash
sudo make uninstall
```

## Cleaning Up

To remove build artifacts:

```bash
make clean
```

## Configuration

ssh_box automatically sets up a configuration directory in your home directory (~/.ssh_box) and stores server information in a servers.json file. This file is managed automatically by the tool.
