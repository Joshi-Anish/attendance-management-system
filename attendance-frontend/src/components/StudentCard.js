export default function StudentCard({ student, onCheckIn, onCheckOut, onDelete }) {
  return (
    <div className="student-card">
      <h3>{student.name}</h3>
      <p>ID: {student.id}</p>
      <p>Address: {student.address}</p>
      <p>Grade: {student.grade}</p>

      <div className="buttons">
        <button onClick={() => onCheckIn(student.id)}>Check In</button>
        <button onClick={() => onCheckOut(student.id)}>Check Out</button>
        <button className="delete" onClick={() => onDelete(student.id)}>
          Delete
        </button>
      </div>
    </div>
  );
}