import { useEffect, useState } from "react";

export default function Students() {
  const [students, setStudents] = useState([]);
  const [history, setHistory] = useState([]);

  const [form, setForm] = useState({
    id: "",
    name: "",
    address: "",
    grade: "",
  });

  // LOAD STUDENTS
  const loadStudents = () => {
    fetch("http://localhost:8080/students")
      .then((res) => res.json())
      .then((data) => setStudents(Array.isArray(data) ? data : []));
  };

  // LOAD HISTORY
  const loadHistory = () => {
    fetch("http://localhost:8080/dashboard")
      .then((res) => res.json())
      .then((data) =>
        setHistory(Array.isArray(data.records) ? data.records : [])
      );
  };

  useEffect(() => {
    loadStudents();
    loadHistory();
  }, []);

  // ADD STUDENT
  const addStudent = () => {
    fetch("http://localhost:8080/students", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(form),
    }).then(() => {
      setForm({ id: "", name: "", address: "", grade: "" });
      loadStudents();
    });
  };

  // DELETE STUDENT
  const deleteStudent = (id) => {
    fetch(`http://localhost:8080/students/delete?id=${id}`, {
      method: "DELETE",
    }).then(loadStudents);
  };

  // CHECK IN
  const checkIn = (id) => {
    fetch("http://localhost:8080/checkin", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ student_id: id }),
    }).then(loadHistory);
  };

  // CHECK OUT
  const checkOut = (id) => {
    fetch("http://localhost:8080/checkout", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ student_id: id }),
    }).then(loadHistory);
  };

  // DOWNLOAD CSV
  const downloadCSV = () => {
    window.open("http://localhost:8080/attendance/download");
  };

  // CLEAR DATA
  const clearData = () => {
    if (!window.confirm("⚠️ This will delete ALL attendance data. Continue?"))
      return;

    fetch("http://localhost:8080/attendance/clear", {
      method: "DELETE",
    }).then(() => {
      loadHistory();
    });
  };

  // ACTIVE CHECK
  const isActive = (id) => {
    return history.some(
      (h) =>
        String(h.student_id) === String(id) &&
        (!h.check_out || h.check_out === "")
    );
  };

  return (
    <div className="container">

      {/* ================= ADD STUDENT ================= */}
      <h2>Add Student</h2>

      <div className="form">
        <input
          placeholder="ID"
          value={form.id}
          onChange={(e) => setForm({ ...form, id: e.target.value })}
        />
        <input
          placeholder="Name"
          value={form.name}
          onChange={(e) => setForm({ ...form, name: e.target.value })}
        />
        <input
          placeholder="Address"
          value={form.address}
          onChange={(e) => setForm({ ...form, address: e.target.value })}
        />
        <input
          placeholder="Grade"
          value={form.grade}
          onChange={(e) => setForm({ ...form, grade: e.target.value })}
        />

        <button onClick={addStudent}>Add Student</button>
      </div>

      {/* ================= STUDENTS ================= */}
      <h2>Students</h2>

      {students.map((s) => (
        <div key={s.id} className="student-card">
          <div style={{ display: "flex", gap: "10px", alignItems: "center" }}>
            <h3 style={{ margin: 0 }}>{s.name}</h3>

            {isActive(s.id) && (
              <span className="online-dot">● Online</span>
            )}
          </div>

          <p>ID: {s.id}</p>
          <p>Address: {s.address}</p>
          <p>Grade: {s.grade}</p>

          <div className="buttons">
            <button onClick={() => checkIn(s.id)}>Check In</button>
            <button onClick={() => checkOut(s.id)}>Check Out</button>
            <button className="delete" onClick={() => deleteStudent(s.id)}>
              Delete
            </button>
          </div>
        </div>
      ))}

      {/* ================= ATTENDANCE HISTORY ================= */}
      <h2>Attendance History</h2>

      {/* 🔥 NEW ACTION BUTTONS */}
      <div style={{ display: "flex", gap: "10px", marginBottom: "10px" }}>
        <button onClick={downloadCSV}>Download CSV</button>
        <button className="delete" onClick={clearData}>
          Clear Data
        </button>
      </div>

      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Check In</th>
            <th>Check Out</th>
            <th>Status</th>
          </tr>
        </thead>

        <tbody>
          {history.map((h, i) => (
            <tr key={i}>
              <td>{h.student_id}</td>
              <td>{h.check_in}</td>
              <td>{h.check_out || "Active"}</td>
              <td>{h.check_out ? "Completed" : "Active"}</td>
            </tr>
          ))}
        </tbody>
      </table>

    </div>
  );
}