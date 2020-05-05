import { DefaultApi, Configuration } from './codegen/index';
import * as axios from 'axios';

const apiPath: string = `http://${location.hostname}:8080/api/v1`
const loginPath: string = '/controller/login';
var loggedIn: boolean = false;

export let ApiClient = new DefaultApi(new Configuration({
	apiKey: "",
	basePath: apiPath
}));

axios.default.interceptors.response.use(config => config, error => {
	if (location.pathname != loginPath)
		window.history.replaceState({}, '', loginPath);
	return error;
});

export function SetApiKey(key: string) {
	ApiClient = new DefaultApi(new Configuration({
		apiKey: key,
		basePath: apiPath
	}));
	loggedIn = true;
}

export function IsLoggedIn() : boolean {
	return loggedIn;
}
