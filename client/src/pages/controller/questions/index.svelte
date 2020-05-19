<script lang="typescript">
	import { onMount } from 'svelte';

	import { ApiClient, codegen } from 'api';

	import { goto } from '@sveltech/routify';

	let loading: boolean = true;

	let questions: codegen.ModelsQuestion[] = [];

	let error: string;

	let currentQuestionExists: boolean = false;

	let deletingQuestionId = undefined;
	let deletingTimeoutId = undefined;

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
			} else {
				questions = response.data;
			}
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

	async function deleteQuestionClicked(question) {
		if (deletingTimeoutId) {
			clearTimeout(deletingTimeoutId);
			deletingTimeoutId = undefined;
		}


		if (deletingQuestionId === question.id) {
			await deleteQuestion(question);
		} else {
			await promptDeleteConfirm(question);
		}
	}

	async function deleteQuestion(question) {
		try {
			question.deleting = true;
			const response = await ApiClient.deleteQuestion(question.id);
			const status = response && response.status;
			if (status === 204) {
				questions = questions.filter(q => q.id !== question.id);
			}
		} catch (e) {
			// TODO: Not sure how to inform user deletion was unsuccessful.
		} finally {
			question.deleting = false;
		}
	}

	async function promptDeleteConfirm(question) {
		deletingQuestionId = question.id;
		deletingTimeoutId = setTimeout(
			() => {
				deletingQuestionId = null;
				deletingTimeoutId = null;
			}, 5000
		);
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
		<div
			class="card grid grid-cols-1 gap-3 max-w-full"
			style="width: 600px;">
			{#if currentQuestionExists}
				<button
					class="button-primary -m-5 mb-0 rounded-t-sm rounded-b-none
					border-0"
					on:click={$goto('/controller')}>
					Back to Current Question
				</button>
			{/if}

			<div class="grid grid-cols-4 gap-1 max-w-full">
				{#each questions as question}
					<button
						class="box-content col-span-3 max-w-full"
						class:button-primary={!question.hasBeenShown}
						class:button-neutral={question.hasBeenShown}
						on:click={selectQuestion(question.id)}>
						{question.text}{question.hasBeenShown ? " ✔" : ""}
					</button>
					<button
						class="button-warning box-content max-w-full"
						on:click={deleteQuestionClicked(question)}
						disabled={question.deleting}>
						{ question.id === deletingQuestionId ? "Confirm" : "Delete" }
					</button>
				{/each}
			</div>

			<button
				class="button-primary -m-5 rounded-b-sm rounded-t-none border-0"
				class:mt-0={questions.length !== 0 || currentQuestionExists}
				on:click={$goto('/controller/questions/add')}>
				Add Questions
			</button>
		</div>
	{/if}
</div>
