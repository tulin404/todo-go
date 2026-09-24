# todo

A simple and lightweight task manager for the terminal, written in Go.

`todo` provides a CLI for managing tasks and a background daemon responsible for scheduling and handling upcoming tasks.

## Features

* Task management directly from the terminal
* Task due dates
* Task completion tracking
* Background daemon
* Unix domain socket IPC
* Persistent task storage
* Task notifications
* Lightweight and dependency-free runtime

## Installation

### Build from source

Requirements:

* Go 1.22+
* Linux
* systemd (for running the daemon as a service)

Clone the repository:

```bash
git clone https://github.com/tulin404/todo-go.git
cd todo-go
```

Build the binary:

```bash
go build -o todo ./cmd/todo
```

Install it somewhere in your `PATH`:

```bash
# Example
sudo mv todo /usr/local/bin/
```

Add the daemon using systemd:

```bash
sudo mv deploy/systemd/todo.service /etc/systemd/system/todo.service
```

Activate and enable the service:

```bash
# Reload systemd
sudo systemctl daemon-reload

# Enable and start the service
sudo systemctl enable --now todo.service
```

Verify if the service is running:

```bash
# Must be active and running
systemctl status todo.service
```

Verify the installation:

```bash
todo --help
```

## Usage

### Add a task

```bash
todo add "Study Go"
```

Tasks can also have a due date:

```bash
todo add "Study Go" --due "2026-09-23 20:00"
```
> `due` accepts any valid [Go duration string](https://pkg.go.dev/time#ParseDuration)

### Help

```bash
todo --help
```

You can also get help for a specific command:

```bash
todo add --help
```

## How it works

The CLI can operate directly on local task data or communicate with the background daemon through Unix domain sockets. The daemon handles scheduling and other long-running operations.

## Development

Clone the repository and enter the project directory:

```bash
git clone https://github.com/tulin404/todo-go.git
cd todo-go
```

Run without building:

```bash
go run ./cmd/todo --help
```

Format the project:

```bash
go fmt ./...
```

## Project/Packages Structure

```text
.
├── cmd/        # CLI commands
├── cli/        # Cobra entry points
├── daemon/     # Daemon logic
├── helpers/    # Shared helper functions  
├── ipc/        # CLI ↔ daemon communication
├── storage/    # Task persistence
├── timeutil/   # Time-related utilities
└── task/       # Task domain logic
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) — CLI framework
- [beeep](https://github.com/gen2brain/beeep) — Desktop notifications

## Contributing

Contributions, issues and suggestions are welcome.

If you find a bug or have an idea for improving the project, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See [`LICENSE`](LICENSE) for details.
