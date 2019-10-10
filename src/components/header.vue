<template>
    <v-app-bar app dark style="width:100%;">
      <v-toolbar-title class="headline text-uppercase">
        <span>MyBookShelf</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
	  <v-form>
		  <v-text-field
			label="search word"
			></v-text-field>
	  </v-form>
		  <v-btn
			  text
			  target="_blank"
			  ><span>検索</span>
		  </v-btn>
      <v-btn
        text
        target="_blank"
      >
	  <span class="mr-2" @click="logout">Logout</span>
      </v-btn>
	</v-app-bar>
        <!-- <header> -->
        <!--         <div class="head&#45;container"> -->
        <!--             <div class="logo">〇</div> -->
        <!--             <div class="header&#45;list"> -->
        <!--                 <ul> -->
        <!--                     <li style="cursor:pointer;">My Bookshelf</li> -->
		<!-- 					<li>B</li> -->
        <!--                     <li>C</li> -->
        <!--                     <li>D</li> -->
        <!--                 </ul> -->
        <!--             </div> -->
        <!--         </div> -->
        <!--         <div class="search&#45;form"> -->
        <!--             <input type="search" name="search" class="search" placeholder="検索ワード"> -->
        <!--             <button type="submit" class="search&#45;btn">検索</button> -->
		<!-- 			<button class="logout&#45;btn">ログアウト</button> -->
        <!--         </div> -->
        <!--         <div class="icon"> -->
		<!-- 			<!&#45;&#45; -->
        <!--             <img class="parsonal&#45;icon" src="../../images/top.png"> -->
		<!-- 			&#45;&#45;> -->
        <!--         </div> -->
        <!--      -->
        <!-- </header> -->
</template>
<script>
export default{
	methods:{
		logout(){
			console.log("logout ok")
			console.log("logout()")
			// firebase.auth().signout().then(()=>
			//  console.log("success for logout!")
			//  )
		}
	}
}
</script>
<style>
