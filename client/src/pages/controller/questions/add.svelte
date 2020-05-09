<script lang="typescript">
	import { goto } from '@sveltech/routify';

	import { ApiClient, codegen } from 'api';

	let questions = [ newQuestion() ];

	function newQuestion(): codegen.ModelsQuestion {
		let question: codegen.ModelsQuestion = {
			answers: [],
			text: ""
		};

		for (let i = 0; i < 8; ++i) {
			question.answers.push(newAnswer());
		}

		return question;
	}

	function newAnswer(): codegen.ModelsAnswer {
		return {
			points: undefined,
			text: ""
		};
	}

	async function submitQuestions() {
		try {
			// TODO: not sure how to inform the user whether the operation was
			// successful.
			const response = await ApiClient.postQuestions(questions);
			const status = response && response.status;
			if (status === 201) {
				$goto('/controller/questions');
			} else if (status === 202) {
				// TODO: Not sure what to do here. Not sure how to tell which
				// questions were created successfully since it doesn't look
				// like the function returns the created questions.
			}
		} catch (e) {
			// TODO: Iono what to do???
		}
	}

	function addQuestion() {
		questions = [ ...questions, newQuestion() ];
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
		<form
			id="questions-form" class="grid grid-cols-1 gap-3"
			on:submit|preventDefault={submitQuestions}>
			{#each questions as question}
				<input
					class="input box-content w-auto max-w-full" type="text"
					required placeholder="Question"
					bind:value={question.text} />
				{#each question.answers as answer, i}
					<div class="grid grid-cols-4 gap-3">
						<input class="input col-span-3 box-content w-auto
							max-w-full" type="text" required
							placeholder="Answer {i}"
							bind:value={answer.text}/>
						<input class="input col-span-1 box-content w-auto
							max-w-full" type="number" required
							placeholder="Points"
							bind:value={answer.points} />
					</div>
				{/each}
			{/each}
		</form>
		<div class="grid grid-cols-3 gap-3">
			<button class="button-warning box-content w-auto w-max-full"
				on:click={() => $goto('/controller/questions')}>
				Cancel
			</button>
			<button class="button-primary box-content w-auto w-max-full"
				on:click={addQuestion}>Add question</button>
			<button class="button-primary box-content w-auto w-max-full"
				type="submit" form="questions-form">Submit</button>
		</div>
	</div>
</div>