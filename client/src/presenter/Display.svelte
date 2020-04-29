<script>
	import {onMount} from 'svelte';

	// This array will be updated every time a question gets added, deleted, or
	// modified in the backend.
	let question = '';
	let answers = [];

	let ws;
	onMount(async () => {
		ws = new WebSocket('ws://localhost:8080/presenter/websocket');
		ws.onmessage = event => handleUpdate(JSON.parse(event.data));
	});

	function handleUpdate(update) {
		switch (update.Type) {
			case 'stateUpdate':
				question = update.Update.Text;
				answers = [];
				update.Update.Answers.forEach((a, i) => {
					answers[i] = a;
				});
		}
	}
</script>

<style>
	h1 {
		color: purple;
	}

	.center-container {
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: center;
	}

	#display {
		position: absolute;
		top: 0;
		left: 0;
		width: 100vw;
		height: 100vh;
		background-color: #001f54;

		padding: 0;
	}

	#ring {
		background-image: url('bg.png');
		background-position: center;
		max-height: 75vw;
		max-width: 95vw;
		min-width: 0;
		min-height: 0;
		height: 95vh;
		width: 120vh;
		border-radius: 50%;
		border: 15px solid #e86f32;

		box-sizing: border-box;
	}

	#container {
		height: 85%;
		width: 85%;

		border: 15px solid #e86f32;
		background-color: #292e36;
		box-sizing: border-box;
		padding: 10px;

		display: grid;
		grid-template-rows: 1fr 1fr 1fr 1fr;
		grid-template-columns: 1fr 1fr;
		grid-auto-flow: column;
		gap: 10px;
	}

	.answer-revealer {
		/* I hardly know her */
      perspective: 1000px;
		box-sizing: border-box;
	}

	.answer-inner {
		position: relative;
		width: 100%;
		height: 100%;
		transition: transform 0.8s;
		transform-style: preserve-3d;
		box-sizing: border-box;
	}

	.answer-revealer:hover .answer-inner {
		transform: rotateX(180deg);
	}

	.answer {
		border: 5px solid #001f54;
		background-color: #222ca0;
		position: absolute;
		height: 100%;
		width: 100%;
		backface-visibility: hidden;
		box-sizing: border-box;
	}

	.answer-unshown {
		transform: rotateX(180deg);
	}

	.full-center {
      padding: 0;
      margin: 0;
      width: 100%;
      height: 100%;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.unshown-answer-container {
		background-color: #544be8;
		width: 30%;
		border-radius: 100%;
		border: 5px solid #0f1961;
		display: flex;
		justify-content: center;
		align-items: center;
		font-size: 3em;
		color: #eee;
		box-sizing: border-box;
	}

	.unshown-answer-container:after {
		content: '';
		display: block;
		padding-bottom: 100%;
	}
</style>

<div id="display" class="center-container">
	<div id="ring" class="center-container">
		<div id="container">
			<div class="answer-revealer">
				<div class="answer-inner">
					<div class="answer answer-unshown">
						<div class="full-center">
							<div class="unshown-answer-container">1</div>
						</div>
					</div>
					<div class="answer answer-shown">le answer</div>
				</div>
			</div>
		</div>
	</div>
</div>

<!-- Example of what the questions will look like -->
<!-- {#each questions as question (question.id)}
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
-->
