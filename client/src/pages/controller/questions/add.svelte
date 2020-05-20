<script lang="typescript">
	import { goto } from '@sveltech/routify';

	import { ApiClient, codegen } from 'api';

	let question: codegen.ModelsQuestion = newQuestion();

	function newQuestion(): codegen.ModelsQuestion {
		return {
			answers: [ newAnswer() ],
			text: ""
		};
	}

	function newAnswer(): codegen.ModelsAnswer {
		return {
			points: undefined,
			text: ""
		};
	}

	/**
	 * @return true if the question is successfully created, false otherwise.
	 */
	async function submitQuestion(): Promise<boolean> {
		try {
			const response = await ApiClient.postQuestions([ question ]);
			const status = response && response.status;
			return status === 201;
		} catch (e) {
			console.error(e);
			return false;
		}
	}

	async function submitAndAddAnother() {
		const success = await submitQuestion();
		if (!success) {
			// TODO: not sure how to inform the user the creation failed.
		} else {
			question = newQuestion();
		}
	}

	async function submitAndBackToQuestions() {
		const success = await submitQuestion();
		if (!success) {
			// TODO: Not sure how to inform the user the creation failed.
		} else {
			$goto('/controller/questions');
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
	<div class="card grid grid-cols-1 gap-3">
		<form id="questions-form" class="grid grid-cols-1 gap-3">
			<input
				class="input box-content w-auto max-w-full" type="text"
				required placeholder="Question"
				bind:value={question.text} />
			{#each question.answers as answer, i}
				<div class="grid grid-cols-4 gap-3">
					<input class="input col-span-3 box-content w-auto
						max-w-full" type="text" required
						placeholder="Answer {i + 1}"
						bind:value={answer.text}/>
					<input class="input col-span-1 box-content w-auto
						max-w-full" type="number" required
						placeholder="Points"
						bind:value={answer.points} />
				</div>
			{/each}
			{#if question.answers.length < 8}
				<!--
					Alternatively we can have a number input box where you type how
					many answers you want, or have a pseudo input field for a new
					question, then when you start typing in it, add another question
					and turn it into a real input box. The problem with this is it
					will defocus once you make that pseudo box a real box and I don't
					want to deal with the complexities of that right now.
				-->
				<button class="button-neutral box-content w-auto w-max-full"
					on:click={() => question.answers = [ ...question.answers, newAnswer() ]}>
					Add another answer
				</button>
			{/if}
		</form>
		<div class="grid grid-cols-3 gap-3">
			<button class="button-warning box-content w-auto w-max-full"
				on:click={() => $goto('/controller/questions')}>
				Cancel
			</button>

			<button class="button-primary box-content w-auto w-max-full"
				on:click={submitAndAddAnother}>
				Submit and add another
			</button>

			<button class="button-primary box-content w-auto w-max-full"
				on:click={submitAndBackToQuestions}>
				Submit
			</button>
		</div>
	</div>
</div>
