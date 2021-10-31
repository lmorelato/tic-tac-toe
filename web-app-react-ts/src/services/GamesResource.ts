import axios from "axios";
import { Game } from "../models/Game";

export default class GamesResource {
  api: string;

  constructor() {
    this.api = "/games";
  }

  public getAllGames() {
    return axios.get(this.api);
  }

  public getGame(sessionId: string) {
    return axios.get(`${this.api}/${sessionId}`);
  }

  public newGame(playerName: string) {
    return axios.post(this.api, { name: playerName });
  }

  public joinGame(sessionId: string, playerName: string) {
    return axios.put(`${this.api}/join/${sessionId}`, {
      name: playerName,
    });
  }

  public play(sessionId: string, index: number, symbol: string) {
    return axios.put(`${this.api}/play/${sessionId}`, {
      index: index,
      symbol: symbol,
    });
  }
}
