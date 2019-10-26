<script>
	import {createEventDispatcher } from "svelte";

	const dispatch = createEventDispatcher();

	let invalidCode;

	function handleSubmit(event) {
		let loginCode = event.target.loginCode.value;

		// TODO: Use login from the server.
		if (loginCode === "secret") {
			dispatch("login", { sessionToken: "16ha051hheef9581" });
		} else {
			invalidCode = true;
		}
	}
</script>

<style>
	.error {
		color: red;
	}
</style>

<form on:submit|preventDefault="{handleSubmit}">
	<label for="login-code-input">Please enter your login code:</label>
	<input type="text" id="login-code-input" placeholder="Login code" name="loginCode" required on:keydown|stopPropagation><br />

	<button type="submit" on:keydown|stopPropagation>Login</button>
</form>

{#if invalidCode}
<div class="error">Invalid login code!</div>
{/if}
