<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { OptimizationResult, WishlistEntry } from '@/types'
import { getOptimization, getWishlist } from '@/api'

const result = ref<OptimizationResult | null>(null)
const wishlist = ref<WishlistEntry[]>([])
const loading = ref(true)

onMounted(async () => {
  const [optRes] = await Promise.all([
    getOptimization(),
    getWishlist(),
  ])
  if (optRes.data) result.value = optRes.data
  loading.value = false
})

function formatPrice(price: number) {
  return price.toLocaleString() + '원'
}
</script>

<template>
  <div class="min-h-screen pb-20">
    <header class="sticky top-0 z-40 bg-bg/80 backdrop-blur-lg border-b border-border-subtle safe-top">
      <div class="max-w-lg mx-auto px-4 py-3">
        <h1 class="text-text text-lg font-semibold">최적 조합</h1>
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

      <!-- Empty wishlist -->
      <div v-else-if="!result" class="text-center pt-16">
        <svg class="w-12 h-12 mx-auto text-text-dim mb-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
          <path d="M3 3v18h18"/><path d="m19 9-5 5-4-4-3 3"/>
        </svg>
        <p class="text-text-muted text-sm">위시리스트에 책을 추가해주세요</p>
      </div>

      <!-- Result -->
      <div v-else class="space-y-4 pb-4">
        <!-- Summary card -->
        <div class="bg-surface-elevated rounded-xl border border-brand/20 p-4 bg-brand/5">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-text-muted text-xs">총 비용</p>
              <p class="text-text text-2xl font-semibold mt-1">{{ formatPrice(result.total_cost) }}</p>
            </div>
            <div class="text-right">
              <p class="text-text-muted text-xs">배송</p>
              <p class="text-text text-lg font-medium mt-1">{{ result.ship_count }}건</p>
            </div>
          </div>
        </div>

        <!-- Items -->
        <div class="space-y-2">
          <p class="text-text text-lg font-semibold">구매 목록</p>
          <div
            v-for="(item, i) in result.items"
            :key="i"
            class="bg-surface-elevated rounded-xl border border-border p-4 flex gap-3"
          >
            <div class="w-14 h-20 shrink-0 rounded-lg overflow-hidden bg-surface">
              <img
                v-if="item.book.cover"
                :src="item.book.cover"
                :alt="item.book.title"
                class="w-full h-full object-cover"
                loading="lazy"
              />
              <div v-else class="w-full h-full flex items-center justify-center text-text-dim text-xs">
                No Cover
              </div>
            </div>
            <div class="flex-1 min-w-0">
              <h3 class="text-text text-sm font-medium line-clamp-2">{{ item.book.title }}</h3>
              <p class="text-text-muted text-xs mt-0.5">{{ item.seller }}</p>
              <div class="flex items-center gap-2 mt-1.5">
                <span class="inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium text-brand-hover bg-brand-muted">{{ item.condition }}</span>
                <span class="text-text-secondary text-sm font-medium">{{ formatPrice(item.price) }}</span>
              </div>
              <p v-if="item.delivery_fee > 0" class="text-text-dim text-xs mt-0.5">
                배송비 {{ formatPrice(item.delivery_fee) }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>
