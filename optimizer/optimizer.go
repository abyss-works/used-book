package optimizer

import (
	"math"
	"sort"

	"github.com/abyss-works/used-book/model"
)

// Config holds optimizer configuration.
type Config struct {
	MaxPerSeller int // 최대 판매자 수 (0 = 제한 없음)
}

// DefaultConfig returns the default optimizer configuration.
func DefaultConfig() Config {
	return Config{
		MaxPerSeller: 0,
	}
}

// Candidate represents one purchase option for a book.
type Candidate struct {
	BookIndex   int    // 위시리스트 내 인덱스
	BookTitle   string
	BookAuthor  string
	SellerName  string
	Price       int
	Condition   string
	DeliveryFee int
}

// Solution represents one possible combination of purchases.
type Solution struct {
	Selections []Candidate
	TotalCost  int
	ShipCount  int // 배송이 발생하는 판매자 수
	Sellers    int // 고유 판매자 수
}

// Optimize finds the optimal combination of used book purchases.
// wishlist: 구매하려는 책 목록
// bookMap: 책 ID → 알라딘 중고 결과 (LookupUsed 결과)
func Optimize(wishlist []model.WishlistEntry, bookMap map[string]*model.AladinUsedResult, cfg Config) *model.OptimizationResult {
	if cfg.MaxPerSeller == 0 {
		cfg.MaxPerSeller = len(wishlist) // 제한 없음
	}

	// 각 책별 구매 후보 생성
	candidates := buildCandidates(wishlist, bookMap)
	if len(candidates) == 0 {
		return &model.OptimizationResult{
			TotalCost: 0,
			Purchases: nil,
		}
	}

	// 최적 조합 탐색
	best := findBestCombination(candidates, cfg)

	// 결과 변환
	result := &model.OptimizationResult{
		TotalCost: best.TotalCost,
		ShipCount: best.ShipCount,
		Sellers:   best.Sellers,
		Purchases: make([]model.PurchasedItem, 0, len(best.Selections)),
	}

	sellerMap := make(map[string]int)
	for _, sel := range best.Selections {
		sellerMap[sel.SellerName]++
		result.Purchases = append(result.Purchases, model.PurchasedItem{
			Title:       sel.BookTitle,
			Author:      sel.BookAuthor,
			SellerName:  sel.SellerName,
			Price:       sel.Price,
			Condition:   sel.Condition,
			DeliveryFee: sel.DeliveryFee,
		})
	}
	result.Sellers = len(sellerMap)

	return result
}

func buildCandidates(wishlist []model.WishlistEntry, bookMap map[string]*model.AladinUsedResult) [][]Candidate {
	all := make([][]Candidate, len(wishlist))
	for i, entry := range wishlist {
		result, ok := bookMap[entry.BookID]
		if !ok || result == nil || len(result.Items) == 0 {
			continue // 이 책은 판매중인 중고가 없음
		}
		// 판매자별 후보 생성
		seen := make(map[string]bool) // 중복 판매자 제거
		for _, item := range result.Items {
			if item.Stock <= 0 {
				continue
			}
			key := item.SellerID + "|" + item.Condition
			if seen[key] {
				continue
			}
			seen[key] = true
			all[i] = append(all[i], Candidate{
				BookIndex:   i,
				BookTitle:   entry.Title,
				BookAuthor:  entry.Author,
				SellerName:  item.SellerName,
				Price:       item.Price,
				Condition:   item.Condition,
				DeliveryFee: item.DeliveryFee,
			})
		}
		// 가격 오름차순 정렬
		sort.Slice(all[i], func(a, b int) bool {
			return all[i][a].Price < all[i][b].Price
		})
	}
	return all
}

