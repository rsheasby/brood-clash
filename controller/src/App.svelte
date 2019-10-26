<script>
	import LoginForm from "./LoginForm.svelte";
	import QuestionsForm from "./QuestionsForm.svelte";

	let sessionToken;

	function handleLogin(event) {
		sessionToken = event.detail.sessionToken;
	}
</script>

<!-- TODO: We probably need a router -->
{#if !sessionToken}
	<LoginForm on:login="{handleLogin}"/>
{:else}
	<QuestionsForm />
{/if}
