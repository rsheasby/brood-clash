const sveltePreprocess = require('svelte-preprocess');

module.exports = {
	preprocess: sveltePreprocess({
		typescript: true,
		postcss: true
	}),
};
