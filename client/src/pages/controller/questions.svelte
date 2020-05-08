<script lang="typescript">
	import { onMount } from 'svelte';

	import { ApiClient, codegen } from 'api';

	import { goto } from '@sveltech/routify';

	let loading: boolean = true;

	let questions: codegen.ModelsQuestion[] = [];

	let error: string;

	let currentQuestionExists: boolean = false;

	onMount(async () => {
		let currentQuestionPromise = ApiClient.getCurrentQuestion();
		try {
			let response = await ApiClient.getAllQuestions();
			let currentQuestion = (await currentQuestionPromise).data;
			if (currentQuestion && currentQuestion.id) {
				questions = response.data.filter(
					(q) => q.id != currentQuestion.id
				);
				currentQuestionExists = true;
			}
			questions = response.data;
		} catch (e) {
			console.error(e);
		}

		loading = false;
	});

	async function selectQuestion(questionId) {
		loading = true;
		await ApiClient.selectQuestion(questionId);
		$goto('/controller');
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
{/if}

<div class="flex-center hw-full p-5 sm:text-xl">
	{#if loading}
		<div class="loading text-6xl" />
	{:else}
		<div class="card grid grid-cols-1 gap-3">
			{#if currentQuestionExists}
				<button
					class="button-primary -m-5 mb-0 rounded-t-sm rounded-b-none
					border-0"
					on:click={$goto('/controller')}>
					Back to Current Question
				</button>
			{/if}
			{#each questions as question}
				<button
					class="button-primary box-content w-64 max-w-full"
					on:click={selectQuestion(question.id)}>
					{question.text}
				</button>
			{/each}
		</div>
	{/if}
</div>
