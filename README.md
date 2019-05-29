# Translator based on a User Story written in GO

[![Go Report Card](https://goreportcard.com/badge/github.com/adrianbrad/translator)](https://goreportcard.com/report/github.com/adrianbrad/translator)
[![codecov](https://codecov.io/gh/adrianbrad/translator/branch/master/graph/badge.svg)](https://codecov.io/gh/adrianbrad/translator)

- [Translator based on a User Story written in GO](#translator-based-on-a-user-story-written-in-go)
  - [User Story](#user-story)
  - [Solutions](#solutions)
    - [Description](#description)
    - [Known Issues and Improvements possible](#known-issues-and-improvements-possible)
    - [CLI](#cli)
      - [Memory storage](#memory-storage)
        - [Usage](#usage)
        - [Run User Story Tests](#run-user-story-tests)
      - [Database storage](#database-storage)
        - [Prerequisites](#prerequisites)
        - [Usage](#usage-1)
        - [Run User Story Tests](#run-user-story-tests-1)
    - [Web API](#web-api)
      - [Exposed endpoints](#exposed-endpoints)
        - [/{langFrom}/{textFrom}/{langTo}](#langfromtextfromlangto)
      - [Memory storage](#memory-storage-1)
        - [Usage](#usage-2)
        - [Run User Story Tests](#run-user-story-tests-2)
      - [Database storage](#database-storage-1)
        - [Prerequisites](#prerequisites-1)
        - [Usage](#usage-3)
        - [Run User Story Tests](#run-user-story-tests-3)
  - [Running tests](#running-tests)

## User Story

**As a user I want to be able to store translations for multiple language-combinations.**

*STORE-REQUEST{DE(Hund)a->EN(dog)}->RESPONSE{OK}*

**As a user I want to be able to get a previous stored translation**

*STORE-REQUEST{DE(Hund)->EN(dog)}->RESPONSE{OK}*

*GET-REQUEST{DE(Hund)->EN} -> RESPONSE{dog}*

**As a user I want to be able to get a previous stored translation in the opposite order**

*STORE-REQUEST{DE(Hund)->EN(dog)} -> RESPONSE{OK}*

*GET-REQUEST{EN(dog)->DE} -> RESPONSE{Hund}*

**As a user I want to get a guess of a translation using other translations if possible**

*STORE-REQUEST{DE(Katze)->ES(gato)} -> RESPONSE{OK}*

*STORE-REQUEST{ES(gato)->EN(cat)} -> RESPONSE{OK}*

*STORE-REQUEST{EN(cat)->FR(chat)} -> RESPONSE{OK}*

*GET-REQUEST{DE(Katze)->EN} -> RESPONSE{cat}*

*GET-REQUEST{DE(Katze)->FR} -> RESPONSE{chat}*

## Solutions

### Description

- Project layout: Go files found at root level can be safely imported by other go packages. In `./cmd` we have very slim `main.go` files where we just run the commands built in `./internal/cmd`.
- Tests: Relevant tests for the `translator` package are found in `./internal/test/` directory, other tests can be found in the same directory as the objects that are tested, example: `./internal/dao`

### Known Issues and Improvements possible

- Sometimes when running all tests (`make test-all`) testing `/translator/internal/dao` freezes thus blocking execution
- Some logging could come in handy
- Maybe add comments in code

### CLI

#### Memory storage

##### Usage

`
make run-cli-memory
`

##### Run User Story Tests

`
make test-us-cli-mem
`

#### Database storage

##### Prerequisites

- A PostgreSQL database and valid credentials to connect passed in the [Makefile](Makefile)


##### Usage

`
make run-cli-db
`

##### Run User Story Tests

`
make test-us-cli-db
`

### Web API

#### Exposed endpoints

##### /{langFrom}/{textFrom}/{langTo}

- Allowed methods: `GET, POST`
  - GET: Retrieves the translation in the `langTo` language for the `textFrom` passed in `langFrom` language.
    - Response body: 
    ```
    {
      "translation" : "{textTo}"
    }
    ```
  - POST: Stores a translation for the `textFrom` in `langFrom` language, in the `langTo` language.
    - Request body:
    ```
    {
      "translation" : "{textTo}"
    }
    ```

#### Memory storage

##### Usage

`
make run-web-memory
`

##### Run User Story Tests

`
make test-us-web-mem
`

#### Database storage

##### Prerequisites

- A PostgreSQL database and valid credentials to connect passed in the [Makefile](Makefile)

##### Usage

`
make run-web-db
`

##### Run User Story Tests

`
make test-us-web-db
`

## Running tests

`make test-all`