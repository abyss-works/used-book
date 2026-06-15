export interface Book {
  id: string
  title: string
  author: string
  isbn: string
  cover: string
  description?: string
  publisher?: string
  pubdate?: string
}

export interface UsedItem {
  seller: string
  price: number
  condition: string
  delivery_fee: number
  stock: number
}

export interface SearchResult {
  book: Book
  used_items: UsedItem[]
}

export interface WishlistEntry {
  id: number
  book: Book
  created_at: string
}

export interface OptimizationResult {
  total_cost: number
  ship_count: number
  items: PurchasedItem[]
}

export interface PurchasedItem {
  book: Book
  seller: string
  price: number
  condition: string
  delivery_fee: number
}

export interface ApiResult<T> {
  data?: T
  error?: string
}
