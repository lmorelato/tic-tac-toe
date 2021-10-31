import { useState } from "react";
import "./App.scss";
import GameBoard from "./components/GameBoard/GameBoard";
import Logs from "./components/Logs";
import Players from "./components/Players";
import StartGame from "./components/StartGame";
import GamesResource from "./services/GamesResource";
import { Game, GameInitialState } from "./models/Game";
import { GameInfoInitialState } from "./models/GameInfo";
import { clone } from "lodash";
import { useInterval } from "usehooks-ts";

const App = () => {
  const restApi = new GamesResource();

  const [game, setGame] = useState(GameInitialState);
  const [gameInfo, setGameInfo] = useState(GameInfoInitialState);
  const [gameStarted, setGameStarted] = useState(false);

  const emptyStringArray: string[] = [];
  const [startGameLogs, setStartGameLogs] = useState(emptyStringArray);
  const [endGameLogs, setEndGameLogs] = useState(emptyStringArray);

  useInterval(
    () => {
      updatePolling();
    },
    gameStarted ? 500 : null
  );

  const onGameLaunchedHandler = async (
    playerName: string,
    sessionId: string
  ) => {
    setGameStarted(true);

    const isNewGame = !sessionId;
    gameInfo.playerSymbol = isNewGame ? "X" : "O";
    setGameInfo(gameInfo);

    const response = isNewGame
      ? await restApi.newGame(playerName)
      : await restApi.joinGame(sessionId, playerName);

    const gameData = response.data as Game;
    setGame(gameData);

    logGameStart(gameData, isNewGame);
  };

  const logGameStart = (gameData: Game, isNewGame: boolean) => {
    const gameLogs = [
      `${gameData.player1.name.toUpperCase()} has started a new game, Session Id: ${
        gameData.sessionId
      }`,
    ];
    if (!isNewGame) {
      gameLogs.unshift(
        `${gameData.player2.name.toUpperCase()} has joined the new game`
      );
    }

    const newStartGameLogs = clone(startGameLogs);
    newStartGameLogs.unshift(...gameLogs);
    setStartGameLogs(newStartGameLogs);
  };

  const onBoardClickedHandler = (index: number, symbol: string) => {
    if (
      game.moves[index] === "" &&
      gameInfo.playerSymbol === game.nextSymbol &&
      !gameInfo.gameOver
    ) {
      setBoardCell(index, symbol);
      restApi.play(game.sessionId, index, symbol);
    }
  };

  const setBoardCell = (index: number, symbol: string) => {
    const newGame = clone(game);
    newGame.moves[index] = symbol;
    setGame(newGame);
  };

  const updatePolling = async () => {
    if (game.sessionId) {
      const waitingPlayer2 = game.player2.name === "";
      const waitingWinner = game.winner.symbol === "";

      const response = await restApi.getGame(game.sessionId);
      const gameData = response.data as Game;
      setGame(gameData);

      gameInfo.gameOver = game.winner.symbol !== "";
      setGameInfo(gameInfo);

      //this.checkTurn(noPlayer2, noWinner);
      checkPlayer2(gameData, waitingPlayer2);
      checkWinner(gameData, waitingWinner);
    }
  };

  const checkPlayer2 = (gameData: Game, waitingPlayer2: boolean) => {
    if (waitingPlayer2 && gameData.player2.name !== "") {
      const joinLog = `${gameData.player2.name.toUpperCase()} has joined the game`;
      // this.showAlertMessage(
      //   joinLog,
      //   "warning"
      // );

      const newStartGameLogs = clone(startGameLogs);
      newStartGameLogs.unshift(joinLog);
      setStartGameLogs(newStartGameLogs);
    }
  };

  const checkWinner = (gameData: Game, waitingWinner: boolean) => {
    if (waitingWinner && gameData.winner.name !== "") {
      const newEndGameLogs = clone(endGameLogs);

      if (gameData.winner.name === "NO_WINNER") {
        // this.showAlertMessage(
        //   `Oh no! It is a tie... Let's play again! =)`,
        //   "info"
        // );
        newEndGameLogs.push("It is a tie!");
      } else {
        if (gameData.winner.symbol === gameInfo.playerSymbol) {
          // this.showAlertMessage(`You Win! =)`, "success");
        } else {
          // this.showAlertMessage(`You Lose! =(`, "danger");
        }
        newEndGameLogs.push(
          `${gameData.winner.name.toUpperCase()} won the game!`
        );
      }

      setEndGameLogs(newEndGameLogs);
    }
  };

  return (
    <div className="App mb-5">
      <div className="header">
        Tic-Tac-Toe React - SessionId: {game.sessionId}
      </div>

      {!gameStarted && (
        <div className="my-4 mx-5 d-flex justify-content-center">
          <StartGame onGameLaunched={onGameLaunchedHandler}></StartGame>
        </div>
      )}

      {gameStarted && (
        <div className="container">
          <Players
            player1={game.player1.name}
            player2={game.player2.name}
          ></Players>
          <div className="row">
            <div className="col mt-2">
              <GameBoard
                board={game.moves}
                playerSymbol={gameInfo.playerSymbol}
                nextSymbol={game.nextSymbol}
                gameOver={gameInfo.gameOver}
                winningPath={game.winningPath}
                onBoardClicked={onBoardClickedHandler}
              ></GameBoard>
            </div>
            <div className="col">
              <Logs
                logs={game.logs}
                startGameLogs={startGameLogs}
                endGameLogs={endGameLogs}
              ></Logs>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default App;
