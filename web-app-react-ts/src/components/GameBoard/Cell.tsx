import "./Cell.scss";
import { BsCircle, BsX } from "react-icons/bs";

const Cell = (props: any) => {
  const index = props.index;
  const isSymbolX = props.symbol === "X";
  const isSymbolO = props.symbol === "O";
  const cursor = !isSymbolX && !isSymbolO ? "pointer" : "default";
  let className =
    "col-4 board-cell d-flex align-items-center justify-content-center";

  if (props.gameOver && props.winningPath.indexOf(index) >= 0) {
    className = className + " blink-bg";
  }

  const clickHandler = () => {
    props.onCellClicked(index);
  };

  return (
    <div
      className={className}
      style={{ cursor: cursor }}
      onClick={clickHandler}
    >
      {isSymbolX && <BsX className="x"></BsX>}
      {isSymbolO && <BsCircle className="o"></BsCircle>}
    </div>
  );
};

export default Cell;
