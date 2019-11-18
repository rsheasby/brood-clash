<script>
	import Home from './Home.svelte';
	import Display from './Display.svelte';
	import page from 'page';

	var component = Home;
	page('/display', () => component = Display);
	page('*', () => component = Home)
	page();
</script>

<svelte:component this={component} />
