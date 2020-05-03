<script>
    import page from 'page';
    import Display from './presenter/Display.svelte';
    import LoginForm from './controller/LoginForm.svelte';
    import QuestionsForm from './controller/QuestionsForm.svelte';
    import Controller from './controller/Controller.svelte';
    import NotFound from './NotFound.svelte';
	 import PleaseWait from './PleaseWait.svelte';

    let component = PleaseWait;

    page('/presenter', () => void (component = Display));
    page('/controller/login', () => void (component = LoginForm));
    page('/controller/add-questions', () => void (component = QuestionsForm));
    page('/controller/controller', () => void (component = Controller));
    page('/controller/*', '/controller/login');
    page('*', () => void (component = NotFound));
    page();
</script>

<svelte:component this={component} />
