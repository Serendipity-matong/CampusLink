import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Products from '../views/Products.vue'
import Tasks from '../views/Tasks.vue'
import Orders from '../views/Orders.vue'
import Profile from '../views/Profile.vue'
import Login from '../views/Login.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: { title: '首页 - CampusLink 校园互联' }
  },
  {
    path: '/products',
    name: 'Products',
    component: Products,
    meta: { title: '二手市场 - CampusLink' }
  },
  {
    path: '/tasks',
    name: 'Tasks',
    component: Tasks,
    meta: { title: '跑腿代购 - CampusLink' }
  },
  {
    path: '/orders',
    name: 'Orders',
    component: Orders,
    meta: { title: '我的订单 - CampusLink' }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: Profile,
    meta: { title: '个人中心 - CampusLink' }
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { title: '登录 - CampusLink' }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  document.title = to.meta.title || 'CampusLink 校园互联'
  next()
})

export default router

