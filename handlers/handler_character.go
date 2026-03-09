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

	query := r.URL.Query()

	// 1) Query param por ID
	idParam := query.Get("id")
	if idParam != "" {

		id, err := strconv.Atoi(idParam)
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
		return
	}

	filtered := h.filterCharacters(query)

	utils.WriteJSON(w, http.StatusOK, filtered)
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

	if newCharacter.Name == "" ||
		newCharacter.FightStyle == "" ||
		newCharacter.Weapon == "" ||
		newCharacter.Speciality == "" {

		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Missing required fields",
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

	if updatedCharacter.Name == "" ||
		updatedCharacter.FightStyle == "" ||
		updatedCharacter.Weapon == "" ||
		updatedCharacter.Speciality == "" {

		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Missing required fields",
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

func (h *CharacterHandler) filterCharacters(query map[string][]string) []models.Character {

	var result []models.Character

	name := ""
	if len(query["name"]) > 0 {
		name = query["name"][0]
	}
	devilFruit := ""
	if len(query["devil_fruit"]) > 0 {
		devilFruit = query["devil_fruit"][0]
	}
	weapon := ""
	if len(query["weapon"]) > 0 {
		weapon = query["weapon"][0]
	}
	speciality := ""
	if len(query["speciality"]) > 0 {
		speciality = query["speciality"][0]
	}
	fightStyle := ""
	if len(query["fight_style"]) > 0 {
		fightStyle = query["fight_style"][0]
	}

	for _, c := range h.Characters {

		if name != "" && !strings.EqualFold(c.Name, name) {
			continue
		}

		if devilFruit != "" && !strings.Contains(strings.ToLower(c.DevilFruit), strings.ToLower(devilFruit)) {
			continue
		}

		if weapon != "" && !strings.Contains(strings.ToLower(c.Weapon), strings.ToLower(weapon)) {
			continue
		}

		if speciality != "" && !strings.Contains(strings.ToLower(c.Speciality), strings.ToLower(speciality)) {
			continue
		}

		if fightStyle != "" && !strings.Contains(strings.ToLower(c.FightStyle), strings.ToLower(fightStyle)) {
			continue
		}

		result = append(result, c)
	}

	return result
}
