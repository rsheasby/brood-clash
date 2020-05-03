import { DefaultApi, Configuration } from 'api';

export let ApiClient = new DefaultApi(new Configuration({
   apiKey: "",
   basePath: `http://${location.hostname}:8080/api/v1`
}));

export function SetApiKey(key: string) {
   ApiClient = new DefaultApi(new Configuration({
      apiKey: key,
      basePath: `http://${location.hostname}:8080/api/v1`
   }));
}
