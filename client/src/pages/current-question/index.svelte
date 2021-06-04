<script lang="typescript">
	import {onMount} from 'svelte';

    let question: string = '';

	let ws;
	onMount(async () => {
		// The websocket stuff is all untyped unfortunately. Swagger can't really model websocket communication.
		ws = new WebSocket(`ws://${location.hostname}:8080/presenter/websocket`);
		ws.onmessage = event => handleUpdate(JSON.parse(event.data));
	});

    function handleUpdate(update) {
        if (update.Type === 'stateUpdate') {
            question = update.Update.text;
        }
    }
</script>
<div class="p-4">
    <h1 class='text-5xl'>{question}</h1>
</div>