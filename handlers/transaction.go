package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	dto "waysbuck/dto/result"
	transactiondto "waysbuck/dto/transaction"
	"waysbuck/models"
	"waysbuck/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	id := int(userInfo["id"].(float64))

	request := new(transactiondto.TransactionRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	transID := int(time.Now().Unix())

	dataTransaction := models.Transaction{
		ID:     transID,
		UserID: id,
		Status: "waiting",
		CreateAt: time.Now(),
	}

	transaction, _ := h.TransactionRepository.GetTransaction(id)
	
	if transaction.Status == "waiting"{
		response := dto.SuccessResult{Code: "Success", Data: transaction}
		json.NewEncoder(w).Encode(response)
	} else {
		data, _ := h.TransactionRepository.CreateTransaction(dataTransaction)
		response := dto.SuccessResult{Code: "Success", Data: data}
		json.NewEncoder(w).Encode(response)
	}
}

func (h *handlerTransaction) FindTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]

	if userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "You're not admin"}
		json.NewEncoder(w).Encode(response)
		return
	}
	transaction, err := h.TransactionRepository.FindTransaction()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataDelete := data.ID
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: dataDelete}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(userInfo["id"].(float64))

	orders, err := h.TransactionRepository.GetOrderByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: orders}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
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

	request := new(transactiondto.StatusTransaction)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }
	// request := transactiondto.MyTransaction{
	// 	Status: r.FormValue("status"),
	// }

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	transaction, _ := h.TransactionRepository.GetTransaction(id)

	transaction.Status = request.Status

	_, err = h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	trans, err := h.TransactionRepository.GetTransaction(transaction.ID)
		if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// order, err := h.TransactionRepository.GetOrderByID(trans.UserID)
		if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// for i, p := range trans.Order {
	// 	Product[i].image = os.Getenv("PATH_FILE") + p.Product.Image
	//   }

	dataTransactions := models.Transaction{
		ID: 			trans.ID,
		Name:			trans.Name,
		Email: 		trans.Email,
		Phone: 		trans.Phone,
		Address: 	trans.Address,
		
		Total:    trans.Total,
		UserID: 	trans.UserID,
		Status: 	trans.Status,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: dataTransactions}
	json.NewEncoder(w).Encode(response)	
}

func (h *handlerTransaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	
	trans, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// order, err := h.TransactionRepository.GetOrderByID(trans.UserID)

		dataTransactions := models.Transaction{
		ID: 			trans.ID,
		Name:			trans.Name,
		Email: 		trans.Email,
		Phone: 		trans.Phone,
		Address: 	trans.Address,
		
		Total:    trans.Total,
		UserID: 	trans.UserID,
		Status: 	trans.Status,
	}


	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: dataTransactions}
	json.NewEncoder(w).Encode(response)
}