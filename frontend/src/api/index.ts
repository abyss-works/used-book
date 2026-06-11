import axios from 'axios'

const api = axios.create({
  baseURL: '/api',
})

export interface Book {
  id: string
  title: string
  author: string
  cover: string
}

export interface SearchResult {
  books: Book[]
  total: number
  start: number
  max: number
}

export interface OptimizationResult {
  total_cost: number
  ship_count: number
  sellers: number
  purchases: Array<{
    title: string
    author: string
    seller_name: string
    price: number
    condition: string
    delivery_fee: number
  }>
}

export async function searchBooks(
  query: string,
  max = 10,
  start = 1,
): Promise<SearchResult> {
  const res = await api.get('/search', {
    params: { q: query, max, start },
  })
  return res.data
}

export async function optimizeWishlist(
  wishlist: Array<{ book_id: string; title: string; author: string }>,
): Promise<OptimizationResult> {
  const res = await api.post('/optimize', { wishlist })
  return res.data
}
