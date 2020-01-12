<script>
    import {onMount} from "svelte";

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
        background-color: #001F54;

        padding: 0;
    }

    #ring {
        background-image: url("bg.png");
        background-position: center;
        max-height: 75vw;
        max-width: 95vw;
        min-width: 0;
        min-height: 0;
        height: 95vh;
        width: 120vh;
        border-radius: 50%;
        border: 15px solid #E86F32;

        box-sizing: border-box;
    }

    #container {
        height: 60%;
        width: 79%;

        border-radius: 15px;
        border: 15px solid #E86F32;
        background-color: #001F54;
        box-sizing: border-box;
        padding: 5px;

        display: flex;
        flex-direction: row;
    }

    .column {
        background-color: #f00;
        flex: 1;
        margin: 10px;
    }
</style>

<div id="display" class="center-container">
    <div id="ring" class="center-container">
        <div id="container">
            <div class="column"></div>
            <div class="column"></div>
        </div>
    </div>
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
