<script>
	// Library imports
	import page from "page";
	
	// Component imports
	import LoginForm from "./LoginForm.svelte";
	import QuestionsForm from "./QuestionsForm.svelte";


	let component = LoginForm;

	page("/", () => page.redirect("/controller"));
	page("/login", () => component = LoginForm);
	page("/create-questions", async () => {
		if (!(await isLoggedIn())) {
			page.redirect("/login");
			return;
		}

		component = QuestionsForm;
	});
	page("/controller", async () => {
		if (!(await isLoggedIn())) {
			page.redirect("/login");
			return;
		}

		if (!(await questionsExist())) {
			page.redirect("/create-questions");
			return;
		}

		component = Controller;
	});

	async function isLoggedIn() {
		const code = window.localStorage.getItem("code");
		return !!(code && (await codeValid(code)));
	}

	async function isCodeValid(code) {
		const response = await axios({
			url: "http://localhost:1323/api/code/valid",
			headers: { "Authorization": `Bearer ${code}` }
		});
		return response.status === 200;
	}

	// NOTE: Assumes a code exists and the code is valid.
	async function questionsExist() {
		const response = await axios({
			url: "http://localhost:1323/api/questions",
			headers: { "Authorization": `Bearer ${code}` }
		});
		if (response.status / 100 !== 2) {
			console.error(response);
			return false;
		}
		console.log(response);
		console.log(response.data);
		const questions = response.data;
		return questions.length > 0;
	}
</script>

<svelte:component this={component} />
