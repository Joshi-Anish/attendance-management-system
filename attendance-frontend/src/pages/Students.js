import { useEffect, useState } from "react";
import StudentCard from "../components/StudentCard";

export default function Students() {
  const [students, setStudents] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/students")
      .then((res) => res.json())
      .then(setStudents);
  }, []);

  const checkIn = (id) => {
    fetch("http://localhost:8080/checkin", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ student_id: id }),
    }).then(() => alert("Checked In"));
  };

  const checkOut = (id) => {
    fetch("http://localhost:8080/checkout", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ student_id: id }),
    }).then(() => alert("Checked Out"));
  };

  return (
    <div className="container">
      <h2>Students</h2>

      {students.map((s) => (
        <StudentCard
          key={s.id}
          student={s}
          onCheckIn={checkIn}
          onCheckOut={checkOut}
        />
      ))}
    </div>
  );
}