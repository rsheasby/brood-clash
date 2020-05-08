<script lang="typescript">
	import { onMount } from 'svelte';
	import { url } from '@sveltech/routify';

	import { ApiClient, codegen } from 'api';

	let questionPromise: Promise<codegen.ModelsQuestion> = new Promise((resolve) => {});

	onMount(async () => {
		questionPromise = loadQuestion();
	});

	async function loadQuestion(): Promise<codegen.ModelsQuestion> {
		try {
			const response = await ApiClient.getCurrentQuestion();
			const status = response && response.status;
			if (status === 200) {
				return response.data;
			} else if (status === 404) {
				history.replaceState({}, '', '/controller/questions');
			} else {
				throw new Error("Something unexpected happened.");
			}
		} catch (e) {
			console.error(e);
			throw new Error("Something unexpected happened.");
		}
	}
</script>

<svelte:head>
	<style>
		body {
			background-color: #001f54;
		}
	</style>
</svelte:head>

{#await questionPromise}
	<p>Loading, please wait...</p>
{:then question}
	<p>{question.text}</p>
	{#each question.answers as answer}
		<p>{answer.points} - {answer.text}</p>
	{/each}
{:catch error}
	<p>{error}</p>
{/await}

<a class="text-blue-500 hover:text-blue-800" href="/controller/questions">
	<button>Back to questions</button>
</a>
