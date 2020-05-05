<script lang="typescript">
	import { goto } from '@sveltech/routify';

	import { ApiClient, SetApiKey } from 'api';

	const UNEXPECTED_BEHAVIOR: string = "Something unexpected happened. Please try again later.";
	const INVALID_CODE: string = "The provided code is invalid.";
	const CODE_TOO_SHORT: string = "The code must be 4 digits.";

	let loading: boolean = false;
	let wrongCode: boolean = false;
	let shakeForm: boolean = false;
	let error: string = undefined;

	async function handleLogin(event) {
		if (loading) {
			return;
		}

		loading = true;
		error = undefined;

		try {
			await login(event.target.code.value);
		} catch (e) {
			console.error(e);

			error = UNEXPECTED_BEHAVIOR;
			loading = false;
		}
	}

	async function login(code: string) {
		// TODO: Temporary delay so you can see the effect of the loading state.
		// await sleep(2000);

		if (!/^\d{4}$/.test(code)) {
			error = CODE_TOO_SHORT;
			loading = false;
			return;
		}

		SetApiKey(code);
		ApiClient.test().then(response => {
			loading = false;
			if (response.status === 204) {
				$goto('/controller');
			} else {
				error = UNEXPECTED_BEHAVIOR;
				return;
			}
		}).catch(error => {
			loading = false;
			if (error.response.status === 401) {
				wrongCode = true;
				shakeForm = true;
				setTimeout(() => { shakeForm = false }, 1000);
				return;
			} else {
				error = UNEXPECTED_BEHAVIOR;
				return;
			}
		});
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
		animation: shake 0.82s cubic-bezier(.36,.07,.19,.97) both;
		transform: translate3d(0, 0, 0);
		backface-visibility: hidden;
		perspective: 1000px;
	}

	@keyframes shake {
		10%, 90% {
		transform: translate(-1px, 0);
		}

		20%, 80% {
		transform: translate(2px, 0);
		}

		30%, 50%, 70% {
		transform: translate(-4px, 0);
		}

		40%, 60% {
		transform: translate(4px, 0);
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
	<form class="bg-gray-400 p-5 rounded max-w-full" class:shake="{shakeForm}" novalidate="novalidate" on:submit|preventDefault="{handleLogin}">
		<input type="text" class="p-3 w-64 max-w-full rounded-t-md border-2 {wrongCode ? "border-red-600" : "border-b"}" placeholder="Code" name="code" minlength="4" maxlength="4" pattern="\d{'{'}4{'}'}" disabled="{loading}" /><br />
		<button type="submit" class="p-3 bg-blue-600 active:bg-blue-700 hover:bg-blue-500 w-64 max-w-full rounded-b-md border-2 border-t-0 text-gray-200">Log in</button>
		{#if !!error}
			<div class="w-64 max-w-full rounded-md bg-red-500 border-red-700 border-2 mt-3 p-3 text-gray-900">{error}</div>
		{/if}
	</form>
</div>