<script>
	const connectionError = "error";

	let ws;
	let connectionStatus = undefined;
	let messages = [];

	function connect(e) {
		ws = new WebSocket("ws://localhost:1323/present");

		connectionStatus = WebSocket.CONNECTING;
		ws.onerror = () => connectionStatus = connectionError;
		ws.onopen = () => connectionStatus = WebSocket.OPEN;
		ws.onclose = () => connectionStatus = WebSocket.CLOSED;
		ws.onmessage = (event) => messages = [...messages, event.data];
	}

	function messageSubmitted(e) {
		ws.send(e.target.message.value);
	}
</script>

<style>
.not-connected {
	color: gray;
}

.connecting {
	color: lightgreen;
}

.connected {
	color: green;
}

.closed {
	color: red;
}

.error {
	color: darkred;
}
</style>

Status:
{#if connectionStatus === undefined}
<span class="not-connected">Not connected</span>
{:else if connectionStatus === WebSocket.CONNECTING}
<span class="connecting">Connecting...</span>
{:else if connectionStatus === WebSocket.OPEN}
<span class="connected">Connected</span>
{:else if connectionStatus === connectionError}
<span class="error">Failed to connect</span>
{:else}
<span class="closed">Closed</span>
{/if}

{#if connectionStatus !== WebSocket.OPEN && connectionStatus !== WebSocket.CONNECTING}
<button on:click={connect}>Connect</button>
{/if}

{#if connectionStatus === WebSocket.OPEN}
<form on:submit|preventDefault={messageSubmitted}>
	<label for="message">Message:</label>
	<input type="text" id="message" name="message">
</form>

<ul>
	{#each messages as message}
	<li>{message}</li>
	{/each}
</ul>
{/if}
