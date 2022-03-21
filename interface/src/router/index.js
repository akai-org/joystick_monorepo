import { createRouter, createWebHistory } from 'vue-router'
import Game from '../interfaces/Game.vue'
import store from '../store/index.js'

const routes = [
  {
    path: '/',
    name: 'Login',
    component: () => import('../views/LoginPage.vue')
  },
  {
    path: '/game',
    name: 'Game',
    component: Game
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach(async (to, from) => {
  console.log(store.state.gui)
  // if (to.name === 'Game' && !store.state.gui) {
  //   return { name: 'Login' }
  // }
})

export default router
