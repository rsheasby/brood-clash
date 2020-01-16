<script>
    export let questions;
</script>

{#if questions}
    {#each questions as question (question.id)}
        {question}
    {/each}
{/if}
