import type { Book, WishlistEntry, OptimizationResult, ApiResult, SearchResponse } from '@/types'

const BASE = ''

async function apiCall<T>(url: string, options?: RequestInit): Promise<ApiResult<T>> {
  try {
    const res = await fetch(BASE + url, {
      headers: { 'Content-Type': 'application/json' },
      ...options,
    })
    if (!res.ok) {
      const text = await res.text().catch(() => '')
      return { error: text || `HTTP ${res.status}` }
    }
    const data = await res.json()
    return { data }
  } catch (e) {
    return { error: e instanceof Error ? e.message : 'Network error' }
  }
}

export async function searchBooks(query: string): Promise<ApiResult<Book[]>> {
  const res = await apiCall<SearchResponse>(`/api/search?q=${encodeURIComponent(query)}`)
  if (res.error) return { error: res.error }
  return { data: res.data?.books || [] }
}

export async function getWishlist(): Promise<ApiResult<WishlistEntry[]>> {
  return apiCall<WishlistEntry[]>('/api/wishlist')
}

export async function addToWishlist(bookId: string): Promise<ApiResult<WishlistEntry>> {
  return apiCall<WishlistEntry>('/api/wishlist', {
    method: 'POST',
    body: JSON.stringify({ id: bookId }),
  })
}

export async function removeFromWishlist(id: number): Promise<ApiResult<null>> {
  return apiCall<null>(`/api/wishlist/${id}`, { method: 'DELETE' })
}

export async function getOptimization(): Promise<ApiResult<OptimizationResult>> {
  return apiCall<OptimizationResult>('/api/optimize')
}

export async function postOptimization(items: string[]): Promise<ApiResult<OptimizationResult>> {
  return apiCall<OptimizationResult>('/api/optimize', {
    method: 'POST',
    body: JSON.stringify({ items }),
  })
}
