import { Log } from "./Log";
import { Player } from "./Player";

export interface Game {
     sessionId: string;
     player1: Player;
     player2: Player;
     moves: string[];
     logs: Log[];
     winner: Player;
     winningPath: number[];
     nextSymbol: string;
}

export const GameInitialState: Game = {
     sessionId: "",
     player1: {
          name: "",
          symbol: "X",
     },
     player2: {
          name: "",
          symbol: "O",
     },
     moves: [],
     logs: [],
     winner: {
          name: "",
          symbol: "",
     },
     winningPath: [],
     nextSymbol: "",
};
