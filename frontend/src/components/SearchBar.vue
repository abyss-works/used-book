<script setup lang="ts">
const props = defineProps<{
  modelValue: string
  placeholder?: string
}>()

const emit = defineEmits<{
  'update:modelValue': [value: string]
  search: [query: string]
}>()

function onInput(e: Event) {
  emit('update:modelValue', (e.target as HTMLInputElement).value)
}

function onSubmit() {
  if (props.modelValue.trim()) {
    emit('search', props.modelValue.trim())
  }
}
</script>

<template>
  <form class="relative" @submit.prevent="onSubmit">
    <svg
      class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-text-muted pointer-events-none"
      viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
    >
      <circle cx="11" cy="11" r="8"/>
      <path d="m21 21-4.35-4.35"/>
    </svg>
    <input
      :value="modelValue"
      @input="onInput"
      type="search"
      :placeholder="placeholder || '책 제목, 작가 검색...'"
      class="w-full pl-12 pr-4 py-3.5 bg-surface-elevated text-text rounded-xl border border-border placeholder:text-text-muted focus:outline-none focus:border-brand/50 focus:ring-1 focus:ring-brand/30 transition-colors duration-200"
      enterkeyhint="search"
    />
  </form>
</template>
