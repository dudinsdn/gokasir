package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dudinsdn/gokasir/internal/usecase"
)

type ProductHandler struct {
	uc usecase.ProductUsecase
}

func NewProductHandler(uc usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("search")
	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "id"
	}
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
	if pageSize < 1 {
		pageSize = 10
	}

	products, total, err := h.uc.ListProducts(filter, sort, page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"data":       products,
		"page":       page,
		"page_size":  pageSize,
		"total":      total,
		"total_page": (total + pageSize - 1) / pageSize,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
