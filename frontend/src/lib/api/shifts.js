// /src/lib/api/shifts.js

const headers = () => ({
  "Content-Type": "application/json",
  Authorization: `Bearer ${localStorage.getItem("token")}`,
});

export async function getAssignedShifts() {
  const res = await fetch("/api/shifts/assigned", {
    headers: headers(),
  });
  return res.json();
}

export async function getAvailableShifts() {
  const res = await fetch("/api/shifts/available", {
    headers: headers(),
  });
  return res.json();
}

export async function requestShift(shiftId) {
  const res = await fetch("/api/shifts/request", {
    method: "POST",
    headers: headers(),
    body: JSON.stringify({ shift_id: shiftId }),
  });
  return res.json();
}

export async function getRequests() {
  const res = await fetch("/api/requests", {
    headers: headers(),
  });
  return res.json();
}
