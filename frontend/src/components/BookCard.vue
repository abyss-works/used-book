<script setup lang="ts">
import type { Book } from '@/types'

const props = defineProps<{
  book: Book
  subtitle?: string
}>()

const emit = defineEmits<{
  click: [book: Book]
}>()
</script>

<template>
  <div
    class="bg-surface-elevated rounded-xl border border-border p-4 flex gap-3 cursor-pointer active:scale-[0.99] transition-transform duration-150"
    @click="emit('click', book)"
  >
    <!-- Cover -->
    <div class="w-16 h-24 shrink-0 rounded-lg overflow-hidden bg-surface">
      <img
        v-if="book.cover"
        :src="book.cover"
        :alt="book.title"
        class="w-full h-full object-cover"
        loading="lazy"
      />
      <div v-else class="w-full h-full flex items-center justify-center text-text-dim text-xs">
        No Cover
      </div>
    </div>

    <!-- Info -->
    <div class="flex-1 min-w-0 flex flex-col justify-between py-0.5">
      <div>
        <h3 class="text-text text-sm font-medium leading-snug line-clamp-2">{{ book.title }}</h3>
        <p class="text-text-muted text-xs mt-1 truncate">{{ book.author }}</p>
      </div>
      <div v-if="subtitle" class="text-text-secondary text-xs mt-1">
        {{ subtitle }}
      </div>
    </div>

    <!-- Chevron -->
    <svg class="w-4 h-4 text-text-dim shrink-0 self-center" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="m9 18 6-6-6-6"/>
    </svg>
  </div>
</template>
