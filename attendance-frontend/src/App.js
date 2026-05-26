import { useEffect, useState } from "react";

function App() {
  const [students, setStudents] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let isMounted = true;

    fetch("http://127.0.0.1:8080/students")
      .then(res => res.json())
      .then(data => {
        console.log("API DATA:", data);

        if (isMounted) {
          setStudents(data);
          setLoading(false);
        }
      })
      .catch(err => {
        console.log("ERROR:", err);
        setLoading(false);
      });

    return () => {
      isMounted = false;
    };
  }, []);

  return (
    <div style={{ padding: "20px" }}>
      <h1>Attendance System</h1>

      <h2>Students</h2>

      {loading ? (
        <p>Loading...</p>
      ) : students.length === 0 ? (
        <p>No students found</p>
      ) : (
        students.map((s) => (
          <div key={s.id}>
            {s.name} ({s.id})
          </div>
        ))
      )}
    </div>
  );
}

export default App;