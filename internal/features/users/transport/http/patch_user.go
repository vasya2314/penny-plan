package users_transport_http

import (
	"net/http"

	"github.com/vasya2314/penny-plan/internal/core/domain"
	core_logger "github.com/vasya2314/penny-plan/internal/core/logger"
	core_http_request "github.com/vasya2314/penny-plan/internal/core/transport/http/request"
	core_http_response "github.com/vasya2314/penny-plan/internal/core/transport/http/response"
	core_http_types "github.com/vasya2314/penny-plan/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name" validate:"omitempty,min=3,max=100"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type PatchUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failde to get userID path value",
		)

		return
	}

	var request PatchUserRequest
	err = core_http_request.DecodeAndValidateRequest(r, &request)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)

		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.FullName.ToDomain(),
		request.PhoneNumber.ToDomain(),
	)
}
