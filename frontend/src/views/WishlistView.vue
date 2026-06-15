<script setup lang="ts">
import { ref, onMounted } from 'vue'
import BookCard from '@/components/BookCard.vue'
import type { Book, WishlistEntry } from '@/types'
import { getWishlist, removeFromWishlist } from '@/api'

const entries = ref<WishlistEntry[]>([])
const loading = ref(true)

onMounted(async () => {
  const res = await getWishlist()
  if (res.data) entries.value = res.data
  loading.value = false
})

async function onRemove(entry: WishlistEntry) {
  await removeFromWishlist(entry.id)
  entries.value = entries.value.filter(e => e.id !== entry.id)
}

function onBookClick(book: Book) {
  // book detail
}
</script>

<template>
  <div class="min-h-screen pb-20">
    <header class="sticky top-0 z-40 bg-bg/80 backdrop-blur-lg border-b border-border-subtle safe-top">
      <div class="max-w-lg mx-auto px-4 py-3">
        <h1 class="text-text text-lg font-semibold">위시리스트</h1>
      </div>
    </header>

    <main class="max-w-lg mx-auto px-4 pt-4">
      <!-- Loading -->
      <div v-if="loading" class="space-y-3 pt-2">
        <div v-for="i in 3" :key="i" class="bg-surface-elevated rounded-xl border border-border p-4 flex gap-3">
          <div class="w-16 h-24 rounded-lg skeleton"/>
          <div class="flex-1 space-y-2 py-1">
            <div class="h-4 skeleton rounded w-3/4"/>
            <div class="h-3 skeleton rounded w-1/2"/>
          </div>
        </div>
      </div>

      <!-- Empty -->
      <div v-else-if="entries.length === 0" class="text-center pt-16">
        <svg class="w-12 h-12 mx-auto text-text-dim mb-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
          <path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/>
        </svg>
        <p class="text-text-muted text-sm">위시리스트가 비어 있습니다</p>
        <p class="text-text-dim text-xs mt-1">검색에서 책을 추가해보세요</p>
      </div>

      <!-- List -->
      <div v-else class="space-y-3 pb-4">
        <p class="text-text-muted text-xs">{{ entries.length }}권</p>
        <div v-for="entry in entries" :key="entry.id" class="group relative">
          <BookCard :book="entry.book" @click="onBookClick" />
          <button
            class="absolute top-2 right-2 w-8 h-8 flex items-center justify-center rounded-full bg-danger/10 text-danger opacity-0 group-hover:opacity-100 transition-opacity duration-150"
            @click="onRemove(entry)"
          >
            <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M3 6h18"/><path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"/>
              <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"/>
            </svg>
          </button>
        </div>
      </div>
    </main>
  </div>
</template>
