// /src/lib/api/shifts.js

const headers = () => ({
  "Content-Type": "application/json",
  Authorization: `Bearer ${localStorage.getItem("token")}`,
});

export async function getAssignedShifts() {
  const res = await fetch(
    `http://localhost:8080/api/shifts/assigned/${localStorage.getItem("id")}`,
    {
      headers: headers(),
    }
  );
  return res.json();
}

export async function getAvailableShifts() {
  const res = await fetch("http://localhost:8080/api/shifts/available", {
    headers: headers(),
  });
  return res.json();
}

export async function requestShift(shiftId) {
  const res = await fetch("http://localhost:8080/api/requests", {
    method: "POST",
    headers: headers(),
    body: JSON.stringify({
      shift_id: shiftId,
      worker_id: parseInt(localStorage.getItem("id"), 10),
    }),
  });
  return res.json();
}

export async function getRequests() {
  const res = await fetch("http://localhost:8080/api/requests/pending", {
    headers: headers(),
  });
  return res.json();
}
