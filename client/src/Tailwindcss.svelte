<style global lang="postcss">
	@tailwind base;

	.loading {
		position: relative;
		display: flex;
	}

	.loading::before {
		content: '';
		position: absolute;
		top: 0;
		bottom: 0;
		left: 0;
		right: 0;
		background-color: inherit;
		border-radius: inherit;
	}

	.loading::after {
		content: '';
		z-index: 10;
		position: absolute;
		top: calc(50% - (1.5em / 2));
		left: calc(50% - (1.5em / 2));
		width: 1.5em;
		height: 1.5em;
		border-radius: 50%;
		border: 0.25em solid #fff;
		border-color: #f7fafc transparent #f7fafc transparent;
		animation: loading 1.5s easeInOutBack infinite;
	}

	@keyframes loading {
		0% {
			transform: rotate(0deg);
		}
		50% {
			transform: rotate(270deg);
		}
		100% {
			transform: rotate(540deg);
		}
	}

	.flex-center {
		@apply flex items-center justify-center;
	}

	.hw-full {
		@apply min-w-full min-h-full;
	}

	.input {
		@apply p-3 border-gray-600 font-semibold rounded-md border-2;
	}

	.button-primary {
		@apply input;
		@apply text-gray-200 bg-blue-600
	}

	.button-primary:hover {
		@apply bg-blue-500
	}

	.button-primary:active {
		@apply bg-blue-700
	}

	.button-warning {
		@apply input;
		@apply text-gray-200 bg-red-600;
	}

	.button-warning:hover {
		@apply bg-red-500;
	}

	.button-warning:active {
		@apply bg-red-700;
	}

	.button-neutral {
		@apply input;
		@apply text-gray-200 bg-gray-600;
	}

	.button-neutral:hover {
		@apply bg-gray-500;
	}

	.button-neutral:active {
		@apply bg-gray-700;
	}

	.card {
		@apply bg-gray-300 p-5 rounded-sm max-w-full;
	}

	@tailwind components;
	@tailwind utilities;

</style>
