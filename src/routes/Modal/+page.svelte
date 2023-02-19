<script>
	import { createEventDispatcher } from 'svelte';
	import Swal from 'sweetalert2';
	import { goto } from '$app/navigation';
	
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
  
	function handleSubmit() {
	  dispatch('submit', inputVal);
	  closeModal();
	  Swal.fire({
					position: 'center',
					icon: 'success',
					title: `Register result : Success `,
					showConfirmButton: true,
				}
			).then((value) =>{
				close();
				goto('/', { state: { foo: 'bar' } });
				
        	})
	}
  </script>
  
  <button on:click={openModal}>Open Modals</button>
  
  <div class="modal" tabindex="-1" role="dialog" style={showModal ? 'display: block;' : ''}>
	<div class="modal-dialog" role="document">
	  <div class="modal-content">
		<div class="modal-header">
		  <h5 class="modal-title">Register Form</h5>
		  <button type="button" class="btn-close" data-dismiss="modal" aria-label="Close" on:click={closeModal}>
			<span aria-hidden="true"></span>
		  </button>
		</div>
		<div class="modal-body">

		  <input type="text" bind:value={inputVal}>

		</div>
		<div class="modal-footer">
		  <button type="button" class="btn btn-secondary" on:click={closeModal}>Close</button>
		  <button type="button" class="btn btn-primary" on:click={handleSubmit}>Save changes</button>
		</div>
	  </div>
	</div>
  </div>