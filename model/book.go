package model

// Book represents a book from Aladin API search results.
type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Isbn     string `json:"isbn"`
	Cover    string `json:"cover"`
	Category string `json:"category"`
}

// UsedItem represents a single used book listing from a seller.
type UsedItem struct {
	SellerID      string `json:"seller_id"`
	SellerName    string `json:"seller_name"`
	Price         int    `json:"price"`
	Condition     string `json:"condition"` // 최상 / 상 / 중
	DeliveryFee   int    `json:"delivery_fee"`
	Stock         int    `json:"stock"`
	Link          string `json:"link"`
}

// Seller summary used by the optimizer.
type Seller struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ItemCount int    `json:"item_count"`
}

// WishlistEntry is a book the user wants to buy.
type WishlistEntry struct {
	ID     int    `json:"id"`
	BookID string `json:"book_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// PurchasedItem is part of an optimization solution.
type PurchasedItem struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	SellerName  string `json:"seller_name"`
	Price       int    `json:"price"`
	Condition   string `json:"condition"`
	DeliveryFee int    `json:"delivery_fee"`
}

// OptimizationResult is the final output of the optimizer.
type OptimizationResult struct {
	TotalCost  int             `json:"total_cost"`
	ShipCount  int             `json:"ship_count"`
	Sellers    int             `json:"sellers"`
	Purchases  []PurchasedItem `json:"purchases"`
}

// SearchRequest for POST /api/optimize
type SearchRequest struct {
	Wishlist []WishlistEntry `json:"wishlist"`
}

// AladinSearchResult from the external API
type AladinSearchResult struct {
	Books  []Book `json:"books"`
	Total  int    `json:"total"`
	Start  int    `json:"start"`
	Max    int    `json:"max"`
}

// AladinUsedResult from ItemLookUp + usedList
type AladinUsedResult struct {
	Book  Book       `json:"book"`
	Items []UsedItem `json:"items"`
}
