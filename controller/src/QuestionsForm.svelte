<script>
	import { createEventDispatcher, onMount } from "svelte";

	const dispatch = createEventDispatcher();

	let questions = [];

	onMount(() => addQuestion());

	function addQuestion() {
		function nextId() {
			let latestQuestion = questions[questions.length - 1];
			if (latestQuestion) {
				return latestQuestion.id + 1;
			} else {
				return 0;
			}
		}

		let question = {
			id: nextId(),
			question: "",
			answers: [
				"", "", "", "",
				"", "", "", ""
			]
		};
		questions = [...questions, question];
	}

	function handleSubmit() {
		dispatch("submit", questions);
	}
</script>

<form on:submit|preventDefault="{handleSubmit}">
	{#each questions as question (question.id)}
	Question:<br />
	<input type="text" bind:value="{question.question}"><br />
	Answers:<br />
		{#each question.answers as answer, index}
		{index + 1}: <input type="text" bind:value="{answer}"><br />
		{/each}
	{/each}

	<button type="submit">Submit</button>
</form>

<button on:click={addQuestion}>Add question</button>
