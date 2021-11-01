export interface GameInfo {
     playerSymbol: string,
     gameOver: boolean,
     timer: any,
}

export const GameInfoInitialState: GameInfo = {
     playerSymbol: "",
     gameOver: false,
     timer: 0,
};
