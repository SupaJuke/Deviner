# Backend

## Installation

### Setting up .env file

An example is provided in the form of .`env.example`. Simply copy and rename it to `.env`. Then, you would need to append the value to the right side of each equal sign.

### Database

PostgreSQL is required to run this application. There are two ways you can do this:

1. install from source code: https://www.postgresql.org/download/
2. install from Docker (Recommended): https://www.docker.com/

If you choose the second option, we have provided a script `/scripts/setup-db.sh` that will automatically set up the database for you once you have started the container. You can use the script by running the following command on root diretory of the repository:

```zsh
./scripts/setup-db.sh <container name>
```

### Golang

**TODO**
