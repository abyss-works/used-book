<template>
  <div class="space-y-6">
    <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h2 class="text-lg font-semibold text-gray-800 mb-4">🔍 도서 검색</h2>
      <div class="flex gap-2">
        <input
          v-model="query"
          type="text"
          placeholder="책 제목, 저자로 검색..."
          class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          @keyup.enter="search"
        />
        <button
          :disabled="loading"
          class="px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
          @click="search"
        >
          {{ loading ? '검색 중...' : '검색' }}
        </button>
      </div>
    </div>

    <div v-if="results.length" class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
      <h2 class="text-lg font-semibold text-gray-800 mb-4">📖 검색 결과</h2>
      <ul class="divide-y divide-gray-200">
        <li
          v-for="book in results"
          :key="book.id"
          class="py-3 flex items-center gap-4"
        >
          <img
            v-if="book.cover"
            :src="book.cover"
            :alt="book.title"
            class="w-12 h-16 object-cover rounded"
          />
          <div class="flex-1 min-w-0">
            <p class="font-medium text-gray-900 truncate">{{ book.title }}</p>
            <p class="text-sm text-gray-500">{{ book.author }}</p>
          </div>
          <button
            @click="addWishlist(book)"
            class="shrink-0 px-3 py-1 text-sm bg-green-100 text-green-700 rounded hover:bg-green-200"
          >
            + 위시리스트
          </button>
        </li>
      </ul>
    </div>

    <div
      v-if="wishlist.length"
      class="bg-white rounded-lg shadow-sm border border-gray-200 p-6"
    >
      <h2 class="text-lg font-semibold text-gray-800 mb-4">📋 위시리스트</h2>
      <ul class="divide-y divide-gray-200 mb-4">
        <li
          v-for="(item, idx) in wishlist"
          :key="idx"
          class="py-2 flex items-center justify-between"
        >
          <div>
            <p class="font-medium text-gray-900">{{ item.title }}</p>
            <p class="text-sm text-gray-500">{{ item.author }}</p>
          </div>
          <button
            @click="wishlist.splice(idx, 1)"
            class="text-red-500 hover:text-red-700 text-sm"
          >
            삭제
          </button>
        </li>
      </ul>
      <button
        :disabled="optimizing"
        class="w-full py-2 bg-purple-600 text-white rounded-lg hover:bg-purple-700 disabled:opacity-50"
        @click="optimize"
      >
        {{ optimizing ? '최적 조합 계산 중...' : '🚀 최적 조합 찾기' }}
      </button>
    </div>

    <div
      v-if="result"
      class="bg-white rounded-lg shadow-sm border border-gray-200 p-6"
    >
      <h2 class="text-lg font-semibold text-gray-800 mb-4">✅ 최적 조합 결과</h2>
      <div class="grid grid-cols-3 gap-4 mb-4 text-center">
        <div class="bg-blue-50 rounded p-3">
          <p class="text-sm text-gray-500">총 비용</p>
          <p class="text-2xl font-bold text-blue-700">{{ result.total_cost.toLocaleString() }}원</p>
        </div>
        <div class="bg-green-50 rounded p-3">
          <p class="text-sm text-gray-500">구매처</p>
          <p class="text-2xl font-bold text-green-700">{{ result.sellers }}곳</p>
        </div>
        <div class="bg-orange-50 rounded p-3">
          <p class="text-sm text-gray-500">배송 횟수</p>
          <p class="text-2xl font-bold text-orange-700">{{ result.ship_count }}회</p>
        </div>
      </div>
      <table class="w-full text-sm">
        <thead>
          <tr class="border-b border-gray-200 text-gray-500">
            <th class="text-left py-2">책</th>
            <th class="text-left py-2">판매자</th>
            <th class="text-right py-2">가격</th>
            <th class="text-right py-2">배송비</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="item in result.purchases"
            :key="item.title + item.seller_name"
            class="border-b border-gray-100"
          >
            <td class="py-2">
              <p class="font-medium text-gray-900">{{ item.title }}</p>
              <p class="text-gray-500">{{ item.author }}</p>
            </td>
            <td class="py-2 text-gray-600">{{ item.seller_name }}</td>
            <td class="py-2 text-right font-medium">{{ item.price.toLocaleString() }}원</td>
            <td class="py-2 text-right text-gray-500">{{ item.delivery_fee.toLocaleString() }}원</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { searchBooks, optimizeWishlist } from '@/api'

const query = ref('')
const loading = ref(false)
const optimizing = ref(false)
const results = ref<any[]>([])
const wishlist = ref<any[]>([])
const result = ref<any>(null)

async function search() {
  if (!query.value.trim()) return
  loading.value = true
  try {
    const data = await searchBooks(query.value, 10, 1)
    results.value = data.books || []
  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }
}

function addWishlist(book: any) {
  if (!wishlist.value.find((w) => w.book_id === book.id)) {
    wishlist.value.push({
      book_id: book.id,
      title: book.title,
      author: book.author,
    })
  }
}

async function optimize() {
  optimizing.value = true
  try {
    const data = await optimizeWishlist(wishlist.value)
    result.value = data
  } catch (e) {
    console.error(e)
    result.value = null
  } finally {
    optimizing.value = false
  }
}
</script>
