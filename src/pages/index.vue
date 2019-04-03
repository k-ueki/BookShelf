<template>
  <div class="main-wrapper">
    <div class="main-container">
      <div class="left-half-page">
        <div class="catch">
          <ul>
            <li><span class="fa fa-check-circle"></span>自分の本棚をカスタマイズしよう</li>
            <li><span class="fa fa-check-circle"></span>友達が読んでいる本をのぞいてみよう</li>
            <li><span class="fa fa-check-circle"></span>本棚をシェアしよう</li>
          </ul>
        </div>
      </div>
      <div class="right-half-page">
        <div class="login-form">
			<input type="email" v-model="email" name="email" class="form" placeholder="メールアドレス">
			<input type="password" v-model="pass" name="pass" class="form" placeholder="パスワード">
			<button class="login-form-btn" @click=login>ログイン</button>
          <!--<a href="#">パスワードを忘れた場合はこちら</a>-->
        </div>
        <div class="catch2">
          <span　class="fa fa-book"></span>
          <h3 class="item">いろいろな本をのぞいてみよう</h3>
          <div class="login-regist-btn">
			  <li>ウチ本を始めよう</li>
			  <router-link to="/signup/"><button class="make-account-btn btn">アカウント作成</button></router-link>
			  <!--
			  <router-link ><button class="login-btn btn">ログイン</button></router-link>
			  -->
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import axios from 'axios'

export default{
	name:"index",
	data(){
		return{
			email:'',
			pass:''
		}
	},
	methods:{
		login(){
			let params = new URLSearchParams();
			params.append('email',this.email);
			params.append('pass',this.pass);

			axios.post("http://localhost:8888/",params)
				.then(response => {

				}).catch(error=>{
					this.errorStatus = "Error: Network Error";
				})
		}
	}
}
</script>
<style>
.main-wrapper{
    position: absolute;
    top:0px;
    left:0px;
	
    background-image:url("../../images/UNADJUSTEDNONRAW_thumb_411.jpg");
    background-size:cover;

 
	/*
    background-color:rgb(129, 194, 8);
    background-color:rgba(0, 0, 0, 0.3);
	*/ 
    color:#ffff;
    
    width:100%;
    height:100%;
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
