<template>
  <div id="main">
  <router-view></router-view>
  </div>
</template>
<script>
import HelloWorld from './components/HelloWorld.vue'
import Header from './components/header.vue'
import Modal from './components/modal.vue'

import axios from "axios"
import Router from 'vue-router'


export default {
  name: 'main',
	data(){
		return {
			msg:"t"
		}
	},
	methods:{
	},
  	components: {
    	HelloWorld,
	  	Header,
	  	Modal
  	}
}

</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
