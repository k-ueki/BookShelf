import Vue from 'vue'
import VueRouter from 'vue-router'

import Index from '@/pages/index'
import Signup from '@/pages/signup'
import Top from '@/pages/top'
import Regist from '@/pages/register'
import CommunityAdd from '@/pages/community/add'

Vue.use(VueRouter)

const router = new VueRouter({
	mode:'history',
	routes:[
		{
			path: '/',
			name:"Index",
			component:Index
		},
		{
			path: '/signup',
			component: Signup
		},
		{
			path: '/top/:uid',
			component: Top
		},
		{
			path: '/regist/:id?',
			name: "Regist",
			component:Regist
		},
		{
			path:'/community/add/',
			name:"CommunityAdd",
			component:CommunityAdd
		}
	]
})

export default router
