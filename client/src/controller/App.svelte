<script>
	// TODO: Since it automatically redirects you to the application if you're
	// logged in, there should be a log out button.
	// TODO: Authorization header should be injected with axios interceptor.

	// Library imports
	import { onMount } from "svelte";
	import page from "page";
	import axios from "axios";

	// Component imports
	import LoginForm from "./LoginForm.svelte";
	import QuestionsForm from "./QuestionsForm.svelte";
	import Controller from "./Controller.svelte";
	import PleaseWait from "./PleaseWait.svelte";


	let component = PleaseWait;
	let isLoading = true;

	onMount(async () => {
		await ((async () => {
			const code = getCode();
			if (!code || !(await isCodeValid(code))) {
				page.redirect("/login");
				return;
			}

			if (!(await questionsExist())) {
				page.redirect("/create-questions");
				return;
			}

			page.redirect("/controller");
		})());
		isLoading = false;
	});

	page("/", () => {
		page.redirect("/controller");
	});
	page("/login", () => {
		component = LoginForm;
	});
	page("/create-questions", () => {
		component = QuestionsForm;
	});
	page("/controller", () => {
		component = Controller;
	});
	page();

	function getCode() {
		return window.localStorage.getItem("code");
	}

	async function isCodeValid(code) {
		try {
			const response = await axios({
				url: "http://localhost:1323/api/code/valid",
				headers: { "Authorization": `Bearer ${code}` }
			});
			return true;
		} catch (e) {
			console.error(e);
			return false;
		}
	}

	// NOTE: Assumes a code exists and the code is valid.
	async function questionsExist() {
		try {
			const response = await axios({
				url: "http://localhost:1323/api/questions",
				headers: { "Authorization": `Bearer ${getCode()}` }
			});
			return response.data.length > 0;
		} catch (e) {
			console.error(e);
			return false;
		}
	}
</script>

{#if !isLoading}
<svelte:component this={component} />
{:else}
<svelte:component this={PleaseWait} />
{/if}
