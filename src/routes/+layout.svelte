<script>
  import { onMount } from 'svelte';
  import posthog from 'posthog-js';
  import "../app.css";


  onMount(() => {
    if (typeof window !== 'undefined') {
      console.log("in poshhog func");
      posthog.init('phc_RHC4yBjDj0XKrwYowFBeGLXljKCz4yK7GzWE7fzMARt', {
        api_host: 'https://us.i.posthog.com',
      });
    }
  });
</script>

<slot />