## Blueprint/Boilerplate For Golang Projects

[![Build Status](https://travis-ci.com/MartinHeinz/go-project-blueprint.svg?branch=master)](https://travis-ci.com/MartinHeinz/go-project-blueprint)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=MartinHeinz_go-project-blueprint&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=MartinHeinz_go-project-blueprint)
[![Test Coverage](https://api.codeclimate.com/v1/badges/ec7ebefe63609984cb5c/test_coverage)](https://codeclimate.com/github/MartinHeinz/go-project-blueprint/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/MartinHeinz/go-project-blueprint)](https://goreportcard.com/report/github.com/MartinHeinz/go-project-blueprint)

### Setting Up
- Replace All Occurrences of `martinheinz/go-project-blueprint` with your username repository name
- Replace All Occurrences of `blueprint` with your desired image name


### Using GitHub Registry

Create and Push:

```bash
docker login docker.pkg.github.com -u <USERNAME> -p <GITHUB_TOKEN>
docker build -t  docker.pkg.github.com/martinheinz/go-project-blueprint/blueprint:latest .
# make container
docker push docker.pkg.github.com/martinheinz/go-project-blueprint/blueprint:latest
# make push
```

Pull and Run:

```bash
docker pull docker.pkg.github.com/martinheinz/go-project-blueprint/blueprint:latest
docker run docker.pkg.github.com/martinheinz/go-project-blueprint/blueprint:latest
```


### Setup new SonarCloud Project

- On _SonarCloud_:
    - Click _Plus_ Sign in Upper Right Corner
    - _Analyze New Project_
    - Click _GitHub app configuration_ link
    - Configure SonarCloud
    - Select Repository and Save
    - Go Back to Analyze Project
    - Tick Newly Added Repository
    - Click Set Up
    - Click Configure with Travis
    - Copy the Command to Encrypt the Travis Token
    - Run `travis encrypt --com <TOKEN_YOU_COPPIED>`
    - Populate the `secure` Field in `.travis.yml` with outputted string
    - Follow steps to populate your `sonar-project.properties`
    - Push
- On Travis CI:
    - Set `DOCKER_USERNAME`
    - Set `DOCKER_PASSWORD` to Your GitHub Registry Token

### Setup CodeClimate
- Go to <https://codeclimate.com/github/repos/new>
- Add Repository
- Go to Test Coverage Tab
- Copy Test Reporter ID
- Go to Travis and Open Settings for Your Repository
- Add Environment Variable: name: `CC_TEST_REPORTER_ID`, value: _Copied from CodeClimate_
