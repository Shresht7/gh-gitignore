# `create`

Create a gitignore file

## Aliases

`new`, `add`, `init`, `clone`

## Description

Create a gitignore file for your project

## Usage

```
gh-gitignore create [flags]
```

## Flags

| Flag              | Type     | Description                       | Default    |
| ----------------- | -------- | --------------------------------- | ---------- |
| `--dest, -d`      | `string` | Destination of the gitignore file | .gitignore |
| `--overwrite, -o` | `bool`   | Overwrite the gitignore file      | false      |

## Examples

```
  gh gitignore create Go
  gh gitignore create Go Python
  gh gitignore create Go Python -d .gitignore
```

## See Also

- [gh-gitignore](./gh-gitignore.md)
- [list](./list.md)
- [view](./view.md)

