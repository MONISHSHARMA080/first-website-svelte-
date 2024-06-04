import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()]
	,server: {
		host: '0.0.0.0', // listen on all IP addresses
		port: 5173, // default Vite dev server port
		hmr: {
		overlay: false,
		},
		},
});
