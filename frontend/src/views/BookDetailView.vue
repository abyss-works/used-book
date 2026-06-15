<script setup lang="ts">
import { useRoute } from 'vue-router'

const route = useRoute()
</script>

<template>
  <div class="min-h-screen pb-20">
    <header class="sticky top-0 z-40 bg-bg/80 backdrop-blur-lg border-b border-border-subtle safe-top">
      <div class="max-w-lg mx-auto px-4 py-3 flex items-center gap-3">
        <button
          class="bg-transparent text-text-secondary font-medium rounded-sm px-2 py-1 hover:bg-white/5 transition-colors duration-150 cursor-pointer"
          @click="$router.back()"
        >
          <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="m15 18-6-6 6-6"/>
          </svg>
        </button>
        <h1 class="text-text text-lg font-semibold truncate">책 상세</h1>
      </div>
    </header>

    <main class="max-w-lg mx-auto px-4 pt-4">
      <p class="text-text-muted text-sm">준비 중입니다</p>
    </main>
  </div>
</template>
