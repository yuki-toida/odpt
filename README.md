# Getting Started

## Dependencies

* [Yarn](https://yarnpkg.com/lang/en/docs/install/)

```sh
brew install yarn
```

* MySQL Client

``` sh
brew install mysql --client-only
```

## Initialize DB

```sh
cd /backend/tools/db
sh init.sh dev
```

## Run

* Backend

```sh
cd /backend
sh run.sh
```

* Frontend

```sh
cd /frontend
yarn dev
```
