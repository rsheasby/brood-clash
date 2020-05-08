<script lang="typescript">
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
		await sleep(1000);

		try {
			await login(event.target.code.value);
		} catch (e) {
			error = true;
			loading = false;
		}
	}

	async function login(code: string) {
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

	function sleep(duration: number): Promise<void> {
		return new Promise((resolve) => {
			setTimeout(() => resolve(), duration);
		});
	}

	function requestAnimationFrame(): Promise<void> {
		return new Promise((resolve) => {
			window.requestAnimationFrame(() => resolve());
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

<div class="flex-center hw-full p-5">
	<form
		class="max-w-full"
		novalidate="novalidate"
		on:submit|preventDefault={handleLogin}>
		<div
			class="card flex flex-row"
			class:shake={error}>
			<input
				type="text"
				class="input rounded-r-none sm:text-xl min-w-0 flex-shrink"
				class:border-red-600={error}
				placeholder="Code"
				name="code"
				minlength="4"
				maxlength="4"
				pattern="\d{'{'}4{'}'}"
				autocomplete="off"
				disabled={loading} />
			<button
				type="submit"
				class="button-primary rounded-l-none border-l-0 sm:text-xl
				flex-none w-24 border-l-0"
				class:loading>
				Log in
			</button>
		</div>
	</form>
</div>
