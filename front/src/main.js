import Vue from 'vue'
import App from './App.vue'
import router from './router'
import firebase from 'firebase'
import { BootstrapVue, BootstrapVueIcons } from 'bootstrap-vue'

Vue.config.productionTip = false

Vue.use(BootstrapVue)
Vue.use(BootstrapVueIcons)


var firebaseConfig = {
  apiKey: process.env.VUE_APP_APIKEY,
  authDomain: process.env.VUE_APP_AUTHDOMAIN,
  databaseURL: process.env.VUE_APP_DATABASEURL,
  projectId: process.env.VUE_APP_PROJECTID,
  storageBucket: process.env.VUE_APP_STORAGEBUCKET,
  messagingSenderId: process.env.VUE_APP_MESSAGINGSENDERID,
  appId: process.env.VUE_APP_APPID
};
// Initialize Firebase
firebase.initializeApp(firebaseConfig);

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
