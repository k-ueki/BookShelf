<template>
	<v-app style="background-color:rgb(245,245,245);">
		<Header/>
        <div class="main-wrapper-top">
			<div class="overlay" v-show="showModal"></div>
			<div class="modal" v-show="showModal">
				<div>
					<div class="modalHeader" style="float:right;">
						<div style="cursor:pointer;" @click="closemodal">X</div>
					</div>
					<div class="modalinfo" align="center" style="float:right;">
						<!--
						-->
						<tr><img class="personalbooksIMG" :src="clickedbook.img_url"></tr>
						<tr class="">{{ clickedbook.title }}</tr>
						<tr class="">{{ clickedbook.author }}</tr>
						<tr class="">{{ clickedbook.price }}円</tr>
						<tr class=""><a :href="clickedbook.rakuten_page_url">楽天ページ</a></tr>
						<!-- {{indextmp}} -->
						<tr class=""><v-btn style="cursor:pointer;color:red;" @click="delBook(clickedbook,indextmp)">削除</v-btn></tr>
					</div>
				</div>
			</div>

            <div class="parsonal-sp">
                <!--top-image-->
				<!--
                <img class="parsonal-icon" href="../../images/{$parsonal_info.image}">
				-->
				<img class="iconSelf" src="../../images/UNADJUSTEDNONRAW_thumb_411.jpg"> 
				<div class="parsonal-name"> {{ name }}</div>
				<div class="">{{userid}}</div>

				<br/>
				<span style="color:rgba(0,0,0,0.4);">Communities</span>
				<div v-for="com in communities">
					<div>
						<router-link :to="{name:'CommunityDetail',params:{com_id:com.id,name:com.name}}">
							<a @click="selectCommunity(com.id)">
								{{com.name}}
							</a>
						</router-link>
					</div>
				</div>
            </div>
            <div class="bookshelf-parsonal-wrapper">
                <!--search-->
                <div class="mini-header">
					<router-link :to="{
						name:'Regist',
						params:{
							id: id 
						}
					}">
						<v-btn
							text
							target="_blank"
							><span>+</span></v-btn></router-link>
					<router-link :to="{
						name:'CommunityAdd' 
						}"
					>
					<v-btn
						text
						target="_blank"
						><span>community add</span></v-btn></router-link>
					<!-- <div class="searchWrap"><input type="search" name="word&#45;search" class="word&#45;search" placeholder="検索ワード"></div> -->
					<div class="searchWrap">
						<v-form>
							<v-text-field
								label="search word"
								></v-text-field>
						</v-form>
					</div>
					<div class="btnWrap"><v-btn class="btn"
							text
							target="_blank"
							>検索</v-btn></div>

				<!-- 	<div class="selectWrap"><select name="refine"> -->
                <!--         <option value="order">新着順</option> -->
                <!--         <option value=""></option> -->
				<!-- 		</select></div> -->
                </div>





				<ul>
					<li class="personalbookWrapper" v-for="(info,index) in booksinfo" @click="bookDetail(info,index)" style="cursor:pointer;width:250px;float:left;">
						<tr><img class="personalbooksIMG" :src="info.img_url"></tr>
						<tr class="personalbooks">{{ info.title }}</tr>
						<tr class="personalbooks">{{ info.author }}</tr>
						<tr class="personalbooks">{{ info.price }}円</tr>
						<!--
						<tr class="personalbooks"><p>楽天ページ</p></tr>
						{{info}}
						-->
					</li>
				</ul>
            </div>
        </div>
	</v-app>
</template>
<script>
import HelloWorld from '../components/HelloWorld.vue'
import Header from '../components/header.vue'
//import Modal from '../components/modal.vue'
						
import firebase from 'firebase'
import router from '../router.js'

import axios from 'axios'

const firebaseConfig = {
  apiKey: process.env.VUE_APP_FIREBASE_APIKEY,
  authDomain: process.env.VUE_APP_FIREBASE_AUTHDOMAIN,
  databaseURL: process.env.VUE_APP_FIREBASE_DATABASEURL,
  projectId: process.env.VUE_APP_FIREBASE_PROJECTID,
  messagingSenderId: process.env.VUE_APP_FIREBASE_MESSAGINGSENDERID,
  appId: process.env.VUE_APP_FIREBASE_APPID
}

