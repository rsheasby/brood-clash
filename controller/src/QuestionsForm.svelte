<script>
	import { createEventDispatcher, onMount } from "svelte";
	import page from "page";
	import axios from "axios";

	const dispatch = createEventDispatcher();

	let questions = [];
	let errors = [];

	onMount(() => addQuestion());

	function addQuestion() {
		let question = {
			description: "",
			answers: []
		};
		for (let i = 0; i < 8; ++i) {
			question.answers.push({
				text: "",
				points: 0
			});
		}
		questions = [...questions, question];
	}

	async function handleSubmit() {
		errors = [];
		if (questions.length < 1) {
			errors = [...errors, "Must have at least question to submit."];
			return;
		}

		for (let question of questions) {
			if (question.description.length < 1) {
				errors = [...errors, "Question cannot be an empty string."];
				return;
			}
			if (question.answers.some(a => a.length < 1)) {
				errors = [...errors, "Answers cannot be empty strings."];
				return;
			}
		}

		try {
			const response = await axios({
				url: "http://localhost:1323/api/questions",
				method: "post",
				headers: { "Authorization": `Bearer ${window.localStorage.getItem("code")}` },
				data: questions
			});
			page.redirect("/controller");
		} catch (e) {
			console.log(e);
			return;
		}

	}
</script>

<style>
.error {
	color: red;
}
</style>

<form on:submit|preventDefault="{handleSubmit}">
	{#each questions as question (question.id)}
		Question:<br />
		<input type="text" bind:value="{question.description}" on:keydown|stopPropagation>
		<br />
		Answers:<br />
		{#each question.answers as answer, index}
			{index + 1}:
			<input type="text" bind:value="{answer.text}" on:keydown|stopPropagation>
			<input type="number" bind:value={answer.points} on:keydown|stopPropagation>
			<br />
		{/each}
	{/each}

	<button type="submit" on:keydown|stopPropagation>Submit</button>
</form>

<button on:click={addQuestion} on:keydown|stopPropagation>Add question</button>

{#each errors as error}
<p class="error">{error}</p>
{/each}
