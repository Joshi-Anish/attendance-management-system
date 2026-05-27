export default function StudentCard({ student, onCheckIn, onCheckOut }) {
  return (
    <div className="student-card">
      <h3>{student.name}</h3>
      <p>ID: {student.id}</p>
      <p>Grade: {student.grade}</p>

      <div className="buttons">
        <button onClick={() => onCheckIn(student.id)}>
          Check In
        </button>
        <button onClick={() => onCheckOut(student.id)}>
          Check Out
        </button>
      </div>
    </div>
  );
}