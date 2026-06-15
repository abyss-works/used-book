<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import SearchBar from '@/components/SearchBar.vue'
import BookCard from '@/components/BookCard.vue'
import type { Book } from '@/types'
import { searchBooks } from '@/api'

const router = useRouter()
const query = ref('')
const results = ref<Book[]>([])
const loading = ref(false)
const searched = ref(false)

async function onSearch(q: string) {
  query.value = q
  loading.value = true
  searched.value = true
  try {
    const res = await searchBooks(q)
    results.value = res.data || []
  } catch {
    results.value = []
  } finally {
    loading.value = false
  }
}

function onBookClick(book: Book) {
  router.push({ name: 'book-detail', params: { id: book.id } })
}
</script>

<template>
  <div class="min-h-screen pb-20">
    <!-- Header -->
    <header class="sticky top-0 z-40 bg-bg/80 backdrop-blur-lg border-b border-border-subtle safe-top">
      <div class="max-w-lg mx-auto px-4 py-3">
        <h1 class="text-text text-lg font-semibold mb-3">중고책 검색</h1>
        <SearchBar
          :model-value="query"
          @update:model-value="query = $event"
          @search="onSearch"
          placeholder="찾는 책 제목을 입력하세요"
        />
      </div>
    </header>

    <!-- Content -->
    <main class="max-w-lg mx-auto px-4 pt-4">
      <!-- Initial state -->
      <div v-if="!searched" class="text-center pt-16">
        <svg class="w-12 h-12 mx-auto text-text-dim mb-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="11" cy="11" r="8"/>
          <path d="m21 21-4.35-4.35"/>
        </svg>
        <p class="text-text-muted text-sm">찾는 책을 검색해보세요</p>
        <p class="text-text-dim text-xs mt-1">알라딘 중고책 매장 검색</p>
      </div>

      <!-- Loading -->
      <div v-else-if="loading" class="space-y-3 pt-2">
        <div v-for="i in 3" :key="i" class="bg-surface-elevated rounded-xl border border-border p-4 flex gap-3">
          <div class="w-16 h-24 rounded-lg skeleton"/>
          <div class="flex-1 space-y-2 py-1">
            <div class="h-4 skeleton rounded w-3/4"/>
            <div class="h-3 skeleton rounded w-1/2"/>
          </div>
        </div>
      </div>

      <!-- Results -->
      <div v-else-if="results.length > 0" class="space-y-3 pb-4">
        <p class="text-text-muted text-xs">{{ results.length }}건 검색됨</p>
        <BookCard
          v-for="r in results"
          :key="r.id"
          :book="r"
          @click="onBookClick"
        />
      </div>

      <!-- Empty -->
      <div v-else class="text-center pt-16">
        <p class="text-text-muted text-sm">검색 결과가 없습니다</p>
      </div>
    </main>
  </div>
</template>
