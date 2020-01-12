<script>
	import axios from "axios";
	import page from "page";

	let invalidCode;

	async function handleSubmit(event) {
		const code = event.target.loginCode.value;
		try {
			const response = await axios({
				url: "http://localhost:1323/api/code/valid",
				headers: { "Authorization": `Bearer ${code}` }
			});
			window.localStorage.setItem("code", code);
			window.location = "/controller";
		} catch (e) {
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
