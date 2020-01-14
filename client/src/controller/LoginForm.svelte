<script>
	import page from "page";

	let invalidCode;

</script>

<style>
	.error {
		color: red;
	}
</style>

<!--<form on:submit|preventDefault="{handleSubmit}">-->
<!--	<label for="login-code-input">Please enter your login code:</label>-->
<!--	<input type="text" id="login-code-input" placeholder="Login code" name="loginCode" required on:keydown|stopPropagation><br />-->

<!--	<button type="submit" on:keydown|stopPropagation>Login</button>-->
<!--</form>-->

{#if invalidCode}
<div class="error">Invalid login code!</div>
{/if}
