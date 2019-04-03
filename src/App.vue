<template>
  <div id="main">
	  <!--    <img alt="Vue logo" src="./assets/logo.png">-->
	  <!--
	  <HelloWorld msg="Welcome to Your Vue.js App"/>
	  -->
  <router-view></router-view>
  </div>
</template>
<script>
import HelloWorld from './components/HelloWorld.vue'
import Header from './components/header.vue'

import axios from "axios"
import Router from 'vue-router'


export default {
  name: 'main',
	data(){
		return {
			msg:"testest"
		}
	},
	methods:{
	},
  components: {
    HelloWorld,
	  Header
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
