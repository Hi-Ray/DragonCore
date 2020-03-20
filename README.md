# DragonCore

DragonCore is a hub for CommunityDragon.
Everything related to CommunityDragon will be interlinked with DragonCore

&nbsp;

## Requirements

- [Node 12+](https://nodejs.org/en/)
- [Go 1.14+](https://golang.org/)
  - [Air](https://github.com/cosmtrek/air)
- [Docker](https://www.docker.com/)
  - [Compose](https://docs.docker.com/compose/install/)

Both Node and Go are required if you want to run the application directly. You can use docker as well.

&nbsp;

## Setup

### Development

**Direct**&nbsp;
Clone the repository and enter it
```bash
git clone https://github.com/CommunityDragon/DragonCore.git
```
Use Air to start live-reload. The config should be pre-configured.
```
air
```
To start up the front-end, enter the front-end directory and install all the dependencies.
```bash
cd frontend/

# either use NPM or Yarn
yarn install
npm install
```
Now you just need to run development mode
```bash
# either use NPM or Yarn
yarn dev
npm run dev
```

&nbsp;

**Docker**&nbsp;
All docker files are set up, all you need to do is start up the containers using docker-compose and the proper target file.
```bash
docker-compose up -f docker-compose.dev.yml
```
Or run it as a daemon
```bash
docker-compose up -f docker-compose.dev.yml -d
```

### Production
Everything on production is run through docker. Just use docker-compose to start up production
```bash
docker-compose up -f docker-compose.prod.yml
```

&nbsp;

## Build
We use docker-images for the production build. Just use the following command
```bash
docker-compose up -f docker-compose.build.yml
```