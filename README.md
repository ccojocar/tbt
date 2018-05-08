# tbt

This is a tool to transfer a binary tree between two remote nodes.

## Installation

You can build the tool by running the command:

```bash
make
```

## Usage

``` bash
$> tbt --help
A tool to transfer a binary tree across two remote nodes

Usage:
  tbt [command]

Available Commands:
  help        Help about any command
  receive     Receive a binary tree from a remote node
  send        Send a binary tree over to a remote node

Flags:
  -h, --help   help for tbt

Use "tbt [command] --help" for more information about a command.

```

The receiver node can be started as follows:

```
$> tbt receive --address localhost:3001
```

You can send now a binary tree from a sender node:

```
$> tbt send --address localhost:3001 1,3,#,#,4,#,#
```

Note that the binary tree is provided as a pre-ordered sequence. The `#` symbol indicates a null termination.

As soon as the binary tree is transmitted, the receiver node should reconstruct again the tree and print it out into the console also as a pre-order sequence.
