<script>
	import '../style/global.css';
	import Bodycontent from './bodycontent.svelte';
	import Carosel from './carosel.svelte';
	import Events from './events.svelte';
	import Flashsale from './flashsale.svelte';
	import Itemcategory from './itemcategory.svelte';
	import Login from './loginmodal/login.svelte';
	import Weeklysale from './weeklysale.svelte';
    import Topnav from './Components/topnav.svelte';
</script>
<Topnav />

<div class="carousel">
	<Carosel />
</div>
<div class="body-content">
	<Bodycontent />
	<Itemcategory />
	<Flashsale />
	<Events />
	<Weeklysale />
	<Login />
</div>
<div class="itemcategory" />

<style>
	* {
		padding: 0;
		margin: 0;
		box-sizing: border-box;
		font-family: Arial, Helvetica, sans-serif;
	}
	.carousel {
		margin-top: 20px;
	}
	.body-content {
		background-color: rgba(0, 0, 0, 0.09);
		height: 100%;
		margin-top: 10px;
	}
</style>
