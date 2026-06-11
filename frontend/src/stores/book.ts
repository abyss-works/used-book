import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Book, OptimizationResult } from '@/api'
import { searchBooks, optimizeWishlist } from '@/api'

export const useBookStore = defineStore('book', () => {
  const wishlist = ref<Book[]>([])
  const result = ref<OptimizationResult | null>(null)
  const loading = ref(false)

  async function search(query: string, max = 10, start = 1) {
    loading.value = true
    try {
      return await searchBooks(query, max, start)
    } finally {
      loading.value = false
    }
  }

  function addToWishlist(book: Book) {
    if (!wishlist.value.find((b) => b.id === book.id)) {
      wishlist.value.push(book)
    }
  }

  function removeFromWishlist(bookId: string) {
    wishlist.value = wishlist.value.filter((b) => b.id !== bookId)
  }

  async function optimize() {
    loading.value = true
    try {
      const items = wishlist.value.map((b) => ({
        book_id: b.id,
        title: b.title,
        author: b.author,
      }))
      result.value = await optimizeWishlist(items)
    } finally {
      loading.value = false
    }
  }

  return { wishlist, result, loading, search, addToWishlist, removeFromWishlist, optimize }
})
