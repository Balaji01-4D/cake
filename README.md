# cake

**cake** is a lightweight TUI for Makefiles.

It turns your `Makefile` into a fast, interactive workflow so you can discover targets, inspect what they run, and execute them without leaving the terminal.

Built with **Bubble Tea v2**, **Bubbles**, and **Lip Gloss**, `cake` is designed to feel snappy, modern, and keyboard-first.

## Why cake?

Working with Makefiles should be **easy as cake**:

- No memorizing obscure targets
- No repeatedly opening the Makefile to inspect recipes
- No guessing what `make <target>` will actually run

## Core experience

`cake` is focused on a clean, practical loop:

- Fuzzy-searchable target list
- Real-time hover preview of target recipes
- One-key execution of the selected target
- Split-pane layout (targets on the left, preview on the right)

## Development status

This project is currently under active development.

You can track planned milestones and scope in [plans.md](./plans.md).

## Quick start

```bash
make build
./bin/cake
```