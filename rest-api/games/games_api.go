package games

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var gamesMapper GamesMap

func HandleRequests() {
	gamesMapper = make(GamesMap)
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.Use(commonMiddleware)

	log.Println("Creating routes")
	log.Println("GET: ", API_GAMES)
	log.Println("GET: ", API_GAMES+"/{sessionId}")
	log.Println("POST: ", API_GAMES)
	log.Println("PUT: ", API_GAMES+"/join/{sessionId}")
	log.Println("PUT: ", API_GAMES+"/play/{sessionId}")

	muxRouter.HandleFunc(API_GAMES, getAllGames).Methods("GET")
	muxRouter.HandleFunc(API_GAMES+"/{sessionId}", getSingleGame).Methods("GET")
	muxRouter.HandleFunc(API_GAMES, createNewGame).Methods("POST")
	muxRouter.HandleFunc(API_GAMES+"/join/{sessionId}", joinGame).Methods("PUT")
	muxRouter.HandleFunc(API_GAMES+"/play/{sessionId}", playGame).Methods("PUT")

	// Setting CORS (not for production!)
	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodOptions},
	})
	handler := c.Handler(muxRouter)

	log.Fatalln(http.ListenAndServe(API_PORT, handler))
	log.Println("Listening to the", API_PORT, "port")
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getAllGames(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: getAllGames")

	json.NewEncoder(w).Encode(gamesMapper.ToList())
}

func getSingleGame(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: getSingleGame")

	// retrieving sessionId
	sessionId := mux.Vars(r)["sessionId"]
	log.Println("SessionId:", sessionId)

	game := gamesMapper[sessionId]
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
	gamesMapper[g.SessionId] = &g

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

	game := gamesMapper[sessionId]
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

	game := gamesMapper[sessionId]
	game.Play(move.Symbol, move.Index)

	json.NewEncoder(w).Encode(game)
}

func (m GamesMap) ToList() []Game {
	log.Println("Converting map of games to list")

	var list []Game
	for _, v := range gamesMapper {
		list = append(list, *v)
	}

	return list
}
