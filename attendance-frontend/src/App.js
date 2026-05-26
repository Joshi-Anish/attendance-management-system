import { useEffect, useState } from "react";
import "./App.css";

const API = "http://127.0.0.1:8080";

function App() {
  const [students, setStudents] = useState([]);
  const [dashboard, setDashboard] = useState({});
  const [history, setHistory] = useState([]);
  const [loading, setLoading] = useState(true);

  const loadAll = async () => {
    try {
      const [sRes, dRes, hRes] = await Promise.all([
        fetch(`${API}/students`),
        fetch(`${API}/dashboard`),
        fetch(`${API}/attendance`)
      ]);

      const s = await sRes.json();
      const d = await dRes.json();
      const h = await hRes.json();

      setStudents(s || []);
      setDashboard(d || {});
      setHistory(h || []);
    } catch (err) {
      console.log("API Error:", err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadAll();
  }, []);

  // CHECK IN
  const checkIn = async (id) => {
    await fetch(`${API}/checkin`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ student_id: id }),
    });
    loadAll();
  };

  // CHECK OUT
  const checkOut = async (id) => {
    await fetch(`${API}/checkout`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ student_id: id }),
    });
    loadAll();
  };

  // STATUS LOGIC (IMPORTANT FIX)
  const getStatus = (studentId) => {
    const records = history.filter(h => h.student_id === studentId);

    if (records.length === 0) return "Not Active";

    const latest = records[0];

    if (!latest.check_out) return "Active 🟢";

    return "Completed 🔴";
  };

  return (
    <div className="container">

      <h1 className="title">HR Attendance Dashboard</h1>

      {/* DASHBOARD */}
      <div className="stats">
        <div className="box">
          <h2>{dashboard.total_students || 0}</h2>
          <p>Total Students</p>
        </div>

        <div className="box">
          <h2>{dashboard.active_students || 0}</h2>
          <p>Active Students</p>
        </div>

        <div className="box">
          <h2>{dashboard.completed_sessions || 0}</h2>
          <p>Completed Sessions</p>
        </div>
      </div>

      {/* STUDENTS */}
      <h2>Students</h2>

      {loading ? (
        <p>Loading...</p>
      ) : (
        <div className="grid">
          {students.map((s) => (
            <div className="card" key={s.id}>
              <h3>{s.name}</h3>
              <p>ID: {s.id}</p>
              <p>Grade: {s.grade}</p>

              <p><b>Status:</b> {getStatus(s.id)}</p>

              <div className="buttons">
                <button onClick={() => checkIn(s.id)}>
                  Check In
                </button>

                <button onClick={() => checkOut(s.id)}>
                  Check Out
                </button>
              </div>
            </div>
          ))}
        </div>
      )}

      {/* HISTORY */}
      <h2 style={{ marginTop: "30px" }}>Attendance History</h2>

      <div className="table">
        <div className="row header">
          <div>Student ID</div>
          <div>Check In</div>
          <div>Check Out</div>
          <div>Status</div>
        </div>

        {history.map((h, i) => (
          <div className="row" key={i}>
            <div>{h.student_id}</div>
            <div>{h.check_in}</div>
            <div>{h.check_out || "Active"}</div>
            <div>{h.check_out ? "Completed" : "Active"}</div>
          </div>
        ))}
      </div>

    </div>
  );
}

export default App;