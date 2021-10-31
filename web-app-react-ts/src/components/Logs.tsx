import { Log } from "../models/Log";
import "./Logs.scss";
import { orderBy } from "lodash";

const Logs = (props: any) => {
  let logs: Log[] = props.logs;
  const hasLogs = logs && logs.length > 0;

  if (hasLogs) {
    logs = orderBy(logs, ['order'], ['desc']);
  }

  return (
    <div className="d-flex justify-content-center">
      <div className="logs">
        <div className="px-3 py-3 title">
          GAME LOGS *********************************************
        </div>
        {props.endGameLogs.map((log: string, index: number) => (
          <div className="line px-3 pt- ml-1" key={index}>
            - Result: {log}
          </div>
        ))}
        {hasLogs &&
          logs.map((log: Log, index: number) => (
            <div className="line px-3 pt- ml-1" key={index}>
              - #{log.order + 1}, Player: {log.player.name.toUpperCase()}, Position{" "}
              {log.index + 1}, Symbol: {log.player.symbol}
            </div>
          ))}
        {props.startGameLogs.map((log: string, index: number) => (
          <div className="line px-3 pt- ml-1" key={index}>
            - {log}
          </div>
        ))}
      </div>
    </div>
  );
};

export default Logs;
