<script>
	import { onMount } from "svelte/types/runtime/internal/lifecycle";
    /**
	 * @type {any[]}
	 */
    let items = [];
    onMount(async () => {
        items = await getItems()
    })


    const getItems = async () => {
        // GET localhost:8080/items/
        return [
            {
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "ฟไกส่ดสรฟ่ไก"
            }, {
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            },{
                url: "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                description: "เสื้อผ้าแฟชั่นผู้ชาย"
            }
        ]
    }

</script>

<div class="spacer-homepage">

</div>
<div class="item-category">
     <div class="">หมวดหมู่</div>
</div>
<div class="category-menu container">
   <div class="topmenu"> 
        <ul class="d-flex">
            {#each items as item, index}
                <li><a href="/" class="tabs item1">
                    <div class="categorytop item-1">
                        <div class="categoryimg item1"></div>
                        <div class="captionimg-1">{item.description}</div>
                    </div>
                </a></li>
            {/each}
        </ul>
    </div>
</div>

<style>
    *{
        box-sizing: border-box;
    }
    .spacer-homepage{
        height: 0.125rem;
    }
    
    .item-category{
        background-color: #fff;
        max-width: 75rem;
        height: 3.875rem;
        margin-top: 1rem;
        margin-left: auto;
        margin-right: auto;
        display: flex;
        align-items: center;
        padding-left: 1rem;
        color: rgba(0,0,0,.54);
        border-bottom: 1px solid rgba(0,0,0,.05);
    }
    .category-menu{
        background-color: #fff;
        max-width: 75rem;
        height: 18.875rem;
        padding: 0;
        
    }
    .topmenu{
        height: 9.375rem;
        border-bottom: 1px solid rgba(0,0,0,.05);
        width: 75rem;
        margin: 0 auto;
    }
    a {
        text-decoration: none;
        color: #000;
        font-size: .875rem;;
    }
    ul{
        padding: 0;
    }
    li {
        list-style: none;
    }
    .categorytop {
        border-right: 1px solid rgba(0,0,0,.05);
        width: 7.5rem;
        height: 9.375rem;
        padding-top: .875rem;
        
    }
    .tabs {
        text-align: center;
    }
    .categorytop:hover{
        box-shadow: 1px 1px 3px rgba(0,0,0,.5);
    }

    .categoryimg {
       width: 5.25rem;
       height: 5.5rem;
       background-size: contain;
       margin: 0 auto;
    }
    .categoryimg.item1{
        background-image: url("");
    }
    .categoryimg.item2{
        background-image: url("https://cf.shopee.co.th/file/398408c2f8ab9de41f1e06e75edb38bc_tn");
    }
    .categoryimg.item3{
        background-image: url("https://cf.shopee.co.th/file/c3283d2ba2ac9b209b3e09c1b1d54279_tn");
    }
    .categoryimg.item4{
        background-image: url("https://cf.shopee.co.th/file/fe6165488b5fc7910e00b99dc7188759_tn");
    }
    .categoryimg.item5{
        background-image: url("https://cf.shopee.co.th/file/ed11c3aa687100c467130e42812e3af1_tn");
    }
    .categoryimg.item6{
        background-image: url("https://cf.shopee.co.th/file/c1c3ef22b08e9e92a820238900edc982_tn");
    }
    .categoryimg.item7{
        background-image: url("https://cf.shopee.co.th/file/b34493d3548c702f96f52cbc8a937f3b_tn");
    }
    .categoryimg.item8{
        background-image: url("https://cf.shopee.co.th/file/d798b34734fdbd0c5b186563453fd9db_tn");
    }
    .categoryimg.item9{
        background-image: url("https://cf.shopee.co.th/file/4b8edfd0d3711df614cee49ce9575384_tn");
    }
    .categoryimg.item10{
        background-image: url("https://cf.shopee.co.th/file/a1e2db8724dcd09ace3371b5d0785140_tn");
    }
    .categoryimg.item11{
        background-image: url("https://cf.shopee.co.th/file/14eb48c96d2af38b43e73bd438cb4418_tn");
    }
    .categoryimg.item12{
        background-image: url("https://cf.shopee.co.th/file/dad0e4f4f94ab4a75392ebf587af5bf0_tn");
    }
    .categoryimg.item13{
        background-image: url("https://cf.shopee.co.th/file/82c330eb4a6202dd2e359cad3a2bc289_tn");
    }
    .categoryimg.item14{
        background-image: url("https://cf.shopee.co.th/file/b680b0e115248f8b7cf4c1dab5240563_tn");
    }
    .categoryimg.item15{
        background-image: url("https://cf.shopee.co.th/file/a879f390f7cfe48a7e9816a3fb489e98_tn");
    }
    .categoryimg.item16{
        background-image: url("https://cf.shopee.co.th/file/f84c1aede5e30931407b670435bfa3db_tn");
    }
    .categoryimg.item17{
        background-image: url("https://cf.shopee.co.th/file/68b87384345976cb3735b129ecec92d7_tn");
    }
    .categoryimg.item18{
        background-image: url("https://cf.shopee.co.th/file/ef3225f54d2d375247837391cf743e93_tn");
    }
    .categoryimg.item19{
        background-image: url("	https://cf.shopee.co.th/file/2b35d04b65cfd4bac5264133f830c3ae_tn");
    }
    .categoryimg.item20{
        background-image: url("https://cf.shopee.co.th/file/6ab5237b2f3716e46ee71d0cea1a23e8_tn");
    }

</style>
