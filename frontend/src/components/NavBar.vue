<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

const tabs = [
  { name: '검색', path: '/', icon: 'search' },
  { name: '위시리스트', path: '/wishlist', icon: 'heart' },
  { name: '최적화', path: '/optimize', icon: 'chart' },
]

function isActive(path: string) {
  return route.path === path
}

function go(path: string) {
  router.push(path)
}
</script>

<template>
  <nav class="fixed bottom-0 left-0 right-0 z-50 bg-surface border-t border-border safe-bottom">
    <div class="flex items-center justify-around h-16 max-w-lg mx-auto">
      <button
        v-for="tab in tabs"
        :key="tab.path"
        :class="[
          'flex flex-col items-center gap-0.5 text-[10px] font-medium transition-colors duration-150 cursor-pointer',
          isActive(tab.path) ? 'text-brand' : 'text-text-dim',
        ]"
        @click="go(tab.path)"
      >
        <!-- Search -->
        <svg v-if="tab.icon === 'search'" class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="11" cy="11" r="8"/>
          <path d="m21 21-4.35-4.35"/>
        </svg>
        <!-- Heart -->
        <svg v-else-if="tab.icon === 'heart'" class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/>
        </svg>
        <!-- Chart -->
        <svg v-else class="w-6 h-6" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M3 3v18h18"/>
          <path d="m19 9-5 5-4-4-3 3"/>
        </svg>
        <span>{{ tab.name }}</span>
      </button>
    </div>
  </nav>
</template>
