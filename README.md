# Playcode

A code playground support a variety of languages.


## Requirements

* PostgreSQL
* Go 1.19+
* Node 16.17+, and Yarn 1.x
* [Hakoniwa Code Runner](https://github.com/souk4711/hakoniwa-code-runner)


## Development

#### 1. Check out the repository

```console
$ git clone git@github.com:souk4711/playcode.git
```

#### 2. Create and setup the database

```console
$ bin/task docker:up
$ bin/setup
```

#### 3. Start the Go server & Vite devserver

```console
$ bin/dev
```

#### 4. Start Hakoniwa Code Runner (optional)

```console
$ cd ../hakoniwa-code-runner
$ make start-server
```


## License

Licensed under the [MIT License](./LICENSE).
