package decoder

import (
	"context"
	dto "demo-go-kit/account/controller/dto"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func DecodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeEmailReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetUserRequest
	vars := mux.Vars(r)

	req = dto.GetUserRequest{
		Id: vars["id"],
	}
	return req, nil
}
