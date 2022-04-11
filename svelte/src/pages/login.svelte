<script lang="ts">

    const email = "test@test.co.jp"
    const password = "test"
    const method = "POST"
    const json = {
        "password": password,
        "email": email,
    }
    const body = JSON.stringify(json)
    const headers = {
        "Content-Type": "application/json"
    
    }
    const login = async () => {
        const response = await fetch("http://localhost:8080/api/login",{headers,method, body})
        console.log(response.body)
    
    }

    const logout = async () => {
        const response = await fetch("http://localhost:8080/api/logout")
        console.log(response.body)
    }
    </script>
    <main>
        <h1>Login</h1>
        <button on:click={login} >Login</button>
        <button on:click={logout}>Logout</button>
    </main>