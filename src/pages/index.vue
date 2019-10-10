<template>
  <div class="main-wrapper" style="background-color:rgba(0,0,0,0.7);">
	  {{id}}
    <div class="main-container">
      <div class="left-half-page"> <div class="catch">
		<ul>
            <li><span class="fa fa-check-circle"></span>自分の本棚をカスタマイズしよう</li>
            <li><span class="fa fa-check-circle"></span>友達が読んでいる本をのぞいてみよう</li>
            <li><span class="fa fa-check-circle"></span>本棚をシェアしよう</li>
          </ul>
        </div>
      </div>
      <div class="right-half-page">
        <!-- <div class="login&#45;form"> -->
		<!-- 	<input type="email" v&#45;model="email" name="email" class="form" placeholder="メールアドレス"> -->
		<!-- 	<input type="password" v&#45;model="pass" name="pass" class="form" placeholder="パスワード"> -->
		<!-- 	<button class="login&#45;form&#45;btn" @click=login>ログイン</button> -->
        <!--   <!&#45;&#45;<a href="#">パスワードを忘れた場合はこちら</a>&#45;&#45;> -->
        <!-- </div> -->
		<div class="catch2">
          <span class="fa fa-book"></span>
          <h3 class="item">いろいろな本をのぞいてみよう</h3>
          <div class="login-regist-btn">
			  <li>ウチ本を始めよう</li>
			  <!-- <router&#45;link to="/signup/"><button class="make&#45;account&#45;btn btn">アカウント作成</button></router&#45;link> -->
			  <!--
			  <router-link ><button class="login-btn btn">ログイン</button></router-link>
			  -->
			  <button @click="signIn" class="make-account-btn btn"><font :icon="['fab', 'google']"></font>SignIn</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import axios from 'axios'
import firebase from 'firebase'
import router from '../router.js'

const firebaseConfig = {
  apiKey: process.env.VUE_APP_FIREBASE_APIKEY,
  authDomain: process.env.VUE_APP_FIREBASE_AUTHDOMAIN,
  databaseURL: process.env.VUE_APP_FIREBASE_DATABASEURL,
  projectId: process.env.VUE_APP_FIREBASE_PROJECTID,
  messagingSenderId: process.env.VUE_APP_FIREBASE_MESSAGINGSENDERID,
  appId: process.env.VUE_APP_FIREBASE_APPID
}

export default{
	name:"index",
	data(){
		return{
			email:'',
			pass:'',
			id:''
		}
	},
	methods:{
		// login(){
		// 	let params = new URLSearchParams();
		// 	params.append('email',this.email);
		// 	params.append('pass',this.pass);
        //
		// 	axios.post("http://localhost:8888/",params)
		// 		.then(response => {
		// 			this.id = response.data
		// 			if(!!response.data && response.data!="ERROR"){
		// 				this.$router.push({path:"/top/",query: {id:response.data}})
		// 			}
		// 			if(response.data=="ERROR"){
		// 				alert("入力エラー")
		// 				this.email=""
		// 				this.pass=""
		// 			}
		// 		}).catch(error=>{
		// 			this.errorStatus = "Error: Network Error";
		// 		})
		// },
		signIn:function(){
			firebase.initializeApp(firebaseConfig);
			const googleProvider = new firebase.auth.GoogleAuthProvider()
       		firebase.auth().signInWithPopup(googleProvider)
				.then(user=>{
					console.log("OK",user)
					firebase.auth().currentUser.getIdToken(true)
						.then(function(idToken) {
							console.log(idToken)
							localStorage.setItem('token',idToken)
						});
					router.push({path:'/top/'})
				}).catch(err=>{
					console.log(err)
				})
		},
		test:function(){
			var user = firebase.auth().currentUser;
			console.log(user)
		}

	}
}
</script>
<style>
.main-wrapper{
    position: absolute;
    top:0px;
    left:0px;
	
    #background-image:url("../../images/UNADJUSTEDNONRAW_thumb_411.jpg");
    #background-image:url("../../images/library-1147815_1920.jpg");
    #background-image:url("../../images/books-1141910_1920.jpg");
    background-image:url("../../images/books-2596809_1920.jpg");
    background-size:cover;

 
	/*
    background-color:rgb(129, 194, 8);
    background-color:rgba(0, 0, 0, 0.7);
	*/
    color:#ffff;
    
    width:100%;
    height:100%;
}
.main-wrapper::before{
    background-color: rgba(0,0,0,0.5);
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    content: ' ';
}
.book-icon{
    font-size:20px;
    color:rgb(129, 194, 8);
}
li{
    list-style: none;
}
.catch{
    position:absolute;
    top:42.5%;
    left:12.5%;
    width:25%;
    margin:auto auto;
    background-color:rgba(255, 255, 255,0);
}
.catch li{
    padding-bottom:15px;
}
.left-half-page{
    margin:auto auto;
}
.right-half-page{
    position: absolute;
    left:50%;
    top:0px;
    z-index:10;
	/*
    background-color:rgba(0, 54, 104,0.5);
	*/
    width:50%;
    height:100%;
    float:right;
}
.form{
    background-color:rgba(255,255,255,0);
    border-color:rgba(255,255,255,0.3);
    float:left;
}
.login-form{
    position: absolute;
    top:0px;
    left:40%;
    
    float:right;
    margin-top:10px;
    
}

.catch2{
    position: absolute;
    top:40%;
    left:37.5%;
    width:25%;
    margin:auto auto;
}
.btn{
    border-radius:50px;
    width:100%;
    margin-top:15px;
}
.login-form-btn{
    position: absolute;
    top:0px;
    left:100%;
    width:25%;
    float:right;

    color:rgb(129, 194, 8);
    border-color:rgb(129, 194, 8);
    background-color:#ffff;
    border-radius:50%;
}
.make-account-btn{
    background-color:rgb(129, 194, 8);
    color:#ffff;
}
.login-btn{
    background-color:#ffff;
    color:rgb(129, 194, 8);
}
</style>
