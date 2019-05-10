<template>
        <div class="main-wrapper-top">
			<!--
			<Modal/>
			-->
			<Header/>

			<!--modal-->
			<div class="overlay" v-show="showModal"></div>
			<div class="modal" v-show="showModal">
				<div>
					<div class="modalHeader" style="float:right;">
						<div style="cursor:pointer;" @click="closemodal">X</div>
					</div>
					<div class="modalinfo" align="center" style="float:right;">
						<!--
						-->
						<tr><img class="personalbooksIMG" :src="clickedbook.ImgUrl"></tr>
						<tr class="">{{ clickedbook.Title }}</tr>
						<tr class="">{{ clickedbook.Author }}</tr>
						<tr class="">{{ clickedbook.Price }}円</tr>
						<tr class=""><a :href="clickedbook.RakutenPageUrl">楽天ページ</a></tr>
						<tr class=""><button style="cursor:pointer;" @click="delBook(clickedbook)">削除</button></tr>
					</div>
				</div>
			</div>

            <div class="parsonal-sp">
                <!--top-image-->
				<!--
                <img class="parsonal-icon" href="../../images/{$parsonal_info.image}">
				-->
				<img class="iconSelf" src="../../images/UNADJUSTEDNONRAW_thumb_411.jpg"> 
				<div class="parsonal-name">{{ id }} - {{ name }}</div>
            </div>
            <div class="bookshelf-parsonal-wrapper">
                <!--search-->
                <div class="mini-header">
					<!--
					<div class="addbook" @click="showModal  = true">+</div>
					-->
					<!--
					<router-link to="/regist/"><button>+</button></router-link>
					-->
					<router-link :to="{
						name:'Regist',
						params:{
							id: id 
						}
					}"><button>+</button></router-link>
					<div class="searchWrap"><input type="search" name="word-search" class="word-search" placeholder="検索ワード"></div>
					<div class="btnWrap"><button class="btn">検索</button></div>
					<div class="selectWrap"><select name="refine">
                        <option value="order">新着順</option>
                        <option value=""></option>
						</select></div>
                </div>
				<ul>
					<li class="personalbookWrapper" v-for="info in booksinfo" @click="bookDetail(info)" style="cursor:pointer;">
						<tr><img class="personalbooksIMG" :src="info.ImgUrl"></tr>
						<tr class="personalbooks">{{ info.Title }}</tr>
						<tr class="personalbooks">{{ info.Author }}</tr>
						<tr class="personalbooks">{{ info.Price }}円</tr>
						<!--
						<tr class="personalbooks"><p>楽天ページ</p></tr>
						{{info}}
						-->
					</li>
				</ul>
				<!--
				<div class="dispInfo">
					<div>
						<div class="iconForTop">
							<img class="iconBook" src="../../images/UNADJUSTEDNONRAW_thumb_411.jpg"> 
						</div>
						<div>exmpple</div>
					</div>
	                <div class="bookshelf">
						<div>hello exmaple!!</div>
						<div>exmaple2</div>
	                </div>
				</div>
				-->
            </div>
        </div>
</template>
<script>
import HelloWorld from '../components/HelloWorld.vue'
import Header from '../components/header.vue'
//import Modal from '../components/modal.vue'

import axios from 'axios'


export default{
	name: 'top',
	data(){
		return{
			id:'',
			name:"",
			showModal:false,
			booksinfo:'',
			clickedbook:'',
			clickedbookID:'',
		}
	},
	created(){
		console.log(this.$route.query)
		var query = Object.assign({}, this.$route.query)
		var params = new URLSearchParams();
		params.append("id",query.id);
		console.log("QUERYID",query.id)
		axios.post("http://localhost:8888/top/?id="+query.id,params)
			.then(res => {
				console.log(res.data)
				this.id = res.data.ID
				this.name = res.data.Name
				this.booksinfo = res.data.Books
			}).catch(err => {
				this.errorStatus = "Error: Network Error";
			})
	},
	methods:{
		bookDetail(info){
			this.clickedbook = info
			this.clickedbookID = info.id
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
		delBook(book){
			var params = new URLSearchParams();
			params.append("delid",book.Id);
			if(confirm(book.Title+" を削除しますか？")){
				axios.post("http://localhost:8888/top/del/",params)
					.then(res => {
						this.showModal=false
					}).catch(err => {

					})
			}
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
    top:55px;
    left:15%;
    width:70%;
    border:4px solid rgb(206,201,100);
}


.parsonal-sp{
    width:20%;
    height:100%;
    float:left;
    border:2px solid rgb(0,0,0);
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
    border:2px solid rgb(65,166,90);
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
	border-bottom:1px solid rgba(0,0,0,0.4);
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
