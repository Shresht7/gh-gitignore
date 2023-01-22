# `gh-gitignore`

A GitHub CLI extension to generate `.gitignore` files.

The `.gitignore` templates are obtained from the GitHub API (https://docs.github.com/en/rest/gitignore).

---

## 📦 Installation

### Requirements

- [GitHub CLI](https://cli.github.com/)

### Install

```sh
gh extension install Shresht7/gh-gitignore
```

---

## 💻 Usage

Invoke the cli extension like so:

```sh
gh gitignore <command>
```

To create a gitignore file:

```sh
gh gitignore create Node
```

## ⌨️ Commands

### `help`

#### Usage

Use the `help` command to explore the cli.

```sh
gh gitignore help
gh gitignore help <command>
gh gitignore <command> --help
```

### `create`

#### Aliases

`new`, `init`, `clone`

#### Usage

To generate a `.gitignore` file use:

```sh
gh gitignore create <language>
```

#### Example

```sh
gh gitignore create Go
gh gitignore create Node
gh gitignore create Rust
```

#### Flags

| Flag         | Description                          |
| ------------ | ------------------------------------ |
| `--dest, -d` | Destination of the `.gitignore` file |

### `list`

#### Usage

To get a list of `.gitignore` templates use:

```sh
gh gitignore list
```

### `view`

#### Aliases

`show`, `get`

#### Usage

To view a particular `.gitignore` template use:

```sh
gh gitignore view <language>
```

#### Example

```sh
gh gitignore view Go
gh gitignore view Node
gh gitignore view Rust
```

---

## 📄 License

[MIT License](./LICENSE)
