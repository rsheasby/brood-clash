<script>
	import { onMount } from "svelte";

	export let name;
	name = "ryan";

	// This array will be updated every time a question gets added, deleted, or
	// modified in the backend.
	let questions = [];

	let ws;
	onMount(async () => {
		ws = new WebSocket("ws://localhost:1323/present");
		ws.onmessage = (event) => questions = JSON.parse(event.data);
	});
</script>

<style>
	h1 {
		color: purple;
	}

	#display {
		position: absolute;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		background-image: url("bg.png");
		display: flex;
		flex-direction: row;
		box-sizing: border-box;
		padding: 10px;
	}

	.column {
		background-color: #f00;
		flex: 1;
		margin: 10px;
	}
</style>

<div id="display">
	<div class="column"></div>
	<div class="column"></div>
</div>

<!-- Example of what the questions will look like -->
{#each questions as question (question.id)}
<div>
	{question.description}
	<ul>
		{#each question.answers as answer (answer.id)}
		{#if answer.revealed}
		<li>({answer.points}) {answer.text}</li>
		{:else}
		<li><em>hidden</em></li>
		{/if}
		{/each}
	</ul>
</div>
{/each}
