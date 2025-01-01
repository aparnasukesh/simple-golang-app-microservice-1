package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc Service
}

func NewHttpHandler(svc Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) Routes(r *gin.RouterGroup) {
	r.POST("/register", h.registerUser)
	r.PUT("/profile/:id", h.updateUserProfile)
	r.DELETE("/profile/:id", h.deleteUserProfile)
	r.GET("/profile/:id", h.getUserProfileDetails)
	r.POST("/methods", h.listUsers)
}

func (h *Handler) registerUser(ctx *gin.Context) {
	userData := User{}
	if err := ctx.BindJSON(&userData); err != nil {
		formattedError := ExtractErrorMessage(err)
		h.responseWithError(ctx, http.StatusBadRequest, errors.New(formattedError))
		return
	}
	if err := ValidateUser(userData); err != nil {
		formattedError := ExtractErrorMessage(err)
		h.responseWithError(ctx, http.StatusBadRequest, errors.New(formattedError))
		return
	}

	userId, err := h.svc.RegisterUser(ctx.Request.Context(), &userData)
	if err != nil {
		formattedError := ExtractErrorMessage(err)
		h.responseWithError(ctx, http.StatusNotFound, errors.New(formattedError))
		return
	}
	h.responseWithData(ctx, http.StatusOK, "signup successfull", userId)
}

func (h *Handler) getUserProfileDetails(ctx *gin.Context) {
	idstr := ctx.Param("id")
	userId, err := strconv.Atoi(idstr)
	if err != nil {
		formattedError := ExtractErrorMessage(err)
		h.responseWithError(ctx, http.StatusInternalServerError, errors.New(formattedError))
	}
	profileDetails, err := h.svc.GetUserProfileDetails(ctx, userId)
	if err != nil {
		h.responseWithError(ctx, http.StatusNotFound, errors.New("profile not found: Unable to retrieve profile details for the user"))
		return
	}

	h.responseWithData(ctx, http.StatusOK, "User profile details retrieved successfully", profileDetails)
}

func (h *Handler) updateUserProfile(ctx *gin.Context) {
	idstr := ctx.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		formattedError := ExtractErrorMessage(err)
		h.responseWithError(ctx, http.StatusInternalServerError, errors.New(formattedError))
		return
	}
	user := &UserProfileDetails{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		formattedError := ExtractErrorMessage(err)
		h.responseWithError(ctx, http.StatusBadRequest, errors.New(formattedError))
		return
	}
	err = h.svc.UpdateUserProfile(ctx, id, *user)
	if err != nil {
		formattedError := ExtractErrorMessage(err)
		h.responseWithError(ctx, http.StatusNotFound, errors.New(formattedError))
		return
	}
	h.response(ctx, http.StatusOK, "update user profile successfull")
}

func (h *Handler) listUsers(ctx *gin.Context) {
	listUserRq := ListUserRequest{}
	if err := ctx.BindJSON(&listUserRq); err != nil {
		formattedError := ExtractErrorMessage(err)
		h.responseWithError(ctx, http.StatusBadRequest, errors.New(formattedError))
		return
	}
	userNames, err := h.svc.ListUsers(ctx, listUserRq)
	if err != nil {
		formattedError := ExtractErrorMessage(err)
		h.responseWithError(ctx, http.StatusInternalServerError, errors.New(formattedError))
		return
	}
	h.responseWithData(ctx, http.StatusOK, "list users successfull", userNames)
}

func (h *Handler) deleteUserProfile(ctx *gin.Context) {
	idstr := ctx.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		formattedError := ExtractErrorMessage(err)
		h.responseWithError(ctx, http.StatusInternalServerError, errors.New(formattedError))
		return
	}
	err = h.svc.DeleteUserProfile(ctx, id)
	if err != nil {
		formattedError := ExtractErrorMessage(err)
		h.responseWithError(ctx, http.StatusInternalServerError, errors.New(formattedError))
		return
	}
	h.response(ctx, http.StatusOK, "user deleted successfully")
}
