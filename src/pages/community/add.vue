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
							<input class="comMembers" placeholder="@name" v-model="addname"/>
							<button class="" @click="count+=1">+{{addname}}</button>
								<!--
							<ul v-show="flag" style="display:block;">
								<li v-bind:class="Pclass" v-for="person in people" style="cursor:pointer;" @mouseover="focus()" @click="select(person)">
									{{person.ID}}  {{person.Name}}
								</li>
							</ul>
								-->
							<div v-bind:class="Pclass" v-show="flag" v-for="person in people" style="cursor:pointer;" @mouseover="focus()" @click="select(person)">{{person.ID}} {{person.Name}}</div>
						</div>
					</div>
				</div>
				<div class="btn-wrapper" style="text-align:center;">
					<button class="create-com-button" style="">決定</button>
				</div>
			</div>
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
			addname:"",
			people:[],
			flag:false,
			Pclass:'',
		}
	},
	created(){
		axios.get("http://localhost:8888/community/add/")
			.then(response =>{
				console.log(response.data)
				this.people = response.data;
				this.flag=true;
			})
	},
	methods:{
		select(person){
			this.addname=person.Name;
			this.flag=false;
		},
		focus(){
			this.Pclass = "changeback"	
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

.changeback{
	background-color:red;
}
</style>
