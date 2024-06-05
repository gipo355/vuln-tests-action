# Hello world docker action with go

This action prints "Hello World" or "Hello" + the name of a person to greet to the log.

## Inputs

## `who-to-greet`

**Required** The name of the person to greet. Default `"World"`.

## Outputs

## `time`

The time we greeted you.

## Example usage

uses: actions/hello-world-docker-action@v1
with:
who-to-greet: 'Mona the Octocat'

## note for nmap

we are using a docker container which allows us to pull and prepare needed tools for the scan.

if we were to use a js action, we would have to install nmap on the runner machine which is not ideal.

golang provides efficient concurrency and minimal image + mem size for containers using alpine.

# todo

- add licenses
- fork vulner repo
- copy vulner scripts in continer
- add more nmap options
- create auto release and docker publish on tag
- generate sarif report
- add cobra and viper
- split github utils in its own lib
- understand golang versioning forlib
- split smiattack cli with nmap in its own repo for action using the github lib and utils

- make a list of all go libs and tools
