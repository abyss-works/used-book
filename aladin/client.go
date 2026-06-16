package aladin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/abyss-works/used-book/model"
)

const (
	defaultBaseURL = "http://www.aladin.co.kr/ttb/api"
	version        = "20131101"
	outputType     = "JS" // JSON
)

var baseURL = defaultBaseURL

func init() {
	if u := os.Getenv("ALADIN_BASE_URL"); u != "" {
		baseURL = u
	}
}


// Client wraps Aladin Open API calls.
type Client struct {
	ttbKey string
	hc     *http.Client
}

// NewClient creates a new Aladin API client.
func NewClient(ttbKey string) *Client {
	return &Client{
		ttbKey: ttbKey,
		hc:     &http.Client{},
	}
}

// SearchBooks searches for books by title/author.
func (c *Client) SearchBooks(query string, max int, start int) (*model.AladinSearchResult, error) {
	if max <= 0 || max > 100 {
		max = 10
	}
	if start <= 0 {
		start = 1
	}

	params := url.Values{}
	params.Set("ttbkey", c.ttbKey)
	params.Set("Query", query)
	params.Set("QueryType", "Title")
	params.Set("MaxResults", fmt.Sprintf("%d", max))
	params.Set("start", fmt.Sprintf("%d", start))
	params.Set("SearchTarget", "Book")
	params.Set("output", outputType)
	params.Set("Version", version)

	u := fmt.Sprintf("%s/ItemSearch.aspx?%s", baseURL, params.Encode())
	resp, err := c.hc.Get(u)
	if err != nil {
		return nil, fmt.Errorf("aladin search request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("aladin search read: %w", err)
	}

	// 알라딘 API는 JS(JSON) 형식에서 최상위 {} 없이 배열로 오기도 함
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("aladin search parse: %w", err)
	}

	var total int
	var alaItems []struct {
		ItemID  int    `json:"itemId"`
		Title   string `json:"title"`
		Author  string `json:"author"`
		Isbn13  string `json:"isbn13"`
		Cover   string `json:"cover"`
		MallType string `json:"mallType"`
	}

	// totalResults, item 배열 찾기
	if totalRaw, ok := raw["totalResults"]; ok {
		var tr struct {
			Total int `json:"total"`
		}
		if err := json.Unmarshal(totalRaw, &tr); err == nil {
			total = tr.Total
		}
	}
	if itemsRaw, ok := raw["item"]; ok {
		if err := json.Unmarshal(itemsRaw, &alaItems); err != nil {
			return nil, fmt.Errorf("aladin search items parse: %w", err)
		}
	}

	books := make([]model.Book, 0, len(alaItems))
	for _, ai := range alaItems {
		books = append(books, model.Book{
			ID:     fmt.Sprintf("%d", ai.ItemID),
			Title:  ai.Title,
			Author: ai.Author,
			Isbn:   ai.Isbn13,
			Cover:  ai.Cover,
		})
	}

	return &model.AladinSearchResult{
		Books:  books,
		Total:  total,
		Start:  start,
		Max:    max,
	}, nil
}

// 알라딘 usedList 응답 구조 (3개 카테고리 요약)
type aladinUsedCategory struct {
	ItemCount int    `json:"itemCount"`
	MinPrice  int    `json:"minPrice"`
	Link      string `json:"link"`
}

type aladinUsedList struct {
	AladinUsed *aladinUsedCategory `json:"aladinUsed"`
	UserUsed   *aladinUsedCategory `json:"userUsed"`
	SpaceUsed  *aladinUsedCategory `json:"spaceUsed"`
}

type aladinItem struct {
	ItemID  int    `json:"itemId"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Isbn13  string `json:"isbn13"`
	Cover   string `json:"cover"`
	SubInfo *struct {
		UsedList *aladinUsedList `json:"usedList"`
	} `json:"subInfo"`
}

// LookupUsed retrieves used book listings for a specific book.
func (c *Client) LookupUsed(bookID string) (*model.AladinUsedResult, error) {
	params := url.Values{}
	params.Set("ttbkey", c.ttbKey)
	params.Set("itemIdType", "ItemId")
	params.Set("ItemId", bookID)
	params.Set("output", outputType)
	params.Set("Version", version)
	params.Set("OptResult", "usedList")

	u := fmt.Sprintf("%s/ItemLookUp.aspx?%s", baseURL, params.Encode())
	resp, err := c.hc.Get(u)
	if err != nil {
		return nil, fmt.Errorf("aladin lookup request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("aladin lookup read: %w", err)
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("aladin lookup parse: %w", err)
	}

	itemsRaw, ok := raw["item"]
	if !ok {
		return nil, fmt.Errorf("aladin lookup: no item in response")
	}

	// item 배열 파싱 시도
	var items []aladinItem
	if err := json.Unmarshal(itemsRaw, &items); err != nil || len(items) == 0 {
		// 단일 객체로 왔을 경우
		var single aladinItem
		if err := json.Unmarshal(itemsRaw, &single); err != nil || single.ItemID == 0 {
			return nil, fmt.Errorf("aladin lookup: item parse failed: %w", err)
		}
		items = []aladinItem{single}
	}

	alaItem := items[0]
	book := model.Book{
		ID:     fmt.Sprintf("%d", alaItem.ItemID),
		Title:  alaItem.Title,
		Author: alaItem.Author,
		Isbn:   alaItem.Isbn13,
		Cover:  alaItem.Cover,
	}

	var used []model.UsedItem
	categories := []struct {
		key  string
		name string
		cond string
	}{
		{"aladinUsed", "알라딘 중고", "상"},
		{"userUsed", "개인 판매자", "중"},
		{"spaceUsed", "알라딘 매장", "상"},
	}

	if alaItem.SubInfo != nil && alaItem.SubInfo.UsedList != nil {
		ul := alaItem.SubInfo.UsedList
		// category pointer map을 만들어서 key로 조회
		catMap := map[string]*aladinUsedCategory{
			"aladinUsed": ul.AladinUsed,
			"userUsed":   ul.UserUsed,
			"spaceUsed":  ul.SpaceUsed,
		}
		for _, cat := range categories {
			if c := catMap[cat.key]; c != nil && c.ItemCount > 0 {
				used = append(used, model.UsedItem{
					SellerName:  cat.name,
					Price:       c.MinPrice,
					Condition:   cat.cond,
					Stock:       c.ItemCount,
					Link:        c.Link,
				})
			}
		}
	}

	return &model.AladinUsedResult{
		Book:  book,
		Items: used,
	}, nil
}
