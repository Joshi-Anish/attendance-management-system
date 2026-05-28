import { useEffect, useState } from "react";

export default function TeacherDashboard() {
  const [students, setStudents] = useState([]);

  const load = async () => {
    const res = await fetch("http://localhost:8080/students");
    const data = await res.json();
    setStudents(Array.isArray(data) ? data : []);
  };

  useEffect(() => {
    load();
  }, []);

  const del = async (id) => {
    await fetch(`http://localhost:8080/students/delete?id=${id}`, {
      method: "DELETE",
    });

    load();
  };

  const checkIn = async (id) => {
    await fetch("http://localhost:8080/checkin", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ student_id: id }),
    });
  };

  const checkOut = async (id) => {
    await fetch("http://localhost:8080/checkout", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ student_id: id }),
    });
  };

  return (
    <div className="container">
      <h2>Teacher Dashboard</h2>

      {students.map((s) => (
        <div className="card" key={s.id}>
          <h3>{s.name}</h3>

          <p>{s.address}</p>
          <p>{s.grade}</p>

          <button onClick={() => checkIn(s.id)}>Check In</button>
          <button onClick={() => checkOut(s.id)}>Check Out</button>
          <button className="delete" onClick={() => del(s.id)}>
            Delete
          </button>
        </div>
      ))}
    </div>
  );
}