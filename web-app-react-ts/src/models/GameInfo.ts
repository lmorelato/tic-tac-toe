export interface GameInfo {
     modalShow: boolean,
     modalAction: string,
     alertShow: boolean,
     alertMessage: string,
     alertVariant: string,
     playerSymbol: string,
     gameOver: boolean,
     timer: any,
}

export const GameInfoInitialState: GameInfo = {
     modalShow: false,
     modalAction: "",
     alertShow: false,
     alertMessage: "",
     alertVariant: "",
     playerSymbol: "",
     gameOver: false,
     timer: 0,
};
