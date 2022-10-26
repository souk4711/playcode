import { createRouter, createWebHistory } from 'vue-router'
import { PlayView } from '@/views'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'play',
      component: PlayView
    }
  ]
})

export default router
