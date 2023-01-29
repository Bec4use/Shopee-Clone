<script>
	import { onMount } from "svelte"


    let second = 59
    let mins = 59
    let hours = 23
    setInterval(() => {
        second -= 1
      if(second==0){
        second=59
    }
    }, 1000)
    setInterval(() => {
        mins -= 1
      if(mins==0){
        mins=59
    }
    }, 60000)
    setInterval(() => {
        hours -= 1
      if(hours=0){
        hours=23
    }
    }, 3600000)
    /**
	 * @type {any[]}
	 */
    let paydayItems = [];
    onMount (async()=>{
        paydayItems = await getPaydayItems()
    })
    const getPaydayItems = async ()=>{
        return[
            {
                id:1,
                price: "24",
                imgUrl: "https://cf.shopee.co.th/file/sg-11134201-23010-9taxyh8y6omv92_tn",
                discount:"73%"
            },
            {
                id:2,
                price: "60",
                imgUrl: "https://cf.shopee.co.th/file/th-50009109-75de64695b4aa1cf6fc532d932ba66d8_tn",
                discount:"80%"
            },
            {
                id:3,
                price: "10,500",
                imgUrl: "https://cf.shopee.co.th/file/f5dbffff9f3761268d1543c51b8796ac_tn",
                discount:"73%"
            },
            {
                id:4,
                price: "38,500",
                imgUrl: "https://cf.shopee.co.th/file/0681d4b6465fdf62237431cdb9c2365c_tn",
                discount:"73%"
            },
            {
                id:5,
                price: "25,900",
                imgUrl: "https://cf.shopee.co.th/file/09e4946e938e74afc0b391bbea764185_tn",
                discount:"73%"
            },
            {
                id:6,
                price: "14,900",
                imgUrl: "https://cf.shopee.co.th/file/fbc4b416919f07d6027d5177b924591d_tn",
                discount:"73%"
            }
        ]
    }
</script>

<div class="spacer-homepage">
</div>
<div class="simple banner container">
    <a href="/"><div class="banner-1"></div></a>
</div>
<div class="flashsale container">
    <div class="flashsale-top">
        <a href="/" class="">
            <div class="flashsale-menu"></div>
        </a>
        <div class="time-zone-menu">
            <span>จบใน</span>
            <div class="ms-2 d-flex">
                <div class="timeitem">{('00'+hours).slice(-2)}</div><h4>:</h4>
                <div class="timeitem">{('00'+mins).slice(-2)}</div><h4>:</h4>
                <div class="timeitem">{('00'+second).slice(-2)}</div>
            </div>
        </div>
        
    </div>
    
    <div class="body-menu">
        <div class="list-flashsale">
            <span>1.25 สินค้าขายดี</span>
            <a href="/">ดูเพิ่มเติม &gt;</a>
        </div>
        <div class="img-flashsale d-flex container">
            {#each paydayItems as item}
            <a href="/" class="flashitem-1">  
                <div class="flashsale1">
                    <div class="flashimg">
                        <img src={item.imgUrl} alt="">
                    </div>
                    <div class="flashimgcover"></div>
                    <div class="price-1">฿{item.price}</div>
                    <div class="discount-1">
                        <span>ลด</span>
                        <span>{item.discount}</span>
                    </div>
                </div>  
            </a>
            {/each}
        </div>
    </div>
    <div class="footer-menu"></div>
</div>
<style>
    *{
        box-sizing: border-box;
    }
    .spacer-homepage{
        height: 0.125rem;
    }
    .simple{
        width: 100%;
    }
    .banner-1{
        background-image:url(https://cf.shopee.co.th/file/23117c6a9868708c4d060bc501f5e258);
        max-width: 1200px;
        height: 110px;
        margin-top: 1rem;
        margin-left: auto;
        margin-right: auto;
        background-size:cover ;
        display: block;
        overflow: hidden;
        position: relative;
    }
    .flashsale-menu{
        background-image: url("https://cf.shopee.co.th/file/46c34fd744ceac9d77d3c0d5444e1437");
        margin: 10px auto 0;
        max-width: 1200px;
        height: 60px;
    }
    .body-menu{
        background-image: url("https://cf.shopee.co.th/file/c829c2b603d9d33904020f0eff18314c");
        margin: 0 auto;
        max-width: 1200px;
        height: 300px;
    }
    .footer-menu{
        background-image: url("https://cf.shopee.co.th/file/8f87fbf4c664cc016f0a6d34e0c2236b");
        margin: 0 auto;
        max-width: 1200px;
        height: 15px;
    }
    .list-flashsale {
        display: flex;
        justify-content: space-between;
        padding: 1rem 2rem;
    }
    .list-flashsale {
        color:  rgb(238, 77, 45);;
    }
    .list-flashsale a{
        text-decoration: none;
        color:  rgb(238, 77, 45);
        font-size: .75rem;
    }
    .img-flashsale{
        margin: 0 auto;
    }
    .flashsale1{
        background-color: #fff;
        height: 14.8rem;
        width: 12.125rem;
        margin: 0 1px;
        padding-top: 0.5rem;
        text-align: center;
    }
    .flashimg img{
        height: 10.125rem;
        width: 10.125rem;
        margin: 1rem auto;
        position: relative;
        background-size: cover;
    }
    .flashitem-1 {
        text-decoration: none;
        color: black;
    }
    .discount-1 {
        position: relative;
        background-color: rgba(255,212,36,.9);
        width: 2.25rem;
        height: 2rem;
        transform: scale(1.4);
        border-radius: 0 0 0 50%;
        left: 150px;
        top: -220px;
        z-index: 99;
    }
    .discount-1 span {
        display: flex;
        font-size: 0.625rem;
        text-align: center;
        justify-content: center;
    }
    .discount-1 span+span {
        color: red;
    }

    .flashimgcover{
        background-image: url("https://cf.shopee.co.th/file/sg-11134004-23010-t16qsa0w5ylv46_tn");
        height: 10.125rem;
        width: 10.125rem;
        margin: 1rem auto;
        position: absolute;
        background-size: cover;
        top: 47.5rem;
    }
    .time-zone-menu{
        width: 10.875rem;
        height: 2.375rem;
        display: flex;
        position: absolute;
        top: 0;
        top: 40.5rem;
        left: 48.5rem;
        color: #fff;
        font-size: 1.25rem;
    }
    @media only screen and (min-width: 1200px)
    {
    .time-zone-menu{
        left: 84.5rem;
    }
    }
    .timeitem{
        background-color: rgb(255, 255, 255);
        color: rgb(238, 77, 45);
        padding: 0.1rem 0.425rem;
        height: 1.875rem;
        font-size: 1.2rem;
        border-radius: 5px;
        margin:0 0.2rem;
    }
</style>
