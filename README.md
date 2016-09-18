## Tracing REST

* Record any actions your system perform, along with any properties that describe the action
* A minimal docker alpine container
* Automatically pushes it to dockerhub if tests pass

## Setup
Run Migrations
```bash
goose up
```

Env vars
```bash
export TRACKING_REST_DB=dbname=tracing_rest_dev sslmode=disable
export TRACKING_REST_PORT=3000
```

## Docker
This repository has automated image builds on hub.docker.com after successfully building and testing. See the `deployment` section of [circle.yml](circle.yml) for details on how this is done. Note that three environment variables need to be set on CircleCI for the deployment to work:

  * DOCKER_EMAIL - The email address associated with the user with push access to the Docker Hub repository
  * DOCKER_USER - Docker Hub username
  * DOCKER_PASS - Docker Hub password (these are all stored encrypted on CircleCI, and you can create a deployment user with limited permission on Docker Hub if you like)

```bash
$ sh buid-container
$ docker run -it -t -p 3000:3000 --name tracing-rest rafaeljesus/tracing-rest
```

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request

## Badges

[![CircleCI](https://circleci.com/gh/rafaeljesus/tracing-rest.svg?style=svg)](https://circleci.com/gh/rafaeljesus/tracing-rest)
[![](https://badge.imagelayers.io/rafaeljesus/tracing-rest:latest.svg)](https://imagelayers.io/?images=rafaeljesus/tracing-rest:latest 'Get your own badge on imagelayers.io')

---

> GitHub [@rafaeljesus](https://github.com/rafaeljesus) &nbsp;&middot;&nbsp;
> Twitter [@rafaeljesus](https://twitter.com/_jesus_rafael)
