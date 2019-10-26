<script>
	import LoginForm from "./LoginForm.svelte";
	import QuestionsForm from "./QuestionsForm.svelte";

	let sessionToken;

	function handleLogin(event) {
		sessionToken = event.detail.sessionToken;
	}

	function handleQuestionsSubmit(event) {
		console.log(event);
	}
</script>

<!-- TODO: We probably need a router -->
{#if !sessionToken}
	<LoginForm on:login="{handleLogin}"/>
{:else}
	<QuestionsForm on:submit="{handleQuestionsSubmit}"/>
{/if}