export default{
	name: 'top',
	data(){
		return{
			id:'',
			name:"",
			showModal:false,
			booksinfo:'',
			clickedbook:'',
			indextmp:'',
			userid:'',
			communities:'',
		}
	},
	created:function(){
		firebase.initializeApp(firebaseConfig);
		firebase.auth().onAuthStateChanged((user) => {
		  if (user) {
		    // サインイン済み
			  console.log("LOGINOK")

			  let uid
			  uid = user.uid

			  axios.get("http://localhost:8888/user/" + uid)
				  .then(res=>{
					  if(res.data){
						axios.get("http://localhost:8888/top/" + uid)
							.then(res => {
								console.log("res",res.data)
								this.name = res.data.name;
								this.booksinfo = res.data.books;
								this.communities = res.data.communities;
								console.log("img",res.data.books[1].img_url)
							}).catch(err => {
								this.errorStatus = "Error: Network Error";
							})
					  }else{
						  let params = new URLSearchParams();
						  params.append('uid',uid);
						  params.append('name',user.displayName);

						  axios.post("http://localhost:8888/user",params)
							  .then(res=>{
								this.name = user.displayName
							  }).catch(err=>{
								this.errorStatus = "Error: Network Error"
							  })
					  }
				  })

		  } else {
		    // サインインしていない状態
		    // サインイン画面に遷移する等
		    // 例:
			  // router.push('/')
			  console.log('not login')
		  }
		});

		// var query = Object.assign({}, this.$route.query)
		// var params = new URLSearchParams();
		// params.append("id",query.id);
		// console.log("QUERYID",query.id)
		// axios.get("http://localhost:8888/top/?uid="+query.id,params)
	},
	methods:{
		bookDetail(info,index){
			this.indextmp = index
			this.clickedbook = info
			console.log("info",info)
			if(!this.showModal){
				this.showModal=true;
			}
//			var params = new URLSearchParams();
//			params.append("id",info.Id);
//			axios.post("http://localhost:8888/top/booksInfo/",params)
//				.then(res => {
//					console.log(res.data)
//				}).catch(err => {
//
//				})
//			console.log(info)
		},
		closemodal(){
			this.showModal=false;
		},
		delBook(book,index){
			this.booksinfo.splice(index,1)
			var params = new URLSearchParams();
			params.append("delid",book.Id);
			if(confirm(book.Title+" を削除しますか？")){
				// axios.post("http://localhost:8888/top/del/",params)
				// 	.then(res => {
				// 		this.showModal=false
				// 	}).catch(err => {
                //
				// 	})
			}
		},
		selectCommunity(id){
			console.log("comid",id)
		}
	},
	components:{
		HelloWorld,
		Header,
		//Modal
	}
}
</script>
<style>
.main-wrapper-top{
    position: absolute;
    top:70px;
    left:15%;
    width:70%;
    /* border:4px solid rgb(206,201,100); */
}


.parsonal-sp{
	background-color:rgb(255,255,255);
    width:20%;
    height:100%;
    float:left;
    /* border:2px solid rgb(0,0,0); */
    text-align:center;
}
.word-search{
	float:right;
	display:inline-block;
}
.btn{
	float:right;
	display:inline-block;
}

.searchWrap{
	display:inline-block;
}
.btnWrap{
	display:inline-block;
	margin-right:10px;
}
.selectWrap{
	display:inline-block;
}


.parsonal-icon{
    size:20px;
    border-radius:10px;
}

.bookshelf-parsonal-wrapper{
    width:78%;
    height:100%;
    float:right;
	background-color:rgb(255,255,255);
    /* border:2px solid rgb(65,166,90); */
}
.mini-header{
    text-align:right;
	width:100%;
}

.iconSelf{
	width:90px;
	height:90px;
	border-radius:50%;
	margin-top:10px;
}
.iconBook{
	padding:10px 10px;
	width:25%;
	height:100%;
}
.addbook{ cursor:pointer;
	float:left;

}
.personalbooks{
	text-align:center;

}
.personalbooksIMG{
}
.personalbookWrapper{
	/* border-bottom:1px solid rgba(0,0,0,0.4); */
	padding:5% 5%;
}

.modal{
	display: flex;
    align-items: center;
    justify-content: center;
    position: fixed;
    z-index: 30;
    top: 25%;
    left: 25%;
	width:50%;
	height:50%;
	background:#ffff;
}
.overlay{
	display: flex;
    align-items: center;
    justify-content: center;
    position: fixed;
    top: 0;
    left: 0;
	z-index:15;
	width:100%;
	height:100%;
	background:rgba(0,0,0,0.5);
}
.modalHeader{
	align:right;
}
</style>
