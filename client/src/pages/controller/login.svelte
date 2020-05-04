<script lang="typescript">
	import { goto } from '@sveltech/routify';

	import { ApiClient, SetApiKey } from 'api';

	const UNEXPECTED_BEHAVIOR: string = "Something unexpected happened. Please try again later.";
	const INVALID_CODE: string = "The provided code is invalid.";

	let loading: boolean = false;
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
		await sleep(2000);

		if (!/^\d{4}$/.test(code)) {
			error = INVALID_CODE;
			loading = false;
			return;
		}

		SetApiKey(code);
		try {
			let response = await ApiClient.test();

			if (response.status === 204) {
				$goto('/controller');
			} else {
				error = UNEXPECTED_BEHAVIOR;
				loading = false;
				return;
			}
		} catch (e) {
			const status = e && e.response && e.response.status;
			if (status === 401) {
				error = INVALID_CODE;
				loading = false;
				return;
			}

			throw e;
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
	form {
		margin: 16px;
	}
</style>

<form on:submit|preventDefault="{handleLogin}">
	<div class="ui action input">
		<input type="text" placeholder="Code" name="code" disabled="{loading}">
		<!-- TODO: the button resizes while loading, maybe set a specific width for it to always be. -->
		<button type="submit" class="ui primary button" class:loading="{loading}">{loading ? 'Loading' : 'Log in'}</button>
	</div>
</form>

{#if error}
	<p>{error}</p>
{/if}