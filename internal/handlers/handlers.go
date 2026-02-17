package handlers

import (
	"context"
	"encoding/json"
	"file-manager/internal/entities"
	"file-manager/internal/logger"
	"file-manager/internal/repository"
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct {
	Ctx  context.Context
	Repo *repository.Postgres
}

type GetRecordsResponse struct {
	Records      []entities.Record `json:"records"`
	TotalRecords int               `json:"total_records"`
	Page         int               `json:"page"`
	Limit        int               `json:"limit"`
	TotalPages   int               `json:"total_pages"`
}

// GetRecordsHandler godoc
// @Summary      Get records
// @Description  Returns paginated list of records
// @Tags         record
// @Accept       json
// @Produce      json
// @Param        page      query int    false "Page number"
// @Param        limit     query int    false "Items per page"
// @Success      200 {object} GetRecordsResponse
// @Failure      500 {string} string
// @Router       /routers [get]
func (h *Handler) GetRecordsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	// GET /records?&page=1&limit=10

	appLogger := h.Ctx.Value(logger.LoggerKey).(*logger.Logger)
	appLogger.Info(h.Ctx, fmt.Sprintf("API Request: %v", query))

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil || limit <= 0 {
		limit = 1
	}

	offset := (page - 1) * limit

	records, err := h.Repo.GetRecordsWithOffset(h.Ctx, limit, offset)
	if err != nil {
		appLogger.Fatal(h.Ctx, fmt.Sprintf("GetRecordsWithOffset: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	recordsAmount, err := h.Repo.CountRecords(h.Ctx)
	if err != nil {
		appLogger.Fatal(h.Ctx, fmt.Sprintf("CountRecords: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	totalPages := recordsAmount / limit

	response := GetRecordsResponse{
		Records:      records,
		TotalRecords: recordsAmount,
		Page:         page,
		Limit:        limit,
		TotalPages:   totalPages,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
