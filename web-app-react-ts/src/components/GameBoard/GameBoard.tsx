import "./GameBoard.scss";
import Cell from "./Cell";

const GameBoard = (props: any) => {
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
            playerSymbol={props.playerSymbol}
            nextSymbol={props.nextSymbol}
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
