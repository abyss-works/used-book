import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/wishlist',
      name: 'wishlist',
      component: () => import('@/views/WishlistView.vue'),
    },
    {
      path: '/optimize',
      name: 'optimize',
      component: () => import('@/views/OptimizedView.vue'),
    },
    {
      path: '/book/:id',
      name: 'book-detail',
      component: () => import('@/views/BookDetailView.vue'),
    },
  ],
})

export default router
