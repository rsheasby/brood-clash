<script lang="typescript">
	import { onMount } from 'svelte';
	import { ApiClient, codegen } from 'api';
	import { goto } from '@sveltech/routify';

	let question: codegen.ModelsQuestion;

	onMount(async () => {
		try {
			question = await loadQuestion();
		} catch (e) {
			console.error(e);
			history.replaceState({}, '', '/controller/questions');
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
				throw new Error('Something unexpected happened.');
			}
		} catch (e) {
			console.error(e);
			throw new Error('Something unexpected happened.');
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

<div class="flex-center hw-full p-5 sm:text-xl">
	{#if !question}
		<div class="loading text-6xl" />
	{:else}
		<div
			class="card grid grid-cols-1 gap-3 max-w-full"
			style="width: 600px;">

			<div
				class="text-gray-200 bg-blue-600 p-3 font-semibold -m-5 mb-0
				rounded-t-sm rounded-b-none">
				{question.text}
			</div>

			{#each question.answers as answer (answer.id)}
				<button
					class:button-neutral={answer.revealed}
					class:button-primary={!answer.revealed}
					class:cursor-default={answer.revealed}
					class="max-w-full hover:bg-gray-600 active:bg-gray-600"
					on:click={() => {
						if (!answer.revealed) reveal(answer);
					}}>
					{answer.points} - {answer.text}{answer.revealed ? ' ✔' : ''}
				</button>
			{/each}

			<button class="button-warning" on:click={wrongAnswer}>
				Wrong answer.
			</button>

			<button
				on:click={$goto('/controller/questions')}
				class="button-primary -m-5 mt-0 rounded-t-none rounded-b-sm">
				Back to questions
			</button>
		</div>
	{/if}
</div>
