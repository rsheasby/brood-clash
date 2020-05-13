<script lang="typescript">
	import { onMount } from 'svelte';
	import { url } from '@sveltech/routify';

	import { ApiClient, codegen } from 'api';

	let question;
	let error;

	onMount(async () => {
		try {
			question = await loadQuestion();
		} catch (e) {
			error = e;
		}
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

	async function reveal(answer: codegen.ModelsAnswer) {
		try {
			if (answer.revealed) {
				return;
			}

			const response = await ApiClient.revealAnswer(answer.id);
			const status = response && response.status;
			if (status === 204) {
				answer.revealed = true;
				question = question;
			} else {
				// TODO: Notify the user that something went wrong.
			}
		} catch (e) {
			console.error(e);
			// TODO: Notify the user that something went wrong.
		}
	}

	async function wrongAnswer() {
		try {
			const response = await ApiClient.incorrectAnswer();
			const status = response && response.status;
			if (status !== 200) {
				// TODO: Notify the user that something went wrong.
			}
		} catch (e) {
			console.error(e);
			// TODO: Notify the user that something went wrong.
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

{#if error}
	<p>{error}</p>
{:else if !question}
	<p>Loading, please wait...</p>
{:else}
	<p style="color: white">{question.text}</p>
	{#each question.answers as answer (answer.id)}
		<p style="color: white">{answer.points} - {answer.text}</p>
		{#if !answer.revealed}
			<button on:click={() => reveal(answer)}>Reveal</button>
		{:else}
			<p>Revealed!</p>
		{/if}
	{/each}
{/if}


<button on:click={wrongAnswer}>Wrong answer.</button>

<a class="text-blue-500 hover:text-blue-800" href="/controller/questions">
	<button>Back to questions</button>
</a>
