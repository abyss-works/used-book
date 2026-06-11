<template>
  <div class="min-h-screen bg-gray-50">
    <header class="bg-white border-b border-gray-200 px-6 py-4">
      <div class="max-w-4xl mx-auto flex items-center gap-3">
        <h1 class="text-xl font-bold text-gray-900">
          <router-link to="/">📚 중고책 최적 조합</router-link>
        </h1>
      </div>
    </header>
    <main class="max-w-4xl mx-auto p-6">
      <router-view />
    </main>
  </div>
</template>
