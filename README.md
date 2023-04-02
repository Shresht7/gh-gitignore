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

![Usage Demonstration](docs/demo.gif)

<div align="right">

[⬆️ Back to top][top]

</div>

## ⌨️ Commands


### `create`

Create a gitignore file

#### Alias

`new`, `add`, `init`, `clone`

#### Usage

```sh
gh-gitignore create [flags]
```

#### Flags

| Flag              | Type     | Description                       | Default    |
| ----------------- | -------- | --------------------------------- | ---------- |
| `--dest, -d`      | `string` | Destination of the gitignore file | .gitignore |
| `--overwrite, -o` | `bool`   | Overwrite the gitignore file      | false      |

#### Examples

```sh
  gh gitignore create Go
  gh gitignore create Go Python
  gh gitignore create Go Python -d .gitignore
```

<div align="right">

[⬆️ Back to top][top]

</div>


### `list`

Lists all gitignore templates



#### Usage

```sh
gh-gitignore list
```



#### Examples

```sh
  gh gitignore list
```

<div align="right">

[⬆️ Back to top][top]

</div>


### `view`

View a gitignore file

#### Alias

`show`, `get`

#### Usage

```sh
gh-gitignore view
```



#### Examples

```sh
  gh gitignore view Go
  gh gitignore view Go Python
```

<div align="right">

[⬆️ Back to top][top]

</div>



---

## 📜 License

This software is licensed under the [MIT License](). See the [LICENSE](./LICENSE) file for details.





[top]: #gh-gitignore
