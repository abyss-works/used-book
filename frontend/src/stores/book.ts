import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Book, OptimizationResult } from '@/types'
import { getWishlist as apiGetWishlist, removeFromWishlist as apiRemoveFromWishlist } from '@/api'

export const useBookStore = defineStore('book', () => {
  const wishlist = ref<Book[]>([])
  const result = ref<OptimizationResult | null>(null)
  const loading = ref(false)

  async function loadWishlist() {
    loading.value = true
    try {
      const res = await apiGetWishlist()
      if (res.data) {
        wishlist.value = res.data.map(e => e.book)
      }
    } finally {
      loading.value = false
    }
  }

  function addToWishlist(book: Book) {
    if (!wishlist.value.find(b => b.id === book.id)) {
      wishlist.value.push(book)
    }
  }

  async function removeFromWishlist(bookId: string) {
    wishlist.value = wishlist.value.filter(b => b.id !== bookId)
  }

  return { wishlist, result, loading, loadWishlist, addToWishlist, removeFromWishlist }
})
