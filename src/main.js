import Vue from 'vue'
import App from './App.vue'
import Signup from './pages/signup.vue'

import axios from 'axios'
import router from './router.js'

Vue.config.productionTip = false

const User = function(){
	this.name = '',
	this.email = '',
	this.password = ''
	//this.file = ''
}

const signup = new Vue({
	el:"#signup",
	render: h => h(Signup),
	data:{
		users:[],
		newUser: new User()
	},
	methods:{
	}
})
new Vue({
	router,
  render: h => h(App),
}).$mount('#app')
