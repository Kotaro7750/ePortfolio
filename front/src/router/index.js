import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'
import Router from 'vue-router'
import firebase from 'firebase'

import SignIn from '@/components/auth/SignIn'
import Ticker from '@/components/ticker/Ticker'
import TickerDetail from '@/components/ticker_detail/TickerDetail'
import Purchase from '@/components/purchase/Purchase'
import Asset from '@/components/share/Asset'

Vue.use(Router)
Vue.use(BootstrapVue)

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

let router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Asset',
      component: Asset,
      meta: { requiresAuth: true }
    },
    {
      path: '/signin',
      name: 'SignIn',
      component: SignIn
    },
    {
      path: '/ticker',
      name: 'Ticker',
      component: Ticker,
      meta: { requiresAuth: true }
    },
    {
      path: '/ticker_detail/:id',
      name: 'TickerDetail',
      component: TickerDetail,
      meta: { requiresAuth: true }
    },
    {
      path: '/purchase',
      name: 'Purchase',
      component: Purchase,
      meta: { requiresAuth: true }
    },
  ],
})

router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  if (requiresAuth) {
    firebase.auth().onAuthStateChanged(function(user) {
      if (user) {
        next()
      } else {
        next({
          path: '/signin',
          query: { redirect: to.fullPath }
        })
      }
    })
  } else {
    next()
  }
})

export default router
