import "./Players.scss";

const Players = (props: any) => {
  const player1: string = props.player1;
  const player2: string = props.player2;

  return (
    <div className="mt-4 mb-5 players">
      <div className="d-flex justify-content-center align-items-center">
        <span className="h1 p1 text-truncate">{player1 || "?"}</span>
        <span className="h2 px-5">vs</span>
        <span className="h1 p2 text-truncate">{player2 || "?"}</span>
      </div>
      <div className="d-flex justify-content-center">
        <hr style={{ width: "630px", margin: 0 }} />
      </div>
    </div>
  );
};

export default Players;
