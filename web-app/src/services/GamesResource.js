import axios from "axios";

export default class GamesResource {
  constructor() {
    this.api = "/games";
  }

  getAllGames() {
    return axios.get(this.api);
  }

  getGame(sessionId) {
    return axios.get(`${this.api}/${sessionId}`);
  }

  newGame(playerName) {
    return axios.post(this.api, { name: playerName });
  }

  joinGame(sessionId, playerName) {
    return axios.put(`${this.api}/join/${sessionId}`, {
      name: playerName,
    });
  }

  play(sessionId, index, symbol) {
    return axios.put(`${this.api}/play/${sessionId}`, {
      index: index,
      symbol: symbol,
    });
  }
}
