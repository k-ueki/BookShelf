<template>
<div id="modal">
	<!--
		<div class="overlay" @click="$emit('close')">
	-->
		<div style="position:relative;">

			<router-link to="/top/?id=14"><div class="del">X</div></router-link>

			<div class="modal-main-content" style="position:absolute;">
				<h3>Create A Community</h3>
				<div class="input-wrapper">
					<div class="comNameflame">Community Name   :   
					<input class="comName" /></div>
					<div class="comMem" style="float:left;">Community Member :   
						<div v-for="c in count" style=";">
							<input class="comMembers"/>
							<button class="" @click="add()">+</button>
						</div>
					</div>
				</div>
				<div class="btn-wrapper" style="text-align:center;">
					<button class="create-com-button" style="">決定</button>
				</div>
			</div>
		</div>
		<div v-show="loading_act" class="search_wrapper">
			検索結果
			<div v-show="loading && loading_act" class="loading">
				Loading...
			</div>
			<ul v-show="!loading">
				<li class="list" v-for="item in items" style="cursor:pointer;" @click="clickItem(item)">	
					<img class="itemImage" :src="item.Item.largeImageUrl">
					<div class="itemInfo">
						<tr><li>{{ item.Item.title }}</li></tr>
						<tr><li>{{ item.Item.author }}</li></tr>
						<tr><li>{{ item.Item.itemPrice }}円</li></tr>
					</div>
				</li>
			</ul>
		</div>
</div>
</template>
<script>
import axios from 'axios'

export default{
//	el:"#modal",
	data(){
		return{
			test:'',
			items:[],
			loading_act:false,
			loading:true,
			count:1,
		}
	},
	created(){
	},
	methods:{
		add(){
			this.count++
		},
		clickItem(item){
			if(confirm("登録しますか？")){
				var params = new URLSearchParams();
				item = item.Item
				params.append("title",item.title);
				params.append("author",item.author);
				params.append("price",item.itemPrice);
				params.append("img_url",item.largeImageUrl);
				params.append("rakuten_page_url",item.itemUrl);
				params.append("user_id",this.$route.params.id);

				axios.post("http://localhost:8888/regist/book/",params)
					.then(res => {

					})
					.catch(err => {

					})
			}else{
				alert("登録をキャンセルしました。")
			}
		},
		searchBook(){
			this.loading_act=true
			var params = new URLSearchParams();
			params.append("booksTitle",this.bookT);

			axios.post("http://localhost:8888/top/bookapi/",params)
				.then(res => {
					this.items = res.data.Items

					this.loading = false
				})
				.catch(err => {

				})
		}	
	}
}
</script>
<style>

.del{
	cursor:pointer;
	float:right;
}
h3 {
  margin: 40px 0 0;
}
.comNameflame,.comMem{
	float:left;
	margin-left:25%;
}
.comName{
	margin-left:10px;
}
.input-wrapper{
	margin-top:10%;
	height:auto;
}
.btn-wrapper{
	margin-top:5%;
	margin-bottom:5%;
	text-align:center;
}
.list{
	width:100%;
	padding:5% 0;
	border-bottom:1px solid black;
}
.itemImage{
	float:left;
}
.itemTitle{
	float:right;
}
.loading{
	text-align:center;
}
.search_wrapper{
	width:75%;
	margin:auto auto;
}
</style>
