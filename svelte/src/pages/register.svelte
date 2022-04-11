<script lang="ts">
    const register = async () => {
        const json = {
            name: "test"
            email: "test@example.com"
            password: "example"
        }
        const response = await fetch("http://localhost:8888/api/auth/register")
    }

</script>


<h1>register</h1>
