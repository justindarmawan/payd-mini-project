<!-- /src/routes/employee/+page.svelte -->
<script>
  import {
    getAssignedShifts,
    getAvailableShifts,
    requestShift,
  } from "$lib/api/shifts";
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";

  let assigned = [];
  let available = [];

  onMount(async () => {
    assigned = await getAssignedShifts();
    available = await getAvailableShifts();
  });

  async function request(id) {
    if (!window.confirm("Are you sure you want to request this shift?")) return;

    try {
      const data = await requestShift(id);

      if (data.error) {
        alert(data.error);
        return;
      }

      available = available.filter((s) => s.id !== id);
    } catch (err) {
      console.error(err);
      alert("Something went wrong while requesting the shift.");
    }
  }

  function logout() {
    localStorage.removeItem("token");
    goto("/login"); // adjust this path to your actual login route
  }
</script>

<div class="container">
  <div class="header">
    <h1>Pending Requests</h1>
    <button class="logout-button" on:click={logout}>Logout</button>
  </div>
  {#each assigned as s}
    <div class="shift">
      <p>{s.date} - {s.start_time} to {s.end_time} - {s.role}</p>
    </div>
  {/each}

  <h2>Available Shifts</h2>
  {#each available as s}
    <div class="shift">
      <p>{s.date} - {s.start_time} to {s.end_time} - {s.role}</p>
      <button on:click={() => request(s.id)}>Request</button>
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
