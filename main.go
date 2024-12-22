package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Player 结构体表示一个玩家
type Player struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Level     string    `json:"level"`
	CreatedAt time.Time `json:"created_at"`
}

// Level 结构体表示一个游戏等级
type Level struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// 假设我们初始化一个玩家数据
var players = []Player{
	{ID: "1", Name: "Player 1", Level: "5", CreatedAt: time.Now()},
	{ID: "2", Name: "Player 2", Level: "3", CreatedAt: time.Now()},
}

// 假设我们初始化一个等级数据
var levels = []Level{
	{ID: "1", Name: "Level 1"},
	{ID: "2", Name: "Level 2"},
}

// 创建玩家
func createPlayer(w http.ResponseWriter, r *http.Request) {
	var player Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 创建一个 ID（这里可以用时间戳或随机数）
	player.ID = fmt.Sprintf("%d", rand.Int())
	player.CreatedAt = time.Now()

	// 添加玩家到内存数据库
	players = append(players, player)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player)
}

// 获取所有玩家
func getPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

// 获取指定玩家
func getPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for _, player := range players {
		if player.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(player)
			return
		}
	}

	http.Error(w, "Player not found", http.StatusNotFound)
}

// 更新玩家
func updatePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedPlayer Player
	err := json.NewDecoder(r.Body).Decode(&updatedPlayer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, player := range players {
		if player.ID == id {
			players[i].Name = updatedPlayer.Name
			players[i].Level = updatedPlayer.Level
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(players[i])
			return
		}
	}

	http.Error(w, "Player not found", http.StatusNotFound)
}

// 删除玩家
func deletePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i, player := range players {
		if player.ID == id {
			players = append(players[:i], players[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Player not found", http.StatusNotFound)
}

// 创建等级
func createLevel(w http.ResponseWriter, r *http.Request) {
	var level Level
	err := json.NewDecoder(r.Body).Decode(&level)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 创建一个 ID（这里可以用时间戳或随机数）
	level.ID = fmt.Sprintf("%d", rand.Int())

	// 添加等级到内存数据库
	levels = append(levels, level)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(level)
}

// 获取所有等级
func getLevels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(levels)
}

func main() {
	r := mux.NewRouter()

	// 玩家相关路由
	r.HandleFunc("/players", getPlayers).Methods("GET")
	r.HandleFunc("/players", createPlayer).Methods("POST")
	r.HandleFunc("/players/{id}", getPlayer).Methods("GET")
	r.HandleFunc("/players/{id}", updatePlayer).Methods("PUT")
	r.HandleFunc("/players/{id}", deletePlayer).Methods("DELETE")

	// 等级相关路由
	r.HandleFunc("/levels", getLevels).Methods("GET")
	r.HandleFunc("/levels", createLevel).Methods("POST")

	// 启动服务器
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}
