# Tic-Tac-Toe game
A basic tic-tac-toe game developed in golang and vue.js v2. 

This game allows two users to play together on different machines.

## Game flow
* Player 1 starts a new game and get back an invitation code to invite someone else to play.
* Player 2 join the game using the shared code.
* The players keep on playing until someone is the winner or we have a tie.

## Installation
You should have golang, vue cli and node installed on your system.

* Running the backend rest api, go to the root app folder:

```
cd web-app
go run .\main.go
```

* Running the backend tests, go to the root app folder:

```
cd rest-api\games\
go test
```

* Running the frontend:
```
cd rest-api
npm i
npm run serve 
```