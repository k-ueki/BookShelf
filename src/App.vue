<template>
  <v-app>
	<router-view></router-view>
  </v-app>
</template>

<script>
import HelloWorld from '@/components/HelloWorld';
import Header from '@/components/header';

export default {
  name: 'App',
  components: {
	  HelloWorld,
	Header,
  },
  data: () => ({
    //
  }),
};
</script>
