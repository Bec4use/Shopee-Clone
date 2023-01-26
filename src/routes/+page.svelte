<script>
    import '../style/global.css'
	import Bodycontent from './bodycontent.svelte';
	import Carosel from './carosel.svelte';
	import Events from './events.svelte';
	import Flashsale from './flashsale.svelte';
	import Followus from './followus.svelte';
	import Itemcategory from './itemcategory.svelte';
	import Navbarbot from './navbarbot.svelte';
</script>
<div class="navbar-wrap sticky-top">
    <div class="navbartopper">
        <div class="container navbar">
            <div class="navbar top-left"> <!--ซ้าย-->
                <ul>
                    <li><a href="/">Seller Centre |</a></li>
                    <li><a href="/">ขายสินค้ากับช้อปปี้ |</a></li>
                    <li><a href="/">ดาวน์โหลด |</a></li>
                </ul>
                <Followus />
            </div> 

            <div class="navbar top-right"> <!---ขวา-->
                <ul>
                    <div class="alarm-menu">
                        <svg xmlns="http://www.w3.org/2000/svg" height="20" width="20"><path d="M4 15.5V14h1V9q0-1.792 1.115-3.177Q7.229 4.438 9 4.104V3q0-.417.292-.708Q9.583 2 10 2t.708.292Q11 2.583 11 3v1.104q1.771.334 2.885 1.708Q15 7.188 15 9v5h1v1.5Zm6-5.875ZM10 18q-.625 0-1.062-.438Q8.5 17.125 8.5 16.5h3q0 .625-.438 1.062Q10.625 18 10 18Zm-3.5-4h7V9q0-1.458-1.021-2.479Q11.458 5.5 10 5.5q-1.458 0-2.479 1.021Q6.5 7.542 6.5 9Z"/></svg>
                        <li><a href="/">การแจ้งเตือน</a></li>
                    </div>
                    <div class="help-menu">
                        <svg xmlns="http://www.w3.org/2000/svg" height="20" width="20"><path d="M10 15q.417 0 .708-.292Q11 14.417 11 14t-.292-.708Q10.417 13 10 13t-.708.292Q9 13.583 9 14t.292.708Q9.583 15 10 15Zm-.75-3.188h1.521q0-.77.135-1.093.136-.323.656-.823.73-.708 1.011-1.208.281-.5.281-1.105 0-1.145-.781-1.864Q11.292 5 10.083 5q-1.062 0-1.843.562-.782.563-1.094 1.521l1.354.563q.188-.584.594-.906.406-.323.948-.323.583 0 .958.333t.375.875q0 .479-.323.854t-.719.729q-.729.667-.906 1.094-.177.427-.177 1.51ZM10 18q-1.646 0-3.104-.625-1.458-.625-2.552-1.719t-1.719-2.552Q2 11.646 2 10q0-1.667.625-3.115.625-1.447 1.719-2.541Q5.438 3.25 6.896 2.625T10 2q1.667 0 3.115.625 1.447.625 2.541 1.719 1.094 1.094 1.719 2.541Q18 8.333 18 10q0 1.646-.625 3.104-.625 1.458-1.719 2.552t-2.541 1.719Q11.667 18 10 18Zm0-1.5q2.708 0 4.604-1.896T16.5 10q0-2.708-1.896-4.604T10 3.5q-2.708 0-4.604 1.896T3.5 10q0 2.708 1.896 4.604T10 16.5Zm0-6.5Z"/></svg>
                        <li><a href="/">ช่วยเหลือ</a></li>
                    </div>
                    <div class="language-menu">
                        <svg xmlns="http://www.w3.org/2000/svg" height="20" width="20"><path d="M10 18q-1.646 0-3.104-.625-1.458-.625-2.552-1.719t-1.719-2.552Q2 11.646 2 10q0-1.667.625-3.115.625-1.447 1.719-2.541Q5.438 3.25 6.896 2.625T10 2q1.667 0 3.115.625 1.447.625 2.541 1.719 1.094 1.094 1.719 2.541Q18 8.333 18 10q0 1.646-.625 3.104-.625 1.458-1.719 2.552t-2.541 1.719Q11.667 18 10 18Zm0-1.562q.354-.355.708-1.323.354-.969.542-2.115h-2.5q.188 1.146.542 2.115.354.968.708 1.323Zm-1.896-.209q-.292-.625-.51-1.437-.219-.813-.365-1.792H4.25q.604 1.188 1.604 2.031 1 .844 2.25 1.198Zm3.792 0q1.25-.354 2.25-1.198 1-.843 1.604-2.031h-2.979q-.146.979-.365 1.792-.218.812-.51 1.437Zm1.062-4.729h3.354q.105-.396.146-.76.042-.365.042-.74t-.042-.75q-.041-.375-.146-.75h-3.354q.042.396.052.781.011.386.011.761t-.011.739q-.01.365-.052.719Zm-4.396 0h2.876q.041-.417.052-.781.01-.365.01-.719t-.01-.729q-.011-.375-.052-.771H8.562q-.041.396-.052.771-.01.375-.01.729t.01.729q.011.375.052.771Zm-4.874 0h3.354q-.042-.396-.052-.781-.011-.386-.011-.761t.011-.739q.01-.365.052-.719H3.688q-.105.396-.146.76-.042.365-.042.74t.042.74q.041.364.146.76ZM12.771 7h2.979q-.604-1.188-1.604-2.031-1-.844-2.25-1.198.292.625.51 1.437.219.813.365 1.792ZM8.75 7h2.5q-.188-1.146-.542-2.115-.354-.968-.708-1.323-.354.355-.708 1.323Q8.938 5.854 8.75 7Zm-4.5 0h2.979q.146-.979.365-1.792.218-.812.51-1.437-1.25.354-2.25 1.198-1 .843-1.604 2.031Z"/></svg>
                        <li><a href="/">ไทย</a></li>
                    </div>
                    
                    <li><a href="/">สมัครใหม่</a></li>
                    <li><a href="/">เข้าสู่ระบบ</a></li>
                </ul>

            </div>
    </div>
    </div>

    <div class="container-wrapper header-with-search-wrapper">
        <Navbarbot />
    </div> <!-- Navbar ด้านล่าง-->
    
