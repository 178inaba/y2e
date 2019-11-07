# y2e

[![CircleCI](https://circleci.com/gh/178inaba/y2e.svg?style=svg)](https://circleci.com/gh/178inaba/y2e)
[![Go Report Card](https://goreportcard.com/badge/github.com/178inaba/y2e)](https://goreportcard.com/report/github.com/178inaba/y2e)

y2e sets environment variables from a YAML file.

## Usage

Create a YAML file with the environment variable name as the key.

```yaml
FOO: foo
BAR: 123
BAZ: true
```

Pass YAML file to `y2e` and execute it as a subcommand of `eval` command.  
After execution, the value of the YAML file is set as an environment variable.

```console
$ eval "$(y2e your_env_var.yaml)"
$ env | grep -e FOO -e BAR -e BAZ
FOO=foo
BAR=123
BAZ=true
```

### Options

- `-verbose`
  - Output detail log to the stderr.

## Install

```console
$ go get -u github.com/178inaba/y2e
```

## License

[MIT](LICENSE)

## Author

[178inaba](https://github.com/178inaba)
