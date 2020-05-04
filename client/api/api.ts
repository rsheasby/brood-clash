import { DefaultApi, Configuration } from './codegen/index';
import * as axios from 'axios';

export let ApiClient = new DefaultApi(new Configuration({
	apiKey: "",
	basePath: `http://${location.hostname}:8080/api/v1`
}));

axios.default.interceptors.response.use(config => config, error => {
	if (location.pathname != "/controller/login")
		window.history.replaceState({}, "", "/controller/login");
	return error;
});

export function SetApiKey(key: string) {
	ApiClient = new DefaultApi(new Configuration({
		apiKey: key,
		basePath: `http://${location.hostname}:8080/api/v1`
	}));
}
