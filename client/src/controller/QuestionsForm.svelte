<script>
	import { onMount } from "svelte";
	import "file-saver";
	import page from "page";
	import axios from "axios";

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

	// TODO: This looks ugly, dunno how to make it look better, but I think it's
	// the right way of doing it.
	let saveQuestions = ((() => {
		try {
			const isFileSaverSupported = !!new Blob;
			if (!isFileSaverSupported) {
				return undefined;
			}

			return () => {
				const questionsJSON = JSON.stringify([questions]);
				const blob = new Blob(questionsJSON, {type: "text/plain"});
				saveAs(blob, "questions.json");
			}
		} catch (e) {
			return undefined;
		}
	})())
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
{#if saveQuestions}
<button on:click={saveQuestions} on:keydown|stopPropagation>Save to file</button>
{/if}

{#each errors as error}
	<p class="error">{error}</p>
{/each}
