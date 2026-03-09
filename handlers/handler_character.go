package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/eldmark/go-http/models"
	"github.com/eldmark/go-http/utils"
)

type CharacterHandler struct {
	FilePath   string
	Characters []models.Character
}

func NewCharacterHandler(path string) *CharacterHandler {

	h := &CharacterHandler{
		FilePath: path,
	}

	h.load()
	return h
}
func (h *CharacterHandler) load() {

	data, err := os.ReadFile(h.FilePath)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &h.Characters)
}
func (h *CharacterHandler) save() {

	data, _ := json.MarshalIndent(h.Characters, "", "  ")
	os.WriteFile(h.FilePath, data, 0644)
}

func (h *CharacterHandler) Ping(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "pong",
	})
}

func (h *CharacterHandler) GetCharacters(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, h.Characters)
}

func (h *CharacterHandler) GetCharacterByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/characters/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid id parameter",
		})
		return
	}

	for _, character := range h.Characters {
		if character.ID == id {
			utils.WriteJSON(w, http.StatusOK, character)
			return
		}
	}

	utils.WriteJSON(w, http.StatusNotFound, map[string]string{
		"error": "Character not found",
	})
}

func (h *CharacterHandler) AddCharacter(w http.ResponseWriter, r *http.Request) {
	var newCharacter models.Character
	err := json.NewDecoder(r.Body).Decode(&newCharacter)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	newCharacter.ID = h.getNextID()
	h.Characters = append(h.Characters, newCharacter)
	h.save()

	utils.WriteJSON(w, http.StatusCreated, newCharacter)
}

func (h *CharacterHandler) getNextID() int {
	maxID := 0
	for _, character := range h.Characters {
		if character.ID > maxID {
			maxID = character.ID
		}
	}
	return maxID + 1
}

func (h *CharacterHandler) UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/characters/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid id parameter",
		})
		return
	}

	var updatedCharacter models.Character
	err = json.NewDecoder(r.Body).Decode(&updatedCharacter)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	for i, character := range h.Characters {
		if character.ID == id {
			updatedCharacter.ID = id
			h.Characters[i] = updatedCharacter
			h.save()
			utils.WriteJSON(w, http.StatusOK, updatedCharacter)
			return
		}
	}

	utils.WriteJSON(w, http.StatusNotFound, map[string]string{
		"error": "Character not found",
	})
}

func (h *CharacterHandler) DeleteCharacter(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/characters/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid id parameter",
		})
		return
	}

	for i, character := range h.Characters {
		if character.ID == id {
			h.Characters = append(h.Characters[:i], h.Characters[i+1:]...)
			h.save()
			utils.WriteJSON(w, http.StatusOK, map[string]string{
				"message": "Character deleted",
			})
			return
		}
	}

	utils.WriteJSON(w, http.StatusNotFound, map[string]string{
		"error": "Character not found",
	})
}
