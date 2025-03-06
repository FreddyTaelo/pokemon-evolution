# pokemon-evolution
pokemon evolution ms

[![Build Status](https://img.shields.io/circleci/build/github/FreddyTaelo/pokemon-evolution/main)](https://circleci.com/gh/FreddyTaelo/pokemon-evolution/tree/main)
[![codecov](https://codecov.io/gh/FreddyTaelo/pokemon-evolution/graph/badge.svg?token=TI8U5EK22W)](https://codecov.io/gh/FreddyTaelo/pokemon-evolution)

# Sample API CAll
```
Swagger
curl http://ec2-13-60-74-36.eu-north-1.compute.amazonaws.com:5001/swagger/index.html

Request ( Get pokemon evolution by ID)

curl http://ec2-13-60-74-36.eu-north-1.compute.amazonaws.com:5001/api/evolution/1

Response

{"id":0,"name":"bulbasaur","evolves_to":[{"id":0,"name":"ivysaur","evolves_to":null,"details":[{"trigger":"level-up","min_level":16}]}]}

```
# Instruction to run
```
Using make file

$ make run

```
TODO

- add extra test
