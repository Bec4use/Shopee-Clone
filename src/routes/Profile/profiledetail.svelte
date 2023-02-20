<script>
// @ts-nocheck


	import { onMount } from 'svelte';
    import Swal from 'sweetalert2';

	let userProfile = [];
    let user = {};
    let Fullname ='';
    let Address ='';

	onMount(async () => {
        const token = localStorage.getItem('token');
        
      const response = await fetch('http://localhost:8080/users/profile', {
      headers: {
        'Authorization': `Bearer ${token}`
      }
    });
		const data = await response.json();
		userProfile = data;
        console.log(userProfile);
        user = userProfile.user[0];
	});

    async function handleChange() {
		const data = { Fullname, Address };
        const token = localStorage.getItem('token');
        
		try {
			const res = await fetch('http://localhost:8080/users/update', {
				method: 'PUT',
				headers: {
                    
					'Authorization': `Bearer ${token}`
				},

				body: JSON.stringify(data)
			});

			const json = await res.json();
			console.log(json);
			if (json.status === 'OK') {
				Swal.fire({
					position: 'center',
					icon: 'success',
					title: json.message ,
					showConfirmButton: true,
				}
			).then((value) =>{
				location.reload();
        	})
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


</script>

<div class="User-title">
    <h1>ข้อมูลของฉัน</h1>
    <p>จัดการข้อมูลส่วนตัวคุณเพื่อความปลอดภัยของบัญชีผู้ใช้นี้</p>
</div>
<div class="user-detail d-flex">
    <div class="">
        <form action="">
            <label for="">ต้องการเปลี่ยนชื่อ ?</label>
            <input type="text" name="" id="" class="form-control" bind:value={ Fullname }>
            <label for="">ต้องการเปลี่ยนที่อยู่ ?</label>
            <input type="text" name="" id="" class="form-control"  bind:value={ Address }>
            <div class="d-flex"><p>ชื่อผู้ใช้  : </p> <label for="">{user.Username}</label></div>
            <div class="d-flex"><p>ชื่อจริง - นามสกุล :</p><label for="">{user.Fullname}</label></div>
            <div class="d-flex"><p>Password : </p><label for="">{user.Password}</label></div>
            <div class="d-flex"><p>ที่อยู่ : </p><label for="">{user.Address}</label></div>
            <div class="d-flex"><p>วัน เดือน ปี เกิด :</p><label for="">{user.Birthday}</label></div>
            <div class="d-flex"><p>หมายเลขโทรศัพท์ : </p><label for="">{user.Tel}</label></div>
            <button type="submit" on:click={handleChange}>บันทึก</button>
        </form>
    </div>
    <div class="input-img">
        <input type="file" name="" id="">
    </div>
</div>


<style>
    .User-title{
        height: 5rem;
        width: 58.375rem;
        border-bottom: 0.0625rem solid #efefef;
        margin: 0 auto;
    }
    .User-title h1{
        font-size: 1.2rem;
        padding-top: 1rem;
    }
    .User-title p {
        font-size: .9rem;
    }
    .user-detail {
        width: 100%;
        margin: 2rem auto;
        display: flex;
        justify-content: center;
        color: rgba(0,0,0,.8);
    }
    form {
        padding-right: 2rem;
    }
    .user-detail p {
        font-size:  .9rem;
    }
    .user-detail label {
        font-size:  .9rem; 
        display: inline-block;
        color: #ee4d2d;
    }

    .user-detail input {
        margin-bottom: 1.5rem;
       
    }
    .user-detail button {
        padding: .5rem 1rem;
        outline: none;
        border: none;
        background: #ee4d2d;
        color: #fff;
        border-radius: 2px;
        margin: 2rem 5rem;
    }
    .input-img {
        border-left: 1px solid rgba(0,0,0,.2);
        padding : 8rem 0;
        padding-left: 1rem;
    }
</style>
