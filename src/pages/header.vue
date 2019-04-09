<template>
        <header>
                <div class="head-container">
                    <div class="logo">MyBook</div>
                    <div class="header-list">
                        <ul>
                            <li>A</li>
                            <li>B</li>
                            <li>C</li>
                            <li>D</li>
                        </ul>
                    </div>
                </div>
                <div class="search-form">
                    <input type="search" name="search" class="search" placeholder="検索ワード">
                    <button type="submit" class="search-btn">検索</button>
                </div>
                <div class="icon">
					<!--
                    <img class="parsonal-icon" src="../../images/top.png">
					-->
					<img class="parsonal-icon" src="../../images/UNADJUSTEDNONRAW_thumb_411.jpg"> 
                </div>
            
        </header>
</template>
<script>
export default{

}
</script>
<style>
header{
    position:fixed;
    top:0%;
    left:0%;
    width:80%;
    height:55px;
    background-color:rgba(253, 188, 46,0.7);
    padding:0 12.5%;

li{
    list-style:none;
    float:left;
    padding-left:30px;
}


.head-container{
    text-align:center;
    
}
.logo{
    font-family:"Monotype Corsiva";
    font-size:35px;
    float:left;
}
.header-list{
    float:left;
    height:55px;
}

.search-form{
    height:55px;
    float:right;
    line-height:30px;
}
.search{
    background-color:rgba(255,255,255,0.2);
    border-color:rgba(0,0,0,0.5);
}
.search-btn{
    background-color:rgba(255,255,255,0);
    color:rgb(0,0,0);
    border-color:rgba(0,0,0,0.3);
}

.icon{
    float:right;
    height:55px;
}
.parsonal-icon{
    width:40px;
    height:40px;
    border-radius:50%;
}
</style>
