<script>
	import { createEventDispatcher, onDestroy } from 'svelte';
    import { slide } from 'svelte/transition';

	const dispatch = createEventDispatcher();
	export const close = () => {
		//save data from modal
		dispatch('close');
	};

	// @ts-ignore
	let modal;

	// @ts-ignore
	const handle_keydown = e => {
		if (e.key === 'Escape') {
			close();
			return;
		}

		if (e.key === 'Tab') {
			// trap focus
			// @ts-ignore
			const nodes = modal.querySelectorAll('*');
			const tabbable = Array.from(nodes).filter(n => n.tabIndex >= 0);

			let index = tabbable.indexOf(document.activeElement);
			if (index === -1 && e.shiftKey) index = 0;

			index += tabbable.length + (e.shiftKey ? -1 : 1);
			index %= tabbable.length;

			tabbable[index].focus();
			e.preventDefault();
		}
	};
</script>

<svelte:window on:keydown={handle_keydown}/>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="modal-background" on:click={close}></div>

<div class="modal rounded-lg z-50" transition:slide role="dialog" aria-modal="true" bind:this={modal}>
	<div class="flex justify-between items-center">
		<slot name="header"></slot>
		<!-- svelte-ignore a11y_consider_explicit_label -->
		<button class="text-white cursor-pointer" on:click={close} id="closebutton">
			<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-x hover:opacity-75 transition" width="24" height="24" viewBox="0 0 24 24" stroke-width="1.5" stroke="white" fill="none" stroke-linecap="round" stroke-linejoin="round">
				<path stroke="none" d="M0 0h24v24H0z" fill="none"/>
				<path d="M18 6l-12 12" />
				<path d="M6 6l12 12" />
			  </svg>
		</button>
	</div>
	<slot></slot>
	<slot name="footer"></slot>
	<!-- svelte-ignore a11y-autofocus -->
</div>

<style>
	.modal-background {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: rgba(0,0,0,0.5);
	}

	.modal {
		position: absolute;
		left: 50%;
		top: 50%;
		width: calc(100vw - 6em);
		max-width: 30em;
		max-height: calc(100vh - 4em);
		overflow: auto;
		transform: translate(-50%,-50%);
		padding: 2em;
		background: rgb(23, 23, 23);
	}

	#closebutton {
		display: block;
		outline: none;
	}
</style>