<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import type { Book, UsedItem } from '@/types'
import { lookupBook, addToWishlist } from '@/api'

const route = useRoute()
const router = useRouter()
const bookId = route.params.id as string

const book = ref<Book | null>(null)
const items = ref<UsedItem[]>([])
const loading = ref(true)
const added = ref(false)

onMounted(async () => {
  const res = await lookupBook(bookId)
  if (res.data) {
    book.value = res.data.book
    items.value = res.data.items || []
  }
  loading.value = false
})

function formatPrice(p: number) {
  return p.toLocaleString() + '원'
}

function conditionColor(c: string) {
  switch (c) {
    case '최상': return 'text-success bg-success/12'
    case '상': return 'text-brand bg-brand-muted'
    case '중': return 'text-warning bg-warning/12'
    default: return 'text-text-muted bg-white/5'
  }
}

async function addToWish() {
  if (!book.value || added.value) return
  await addToWishlist(book.value.id)
  added.value = true
}
</script>

<template>
  <div class="min-h-screen pb-20">
    <!-- Header -->
    <header class="sticky top-0 z-40 bg-bg/80 backdrop-blur-lg border-b border-border-subtle safe-top">
      <div class="max-w-lg mx-auto px-4 py-3 flex items-center gap-3">
        <button
          class="bg-transparent text-text-secondary font-medium rounded-sm px-2 py-1 hover:bg-white/5 transition-colors duration-150 cursor-pointer"
          @click="router.back()"
        >
          <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="m15 18-6-6 6-6"/>
          </svg>
        </button>
        <h1 class="text-text text-lg font-semibold truncate">책 상세</h1>
      </div>
    </header>

    <main class="max-w-lg mx-auto px-4 pt-4">
      <!-- Loading -->
      <div v-if="loading" class="space-y-4 pt-4">
        <div class="flex gap-4">
          <div class="w-24 h-36 skeleton rounded-xl"/>
          <div class="flex-1 space-y-3">
            <div class="h-5 skeleton rounded w-3/4"/>
            <div class="h-4 skeleton rounded w-1/2"/>
            <div class="h-3 skeleton rounded w-2/3"/>
          </div>
        </div>
      </div>

      <!-- Content -->
      <template v-else-if="book">
        <!-- Book info -->
        <div class="flex gap-4 pb-4 border-b border-border-subtle">
          <div class="w-24 h-36 shrink-0 rounded-xl overflow-hidden bg-surface">
            <img
              v-if="book.cover"
              :src="book.cover"
              :alt="book.title"
              class="w-full h-full object-cover"
            />
            <div v-else class="w-full h-full flex items-center justify-center text-text-dim text-xs">
              No Cover
            </div>
          </div>
          <div class="flex-1">
            <h1 class="text-text text-lg font-semibold leading-snug">{{ book.title }}</h1>
            <p class="text-text-muted text-sm mt-1">{{ book.author }}</p>
            <p v-if="book.publisher" class="text-text-dim text-xs mt-1">{{ book.publisher }}</p>
            <p class="text-text-dim text-xs mt-0.5">ISBN: {{ book.isbn }}</p>
            <button
              v-if="!added"
              class="bg-brand text-bg font-semibold rounded-sm px-5 py-2.5 active:scale-[0.97] transition-all duration-150 cursor-pointer"
              @click="addToWish"
            >
              위시리스트 담기
            </button>
            <span v-else class="inline-flex items-center gap-1 mt-3 text-success text-sm font-medium">
              <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M20 6 9 17l-5-5"/>
              </svg>
              위시리스트에 추가됨
            </span>
          </div>
        </div>

        <!-- Used items -->
        <div class="pt-4">
          <h2 class="text-text text-base font-semibold mb-3">중고 매장 {{ items.length > 0 ? `(${items.length})` : '' }}</h2>

          <div v-if="items.length === 0" class="text-center pt-8">
            <p class="text-text-muted text-sm">중고 매물이 없습니다</p>
            <p class="text-text-dim text-xs mt-1">알라딘 중고 매장에 등록된 상품이 없습니다</p>
          </div>

          <div v-else class="space-y-2">
            <div
              v-for="(item, i) in items"
              :key="i"
              class="bg-surface-elevated rounded-xl border border-border p-4"
            >
              <div class="flex items-start justify-between">
                <div>
                  <p class="text-text text-sm font-medium">{{ item.sellerName }}</p>
                  <div class="flex items-center gap-2 mt-1.5">
                    <span class="inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium"
                      :class="conditionColor(item.condition)">
                      {{ item.condition }}
                    </span>
                    <span class="text-text-secondary text-sm font-semibold">{{ formatPrice(item.price) }}</span>
                  </div>
                  <p v-if="item.deliveryFee > 0" class="text-text-dim text-xs mt-1">
                    배송비 {{ formatPrice(item.deliveryFee) }}
                  </p>
                  <p v-else class="text-text-dim text-xs mt-1">무료배송</p>
                </div>
                <div class="text-right shrink-0">
                  <p class="text-text-dim text-xs">재고 {{ item.stockCount }}</p>
                  <a
                    v-if="item.link"
                    :href="item.link"
                    target="_blank"
                    rel="noopener"
                    class="inline-flex items-center gap-1 mt-2 text-brand text-xs font-medium hover:underline"
                  >
                    구매하기
                    <svg class="w-3 h-3" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M7 7h10v10"/><path d="M7 17 21 3"/>
                    </svg>
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </template>

      <!-- Not found -->
      <div v-else class="text-center pt-16">
        <p class="text-text-muted text-sm">책 정보를 불러올 수 없습니다</p>
      </div>
    </main>
  </div>
</template>
