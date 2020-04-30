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

	function updateAnswer(answer) {
		answers.forEach((a, i) => {
			if (a.ID === answer.ID) {
				answers[i] = answer;
				return;
			}
		});
	}

	function handleUpdate(update) {
		switch (update.Type) {
			case 'stateUpdate':
				question = update.Update.Text;
				answers = [];
				update.Update.Answers.forEach((a, i) => {
					answers[i] = a;
				});
				break;
			case 'revealAnswer':
				updateAnswer(update.Update);
				break;
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
		padding: 2%;

		display: grid;
		grid-template-rows: 1fr 1fr 1fr 1fr;
		grid-template-columns: 1fr 1fr;
		grid-auto-flow: column;
		gap: 2%;
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

	.reveal .answer-inner {
		transform: rotateX(180deg);
	}

	.answer {
		border: 10px solid #001f54;
		background-color: #222ca0;
		position: absolute;
		height: 100%;
		width: 100%;
		box-sizing: border-box;
	}

	.answer-shown {
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
		position: relative;
		width: 30%;
		border-radius: 100%;
		border: 10px solid #0f1961;
		display: flex;
		justify-content: center;
		align-items: center;
		box-sizing: border-box;
	}

	.unshown-answer-container:after {
		content: '';
		display: block;
		padding-bottom: 100%;
	}

	.unshown-answer-number {
		width: 100%;
	}

	.unshown-answer-number text {
		font-size: 400%;
		font-family: Verdana, Geneva, Tahoma, sans-serif;
		font-weight: bolder;
		fill: #eee;
	}

	.answer-shown {
		display: flex;
		flex-direction: row;
	}

	.answer-points-container {
		width: 33%;
		position: relative;
      background-color: #544be8;
      border-right: 10px solid #001f54;
	}

	.answer-points text {
		font-size: 300%;
		font-family: Verdana, Geneva, Tahoma, sans-serif;
		font-weight: bolder;
		fill: #eee;
	}
   
   .answer-points {
      width: 100%;
      height: 100%;
   }
</style>

<div id="display" class="center-container">
	<div id="ring" class="center-container">
		<div id="container">
			{#each answers as answer, i}
				<div class="answer-revealer" class:reveal={answer.Revealed}>
					<div class="answer-inner">
						<div class="answer answer-unshown full-center">
							<div class="unshown-answer-container">
								<svg
									viewBox="0 0 100 100"
									class="unshown-answer-number">
									<text x="27.5" y="75">{i + 1}</text>
								</svg>
							</div>
						</div>
						<div class="answer answer-shown">
							<div class="answer-points-container">
								<svg
									viewBox="0 0 100 100"
									class="answer-points">
									<text x="50" y="70" text-anchor="middle">{answer.Points}</text>
								</svg>
							</div>
							<div class="answer-text">{answer.Text}</div>
						</div>
					</div>
				</div>
			{/each}
		</div>
	</div>
</div>
