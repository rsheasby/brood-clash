<script lang="typescript">
	import { onMount } from 'svelte';

	import { ApiClient } from 'api';

	import { goto } from '@sveltech/routify';

	let loading = true;

	let questions: any[];

	let error: string;

	onMount(async () => {
		try {
			let response = await ApiClient.getAllQuestions();
			questions = response.data;
		} catch (e) {
		}

		loading = false;
	});

	async function selectQuestion(questionId) {
		loading = true;
		await ApiClient.selectQuestion(questionId);
		$goto('/controller');
	}
</script>

{#if error}
	<p>{error}</p>
{/if}

{#if loading}
	<p>Loading, please wait...</p>
{:else}
	<ul>
	{#each questions as question}
		<li><div class="text-blue-500 hover:text-blue-800" on:click="{selectQuestion(question.ID)}">{question.Text}</div></li>
	{/each}
	</ul>
{/if}
