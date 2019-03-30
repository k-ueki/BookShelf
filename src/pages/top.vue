<template>
        <div class="main-wrapper-top">
	<Header/>
            <div class="parsonal-sp">
                <!--top-image-->
				<!--
                <img class="parsonal-icon" href="../../images/{$parsonal_info.image}">
				-->
				<img class="iconSelf" src="../../images/UNADJUSTEDNONRAW_thumb_411.jpg"> 
				<div class="parsonal-name">{{ name }}</div>

            </div>
            <div class="bookshelf-parsonal-wrapper">
                <!--search-->
                <div class="mini-header">
					<div class="searchWrap"><input type="search" name="word-search" class="word-search" placeholder="検索ワード"></div>
					<div class="btnWrap"><button class="btn">検索</button></div>
					<div class="selectWrap"><select name="refine">
                        <option value="order">新着順</option>
                        <option value=""></option>
						</select></div>
                </div>
				<div class="dispInfo">
					<!--display personal information-->
					<div>
						<div class="iconForTop">
							<img class="iconBook" src="../../images/UNADJUSTEDNONRAW_thumb_411.jpg"> 
						</div>
						<div>exmpple</div>
					</div>
	                <div class="bookshelf">
	                    <!--images-->
						<div>hello exmaple!!</div>
						<div>exmaple2</div>
	                </div>
				</div>
            </div>
        </div>
</template>
<script>
import Header from '../components/header.vue'
import axios from 'axios'

export default{
	name: 'top',
	data(){
		return{
			name:"ueki"
		}
	},
	methods:{

	},
	components:{
		Header
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
    width:79%;
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
}
.iconBook{
	padding:10px 10px;
	width:25%;
	height:100%;
}
</style>
