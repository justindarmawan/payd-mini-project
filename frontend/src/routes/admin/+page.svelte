<!-- /src/routes/admin/+page.svelte -->
<script>
  import { getRequests } from "$lib/api/shifts";
  import { onMount } from "svelte";

  let requests = [];

  onMount(async () => {
    requests = await getRequests();
  });

  async function approveRequest(id) {
    await fetch(`/api/requests/${id}/approve`, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    });
    requests = requests.filter((r) => r.id !== id);
  }

  async function rejectRequest(id) {
    await fetch(`/api/requests/${id}/reject`, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    });
    requests = requests.filter((r) => r.id !== id);
  }
</script>

<div class="admin-container">
  <h1>Pending Requests</h1>
  {#each requests as req}
    <div class="card">
      <p>
        <strong>{req.username}</strong> requested shift on {req.date} from {req.start_time}
        to {req.end_time}
      </p>
      <button on:click={() => approveRequest(req.request_id)}>Approve</button>
      <button on:click={() => rejectRequest(req.request_id)}>Reject</button>
    </div>
  {/each}
</div>

<style lang="scss">
  .admin-container {
    padding: 2rem;
  }
  .card {
    border: 1px solid #ddd;
    border-radius: 0.5rem;
    padding: 1rem;
    margin-bottom: 1rem;
  }
  button {
    margin-right: 1rem;
  }
</style>
