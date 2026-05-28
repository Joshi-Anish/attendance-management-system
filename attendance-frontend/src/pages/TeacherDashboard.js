import { useEffect, useState } from "react";

export default function TeacherDashboard() {
  const [students, setStudents] = useState([]);
  const [history, setHistory] = useState([]);

  const [form, setForm] = useState({
    id: "",
    name: "",
    address: "",
    grade: "",
  });

  const loadStudents = async () => {
    const res = await fetch("http://localhost:8080/students");
    const data = await res.json();
    setStudents(Array.isArray(data) ? data : []);
  };

  const loadHistory = async () => {
    const res = await fetch("http://localhost:8080/dashboard");
    const data = await res.json();
    setHistory(Array.isArray(data.records) ? data.records : []);
  };

  useEffect(() => {
    loadStudents();
    loadHistory();
  }, []);

  const addStudent = async () => {
    await fetch("http://localhost:8080/students", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(form),
    });

    setForm({ id: "", name: "", address: "", grade: "" });
    loadStudents();
  };

  const deleteStudent = async (id) => {
    await fetch(`http://localhost:8080/students/delete?id=${id}`, {
      method: "DELETE",
    });

    loadStudents();
  };

  const checkIn = async (id) => {
    await fetch("http://localhost:8080/checkin", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ student_id: id }),
    });

    loadHistory();
  };

  const checkOut = async (id) => {
    await fetch("http://localhost:8080/checkout", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ student_id: id }),
    });

    loadHistory();
  };

  const isActive = (id) =>
    history.some(
      (h) =>
        String(h.student_id) === String(id) &&
        (!h.check_out || h.check_out === "")
    );

  const clearData = async () => {
    if (!window.confirm("Delete all attendance?")) return;

    await fetch("http://localhost:8080/attendance/clear", {
      method: "DELETE",
    });

    loadHistory();
  };

  const download = () => {
    window.open("http://localhost:8080/attendance/download");
  };

  return (
    <div className="container">

      <h2>Teacher Dashboard</h2>

      {/* ADD STUDENT */}
      <div className="card">
        <h3>Add Student</h3>

        <div className="form">
          <input placeholder="ID"
            value={form.id}
            onChange={(e) => setForm({ ...form, id: e.target.value })}
          />
          <input placeholder="Name"
            value={form.name}
            onChange={(e) => setForm({ ...form, name: e.target.value })}
          />
          <input placeholder="Address"
            value={form.address}
            onChange={(e) => setForm({ ...form, address: e.target.value })}
          />
          <input placeholder="Grade"
            value={form.grade}
            onChange={(e) => setForm({ ...form, grade: e.target.value })}
          />

          <button onClick={addStudent}>Add Student</button>
        </div>
      </div>

      {/* STUDENTS */}
      <h2>Students</h2>

      {students.map((s) => (
        <div key={s.id} className="card">

          <h3>
            {s.name} {isActive(s.id) && <span className="online">● Online</span>}
          </h3>

          <p>ID: {s.id}</p>
          <p>{s.address}</p>
          <p>{s.grade}</p>

          <button onClick={() => checkIn(s.id)}>Check In</button>
          <button onClick={() => checkOut(s.id)}>Check Out</button>
          <button className="delete" onClick={() => deleteStudent(s.id)}>
            Delete
          </button>
        </div>
      ))}

      {/* HISTORY */}
      <h2>Attendance History</h2>

      <button onClick={download}>Download</button>
      <button onClick={clearData}>Clear</button>

      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>In</th>
            <th>Out</th>
            <th>Status</th>
          </tr>
        </thead>

        <tbody>
          {history.map((h, i) => (
            <tr key={i}>
              <td>{h.student_id}</td>
              <td>{h.check_in}</td>
              <td>{h.check_out || "Active"}</td>
              <td>{h.check_out ? "Done" : "Active"}</td>
            </tr>
          ))}
        </tbody>
      </table>

    </div>
  );
}