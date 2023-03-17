# Pooe-Guessing-Game

A project I should have made 2 years ago for Agoda internship.

> **Fullstack Guessing Game** - Create web application which can do following  
> Web - React  
> Backend - Golang
>
> Task:
>
> - Web:
>   - Login if not authenticated
>   - Guess the number
>   - Do API call to backend to guess the number
> - Backend Service:
>   - API endpoints
>   - /login
>     - Very simple yes / no with password combination
>     - Return "token"
>   - /guess
>     - Access to this endpoint needed to be authenticated via token returned from login
>     - Guess the hidden number - if correct, return HTTP 201 and regenerate the number
>   - RESTful
>   - Your response should be in form of JSON format
>   - Responses should have CRUD functionality
>
> Bonus (for challenge):
>
> - Web
>   - Use React.js context for authentication
> - Backend Service
>   - Use of middleware for authentication
>   - If we wanted to hide the guess data by not using GET, can we use other method to do so ?

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
