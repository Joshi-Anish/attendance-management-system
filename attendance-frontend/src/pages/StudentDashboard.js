import { useEffect, useState } from "react";

export default function StudentDashboard() {
  const [history, setHistory] = useState([]);

  const user = JSON.parse(localStorage.getItem("user"));

  useEffect(() => {
    fetch("http://localhost:8080/dashboard")
      .then((res) => res.json())
      .then((data) =>
        setHistory(Array.isArray(data.records) ? data.records : [])
      );
  }, []);

  const myHistory = history.filter(
    (h) => String(h.student_id) === String(user?.id)
  );

  return (
    <div className="container">
      <div className="card">
        <h2>Student Dashboard</h2>
        <p>Welcome {user?.username}</p>
      </div>

      <div className="card">
        <h3>My Attendance</h3>

        <table>
          <thead>
            <tr>
              <th>Check In</th>
              <th>Check Out</th>
              <th>Status</th>
            </tr>
          </thead>

          <tbody>
            {myHistory.map((h, i) => (
              <tr key={i}>
                <td>{h.check_in}</td>
                <td>{h.check_out || "Active"}</td>
                <td>{h.check_out ? "Done" : "Active"}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}