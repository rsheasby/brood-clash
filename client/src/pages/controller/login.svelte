<script lang="typescript">
	import { tick } from 'svelte';
	import { goto } from '@sveltech/routify';

	import { ApiClient, SetApiKey } from 'api';

	let loading: boolean = false;
	let error: boolean = false;

	async function handleLogin(event) {
		if (loading) {
			return;
		}

		loading = true;
		error = false;
		await tick();

		try {
			await login(event.target.code.value);
		} catch (e) {
			error = true;
			loading = false;
		}
	}

	async function login(code: string) {
		if (!/^\d{4}$/.test(code)) {
			error = true;
			loading = false;
			return;
		}

		await sleep(500);

		SetApiKey(code);
		try {
			let response = await ApiClient.test();
			if (response.status === 204) {
				$goto('/controller');
			} else {
				loading = false;
				error = true;
				return;
			}
		} catch (e) {
			loading = false;
			error = true;
			return;
		}
	}

	// TODO: Temporary function so you can see the effects of the loading state.
	function sleep(duration: number): Promise<void> {
		return new Promise((resolve) => {
			setTimeout(() => resolve(), duration);
		});
	}
</script>

<style>
	.shake {
		animation: shake 0.82s cubic-bezier(0.36, 0.07, 0.19, 0.97) both;
		transform: translate3d(0, 0, 0);
		backface-visibility: hidden;
		perspective: 1000px;
	}

	@keyframes shake {
		10%,
		90% {
			transform: translate(-3px, 0);
		}

		20%,
		80% {
			transform: translate(4px, 0);
		}

		30%,
		50%,
		70% {
			transform: translate(-6px, 0);
		}

		40%,
		60% {
			transform: translate(6px, 0);
		}
	}
</style>

<svelte:head>
	<style>
		body {
			background-color: #001f54;
		}
	</style>
</svelte:head>

<div class="flex items-center justify-center h-full w-full p-5">
	<form
		class="max-w-full"
		novalidate="novalidate"
		on:submit|preventDefault={handleLogin}>
		<div
			class="bg-gray-400 p-5 rounded max-w-full flex flex-row"
			class:shake={error}>
			<input
				type="text"
				class="p-3 sm:text-xl min-w-0 border-gray-500 flex-shrink
				rounded-l-md border-2" class:border-red-600="{error}"
				placeholder="Code"
				name="code"
				minlength="4"
				maxlength="4"
				pattern="\d{'{'}4{'}'}"
				disabled={loading} />
			<button
				type="submit"
				class="p-3 sm:text-xl border-gray-500 bg-blue-600
				active:bg-blue-700 hover:bg-blue-500 flex-none w-24 rounded-r-md
				border-2 border-l-0 text-gray-200" class:loading="{loading}">
				Log in
			</button>
		</div>
	</form>
</div>
