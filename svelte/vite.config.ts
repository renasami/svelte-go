import { defineConfig } from 'vite';
import { svelte } from '@sveltejs/vite-plugin-svelte';
import Pages from 'vite-plugin-pages-svelte';

export default defineConfig({
  plugins: [svelte(), Pages()],
});
