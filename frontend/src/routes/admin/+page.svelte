<script>
  import { getRequests } from "$lib/api/shifts";
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";

  let requests = [];
  let isEditing = null;

  let form = {
    date: "",
    start_time: "",
    end_time: "",
    role: "",
    location: "",
  };

  onMount(async () => {
    requests = await getRequests();
  });

  async function approveRequest(id) {
    if (!window.confirm("Approve this request?")) return;

    try {
      const res = await fetch(
        `http://localhost:8080/api/requests/${id}/approve`,
        {
          method: "POST",
          headers: {
            Authorization: `Bearer ${localStorage.getItem("token")}`,
          },
        }
      );

      if (!res.ok) {
        const data = await res.json();
        alert(data.error || "Failed to approve request.");
        return;
      }
      requests = requests.filter((r) => r.request_id !== id);
    } catch (err) {
      console.error(err);
      alert("Something went wrong while approving the request.");
    }
  }
  async function rejectRequest(id) {
    if (!window.confirm("Reject this request?")) return;
    await fetch(`http://localhost:8080/api/requests/${id}/reject`, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    });
    requests = requests.filter((r) => r.request_id !== id);
  }

  async function createRequest() {
    try {
      const res = await fetch("http://localhost:8080/api/shifts", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        body: JSON.stringify(form),
      });

      if (!res.ok) {
        const data = await res.json();
        alert(data.error || "Failed to create request.");
        return;
      }

      const newReq = await res.json();
      requests = [...requests, newReq];
      form = { date: "", start_time: "", end_time: "", role: "", location: "" };
    } catch (err) {
      console.error(err);
      alert("Something went wrong while creating the request.");
    }
  }

  function editRequest(req) {
    isEditing = req.request_id;
    form = { ...req };
  }

  async function updateRequest() {
    try {
      const res = await fetch(`http://localhost:8080/api/shifts/${isEditing}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        body: JSON.stringify(form),
      });

      if (!res.ok) {
        const data = await res.json();
        alert(data.error || "Failed to update request.");
        return;
      }

      const updated = await res.json();
      requests = requests.map((r) =>
        r.request_id === isEditing ? updated : r
      );
      isEditing = null;
      form = { date: "", start_time: "", end_time: "", role: "", location: "" };
    } catch (err) {
      console.error(err);
      alert("Something went wrong while updating the request.");
    }
  }

  async function deleteRequest(id) {
    if (!window.confirm("Delete this request?")) return;

    try {
      const res = await fetch(`http://localhost:8080/api/shifts/${id}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
      });

      if (!res.ok) {
        const data = await res.json();
        alert(data.error || "Failed to delete request.");
        return;
      }

      requests = requests.filter((r) => r.request_id !== id);
    } catch (err) {
      console.error(err);
      alert("Something went wrong while deleting the request.");
    }
  }

  function logout() {
    localStorage.removeItem("token");
    goto("/login");
  }
</script>

<div class="admin-container">
  <div class="header">
    <h1>Manage Requests</h1>
    <button class="logout-button" on:click={logout}>Logout</button>
  </div>

  <form
    on:submit|preventDefault={isEditing ? updateRequest : createRequest}
    class="request-form"
  >
    <input type="date" bind:value={form.date} required />
    <input type="time" bind:value={form.start_time} required />
    <input type="time" bind:value={form.end_time} required />
    <input type="text" placeholder="Role" bind:value={form.role} required />
    <input
      type="text"
      placeholder="Location"
      bind:value={form.location}
      required
    />
    <button type="submit">{isEditing ? "Update" : "Create"}</button>
    {#if isEditing}
      <button
        type="button"
        on:click={() => {
          isEditing = null;
          form = {
            date: "",
            start_time: "",
            end_time: "",
            role: "",
            location: "",
          };
        }}
      >
        Cancel
      </button>
    {/if}
  </form>

  {#each requests as req}
    <div class="card">
      <p>
        <strong>{req.username}</strong> requested shift on {req.date} from {req.start_time}
        to {req.end_time}
        ({req.role} @ {req.location})
      </p>
      <button on:click={() => approveRequest(req.request_id)}>Approve</button>
      <button on:click={() => rejectRequest(req.request_id)}>Reject</button>
      <button on:click={() => editRequest(req)}>Edit</button>
      <button on:click={() => deleteRequest(req.request_id)}>Delete</button>
    </div>
  {/each}
</div>

<style lang="scss">
  .admin-container {
    padding: 2rem;
  }

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
  }

  .logout-button {
    background-color: #e74c3c;
    color: white;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 0.5rem;
    cursor: pointer;
  }

  .request-form {
    display: flex;
    gap: 0.5rem;
    flex-wrap: wrap;
    margin-bottom: 2rem;
  }

  .request-form input {
    padding: 0.5rem;
    border: 1px solid #ccc;
    border-radius: 0.25rem;
  }

  .request-form button {
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 0.25rem;
    cursor: pointer;
    background-color: #3498db;
    color: white;
  }

  .card {
    border: 1px solid #ddd;
    border-radius: 0.5rem;
    padding: 1rem;
    margin-bottom: 1rem;
  }

  button {
    margin-right: 0.5rem;
  }
</style>
