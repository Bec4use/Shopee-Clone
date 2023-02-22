<script lang="ts">
	import Swal from 'sweetalert2';
	import { goto } from '$app/navigation';
	import { createEventDispatcher } from 'svelte';

	let password = '';
	let username = '';
	let address = '';
	let birthday = '';
	let fullname = '';
	let tel = '';
	let showLoginForm = false;
	let showRegisterForm = false;
	let registerbtn = false;
	let loginbtn = false;
	

	
	const dispatch = createEventDispatcher();
	let showModal = false;
	let inputVal = '';
  
	function clearModal() {
	  inputVal = '';
	}
  
	function closeModal() {
	  showModal = false;
	  clearModal();
	}
  
	function openModal() {
	  showModal = true;
	}
	
	async function handleSubmit() {
		const data = { password, username, address, birthday, tel, fullname };

		try {
			const res = await fetch('http://localhost:8080/register', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},

				body: JSON.stringify(data)
			});

			const json = await res.json();
			console.log(json);
			if (json.Status === 'OK') {
				Swal.fire({
					position: 'center',
					icon: 'success',
					title: `Register result : Success `,
					showConfirmButton: true,
				}
			)
			} else {
				Swal.fire({
					icon: 'error',
					title: 'Oops...',
					text: 'Register result : Failed!',
					footer: '<a href="">Why do I have this issue?</a>'
				});
			}
		} catch (error) {
			console.error(error);
		}
	}
	async function handleLogin() {
		dispatch('submit', inputVal);
	 	 closeModal();
		const data = { password, username };

		try {
			const res = await fetch('http://localhost:8080/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},

				body: JSON.stringify(data)
			});

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
				close();
				localStorage.setItem('token',json.token)
				goto('/Profile', { state: { foo: 'bar' } });
				
        	})
			} else {
				Swal.fire({
					icon: 'error',
					title: 'Oops...',
					text: 'Something went wrong!',
					footer: '<a href="">Why do I have this issue?</a>'
				});
			}
		} catch (error) {
			console.error(error);
		}
	}

	function toggleLoginForm() {
		showLoginForm = !showLoginForm;
		showRegisterForm = false;
		loginbtn = true;
		registerbtn = false;
	}
	function toggleRegisterForm() {
		showRegisterForm = !showRegisterForm;
		showLoginForm = false;
		registerbtn = true;
		loginbtn = false;
	}
</script>
<li><a href="/" on:click={openModal}>เข้าสู่ระบบ</a></li>

