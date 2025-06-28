<script>
  let videos = [];
  let error = '';

  async function fetchVideos() {
    try {
      const res = await fetch('http://localhost:8080/videos');
      if (res.ok) {
        videos = await res.json();
      } else {
        error = "Failed to load videos";
      }
    } catch (e) {
      error = e;
    }
  }
  fetchVideos();
</script>

{#if error}
  <div>{error}</div>
{/if}

<ul class="list">
  {#each videos as video}
    <li class="item">
      <video src={`http://localhost:8080${video.url}`} controls width="320"></video>
      <div class="video-content">
        <div class="description-container">
          <strong><span class="darker-text">Name:</span> {video.name}</strong>
        </div>
        <div class="btn-container">
          <a class="download-btn" href={`http://localhost:8080${video.url}`} download>Download</a>
        </div>
      </div>
    </li>
  {/each}
</ul>

<style>
  .list {
    justify-content: center;
    display: flex;
    flex-wrap: wrap;
    gap: 15px;
  }
  .item {
    background-color: rgb(32, 32, 32);
    color: #fff;
    padding: 15px;
    border-radius: 6px;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .video-content {
    display: flex;
    justify-content: space-between;
    flex-direction: column;
    height: 80px;
    width: 100%;
  }
  .darker-text {
    filter: brightness(.8);
  }
  .description-container {
    padding: 5px;
    padding-left: 15px;
  }
  .btn-container {
    width: 100%;
    display: flex;
    justify-content: flex-start;
    align-items: flex-end;
  }
  .download-btn {
    background-color: white;
    margin-left: 15px;
    border: 0;
    font-weight: bold;
    border-radius: 6px;
    padding: 10px;
    text-decoration: none;
    color: black;
  }
  .download-btn:hover {
    filter: brightness(.9);
    cursor: pointer;
  }
</style>