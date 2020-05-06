<script lang="typescript">
	import {onMount} from 'svelte';
	import {url} from '@sveltech/routify';

	import { ApiClient } from 'api';

	let question: any;

	let loading: boolean = true;

	onMount(async () => {
		try {
			let response = await ApiClient.getCurrentQuestion();
			question = response.data;
		} catch (e) {
			const status = e && e.response && e.response.status;
			if (status !== 404) {
				console.error(e);
			}
		}

		loading = false;
	});
</script>

{#if loading}
	<p>Loading, please wait...</p>
{:else if !question}
	<!-- TODO: alternatively, we could automatically redirect. -->
	<p>No question selected. Please go to the <a class="text-blue-500 hover:text-blue-800" href="{$url('/controller/questions')}">questions page</a> to select one.</p>
{:else}
	<p>{question.Text}</p>
	{#each question.Answers as answer}
		<p>{answer.Points} - {answer.Text}</p>
	{/each}
{/if}

<a class="text-blue-500 hover:text-blue-800" href="{$url('/controller/questions')}">
	<button>Back to questions</button>
</a>

