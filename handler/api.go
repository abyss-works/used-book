package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/abyss-works/used-book/aladin"
	"github.com/abyss-works/used-book/model"
	"github.com/abyss-works/used-book/optimizer"
)

type Handler struct {
	aladin *aladin.Client
}

func New(aladinClient *aladin.Client) *Handler {
	return &Handler{aladin: aladinClient}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/api/search", h.handleSearch)
	mux.HandleFunc("/api/lookup", h.handleLookup)
	mux.HandleFunc("/api/optimize", h.handleOptimize)
	mux.HandleFunc("/api/health", h.handleHealth)
}

// GET /api/search?q={query}&max=10&start=1
func (h *Handler) handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "GET만 허용")
		return
	}

	q := r.URL.Query().Get("q")
	if strings.TrimSpace(q) == "" {
		writeError(w, http.StatusBadRequest, "검색어(q)가 필요합니다")
		return
	}

	max, _ := strconv.Atoi(r.URL.Query().Get("max"))
	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	result, err := h.aladin.SearchBooks(q, max, start)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "알라딘 검색 실패: "+err.Error())
		return
	}

	writeJSON(w, http.StatusOK, result)
}

// GET /api/lookup?id={bookId}
func (h *Handler) handleLookup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "GET만 허용")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "도서 ID(id)가 필요합니다")
		return
	}

				result, err := h.aladin.LookupUsed(id)
				if err != nil {
					writeError(w, http.StatusInternalServerError, "알라딘 조회 실패: "+err.Error())
					return
				}

				writeJSON(w, http.StatusOK, result)
}

// POST /api/optimize
// Body: { "wishlist": [ { "book_id": "...", "title": "...", ... } ] }
// 또는 GET /api/optimize?ids=bookId1,bookId2&titles=title1,title2
func (h *Handler) handleOptimize(w http.ResponseWriter, r *http.Request) {
	var entries []model.WishlistEntry

	if r.Method == http.MethodPost {
		var req model.SearchRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "JSON 파싱 실패: "+err.Error())
			return
		}
		entries = req.Wishlist
	} else if r.Method == http.MethodGet {
		ids := strings.Split(r.URL.Query().Get("ids"), ",")
		titles := strings.Split(r.URL.Query().Get("titles"), ",")
		for i, id := range ids {
			id = strings.TrimSpace(id)
			if id == "" {
				continue
			}
			title := ""
			if i < len(titles) {
				title = strings.TrimSpace(titles[i])
			}
			entries = append(entries, model.WishlistEntry{
				BookID: id,
				Title:  title,
			})
		}
	} else {
		writeError(w, http.StatusMethodNotAllowed, "GET 또는 POST만 허용")
		return
	}

	if len(entries) == 0 {
		writeError(w, http.StatusBadRequest, "최소 1개의 위시리스트 항목이 필요합니다")
		return
	}

	// 각 책의 중고 정보 조회
	bookMap := make(map[string]*model.AladinUsedResult)
	for _, entry := range entries {
		if _, ok := bookMap[entry.BookID]; ok {
			continue
		}
		result, err := h.aladin.LookupUsed(entry.BookID)
		if err != nil {
			writeError(w, http.StatusInternalServerError,
				"도서 조회 실패 ("+entry.Title+"): "+err.Error())
			return
		}
		bookMap[entry.BookID] = result
	}

	// 최적화 실행
	result := optimizer.OptimizeAll(entries, bookMap)

	writeJSON(w, http.StatusOK, result)
}

// GET /api/health
func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
