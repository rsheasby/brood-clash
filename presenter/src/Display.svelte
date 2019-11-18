<script>
	export let name;
	name = "ryan";
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
