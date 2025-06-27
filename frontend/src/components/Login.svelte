<script>
  import { token } from '../lib/auth.js';
  let username = '';
  let password = '';
  let error = '';

  async function login() {
    error = '';
    const res = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password })
    });
    if (res.ok) {
      const data = await res.json();
      token.set(data.token);  // Assume backend sends { token: "..." }
    } else {
      error = 'Login failed.';
    }
  }
</script>

<form on:submit|preventDefault={login}>
  <input placeholder="Username" bind:value={username} required>
  <input type="password" placeholder="Password" bind:value={password} required>
  <button>Login</button>
  {#if error}<div class="error">{error}</div>{/if}
</form>