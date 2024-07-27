import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/login'
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/Login.vue')
    },
    {
      path: '/sit',
      name: 'sit',
      component: () => import('@/views/sit.vue')
    },
    {
      path: '/:catchAll(.*)',
      meta: {
        closeTab: true,
      },
      component: () => import('@/views/404.vue')
    }
  ]
})

export default router
