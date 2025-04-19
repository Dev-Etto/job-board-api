package handler

import (
	"fmt"
)

type CreateOpeningRequest struct {
	Role string `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote *bool `json:"remote"`
	Link string `json:"link"`
	Salary int64 `json:"salary"`
}

type UpdateOpeningRequest struct {
	Role string `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote *bool `json:"remote"`
	Link string `json:"link"`
	Salary int64 `json:"salary"`
}


func (request * UpdateOpeningRequest) Validate() error {
	if request.Role != "" || request.Company != "" || request.Link != "" || request.Location != "" || request.Remote != nil || request.Salary > 0 {
		return nil
	}
	
	return fmt.Errorf("at least one valid field must be provided")
}

func errParamIsRequered(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}


func(request *CreateOpeningRequest) Validate() error {

	if request.Role == "" && request.Company == "" && request.Link == "" && request.Location == "" && request.Remote == nil && request.Salary <= 0 {
		return fmt.Errorf("malformed request body")
	}

	if request.Role == "" {
		return errParamIsRequered("role", "string")
	}

	if request.Company == "" {
		return errParamIsRequered("company", "string")
	}

	if request.Location == "" {
		return errParamIsRequered("location", "string")
	}

	if request.Link == "" {
		return errParamIsRequered("link", "string")
	}

	if request.Remote == nil {
		return errParamIsRequered("remote", "string")
	}

	if request.Salary <= 0  {
		return errParamIsRequered("salary", "string")
	}

	return nil
}
