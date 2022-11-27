package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	dto "waysbuck/dto/result"
	transactiondto "waysbuck/dto/transaction"
	"waysbuck/models"
	"waysbuck/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	request := new(transactiondto.TransactionRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "cek dto"}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "error validation"}
		json.NewEncoder(w).Encode(response)
		return
	}

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(userInfo["id"].(float64))

	orders, err := h.TransactionRepository.GetOrderByID(userID)
	
	fmt.Println(orders)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Transaction not found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	var Total = 0
	for _, i := range orders {
		Total += i.Price
	}


	dataTransaction := models.Transaction{
		Name:	request.Name,
		Email: request.Email,
		Phone: request.Phone,
		PosCode: request.PosCode,
		Address: request.Address,
		Total:     Total,
	}

	transaction, err := h.TransactionRepository.CreateTransaction(dataTransaction)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Transaction Failed!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	transactions, _ := h.TransactionRepository.GetTransaction(transaction.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: transactions}
	json.NewEncoder(w).Encode(response)
}