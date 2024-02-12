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

{{ template "BackToTop" }}

## ⌨️ Commands

{{ range .Commands }}
{{- template "Command" . -}}
{{ end }}

---

## 📜 License

This software is licensed under the [MIT License](). See the [LICENSE](./LICENSE) file for details.

<!-- ----- -->
<!-- LINKS -->
<!-- ----- -->

[top]: #gh-gitignore
