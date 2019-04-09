<template>
<transition name="modal">
		<div class="overlay" @click="$emit('close')">
			<div class="del">X</div>
			<div class="modal-main-content">
				<h3>Resisration A Book</h3>
				<div class="input-wrapper">
				<div class="them">Book's Title : </div>
				<input class="booktitle" v-model="bookT"/>
				</div>
				<div class="btn-wrapper">
					<button class="regisBtn" @click=regisBook>登録</button>
				</div>
			</div>
		</div>
</transition>
</template>
<script>
import axios from 'axios'

export default{
	methods:{
		regisBook(){
			var params = new URLSearchParams();
			params.append("booksTitle",this.bookT);
			
			axios.post("http://localhost:8888/top/book/",params)
				.then(res => {
					console.log("OKOK")
				})
				.catch(err => {

				})
			console.log(this.bookT)
			console.log("OKOK")
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
	margin-bottom:10%;
	text-align:center;
}
</style>
