<script>
	import { onMount } from "svelte";


    /**
	 * @type {any[]}
	 */
    let itemCategory1 = [];
    /**
	 * @type {any[]}
	 */
    let itemCategory2 = [];
    onMount (async()=> {
        itemCategory1 = await getItemCategory1()
        itemCategory2 = await getItemCategory2()
    })
    const getItemCategory1 = async ()=> {
        return [
            {   
                id:1,
                imgUrl : "https://cf.shopee.co.th/file/f06f0abdf6bb5a65fdc6edf73101b820_tn",
                caption : "เสื้อผ้าแฟชั่นผู้ชาย"
            },
            {   
                id:2,
                imgUrl : "https://cf.shopee.co.th/file/398408c2f8ab9de41f1e06e75edb38bc_tn",
                caption : "มือถือและอุปกรณ์เสริม"
            },
            {   
                id:3,
                imgUrl : "https://cf.shopee.co.th/file/c3283d2ba2ac9b209b3e09c1b1d54279_tn",
                caption : "อาหารเสริมและผลิตภัณฑ์สุขภาพ"
            },
            {   
                id:4,
                imgUrl : "https://cf.shopee.co.th/file/fe6165488b5fc7910e00b99dc7188759_tn",
                caption : "นาฬิกาและแว่นตา"
            },
            {   
                id:5,
                imgUrl : "https://cf.shopee.co.th/file/ed11c3aa687100c467130e42812e3af1_tn",
                caption : "รองเท้าผู้ชาย"
            },
            {   
                id:6,
                imgUrl : "https://cf.shopee.co.th/file/c1c3ef22b08e9e92a820238900edc982_tn",
                caption : "คอมพิวเตอร์และแล็ปท็อป"
            },
            {   
                id:7,
                imgUrl : "https://cf.shopee.co.th/file/b34493d3548c702f96f52cbc8a937f3b_tn",
                caption : "กล้องและอุปกรณ์ถ่ายภาพ"
            },
            {   
                id:8,
                imgUrl : "https://cf.shopee.co.th/file/d798b34734fdbd0c5b186563453fd9db_tn",
                caption : "กีฬาและกิจกรรมกลางแจ้ง"
            },
            {   
                id:9,
                imgUrl : "https://cf.shopee.co.th/file/4b8edfd0d3711df614cee49ce9575384_tn",
                caption : "สื่อบันเทิงภายในบ้าน"
            },
            {   
                id:10,
                imgUrl : "https://cf.shopee.co.th/file/a1e2db8724dcd09ace3371b5d0785140_tn",
                caption : "เกมและอุปกรณ์เสริม"
            },
        ]
    }
const getItemCategory2 = async ()=> {
    return [
        {   
                id:11,
                imgUrl : "https://cf.shopee.co.th/file/14eb48c96d2af38b43e73bd438cb4418_tn",
                caption : "เสื้อผ้าแฟชั่นผู้หญิง"
            },
            {   
                id:12,
                imgUrl : "https://cf.shopee.co.th/file/dad0e4f4f94ab4a75392ebf587af5bf0_tn",
                caption : "ความงามและของใช้ส่วนตัว"
            },
            {   
                id:13,
                imgUrl : "https://cf.shopee.co.th/file/82c330eb4a6202dd2e359cad3a2bc289_tn",
                caption : "ของเล่น สินค้าแม่และเด็ก"
            },
            {   
                id:14,
                imgUrl : "https://cf.shopee.co.th/file/b680b0e115248f8b7cf4c1dab5240563_tn",
                caption : "เครื่องใช้ในบ้าน"
            },
            {   
                id:15,
                imgUrl : "https://cf.shopee.co.th/file/a879f390f7cfe48a7e9816a3fb489e98_tn",
                caption : "กระเป๋า"
            },
            {   
                id:16,
                imgUrl : "https://cf.shopee.co.th/file/f84c1aede5e30931407b670435bfa3db_tn",
                caption : "รองเท้าผู้หญิง"
            },
            {   
                id:17,
                imgUrl : "https://cf.shopee.co.th/file/68b87384345976cb3735b129ecec92d7_tn",
                caption : "เครื่องประดับ"
            },
            {   
                id:18,
                imgUrl : "https://cf.shopee.co.th/file/ef3225f54d2d375247837391cf743e93_tn",
                caption : "อาหารและเครื่องดื่ม"
            },
            {   
                id:19,
                imgUrl : "https://cf.shopee.co.th/file/2b35d04b65cfd4bac5264133f830c3ae_tn",
                caption : "เครื่องใช้ไฟฟ้าภายในบ้าน"
            },
            {   
                id:20,
                imgUrl : "https://cf.shopee.co.th/file/6ab5237b2f3716e46ee71d0cea1a23e8_tn",
                caption : "สัตว์เลี้ยง"
            },
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
            {#each itemCategory1 as item }       
            <li><a href="/" class="tabs item1">
                <div class="categorytop item-1">
                    <div class="categoryimg">
                        <img src={item.imgUrl} alt="">
                    </div>
                    <div class="captionimg-1">{item.caption}</div>
                </div>
            </a></li>
            {/each}
        </ul>
    </div>
    <div class="topmenu"> 
        <ul class="d-flex">
            {#each itemCategory2 as item2 } 
            <li><a href="/" class="tabs item1">
                <div class="categorytop item-1">
                    <div class="categoryimg">
                        <img src="{item2.imgUrl}" alt="">
                    </div>
                    <div class="captionimg-1">{item2.caption}</div>
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

    .categoryimg img{
       width: 5.25rem;
       height: 5.5rem;
       background-size: contain;
       margin: 0 auto;
    }

</style>
