<script>
	import {onMount} from 'svelte';

	// This array will be updated every time a question gets added, deleted, or
	// modified in the backend.
	let question = '';
	let answers = [];
	let wrongAnswer = false;

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

	function displayX() {
		wrongAnswer = true;
		document.getElementById('buzzer').play();
		setTimeout(() => {
			wrongAnswer = false;
		}, 1000);
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
				document.getElementById("success").play();
				updateAnswer(update.Update);
				break;
			case 'incorrectAnswer':
				displayX();
				break;
		}
	}
</script>

<style>
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
		min-width: 0;
		min-height: 0;
		max-height: 70vw;
		max-width: 97vw;
		height: 95vh;
		width: 132vh;
		border-radius: 50%;
		border: 15px solid #e86f32;

		box-sizing: border-box;
	}

	#container {
		height: 85%;
		width: 85%;

		position: relative;

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
		flex-basis: auto;
		flex-shrink: 0;
		width: 31%;
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

	.answer-text {
		flex-basis: auto;
		flex-grow: 1;
		width: 0;
		height: 100%;
		color: #eee;
		font-family: Verdana, Geneva, Tahoma, sans-serif;
		font-weight: bolder;
		padding: 0 10px;
		box-sizing: border-box;
	}

	.answer-text span {
		max-height: 100%;
		max-width: 100%;
		overflow-wrap: break-word;
		overflow: hidden;
		display: -webkit-box;
		-webkit-line-clamp: 3;
		-webkit-box-orient: vertical;
		text-overflow: ellipsis;
		text-align: center;
	}

	@media (max-aspect-ratio: 132/95) {
		.answer-text {
			font-size: calc((95vw / 100 * 2.6) - 2.5pt);
		}
	}

	@media (min-aspect-ratio: 132/95) {
		.answer-text {
			font-size: calc((132vh / 100 * 2.6) - 2.5pt);
		}
	}

	.hidden {
		display: none !important;
	}

	.wrong-answer-display {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		z-index: 999;

		display: flex;
		align-items: center;
		justify-content: center;

		padding: 10%;
	}

	.wrong-answer-x {
		height: 100%;
		fill: #f00;
		font-size: 15500%;
		font-family: Verdana, Geneva, Tahoma, sans-serif;
		font-weight: bolder;
		background-color: #0f1961;
		border: 15px solid #f00;
	}
</style>

<div id="display" class="center-container">
	<div id="ring" class="center-container">
		<div id="container">
			<audio id="success" preload="auto">
				<source src="success.mp3" type="audio/mpeg"/>
			</audio>
			{#each answers as answer, i}
				<div class="answer-revealer" class:reveal={answer.Revealed}>
					<div class="answer-inner">
						<div class="answer answer-unshown full-center">
							<div class="unshown-answer-container">
								<svg
									viewBox="0 0 100 100"
									class="unshown-answer-number">
									<text
										x="50"
										y="50"
										text-anchor="middle"
										dominant-baseline="central">
										{i + 1}
									</text>
								</svg>
							</div>
						</div>
						<div class="answer answer-shown">
							<div class="answer-points-container">
								<svg
									viewBox="0 0 100 100"
									class="answer-points">
									<text
										x="50"
										y="50"
										text-anchor="middle"
										dominant-baseline="central">
										{answer.Points ? answer.Points : ''}
									</text>
								</svg>
							</div>
							<div
								class="answer-text full-center"
								class:answer-text-revealed={answer.Revealed}>
								<span>{answer.Text ? answer.Text : ''}</span>
							</div>
						</div>
					</div>
				</div>
			{/each}
			<audio id="buzzer" preload="auto">
				<source src="buzzer.mp3" type="audio/mpeg"/>
			</audio>
			<div class="wrong-answer-display" class:hidden={!wrongAnswer}>
				<svg viewBox="0 0 2000 2000" class="wrong-answer-x">
					<text
						x="1000"
						y="935"
						text-anchor="middle"
						dominant-baseline="central">
						X
					</text>
				</svg>
			</div>
		</div>
	</div>
</div>
