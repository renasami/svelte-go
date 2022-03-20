<script lang="ts">
	import { Router, Link, Route } from "svelte-routing";
	import Home from "./components/pages/Home.svelte"
	import  Login from "./components/pages/Login.svelte";
	import  Register from "./components/pages/Register.svelte";

</script>

<Router>
	<nav>
	  <Link to="/">home</Link>
	  <Link to="/login">about</Link>
	  <Link to="/register">about</Link>
	</nav>
	<main>
	  <Route path="" component={Home} />
	  <Route path="/login" component={Login} />
	  <Route path="/register" component={Register} />
	</main>
</Router>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	/* h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	} */

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>