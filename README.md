# Testing Go os.Create

During my work on kaniko, I got to the issue, that the root user can't run os.Create on files, that were already created and not owned by root.

To reproduce, use the following:

1. Inside docker:
```console
$ docker build . -t go-os-create
$ docker run --user root go-os-create
```

2. On Linux:
```console
$ sh ./run-twice.sh
```

This only occured on linux, not on darwin.
