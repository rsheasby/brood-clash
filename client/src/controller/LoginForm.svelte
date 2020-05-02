<script>
	import * as api from 'brood_clash';
	import * as auth from '../auth.js';

	let loading = false;
	let error = false;

	async function handleLogin(event) {
		await login(event.target.code.value);
	}

	async function login(code) {
		if (loading) {
			return;
		}

		loading = true;
		error = undefined;

		try {
			auth.setCode(code);
		} catch (e) {
			error = "Invalid code.";
			loading = false;
			return;
		}

		// ApiClient.instance.authentications.ApiKey.apiKey = auth.getCode();
		// const client = new DefaultApi();

		try {
			await client.Test();
			window.location = '/controller/add-questions';
		} catch (e) {
			if (e.status === 403) {
				error = "Invalid code.";
			} else {
				error = "Something unexpected happened. Please refresh and try again.";
			}
			loading = false;
		}
	}
</script>

<form on:submit|preventDefault="{handleLogin}">
	<label for="name">Code</label>
   <div class="ui input">
      <input type="text" name="code" disabled="{loading}">
   </div>
</form>

{#if loading}
	<p>Loading...</p>
{/if}

{#if error}
	<p>{error}</p>
{/if}