<template>
<v-app id="modal">
		<div>
			<router-link to="/top"><div class="del">X</div></router-link>

			<div class="modal-main-content">
				<h3>Resister A Book</h3>
				<div class="input-wrapper">
					<v-form>
					<span>Book's Title : </span>
					<v-text-field 
						v-model="bookTitle"
						label="search word"></v-text-field>
					</v-form>
				</div>
				<div class="btn-wrapper">
					<v-btn class="regisBtn" @click=searchBook
						target="_blank">検索</v-btn>
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
					<br/>
				</li>
			</ul>
		</div>
</v-app>
</template>
<script>
import axios from 'axios'

export default{
	data(){
		return{
			test:'',
			items:[],
			loading_act:false,
			loading:true,
		}
	},
	methods:{
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
						alert("登録が完了しました。")
					})
					.catch(err => {
						alert("登録に失敗しました。")
					})
			}else{
				alert("登録をキャンセルしました。")
			}
		},
		searchBook(){
			this.loading_act=true

			axios.get("http://localhost:8888/books/" + this.bookTitle)
				.then(res => {
					console.log(res.data.Books)
					this.items = res.data.Books
					this.loading = false
				})
				.catch(err => {

				})
		}	
	}
}
</script>
<style>
#modal{
	/*
	display:none;
	*/
}
/*
.overlay {
    background: rgba(0, 0, 0, .8);
    position: fixed;
    width: 100%;
    height: 100%;
    left: 0;
    top: 0;
    right: 0;
    bottom: 0;
    z-index: 900;
    transition: all .5s ease;
}
.modal-enter,
.modal-leave-active {
    opacity: 0;
}
*/

.del{
	cursor:pointer;
	float:right;
}
h3 {
  margin: 40px 0 0;
}
.them{
	float:left;
	margin-left:25%;
}
.booktitle{
	margin-left:10px;
}
.input-wrapper{
	margin-top:10%;
}
.btn-wrapper{
	margin-top:5%;
	margin-bottom:5%;
	text-align:center;
}
.list{
	/*
	margin:5% 0;
	*/
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
