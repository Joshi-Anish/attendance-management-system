import { useEffect, useState } from "react";

export default function Dashboard() {
  const [data, setData] = useState(null);

  useEffect(() => {
    fetch("http://localhost:8080/dashboard")
      .then((res) => res.json())
      .then(setData)
      .catch(console.error);
  }, []);

  if (!data) return <p>Loading...</p>;

  return (
    <div className="container">
      <h2>Dashboard</h2>

      <div className="cards">
        <div className="card">
          <h4>Total Students</h4>
          <h2>{data.total_students}</h2>
        </div>

        <div className="card">
          <h4>Active Students</h4>
          <h2>{data.active_students}</h2>
        </div>

        <div className="card">
          <h4>Completed Sessions</h4>
          <h2>{data.completed_sessions}</h2>
        </div>
      </div>

      <h3>Attendance History</h3>

      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Check In</th>
            <th>Check Out</th>
          </tr>
        </thead>
        <tbody>
          {data.records.map((r, i) => (
            <tr key={i}>
              <td>{r.student_id}</td>
              <td>{r.check_in}</td>
              <td>{r.check_out || "Active"}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}