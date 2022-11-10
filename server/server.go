package server

import (
	"connectDB/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Craete User, inset into db
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var user user
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct!"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (name, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement!"))
		return
	}
	defer statement.Close()

	insert, erro := statement.Exec(user.Name, user.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar statement!"))
		return
	}
	idInsert, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter id inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInsert)))
}

// Grab all users db
func SearchUSers(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Connect()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	lines, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
		return
	}
	defer lines.Close()

	var users []user
	for lines.Next() {
		var user user
		if erro := lines.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuários!"))
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para json!"))
		return
	}
}

// Grab one user db
func SearchUSer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, erro := strconv.ParseUint(params["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter param inter"))
		return
	}
	db, erro := database.Connect()
	if erro != nil {
		w.Write([]byte("Erro ao Conectar no banco"))
		return
	}
	line, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário"))
		return
	}

	var user user
	if line.Next() {
		if erro := line.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	}
	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para json"))
		return
	}
}

// Update user db
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parms := mux.Vars(r)

	ID, erro := strconv.ParseUint(parms["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter param para interior"))
		return
	}
	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler requisição"))
		return
	}

	var user user
	if erro := json.Unmarshal(bodyRequest, &user); erro != nil {
		w.Write([]byte("Erro ao converter user para struct"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao DB"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set name = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(user.Name, user.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o user"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Remove user to db
func DeletUser(w http.ResponseWriter, r *http.Request) {
	parms := mux.Vars(r)
	ID, erro := strconv.ParseUint(parms["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter params"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao DB"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
