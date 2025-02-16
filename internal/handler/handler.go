package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/heismyke/local_business_booking_app/internal/service"
	"github.com/heismyke/local_business_booking_app/types"
)


type Handler struct{
  svc service.Service
}

func New(service service.Service) *Handler{
  return &Handler{
    svc : service,
  }
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request){
  var req types.RegisterUser

  err := json.NewDecoder(r.Body).Decode(&req)
  if err != nil {
    http.Error(w, "Invalid reqest payload", http.StatusBadRequest)
    return
  }

  fmt.Println(req)
  
  user, err := h.svc.CreateUserService(r.Context(), &req)

  if err != nil{
    http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
  }

  resp, err := json.Marshal(&user)
  if err != nil {
    http.Error(w, "Error encoding response", http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  w.Write(resp)


}
