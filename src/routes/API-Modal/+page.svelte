<script>
// @ts-nocheck
    import Swal from 'sweetalert2';
	import { goto } from '$app/navigation';
	import avatar from '../img/Avatar.svg';
	import shop from '../img/shop.svg';
	let password = '';
	let username = '';


    function handleClick() {
        let parent = this.parentNode.parentNode;
        parent.classList.add('focus');
    }
    function handleBlur() {
        let parent = this.parentNode.parentNode;
        if(this.value == ""){
            // document.getElementById('user').classList.remove('focus');
            parent.classList.remove('focus');
        }
    }

    async function handleLogin() {

		const data = { password, username };

		try {
			const res = await fetch('http://localhost:8080/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},

				body: JSON.stringify(data)
			});
            console.log(username);
            console.log(password);
			const json = await res.json();
			console.log(json);
			if (json.Status === 'OK') {
				Swal.fire({
					position: 'center',
					icon: 'success',
					title: `Welcome back : ${json.Username}`,
					showConfirmButton: true,
					
				}
			).then((value) =>{
				localStorage.setItem('token',json.token)
				goto('/Profile', { state: { foo: 'bar' } });
				
        	})
			} else {
				Swal.fire({
					icon: 'error',
					title: 'Oops...',
					text: json.error,
					footer: '<a href="">Why do I have this issue?</a>'
				});
			}
		} catch (error) {
			console.error(error);
		}
	}

</script>

<div class="containers">
	<div class="img">
		<img src={shop} alt="" />
	</div>
	<div class="login-container">
		<form action="">
			<img src={avatar} alt="" class="avatar"/>
			<h2>Welcome</h2>
			<!-- svelte-ignore a11y-click-events-have-key-events -->
			<div class="input-div one" id="user">
				<div class="i">
					<svg xmlns="http://www.w3.org/2000/svg" height="40" viewBox="0 96 960 960" width="40"><path d="M480 575.333q-66 0-109.667-43.666Q326.667 488 326.667 422t43.666-109.666Q414 268.667 480 268.667t109.667 43.667Q633.333 356 633.333 422t-43.666 109.667Q546 575.333 480 575.333ZM160 896V796q0-36.666 18.5-64.166T226.667 690Q292 659.667 354.333 644.5 416.667 629.334 480 629.334t125.333 15.5q62 15.5 127.334 45.166Q763 704.334 781.5 731.834 800 759.334 800 796v100H160Z"/></svg>
				</div>
				<div>
					<h5>Username</h5>
					<input type="text" class="input" on:focus={handleClick} on:blur={handleBlur} bind:value={username}/>
				</div>
			</div>
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <div class="input-div two" id="password">
				<div class="i">
					<svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 96 960 960" width="24"><path d="M240 976q-33 0-56.5-23.5T160 896V496q0-33 23.5-56.5T240 416h40v-80q0-83 58.5-141.5T480 136q83 0 141.5 58.5T680 336v80h40q33 0 56.5 23.5T800 496v400q0 33-23.5 56.5T720 976H240Zm0-80h480V496H240v400Zm240-120q33 0 56.5-23.5T560 696q0-33-23.5-56.5T480 616q-33 0-56.5 23.5T400 696q0 33 23.5 56.5T480 776ZM360 416h240v-80q0-50-35-85t-85-35q-50 0-85 35t-35 85v80ZM240 896V496v400Z"/></svg>
				</div>
				<div>
					<h5>Password</h5>
					<input type="password" class="input" on:focus={handleClick} on:blur={handleBlur} bind:value={password}/>
				</div>
			</div>
            <a href="/">Forget Password?</a>
            <input type="button" class="btn" on:click={handleLogin} value="Login">
		</form>
	</div>
</div>

<style>
    @import url('https://fonts.googleapis.com/css2?family=Poppins:wght@600&display=swap');
    
    *{
        padding: 0;
        margin: 0;
        box-sizing: border-box;
        font-family: 'Poppins', sans-serif;
    }
    .containers {
        width: 100vw;
        height: 100vh;
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        grid-gap: 7rem;
        padding: 0 2rem;
        margin: 0 auto;
    }
    .img {
        display: flex;
        justify-content: flex-end;
        align-items: center;
    }
    .img img {
        width: 500px;
    }
    .login-container {
        display: flex;
        align-items: center;
        text-align: center;
    }
    .i svg {
        fill: #d9d9d9;
        transition: .3s;
    }
    .avatar {
        width: 100px;
    }
    form h2{
        font-size: 2.9rem;
        text-transform: uppercase;
        margin: 15px 0;
        color: #333;
    }
    .input-div {
        display: grid;
        position: relative;
        grid-template-columns: 7%  93%;
        margin: 25px 0;
        padding: 5px 0;
        border-bottom: 2px solid #d9d9d9;
    }
    .input-div::after, .input-div::before {
        content: '';
        position: absolute;
        bottom: -2px;
        width: 0;
        height: 2px;
        background-color: #ee4d2d;
        transition: .3s;
    }
 
    .input-div::after{
        right: 50%;
    }
    .input-div::before{
        left: 50%;
    }
    .input-div.one  {
        margin-top: 0;
    }
    .input-div.two {
        margin-bottom: 4px;
    }
    .i{
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .input-div > div {
        position: relative;
        height: 45px;
    }
    .input-div > div h5{
        position: absolute;
        left: 10px;
        top: 50%;
        transform: translateY(-50%);
        color: #999;
        font-size: 18px;
        transition: .3s;
        
    }
    .input-div.focus .i svg{
        fill: #ee4d2d;
    }
    .input-div.focus  div h5{
        top: -5px;
        font-size: 15px;
    }
    /* .input-div.focus::after, .input-div.focus::before{
        width: 50%;
    } */

    .input {
        position: absolute;
        width: 100%;
        height: 100%;
        top: 0;
        left: 0;
        border: none;
        outline: none;
        background: none;
        padding: 0.5rem 0.7rem;
        font-size: 1.2rem;
        color: #555;
        font-family: 'Poppins', sans-serif;
    }
    a {
        display: block;
        text-align: right;
        text-decoration: none;
        color: #999;
        font-size: 0.9rem;
        transition: .3s;
    }
    a:hover {
        color: #ee4d2d;
    }
    .btn {
        display: block;
        width: 100%;
        height: 50px;
        border-radius: 25px;
        margin: 1rem 0;
        font-size: 1.2rem;
        outline: none;
        border: none;
        background-image: linear-gradient(to right,#e09a8c,#e78773,#eb4d2e);
        cursor: pointer;
        text-transform: uppercase;
        color: #fff;
        background-size: 200%;
        transition: .5s;
    }
    .btn:hover {
        background-position: right;
    }
</style>
