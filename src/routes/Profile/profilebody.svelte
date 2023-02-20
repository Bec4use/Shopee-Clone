<script>
	import Profiledetail from './profiledetail.svelte';
	let x = 'https://cf.shopee.co.th/file/ea97b2cfd8f08ba5f5883c8e3c6a3967_tn';
</script>

<div class="containerx d-flex">
	<div class="menu-profile">
		<div class="imgprofile d-flex">
			<img src={x} alt="" />
			<div class="username">
				<div>JoeNetipong</div>
				<a href="/"
					><svg xmlns="http://www.w3.org/2000/svg" height="20" viewBox="0 96 960 960" width="20"
						><path
							d="M794 477 589 271l43-44q32-32 75.5-33t79.5 33l53 53q35 33 31.5 78T837 434l-43 43Zm-58 57L303 968H97V763l434-433 205 204Z"
						/></svg
					>
					แก้ไขข้อมูลส่วนตัว</a
				>
			</div>
		</div>
		<div class="menulist">
			<a href="/"><img src="https://cf.shopee.co.th/file/ba61750a46794d8847c3f463c5e71cc4" alt="">
                <span>บัญชีของฉัน</span></a>
			<div class="user-menu">
				<a href="/"><span>ประวัติ</span></a>
				<a href="/"><span>บัญชีธนาคาร&บัตร</span></a>
				<a href="/"><span>ที่อยู่</span></a>
				<a href="/"><span>เปลี่ยนรหัสผ่าน</span></a>
				<a href="/"><span>การตั้งค่าความเป็นส่วนตัว</span></a>
				<a href="/"><span>ตั้งค่าการแจ้งเตือน</span></a>
			</div>
            <a href="/"><img src="https://cf.shopee.co.th/file/f0049e9df4e536bc3e7f140d071e9078" alt="">
                <span>การซื้อของฉัน</span></a>
            <a href="/"><img src="https://cf.shopee.co.th/file/e10a43b53ec8605f4829da5618e0717c" alt="">
                <span>การแจ้งเตือนของฉัน</span></a>
            <a href="/"><img src="https://cf.shopee.co.th/file/84feaa363ce325071c0a66d3c9a88748" alt="">
                <span>โค้ดส่วนลดของฉัน</span></a>
            <a href="/"><img src="https://cf.shopee.co.th/file/a0ef4bd8e16e481b4253bd0eb563f784" alt="">
                <span>Shopee Coins ของฉัน</span></a>
		</div>
	</div>
	<div class="update-profile">
        <Profiledetail />
    </div>
</div>

<style>
	.containerx {
		height: 38.75rem;
		width: 75rem;
		padding-top: 20px;
		padding-bottom: 50px;
		margin: 0 auto;
	}
	.menu-profile {
		width: 11.25rem;
		height: 100%;
		margin-right: 2rem;
	}
	.update-profile {
		width: 62rem;
		background-color: #fff;
		box-shadow: 1px 1px 1px 1px rgba(0, 0, 0, 0.2);
		border-radius: 3px;
	}
	.imgprofile {
		width: 100%;
		height: 5.4rem;
		border-bottom: 1px solid #888888;
		padding: 15px 0;
        margin-bottom: 1rem;
	}
	.imgprofile img {
		width: 3rem;
		height: 3rem;
		background-size: contain;
		border-radius: 50%;
		border: 1px solid rgba(0, 0, 0, 0.2);
		margin-right: 0.5rem;
	}
	.username {
		font-size: 0.9rem;
		font-weight: bolder;
	}
	svg {
		height: 1rem;
		width: 1rem;
		fill: #888888;
	}
	.username a {
		font-size: 0.75rem;
		color: #888888;
		text-decoration: none;
	}
    .menulist {
        display: grid;
		grid-template-columns: 1;
    }
    .menulist a {
        text-decoration: none;
        margin: .4rem 0;
        font-size: .9rem;
        color: #000;
    }
    .menulist img {
        width: 1.2rem;
        height: 1.2rem;
        margin-right: .5rem;
    }
	.user-menu {
		display: grid;
		grid-template-columns: 1;
        text-decoration: none;
	}
	.user-menu a {
		margin-left: 2rem;
        text-decoration: none;
        color: #888888;
        font-size: .9rem;
	}
</style>