</div>
<div class="carousel">
    <Carosel />
</div>
<div class="body-content">
    <Bodycontent />
    <Itemcategory />
    <Flashsale />
    <Events />
</div>
<div class="itemcategory">
    
</div>
<style>
    *{
        padding: 0;
        margin: 0;
        box-sizing: border-box;
        font-family: Arial, Helvetica, sans-serif;
    }
    li {
        list-style: none;
    }
    li a{
        text-decoration: none;
        color: #fff;
        font-size: .8125rem;
        font-weight: lighter;
        padding: 0.25rem;
    }
    .navbartopper{
        background-color: #ee4d2d;
    }
    .container{
        margin-left: auto;
        margin-right: auto;
        background-color: #ee4d2d;
        margin-top: 0;
        max-width: 1200px;
        height: 34px;
    }
    .navbar{
        display: flex;
        align-items: center;
        justify-content: space-between;
    }
    .navbar.top-left {
        display: flex;
    }
    ul {
        display: flex;
        margin: 0px 10px;
    }
    .navbar-wrap{
        background-color: #fff;
    }
    svg{
        fill: #fff;
        height: 20px;
    }
    .alarm-menu {
        display: flex;
    }
    .help-menu {
        display: flex;
    }
    .language-menu{
        display: flex;
    }
    .carousel {
        margin-top: 20px;
    }
    .body-content {
        background-color: rgba(0,0,0,.09);
        height: 100%;
        margin-top: 10px;
    }
</style>