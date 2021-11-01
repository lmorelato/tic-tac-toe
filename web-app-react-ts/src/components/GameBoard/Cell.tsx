import "./Cell.scss";
import { BsCircle, BsX } from "react-icons/bs";

type CellProps = {
  index: number;
  symbol: string;
  gameOver: boolean;
  winningPath: number[];
  onCellClicked: (index: number) => void;
}; 

const Cell = (props: CellProps) => {
  const index = props.index;
  const isSymbolX = props.symbol === "X";
  const isSymbolO = props.symbol === "O";
  const cursor = !isSymbolX && !isSymbolO ? "pointer" : "default";
  let css = "col-4 board-cell d-flex align-items-center justify-content-center";

  // let's animate the winning path
  if (props.gameOver && props.winningPath.indexOf(index) >= 0) {
    css = css + " blink-bg";
  }

  const clickHandler = () => {
    props.onCellClicked(index);
  };

  return (
    <div className={css} style={{ cursor: cursor }} onClick={clickHandler}>
      {isSymbolX && <BsX className="x"></BsX>}
      {isSymbolO && <BsCircle className="o"></BsCircle>}
    </div>
  );
};

export default Cell;
