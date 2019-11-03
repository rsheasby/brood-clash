<script>
	import { createEventDispatcher } from "svelte";
	import axios from "axios";

	const dispatch = createEventDispatcher();

	let invalidCode;

	async function handleSubmit(event) {
		const loginCode = event.target.loginCode.value;
		const response = await axios({
			url: "http://localhost:3000/login",
			method: "POST",
			data: { LoginCode: loginCode }
		});
		console.log(response);
		/*if (response && response.SessionToken && ) {
			dispatch("login", { sessionToken: response.SessionToken });
		} else {
			invalidCode = true;
		}*/
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
