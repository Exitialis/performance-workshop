package handlers

import (
	"context"
	"encoding/json"
	phones_gateway "github.com/exitialis/workshop/homework/complex/internal/gateway"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	phones phones_gateway.PhonesGateway
}

type Request struct {
	PhoneID int64
	UserID int64
	CategoryID int64
	ItemID int64
	PhoneDisplayLoc string
}

type Response struct {
	Phone string
	Type int64
}

func New(
	phonesService phones_gateway.PhonesGateway,
) *Handler {
	return &Handler{
		phones: phonesService,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-type", "application/json")
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var request Request
	err = json.Unmarshal(body, &request)
	if err != nil {
		h.writeError(err, w)
		return
	}

	getPhoneIn := phones_gateway.GetPhoneIn{
		PhoneID:         request.PhoneID,
		UserID:          request.UserID,
		CategoryID:      request.CategoryID,
		ItemID:          request.ItemID,
		PhoneDisplayLoc: request.PhoneDisplayLoc,
	}

	err = h.phones.Validate(getPhoneIn)

	if err != nil {
		h.writeError(err, w)
		return
	}

	response, err := h.phones.GetPhone(ctx, getPhoneIn)

	if err != nil {
		h.writeError(err, w)
		return
	}

	out := Response{
		Phone: response.Phone,
		Type:  response.Type,
	}

	respBytes, err := json.Marshal(out)
	if err != nil {
		h.writeError(err, w)
		return
	}

	_, _ = w.Write(respBytes)
}

type ErrorResponse struct {
	Code int64
	Error string
}

func (h *Handler) writeError(err error, w http.ResponseWriter) {
	resp := ErrorResponse{
		Code:  500,
		Error: err.Error(),
	}

	r, err := json.Marshal(resp)
	if err != nil {
		rr, _ := json.Marshal(ErrorResponse{
			Code:  500,
			Error: "Internal error",
		})

		_, _ = w.Write(rr)
	}

	_, _ = w.Write(r)
}