func findBestCombination(candidates [][]Candidate, cfg Config) Solution {
	// 책 수
	n := len(candidates)

	// 각 책별 candidate 수
	sizes := make([]int, n)
	totalCombos := 1
	for i, c := range candidates {
		sizes[i] = len(c)
		if sizes[i] == 0 {
			totalCombos = 0 // 이 책은 구매 불가
		} else {
			totalCombos *= sizes[i]
		}
	}

	if totalCombos == 0 {
		return Solution{}
	}

	if totalCombos > 100000 {
		// 조합이 너무 많으면 휴리스틱 사용
		return findBestCombinationHeuristic(candidates, cfg)
	}

	// 모든 조합 탐색 (완전 탐색)
	best := Solution{}
	bestSet := false
	indices := make([]int, n)

	for {
		// 현재 조합 평가
		sel := evaluate(indices, candidates, cfg)
		cost := totalCost(sel)
		if !bestSet || cost < best.TotalCost || (cost == best.TotalCost && len(sel) > len(best.Selections)) {
			best = Solution{
				Selections: copySelections(sel),
				TotalCost:  cost,
				ShipCount:  countSellers(sel),
				Sellers:    countSellers(sel),
			}
			bestSet = true
		}

		// 다음 조합
		carry := 1
		for i := n - 1; i >= 0 && carry > 0; i-- {
			indices[i]++
			if indices[i] >= sizes[i] {
				indices[i] = 0
				carry = 1
			} else {
				carry = 0
			}
		}
		if carry > 0 {
			break
		}
	}

	return best
}

func findBestCombinationHeuristic(candidates [][]Candidate, cfg Config) Solution {
	n := len(candidates)
	best := Solution{}
	bestSet := false

	// 각 책별 최저가를 기본으로 시작
	sel := make([]Candidate, 0, n)
	for _, cands := range candidates {
		if len(cands) > 0 {
			sel = append(sel, cands[0])
		}
	}
	best = Solution{
		Selections: copySelections(sel),
		TotalCost:  totalCost(sel),
		ShipCount:  countSellers(sel),
		Sellers:    countSellers(sel),
	}
	bestSet = true

	// 같은 판매자로 통합 시도 (배송비 절약)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if len(candidates[i]) == 0 || len(candidates[j]) == 0 {
				continue
			}
			// i와 j를 같은 판매자의 상품으로 교체 시도
			ci := candidates[i][0]
			cj := candidates[j][0]
			ci.SellerName = cj.SellerName
			sel[i] = ci
			cost := totalCost(sel)
			if cost < best.TotalCost {
				best = Solution{
					Selections: copySelections(sel),
					TotalCost:  cost,
					ShipCount:  countSellers(sel),
					Sellers:    countSellers(sel),
				}
			}
		}
	}

	return best
}

func evaluate(indices []int, candidates [][]Candidate, cfg Config) []Candidate {
	n := len(indices)
	sel := make([]Candidate, 0, n)
	for i := range indices {
		if indices[i] < len(candidates[i]) {
			sel = append(sel, candidates[i][indices[i]])
		}
	}
	return sel
}

func totalCost(sel []Candidate) int {
	// 판매자별 배송비 누적
	type sellerKey struct{ name string }
	deliveryPaid := make(map[string]bool)
	total := 0
	for _, c := range sel {
		total += c.Price
		key := c.SellerName
		if !deliveryPaid[key] {
			total += c.DeliveryFee
			deliveryPaid[key] = true
		}
	}
	return total
}

func countSellers(sel []Candidate) int {
	seen := make(map[string]bool)
	for _, c := range sel {
		seen[c.SellerName] = true
	}
	return len(seen)
}

func copySelections(sel []Candidate) []Candidate {
	out := make([]Candidate, len(sel))
	copy(out, sel)
	return out
}

// OptimizeAll tries different permutations and returns the best one.
func OptimizeAll(wishlist []model.WishlistEntry, bookMap map[string]*model.AladinUsedResult) *model.OptimizationResult {
	best := Optimize(wishlist, bookMap, DefaultConfig())
	return best
}

// RoundRobinShippingCost calculates average shipping for estimation.
func RoundRobinShippingCost(items []model.UsedItem, perSellerMax int) int {
	if len(items) == 0 {
		return 0
	}
	sellerFees := make(map[string]int)
	for _, item := range items {
		sellerFees[item.SellerName] = item.DeliveryFee
	}
	total := 0
	count := 0
	for _, fee := range sellerFees {
		total += fee
		count++
	}
	if count == 0 {
		return 0
	}
	return int(math.Ceil(float64(total) / float64(count)))
}
