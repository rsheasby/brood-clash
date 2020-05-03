<script lang="typescript">
	import {onMount} from 'svelte';
	import {ApiClient, SetApiKey} from '../api';

	let questions;
	let currentQuestion;
	let loading = true;
	let client;
	let done = false;

	onMount(async () => {
		SetApiKey('1234');
		ApiClient.test().then(result => {
			console.log(result);
		});
	});

	//  function nextQuestion() {
	//  	const nextId = currentQuestion ? currentQuestion.id + 1 : 1;
	//      currentQuestion = questions.find(q => q.id === nextId);
	//      if (currentQuestion === undefined) {
	//      	done = true;
	//      }
	//  }

	//  async function revealAnswer(answerId) {
	//  	const answerIndex = currentQuestion.answers.findIndex(a => a.id === answerId);
	//  	if (answerIndex !== -1) {
	//          try {
	//              currentQuestion.answers[answerIndex].revealed = true;
	//              await client.revealAnswer(currentQuestion.id, answerId);
	//          } catch (e) {
	//          	currentQuestion.answers[answerIndex].revealed = false;
	//          }
	//  	}
	//  }
</script>

{#if loading}
	<p>Loading, please wait...</p>
{/if}
<!-- {#if !done}
    {#if currentQuestion}
        Question: {currentQuestion.text} <br />
        {#each currentQuestion.answers as answer (answer.id)}
            Answer: {answer.text} -->
<!-- TODO: This button should be disabled when the answer is already revealed but I can't get it to work. -->
<!-- {#if !answer.revealed}
                <button on:click|preventDefault="{() => revealAnswer(answer.id)}" disabled="{loading}">Reveal</button>
            {:else}
                <span style="color: green;">Revealed!</span>
            {/if}
            <br />
        {/each}
    {/if}
    <button on:click|preventDefault="{nextQuestion}" disabled="{loading}">Next</button>
{:else}
    No more questions.
{/if} -->
