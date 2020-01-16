<script>
    import page from 'page';
    import Display from './presenter/Display.svelte';
    import LoginForm from './controller/LoginForm.svelte';
    import QuestionsForm from './controller/QuestionsForm.svelte';
    import Controller from './controller/Controller.svelte';

    let component = LoginForm;

    page('/presenter', () => void (component = Display));
    page('/controller/login', () => void (component = LoginForm));
    page('/controller/add-questions', () => void (component = QuestionsForm));
    page('/controller/controller', () => void (component = Controller));
    page();
</script>

<svelte:component this={component} />
