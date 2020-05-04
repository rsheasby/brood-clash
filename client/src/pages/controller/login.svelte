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
	.container {
		width: 100%;
		height: 100%;

		display: flex;
		justify-content: center;
		align-items: center;

		background-color: #001f54;
	}

	#submit-button {
		width: 135px;
	}

	.shake {
		animation: shake 0.82s cubic-bezier(.36,.07,.19,.97) both;
		transform: translate3d(0, 0, 0);
		backface-visibility: hidden;
		perspective: 1000px;
	}

	@keyframes shake {
		10%, 90% {
		transform: translate3d(-1px, 0, 0);
		}

		20%, 80% {
		transform: translate3d(2px, 0, 0);
		}

		30%, 50%, 70% {
		transform: translate3d(-4px, 0, 0);
		}

		40%, 60% {
		transform: translate3d(4px, 0, 0);
		}
	}
</style>

<div class="container">
	<form on:submit|preventDefault="{handleLogin}">
		<div class="ui massive right action input {error ? 'error shake' : ''}">
			<input type="text" placeholder="Code" name="code" minlength="4" maxlength="4" pattern="\d{'{'}4{'}'}" disabled="{loading}" />
			<button id="submit-button" type="submit" class="ui massive primary button" class:loading="{loading}">{loading ? 'Loading' : 'Log in'}</button>
		</div>
	</form>
</div>