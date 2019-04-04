import Vue from 'vue'
import VueRouter from 'vue-router'

import Index from '@/pages/index'
import Signup from '@/pages/signup'
import Top from '@/pages/top'

Vue.use(VueRouter)

const router = new VueRouter({
	mode:'history',
	routes:[
		{
			path: '/',
			component:Index
		},
		{
			path: '/signup',
			component: Signup
		},
		{
			path: '/top/:id',
			component: Top
		}
	]
})

export default router
