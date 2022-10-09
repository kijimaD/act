[![⚗️Check](https://github.com/kijimaD/act/actions/workflows/check.yml/badge.svg)](https://github.com/kijimaD/act/actions/workflows/check.yml)

# act

<img src="https://user-images.githubusercontent.com/11595790/193450591-6b681517-3b5a-4dd4-ac04-5dce9b209882.png" width="40%" align=right>

act is curating your github activity tool.

working example: [kijimaD/central](https://github.com/kijimaD/central)

# Install

```sh
$ go install github.com/kijimaD/act@main
```

# How to use

set `.act.yml`

```yml
userId: kijimaD
outType: file
outPath: ./README.md
commit: true
push: false
```

and prepare GitHub API token, run

```shell
$ GH_TOKEN=<your_token> act
```

# Docker run

```shell
docker run --rm \
           -e GH_TOKEN=<API Token> \
           -v "${PWD}":/workdir \
           ghcr.io/kijimad/act:latest
```
