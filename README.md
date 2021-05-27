# HostCLI POC

Extend your command-line interface with any binary!

## Build

1. Run `make`
2. Run `./host -h` to list available commands

## POC

The `git-helper copy` command maps to `git clone` under the hood using a custom plugin configuration defined in [./plugins/git.yaml](./plugins/git.yaml).

- This shows that it is possible to control another binary
- Gives native completion support and a strong interface

Not yet covered:

- Handling of the event where the plugin binary does not exist - should we download it for the user?
- Further investigations needed
