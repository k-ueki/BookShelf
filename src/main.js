import Vue from 'vue'
import App from './App.vue'
import Signup from './pages/signup.vue'

import axios from 'axios'
import router from './router.js'

Vue.config.productionTip = false

new Vue({
	router,
  render: h => h(App),
}).$mount('#app');
