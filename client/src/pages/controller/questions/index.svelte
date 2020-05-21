<script lang="typescript">
	import { onMount } from 'svelte';
	import { goto } from '@sveltech/routify';

	import { ApiClient, codegen } from 'api';

	let loading: boolean = true;

	let questions: codegen.ModelsQuestion[] = [];

	let currentQuestionExists: boolean = false;

	let deletingQuestionId = undefined;
	let deletingTimeoutId = undefined;

	let resetGameTimeoutId = undefined;

	let questionsExist: boolean = false;
	$: questionsExist = questions.length > 0 || currentQuestionExists;

	onMount(init);

	async function init() {
		loading = true;

		questions = [];
		currentQuestionExists = false;

		deletingQuestionId = undefined;
		if (deletingTimeoutId) {
			clearTimeout(deletingTimeoutId);
		}
		deletingTimeoutId = undefined;

		if (resetGameTimeoutId) {
			clearTimeout(resetGameTimeoutId);
		}
		resetGameTimeoutId = undefined;

		try {
			const currentQuestionPromise = getCurrentQuestion();

			let response = await ApiClient.getAllQuestions();

			let currentQuestion = await currentQuestionPromise;
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
	}

	async function getCurrentQuestion() {
		let response = await ApiClient.getCurrentQuestion();
		let status = response && response.status;
		if (status === 200) {
			return response.data;
		} else if (status === 404) {
			return null;
		} else {
			throw {
				message: 'Could not retrieve current question',
				response: response
			};
		}
	}

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
				questions = questions.filter((q) => q.id !== question.id);
			}
		} catch (e) {
			// TODO: Not sure how to inform user deletion was unsuccessful.
		} finally {
			question.deleting = false;
		}
	}

	async function promptDeleteConfirm(question) {
		deletingQuestionId = question.id;
		deletingTimeoutId = setTimeout(() => {
			deletingQuestionId = null;
			deletingTimeoutId = null;
		}, 5000);
	}

	async function handleReset() {
		if (resetGameTimeoutId) {
			clearTimeout(resetGameTimeoutId);
			resetGameTimeoutId = null;
			await reset();
		} else {
			promptResetConfirm();
		}
	}

	async function reset() {
		const response = await ApiClient.resetGameState();
		const status = response && response.status;
		if (status === 204) {
			await init();
		}
	}

	async function promptResetConfirm() {
		resetGameTimeoutId = setTimeout(() => {
			resetGameTimeoutId = null;
		}, 5000);
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

			{#each questions as question}
				<div class="grid grid-cols-4 gap-0 max-w-full">
					<button
						class="box-content col-span-3 max-w-full rounded-r-none
						mr-0"
						class:button-primary={!question.hasBeenShown}
						class:button-neutral={question.hasBeenShown}
						on:click={selectQuestion(question.id)}>
						{question.text}{question.hasBeenShown ? ' ✔' : ''}
					</button>
					<button
						class="button-warning box-content max-w-full
						rounded-l-none ml-0"
						on:click={deleteQuestionClicked(question)}
						disabled={question.deleting}>
						{question.id === deletingQuestionId ? 'Confirm' : 'Delete'}
					</button>
				</div>
			{/each}

			<div class="grid gap-0 max-w-full {questionsExist ? 'grid-cols-2' : 'grid-cols-1'}">
				<button
					class="button-primary -m-5 rounded-b-sm rounded-t-none border-0
						{questionsExist ? 'mt-0 mr-0' : ''}"
					on:click={$goto('/controller/questions/add')}>
					Add Questions
				</button>
				
				{#if questionsExist}
					<button
						class="button-warning -m-5 rounded-b-sm rounded-t-none border-0 mt-0 ml-0"
		  				on:click={handleReset}>
						{resetGameTimeoutId ? 'Confirm reset' : 'Reset questions' }
					</button>
				{/if}
			</div>
		</div>
	{/if}
</div>
