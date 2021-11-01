import "./GameBoard.scss";
import Cell from "./Cell";

type GameBoardProps = {
  board: string[];
  playerSymbol: string;
  nextSymbol: string;
  gameOver: boolean;
  winningPath: number[];
  onBoardClicked: (index: number, symbol: string) => void;
}; 

const GameBoard = (props: GameBoardProps) => {
  const cellClickedHandler = (index: number) => {
    props.onBoardClicked(index, props.playerSymbol);
  };

  return (
    <div className="d-flex justify-content-center">
      <div className="row board">
        {props.board.map((value: string, index: number) => (
          <Cell
            symbol={value}
            index={index}
            key={index}
            gameOver={props.gameOver}
            winningPath={props.winningPath}
            onCellClicked={cellClickedHandler}
          ></Cell>
        ))}
      </div>
    </div>
  );
};

export default GameBoard;
