# act

act is curating your github activity tool.

# Install

```sh
$ go install github.com/kijimaD/gclone@main
```

# How to use

set `.act.yml`

```yaml
outType: file
outFile: ./README.md
```

and prepare GitHub API token, run

```shell
$ GH_TOKEN=your_token... act
```

# Docker run

```shell
docker run --rm
           -it
           -v "${PWD}":/workdir \
           ghcr.io/kijimad/act:latest
```
