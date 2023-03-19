# Pooe Guessing Game

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
>     - /login
>       - Very simple yes / no with password combination
>       - Return "token"
>     - /guess
>       - Access to this endpoint needed to be authenticated via token returned from login
>       - Guess the hidden number - if correct, return HTTP 201 and regenerate the number
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

Please go to `go`/`react` and follow the instructions in each folder to run back and frontend respectively