<div class="modal" tabindex="-1" role="dialog" style={showModal ? 'display: block;' : ''}>
	<div class="modal-dialog" role="document">
		<div class="modal-content">
			<button class="btn-close" data-bs-dismiss="modal" aria-label="Close" on:click={closeModal} />
			<div class="modal-header">
				{#if loginbtn}
					<div class="" id="btn-login" />
				{/if}
				{#if registerbtn}
					<div class="" id="btn-register" />
				{/if}
				<button type="button" class="toggle-btn" on:click={toggleLoginForm}>Login</button>
				<button type="button" class="toggle-btn" on:click={toggleRegisterForm}>Register</button>
			</div>
			{#if showLoginForm}
				<form class="input-group" id="login">
					<div class="modal-body">
						<div class="form-group">
							<label for="username">Username</label>
							<input type="text" class="" bind:value={username} required />
						</div>
						<div class="form-group">
							<!-- svelte-ignore a11y-label-has-associated-control -->
							<label>Password</label>
							<input type="password" class="" bind:value={password} />
						</div>
						<div class="modal-footer">
							<button type="button" class="cancel-btn" on:click={closeModal}>ยกเลิก</button>
							<button type="button" class="submit-btn" on:click={handleLogin}>เข้าสู่ระบบ</button>
						</div>
					</div>
				</form>
			{/if}
			{#if showRegisterForm}
				<form class="input-group" id="register" on:submit|preventDefault={handleSubmit}>
					<div class="modal-body">
						<div class="form-group">
							<!-- svelte-ignore a11y-label-has-associated-control -->
							<label>Username</label>
							<input type="text" class="" bind:value={username} required />
						</div>
						<div class="form-group">
							<!-- svelte-ignore a11y-label-has-associated-control -->
							<label>Password</label>
							<input type="password" class="" bind:value={password} />
						</div>
						<div class="form-group">
							<!-- svelte-ignore a11y-label-has-associated-control -->
							<label>Fullname</label>
							<input type="text" class="" bind:value={fullname} />
						</div>
						<div class="form-group">
							<!-- svelte-ignore a11y-label-has-associated-control -->
							<label>Address</label>
							<input type="text" class="" bind:value={address} />
						</div>
						<div class="form-group">
							<!-- svelte-ignore a11y-label-has-associated-control -->
							<label>Birth Day</label>
							<input type="date" class="" bind:value={birthday} />
						</div>
						<div class="form-group">
							<!-- svelte-ignore a11y-label-has-associated-control -->
							<label>Tel</label>
							<input type="text" class="" bind:value={tel} />
						</div>
						<div class="modal-footer">
							<button type="button" class="cancel-btn" on:click={closeModal}>ยกเลิก</button>
							<button type="button" class="submit-btn" on:click={handleSubmit}>สมัครสมาชิก</button>
							<!-- <button type="submit" class="submit-btn" on:click={handleSubmit}>สมัครสมาชิก</button> -->
						</div>
					</div>
				</form>
			{/if}
		</div>
	</div>
</div>

<style>
	@import url('https://fonts.googleapis.com/css2?family=Dosis&display=swap');
	.modal-content {
		border: 1px solid rgb(238, 77, 45);
		background-color: rgba(0, 0, 0, 0.7);
		box-shadow: 0 0 20px 9px #ff61241f;
		font-family: 'Dosis', sans-serif;
	}
	label {
		color: #fff;
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
	.btn-close {
		margin: 1rem 1rem;
		right: 0;
		position: absolute;
	}
	.modal-footer {
		margin: 1rem;
	}
	.modal-header {
		width: 245px;
		height: 50px;
		margin: 35px auto;
		position: relative;
		box-shadow: 0 0 20px 9px #ff61241f;
		border-radius: 30px;
    	color: #fff;
	}
	.modal-header button {
		color: #fff;
		font-size: 1.2rem;
		font-weight: bold;
	}
	.toggle-btn {
		padding: 10px 30px;
		cursor: pointer;
		background: transparent;
		border: 0;
		outline: none;
		position: relative;
	}
	#btn-login {
		top: 0;
		left: 0;
		position: absolute;
		width: 130px;
		height: 100%;
		background: linear-gradient(to right, #ff105f, #ffad06);
		border-radius: 30px;
		transition: 2s;
	}
	#btn-register {
		top: 0;
		left: 7.5rem;
		position: absolute;
		width: 125px;
		height: 100%;
		background: linear-gradient(to right, #ff105f, #ffad06);
		border-radius: 30px;
		transition: 2s;
	}
	.modal-body {
		width: 100%;
		padding: 10px auto;
		margin: 5px 0;
	}
	.input-group {
		position: relative;
		top: 5rem;
		transition: 2s;
	}
	input {
		width: 100%;
		padding: 0.5rem 01rem;
		margin: 5px 0;
		border-top: 0;
		border-left: 0;
		border-right: 0;
		border-bottom: 1px solid #999;
		outline: none;
		background: transparent;
		border-radius: 10px;
		color: rgb(238, 77, 45);
		font-weight: bolder;
	}
	.submit-btn {
		width: 50%;
		padding: 1rem 2rem;
		display: block;
		background: linear-gradient(to right, #ff105f, #ffad06);
		border: 0;
		outline: none;
		border-radius: 2rem;
	}
	.cancel-btn {
		width: 30%;
		padding: 1rem 1.5rem;
		display: block;
		border: 0;
		outline: none;
		border-radius: 2rem;
		background: linear-gradient(to right, #313030, #b6b5b2);
	}
	#login {
		left: 0rem;
	}
</style>
