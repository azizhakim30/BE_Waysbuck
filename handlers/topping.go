package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	dto "waysbuck/dto/result"
	toppingdto "waysbuck/dto/topping"
	"waysbuck/models"
	"waysbuck/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerTopping struct {
	ToppingRepository repositories.ToppingRepository
}

// Create `path_file` Global variable here ...
// var path_file = os.Getenv("PATH_FILE")

func HandlerTopping(ToppingRepository repositories.ToppingRepository) *handlerTopping {
	return &handlerTopping{ToppingRepository}
}

func (h *handlerTopping) FindTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	Toppings, err := h.ToppingRepository.FindTopping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	for i, p := range Toppings {
		Toppings[i].Image = os.Getenv("PATH_FILE") + p.Image
	  }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: Toppings}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTopping) GetTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var topping models.Topping
	topping, err := h.ToppingRepository.GetTopping(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	topping.Image = os.Getenv("PATH_FILE") + topping.Image

	toppings, err := h.ToppingRepository.GetTopping(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: toppings}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTopping) CreateTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Get dataFile from midleware and store to filename variable here ...
	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string) // add this code

	price, _ := strconv.Atoi(r.FormValue("price"))
	qty, _ := strconv.Atoi(r.FormValue("qty"))
	request := toppingdto.CreateTopping{
		Title:       r.FormValue("title"),
		Price:      price,
		Qty:        qty,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping := models.Topping{
		Title:  request.Title,
		Price:  request.Price,
		Image:  filename,
		Qty:    request.Qty,
		UserID: userId,
	}

	topping, err = h.ToppingRepository.CreateTopping(topping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping, _ = h.ToppingRepository.GetTopping(topping.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: topping}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTopping) UpdateTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]
	userID := int(userInfo["id"].(float64))

	if userID != id && userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "You're not admin"}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string) // add this code

	price, _ := strconv.Atoi(r.FormValue("price"))
	qty, _ := strconv.Atoi(r.FormValue("qty"))
	request := toppingdto.UpdateTopping{
		Title:       r.FormValue("title"),
		Price:      price,
		Qty:        qty,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	topping, _ := h.ToppingRepository.GetTopping(id)

	// product.Title = request.Title
	// product.Price = request.Price
	// product.Qty = request.Qty
	
	if request.Title != "" {
		topping.Title = request.Title
	}

	if request.Price != 0 {
		topping.Price = request.Price
	}

	if request.Qty != 0 {
		topping.Qty = request.Qty
	}
	
	if filename != "false" {
		topping.Image = filename
	}

	topping, err = h.ToppingRepository.UpdateTopping(topping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: topping}
	json.NewEncoder(w).Encode(response)	
}

func (h *handlerTopping) DeleteTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]


	if userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "You're not admin"}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping, err := h.ToppingRepository.GetTopping(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	_, err = h.ToppingRepository.DeleteTopping(topping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data := models.Topping{
		ID: topping.ID,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: convertResponseTopping(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseTopping(u models.Topping) toppingdto.DeleteResponse {
	return toppingdto.DeleteResponse{
		ID: 			u.ID,
	
	}
}