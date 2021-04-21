# User-Role Hierarchy Manager

## Prerequisites:
* GoLang is properly installed
* make is properly installed

## Assumptions:
* No Database needed. All solution is managed InMemory
* Multiple Users could be assigned to the same Role
* While retrieving SubOrdinates, all users assigned to particular role will be included

## Run Tests
Clean project
```bash
make clean
```
Build
```bash
make build
```
Run test cases
```bash
make test
```