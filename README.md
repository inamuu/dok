

# dok

`dok` is a CLI tool designed to simplify Docker command usage via interactive selection powered by `peco`.

## Features

- `dok ps` — Interactively select whether to view running or all containers.
- `dok exec` — Select a container and execute a command (bash, sh, etc.) using interactive prompts.
- `dok run` — Launch containers with interactive selection of image, port, and optional commands.
- `dok rm` — Remove containers by selecting from a list.
- `dok rmi` — Remove images via selection.
- `dok stop` — Stop containers via selection.
- `dok start` — Start containers from selected images.

## Installation

```
go install github.com/inamuu/dok@latest
```

## Configuration

Upon `dok init`, the following config file is created at `~/.dok.config`.  
You can chose a command when you add commands to this file.

```ini
[Commands]
bash
sh
ls -la
ps
```

You can modify this file to add custom commands used in `dok exec` or `dok run`.

## Dependencies

- [Go](https://golang.org/doc/install)
- [peco](https://github.com/peco/peco)

## License

MIT