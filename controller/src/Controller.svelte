<script>
	import { onMount } from "svelte";
	import axios from "axios";

	let questions;

	let errorMessage;

	onMount(async () => {
		getQuestions();
	});

	function handleAnswerClick(question, answer) {
		if (!answer.revealed) {
			revealQuestion(question.id, answer.id);
		}
	}

	async function revealQuestion(questionId, answerId) {
		let question = questions.find(q => q.id === questionId);
		if (question == null) {
			console.error(`Invalid questionId: ${questionId}`);
			return;
		}
		let answer = question.answers.find(a => a.id === answerId);
		if (answer == null) {
			console.error(`Invalid answerId: ${answerId}`);
			return;
		}
		answer.revealed = true;
		questions = questions;

		try {
			const response = await axios({
				url: `http://localhost:1323/api/questions/${questionId}/answers/${answerId}/reveal`,
				method: "put",
				headers: { "Authorization": `Bearer ${window.localStorage.getItem("code")}`}
			});
		} catch (e) {
			console.log(e);
			errorMessage = "Failed to reveal question.";
			questions = null;
			getQuestions();
		}
	}

	async function getQuestions() {
		try {
			const response = await axios({
				url: "http://localhost:1323/api/questions",
				headers: { "Authorization": `Bearer ${window.localStorage.getItem("code")}`}
			});
			questions = response.data;
		} catch (e) {
			console.log(e);
			errorMessage = "Failed to retrieve questions.";
			questions = null;
		}
	}
</script>

<style>
.error {
	color: red;
}

.not-revealed {
	background-color: orange;
}

.revealed {
	background-color: green;
}

.clickable {
	position: relative;
}

.clickable:hover > .overlay {
    width: 100%;
    height: 100%;
    background-color: #000;
    opacity: 0.5;
	position: absolute;
	top: 0;
	left: 0;
	right: 0;
	bottom: 0;
}
</style>

{#if errorMessage}
	<span class="error">{errorMessage}</span>
{:else if questions}
	{#each questions as question (question.id)}
		<span>{question.description}</span>
		{#each question.answers as answer (answer.id)}
		<div class="{answer.revealed ? 'revealed' : 'clickable not-revealed'}" on:click={() => handleAnswerClick(question, answer)}>
				<div class="overlay"></div>
				({answer.points}) {answer.text}
			</div>
		{/each}
	{/each}
{:else}
	<p>Please wait, we're loading the questions. If this takes more than a few seconds, try reloading the page.</p>
{/if}

