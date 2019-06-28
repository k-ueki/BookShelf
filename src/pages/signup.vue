<template>
        <div id="signup" class="signup-form">
			<h2>アカウント作成</h2>
			<div class="name-form">
				<li class="required">*必須</li>
				<input v-model="name" type="text" name="name" class="name" placeholder="名前" value="">
			</div>
			<div class="userid-form">
				<li class="required">*必須</li>
				<input v-model="userid" type="text" name="userid" class="name" placeholder="ユーザーID" value="">
			</div>
			<div class="pass-form">
				<li class="required">*必須</li>
				<input v-model="password" type="password" name="pass" class="pass" placeholder="パスワード" value="">
			</div>
			<div class="email-form">
				<li class="required">*必須</li>
				<input v-model="email" type="email" name="email" class="email" placeholder="メールアドレス" value="">
			</div>
			<div class="image-form">
				<!--
				<input v-model="newUser.file" type="file" name="image" class="image">
				-->
			</div>
			<div class="btnform"><button type="submit" class="signup-btn" v-on:click=signUp>登録</button></div>
		</div>
</template>
<script src="../main.js"></script>
<script>
import Vue from 'vue'
import axios from 'axios'

const newUser = function(){
	this.name = '',
	this.email = '',
	this.password = '' //	//this.file = ''
}

export default{
	pops:['name','password','email','file'],
	name: "signup",
	data(){
		return{
				name:'',
				userid:'',
				email:'',
				password:''
				//file:''
		}
	},
	methods:{
		signUp(){
			if(this.name=="" || this.userid=="" || this.password=="" || this.email==""){
				alert("入力エラー")
			}else{
				let params = new URLSearchParams();
				params.append("name",this.name);
				params.append("userid",this.userid);
				params.append("password",this.password);
				params.append("email",this.email);
				//params.append("file",this.file);

				axios.post("http://localhost:8888/signup/",params)
				.then(response => {
					if(response){
						if(response.data === "このメールアドレスはすでに登録されています"){
							alert(response.data)
						}else if(response.data === "このユーザーIDはすでに使用されています"){
							alert(response.data)
						}else if(response.data ==="このメールアドレスとユーザーIDはすでに使用されています"){
							alert(response.data)
						}else{
							alert("登録が成功しました。MyPageに移行します。")
							this.$router.push({path:"/top/",query: {id:response.data.ID}})
						}
					}else{
						alert("登録に失敗しました。再度登録をお願いします。")
					}
				}).catch(error => {
					this.errorStatus = "Error: Network Error";
				})
			}
		},
	}
}
</script>

<style>
body{
    /*background-color:rgba(0,54,104,0.9);*/
    
}
h2{
    color:white;
    /*border-bottom:1px solid white;*/
}
.required{
    list-style:none;
    margin-top:0px;
    height:15px;
    color:red;
    padding-bottom:5px;
}

.signup-form{
    position: absolute;
    top:20%;
    left:37.5%;

    height:60%;
    width:25%;

    border:1px solid rgb(0,0,0);
    background-color:rgba(0,54,104,0.7)
}

.name-form{
    border-bottom:1px solid rgb(0,0,0);
    /*padding-top:10%;*/
}
.name{
    border:none;
    background-color:rgba(255,255,255,0);
    font-size:20px;
    color:rgb(255,255,255);
}

.pass-form{
    border-bottom:1px solid rgb(0,0,0);
    padding-top:10%;
}
.pass{
    border:none;
    background-color:rgba(255,255,255,0);
    font-size:20px;
    color:rgb(255,255,255);
}

.email-form{
    border-bottom:1px solid rgb(0,0,0);
    padding-top:10%;
}
.email{
    border:none;
    background-color:rgba(255,255,255,0);
    font-size:20px;
    color:rgb(255,255,255);
}

.image-form{
    padding-top:10%;
    border-bottom:1px solid rgb(0,0,0);
}

.btnform{
    padding-top:5%;
    text-align:center;
}
.signup-btn{
    font-size:10px;
    /*display:none; */
}
</style>
