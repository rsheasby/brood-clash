<script>
    import { onMount } from 'svelte';
    import { ApiClient, DefaultApi, Question } from 'brood_clash';

    let questions = [];
    let loading = true;
    let client;

    onMount(async () => {
        ApiClient.instance.authentications.ApiKey.apiKey = window.localStorage.getItem('code');
    	client = new DefaultApi();
        questions = await client.getQuestions();
        loading = false;
    });

    function addQuestion() {
        if (!questions) {
            questions = [];
        }

        const question = {
        	text: '',
        	answers: []
        };
        for (let i = 0; i < 8; ++i) {
            const answer = {
            	text: '',
            	points: undefined
            };
            question.answers.push(answer);
        }
        questions = [ ...questions, question ];
    }

    async function uploadQuestions() {
    	loading = true;
    	try {
            await client.createQuestions(questions);
    	} catch (e) {
    		console.error("Could not upload questions:", e);
    	}
        // TODO: we should prompt the user if they want to go to the controller
        // after questions have successfully been uploaded.
        loading = false;
    }
</script>

{#if loading}
    <p>Loading, please wait...</p>
{/if}
<!--
    TODO: we can maybe look at this to clean up the page and make it a
    little more usable: https://alligator.io/css/collapsible/
-->
{#each questions as question}
    Question:
    <input type="text" bind:value="{question.text}" disabled="{loading}"> <br />
    {#each question.answers as answer}
        Answer:
        <input type="text" bind:value="{answer.text}" disabled="{loading}">
        Points:
        <input type="number" bind:value="{answer.points}" disabled="{loading}"> <br />
    {/each}
{/each}
<button on:click|preventDefault="{addQuestion}" disabled="{loading}">Add question.</button>
<button on:click|preventDefault="{uploadQuestions}" disabled="{loading}">Upload.</button>
