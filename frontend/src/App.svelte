<script>
  import { token } from './lib/auth.js';
  import Login from './components/Login.svelte';
  import Upload from './components/Upload.svelte';
  import VideoList from './components/VideoList.svelte';

  let page = 'list';

  $: authed = $token.length > 0;

  function logout() { token.set(""); }
</script>

<nav class="nav">
  {#if authed}
    <button class="btn" on:click={() => page = 'list'}>VIDEOS</button>
    <button class="btn" on:click={() => page = 'upload'}>UPLOAD</button>
    <button class="btn" on:click={logout}>LOGOUT</button>
  {/if}
</nav>

{#if !authed}
  <Login />
{:else if page === 'upload'}
  <Upload />
{:else}
  <VideoList />
{/if}

<style>
  .nav {
    background-color: rgb(22, 22, 22);
    display: flex;
    justify-content: center;
    padding-top: 10px;
    padding-bottom: 10px;
    flex-wrap: wrap;
    width: 100%;
    gap: 15px;
  }
  .btn {
    background-color: white;
    border: 0;
    font-weight: bold;
    border-radius: 6px;
    padding: 10px;
  }
  .btn:hover {
    filter: brightness(.9);
    cursor: pointer;
  }
</style>