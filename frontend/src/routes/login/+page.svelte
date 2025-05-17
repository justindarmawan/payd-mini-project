<script>
  import { login } from "$lib/api/login";
  import { goto } from "$app/navigation";

  let username = "";
  let password = "";
  let error = "";

  async function handleLogin() {
    try {
      const res = await login(username, password);
      goto(res.role === "admin" ? "/admin" : "/worker");
    } catch (err) {
      error = "Invalid username or password";
    }
  }
</script>

<div class="login-container">
  <div class="login-card">
    <h2>Login</h2>
    {#if error}
      <div class="error">{error}</div>
    {/if}
    <input type="text" placeholder="Username" bind:value={username} />
    <input type="password" placeholder="Password" bind:value={password} />
    <button on:click={handleLogin}>Login</button>
  </div>
</div>

<style lang="scss">
  .login-container {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100vh;
    background: #f9fafb;
  }

  .login-card {
    background: white;
    padding: 2rem;
    border-radius: 1rem;
    box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 350px;
  }

  h2 {
    text-align: center;
    margin-bottom: 1.5rem;
  }

  input {
    width: 100%;
    padding: 0.75rem;
    margin-bottom: 1rem;
    border: 1px solid #ccc;
    border-radius: 0.5rem;
  }

  button {
    width: 100%;
    padding: 0.75rem;
    background-color: #007bff;
    color: white;
    font-weight: bold;
    border: none;
    border-radius: 0.5rem;
    cursor: pointer;
  }

  .error {
    color: red;
    margin-bottom: 1rem;
    text-align: center;
  }
</style>
