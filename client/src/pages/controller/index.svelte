<script lang="typescript">
	import { onMount } from 'svelte';
	import { url } from '@sveltech/routify';

	import { ApiClient, codegen } from 'api';

	let question: codegen.ModelsQuestion;

	let loading: boolean = true;

	onMount(async () => {
		try {
			let response = await ApiClient.getCurrentQuestion();
			const status = response && response.status;
			if (status === 200) {
				question = response.data;
			} else if (status === 404) {
				// This is so the back button works correctly,
				// instead of infinitely redirecting forward again
				history.replaceState({}, '', '/controller/questions');
			}
		} catch (e) {
			console.error(e);
		}

		loading = false;
	});
</script>

<svelte:head>
	<style>
		body {
			background-color: #001f54;
		}
	</style>
</svelte:head>

{#if loading}
	<p>Loading, please wait...</p>
{:else if question}
	<p>{question.text}</p>
	{#each question.answers as answer}
		<p>{answer.points} - {answer.text}</p>
	{/each}
{/if}

<a class="text-blue-500 hover:text-blue-800" href="/controller/questions">
	<button>Back to questions</button>
</a>
