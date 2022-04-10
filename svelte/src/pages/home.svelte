<script lang="ts">
    import { Route } from "svelte-router-spa"
    export let currentRoute
    const params = { validCheck: true }
    console.log(currentRoute)
</script>

<h1>home</h1>

<section class="section">
    <Route {currentRoute}  />
    <p>Route params are: {currentRoute.name} and {currentRoute.path}</p>
  </section>