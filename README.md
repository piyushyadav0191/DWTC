# Daily Work Tracker

A command-line daily work tracker tool built using Golang. This tool helps you manage your daily tasks efficiently by allowing you to add, delete, list, and mark tasks as complete.

## Prerequisites

- Golang installed on your system.
- The `make` utility.

## Installation

1. Fork and clone the repository:

```bash
git clone https://github.com/piyushyadav0191/Daily-Work-Tracker-CLIgit
cd daily-work-tracker
```

2. Build the project:

```bash
make build
```

## Usage

- Add a new task:

```bash
./todo <write_your_work>
```

- Delete a task:

```bash
./todo -del=<ID_NUMBER>
```

- List all tasks:

```bash
./todo -list
```

- Mark a task as complete:

```bash
./todo -complete=<ID_NUMBER>
```

## Example

```bash
./todo "Prepare presentation for meeting"
./todo "Review code changes"
./todo -list
./todo -complete=2
./todo -list
```

## Contributing

Contributions are welcome! Feel free to open a pull request.
