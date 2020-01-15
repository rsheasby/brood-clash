<script>
    import page from 'page';
    import Display from './presenter/Display.svelte';
    import LoginForm from './controller/LoginForm.svelte'

    var component = LoginForm;

    page("/presenter", () => component = Display);
    page();
</script>

<svelte:component this={component} />
