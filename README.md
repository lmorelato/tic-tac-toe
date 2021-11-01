# Tic-Tac-Toe game

A basic tic-tac-toe game developed in golang and vue.js v2.

This game allows two users to play together on different machines.

**UPDATE: A new version with small improvments and the frontend developed in React was recently added!**


## Game flow

- Player 1 starts a new game and get back an invitation code to invite someone else to play.
- Player 2 join the game using the shared code.
- The players keep on playing until someone is the winner or we have a tie.


- Vue frontend
![game flow](https://github.com/lmorelato/tic-tac-toe/blob/main/docs/tic-tac-toe.gif "Game Flow")


- **NEW React frontend**
![game flow](https://github.com/lmorelato/tic-tac-toe/blob/main/docs/ttt-react.gif "Game Flow")

## Installation

You should have golang, vue cli and node installed on your system.

- Running the backend rest api, go to the root app folder:

```
cd rest-api
go run .\main.go
```

- Running the backend tests, go to the root app folder:

```
cd rest-api\games\
go test
```

- Running the frontend in Vue:

```
cd web-app
npm i
npm run serve
```

- **Running the frontend in React:**

```
cd web-app-react-ts
npm i
npm run start
```

## API Endpoints

- Get All Games - GET: /api/games
- Get Single Game - GET: /api/games/{sessionId}
- Create New Game - POST: /api/games
- Play Game - PUT: /api/games/join/{sessionId}
- Join Game - PUT: /api/games/play/{sessionId}

A postman config file can be found [here](https://github.com/lmorelato/tic-tac-toe/blob/main/docs/tic-tac-toe.postman_collection.json) to help with api tests.
