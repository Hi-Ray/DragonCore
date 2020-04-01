# DragonCore

DragonCore is a hub for CommunityDragon.
Everything related to CommunityDragon will be interlinked with DragonCore

&nbsp;

## Structure

```
ROOT
 ├─ .external/          # External repositories
 ├─ build/              # All build files for CI/CD and local development
 ├─ frontend/           # All front-end libraries and projects
 │   ├─ main/           # Main front-end using Nuxt
 │   ├─ auth/           # Authentication front-end using plain Vue
 │   └─ developer/      # Developer front-end using Nuxt
 │
 ├─ lib/                # The API library
 │   └─ models/         # The API models
 │
 ├─ modules/            # The API is split up in modules, this is the directory for it
 ├─ ...
 ├─ lerna.json          # Lerna configuration for managing all front-end projects
 ├─ package.json        # Global NPM dependencies for all front-end projects
 ├─ go.mod              # Dependencies of the API
 └─ main.go             # Entrypoint of the API
```

&nbsp;

## Requirements

- [Go 1.14+](https://golang.org/) (optional)
- [Node 12+](https://nodejs.org/en/) (optional)
  - [Yarn](https://yarnpkg.com/)
- [Docker](https://www.docker.com/)
  - [Compose](https://docs.docker.com/compose/install/)

Docker and Docker-Compose are used for running all the applications.
Node and Go are optional, as the only use is to install other dependencies
using a CLI. If you do want to install a dependency for node, do make sure 
to use Yarn.

&nbsp;

## Setup

### Development

All docker files are set up, all you need to do is start up the containers using docker-compose and the proper target file.
```bash
docker-compose up
```
Or run it as a daemon
```bash
docker-compose up -d
```

&nbsp;

### Production
**(not set up yet)**

We use docker-images for the production build. Just use the following command (keep 
in mind, a proxy is set up within docker. You'll need to change out the domains if 
you want to run production yourself)
```bash
docker-compose up -f docker-compose.build.yml
```
You can now spin up the production images
```bash
docker-compose up -f docker-compose.prod.yml -d
```
