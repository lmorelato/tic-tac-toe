package games

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var GamesMapper GamesMap

func HandleRequests() {
	GamesMapper = make(GamesMap)
	apiRouter := mux.NewRouter().StrictSlash(true)

	log.Println("Creating routes")
	apiRouter.HandleFunc(API_GAMES, getAllGames).Methods("GET")
	log.Println("GET: ", API_GAMES)

	apiRouter.HandleFunc(API_GAMES+"/{sessionId}", getSingleGame).Methods("GET")
	log.Println("GET: ", API_GAMES+"/{sessionId}")

	apiRouter.HandleFunc(API_GAMES, createNewGame).Methods("POST")
	log.Println("POST: ", API_GAMES)

	apiRouter.HandleFunc(API_GAMES+"/join/{sessionId}", joinGame).Methods("PUT")
	log.Println("PUT: ", API_GAMES+"/join/{sessionId}")

	apiRouter.HandleFunc(API_GAMES+"/play/{sessionId}", playGame).Methods("PUT")
	log.Println("PUT: ", API_GAMES+"/play/{sessionId}")

	log.Println("Listening to the", API_PORT, "port")
	log.Fatalln(http.ListenAndServe(API_PORT, apiRouter))
}

// REST API
func getAllGames(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: getAllGames")

	json.NewEncoder(w).Encode(GamesMapper.ToList())
}

func getSingleGame(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: getSingleGame")

	// retrieving sessionId
	sessionId := mux.Vars(r)["sessionId"]
	log.Println("SessionId:", sessionId)

	game := GamesMapper[sessionId]
	json.NewEncoder(w).Encode(game)
}

func createNewGame(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: createNewGame")

	// retrieving body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Request body:", string(reqBody))

	var p Player
	json.Unmarshal(reqBody, &p)

	g := NewGame(p.Name)
	GamesMapper[g.SessionId] = &g

	json.NewEncoder(w).Encode(g)
}

func joinGame(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: joinGame")

	// retrieving sessionId
	sessionId := mux.Vars(r)["sessionId"]
	log.Println("SessionId:", sessionId)

	// retrieving body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Request body:", string(reqBody))

	var p Player
	json.Unmarshal(reqBody, &p)

	game := GamesMapper[sessionId]
	game.SetPlayer2(p.Name)

	json.NewEncoder(w).Encode(game)
}

func playGame(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: playGame")

	// retrieving sessionId
	sessionId := mux.Vars(r)["sessionId"]
	log.Println("SessionId:", sessionId)

	// retrieving body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Request body:", string(reqBody))

	var move PlayerMove
	json.Unmarshal(reqBody, &move)

	game := GamesMapper[sessionId]
	game.Play(move.Symbol, move.Index)

	json.NewEncoder(w).Encode(game)
}

func (m GamesMap) ToList() []Game {
	log.Println("Converting map of games to list")

	var list []Game
	for _, v := range GamesMapper {
		list = append(list, *v)
	}

	return list
}
