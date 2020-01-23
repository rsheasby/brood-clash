<script>
    import { ApiClient, DefaultApi } from 'brood_clash';

    let loading = true;
    let invalidCode = false;

    async function login(code) {
    	loading = true;
    	invalidCode = false;

    	code = String(code);
    	if (!/^\d{4}$/.test(code)) {
    		loading = false;
    		invalidCode = true;
    		return;
    	}

        ApiClient.instance.authentications.ApiKey.apiKey = code;
    	window.localStorage.setItem('code', code);
    	try {
            const client = new DefaultApi();
            await client.authTest();
            window.location = '/controller/add-questions';
    	} catch (e) {
            invalidCode = true;
            loading = false;
    	}
    }
</script>

<form on:submit|preventDefault="{(event) => login(event.target.code.value)}">
    <input type="text" name="code">
    <label for="name">Code</label>
</form>
