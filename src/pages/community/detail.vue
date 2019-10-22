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
						<tr class=""><v-btn style="cursor:pointer;color:red;" @click="delBook(clickedbook,indextmp)">削除</v-btn></tr>
					</div>
				</div>
			</div>

			<div class="modal" v-show="showModalMemAdd">
				<!-- <div style="float:right;"> -->
				<!-- 		<div style="cursor:pointer;" @click="closemodalMemAdd">X</div> -->
				<!-- </div> -->
				<div style="height:100%;width:100%;">
					<div class="modalHeader" style="float:right;height:5%;">
						<div style="cursor:pointer;" @click="closemodalMemAdd">X</div>
					</div>
					<div style="float:left;height:50%;width:30%;" class="addition-sp">
						<div v-for="">
							<input placeholder="name" v-model="addname">
							<v-btn @click="plus">+</v-btn>
						</div>
					</div>
					<div style="float:left;border:solid 2px red;height:95%;width:50%;" class="add-list-sp">
						<div style="border-bottom=solid 2px red;">
							<table>
								<th>Id</th>
								<th>Name</th>
								<tr v-for="co in count">
									<td>{{ids[co]}}</td>
									<td>{{names[co]}}</td>
								</tr>
							</table>
						</div>
						<v-btn><span>決定</span></v-btn>
					</div>
				</div>
			</div>

            <div class="parsonal-sp">
				<img class="iconSelf" src="../../../images/library-1147815_1920.jpg"> 
				<div class="com-name" style="font-size:25;"> {{ name }} </div>

				<br/>
				<span style="color:rgba(0,0,0,0.4);">members</span>
				<div v-for="mem in members">
					<div>
						<a>
							<!-- <img class="mem&#45;icon" :src="mem.photo_url">  -->
						</a>
						<a>
							{{mem.name}}
						</a>
					</div>
				</div>
            </div>

            <div class="bookshelf-parsonal-wrapper">
                <div class="mini-header">
					<v-btn
						text
						target="_blank"
						@click="memAddition"
						><span>member add</span></v-btn>
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
                </div>

				<ul>
					<li class="personalbookWrapper" v-for="(info,index) in booksinfo" @click="bookDetail(info,index)" style="cursor:pointer;width:250px;float:left;">
						<tr><img class="personalbooksIMG" :src="info.img_url"></tr>
						<tr class="personalbooks">{{ info.title }}</tr>
						<tr class="personalbooks">{{ info.author }}</tr>
						<tr class="personalbooks">{{ info.price }}円</tr>
					</li>
				</ul>
            </div>
        </div>
	</v-app>
</template>
<script>
import HelloWorld from '../../components/HelloWorld.vue'
import Header from '../../components/header.vue'
						
import firebase from 'firebase'
import router from '../../router.js'

import axios from 'axios'

export default{
	name: 'top',
	data(){
		return{
			id:'',
			name:"",
			showModal:false,
			showModalMemAdd:false,
			booksinfo:'',
			clickedbook:'',
			indextmp:'',
			userid:'',
			members:[],
			count:1,
			names:[],
			ids:[],
			addname:"",
		}
	},
	created:function(){
		axios.get("http://localhost:8888/top/community/"+this.$route.params['com_id'])
			.then(res=>{
				this.name = res.data.community.name
				this.members = res.data.users
				this.booksinfo = res.data.books
			}).catch(res=>{
				console.log("err",res)
			})
	},
	methods:{
		bookDetail(info,index){
			this.indextmp = index
			this.clickedbook = info
			console.log("info",info)
			if(!this.showModal){
				this.showModal=true;
			}
		},
		memAddition(){
			if(!this.showModalMemAdd){
				this.showModalMemAdd=true
			}
		},
		closemodal(){
			this.showModal=false;
		},
		closemodalMemAdd(){
			this.showModalMemAdd=false;
			this.count=1;
		},
		delBook(book,index){
			this.booksinfo.splice(index,1)
			var params = new URLSearchParams();
			params.append("delid",book.Id);
			if(confirm(book.Title+" を削除しますか？")){
				// axios.post("http://localhost:8888/top/del/",params)
			}
		// },
		// countMinus(){
		// 	if(this.count>1){
		// 		this.count--;
		// 	}else{
		// 		//this.name=""
		// 	}
		},
		plus(){
			if(this.addname==""){
				alert("no name")
			}else{
				this.names[this.count] = this.addname
				this.ids[this.count]=this.count//暫定
				this.count++
				this.addname=""
				console.log(this.names)
			}
		}
	},
	components:{
		HelloWorld,
		Header,
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
	border:solid 3px red;
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
	background-color:rgba(144,24,38,0.5);
}

.mem-icon{
	width:30px;;
	height:30px;
	border-radius:50%;
	margin-top:10px;
}

.add-list-sp{
    width:100%;
    height:100%;
    float:right;
	/* background-color:rgb(255,255,255); */
	background-color:rgba(25,177,140,0.5);
}

.addition-sp{
	/* background-color:rgb(255,255,255); */
	background-color:rgba(144,190,240,0.5);
    width:100%;
    height:100%;
    float:left;
    text-align:center;
}
</style>
