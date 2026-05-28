import { Link } from "react-router-dom";

export default function Header() {
  const user = JSON.parse(localStorage.getItem("user"));

  const logout = () => {
    localStorage.removeItem("user");
    window.location.href = "/";
  };

  return (
    <div className="header">
      <h2>Attendance System</h2>

      <nav>
        {!user && (
          <>
            <Link to="/">Login</Link>
            <Link to="/register">Register</Link>
          </>
        )}

        {user?.role === "student" && <Link to="/student">Dashboard</Link>}
        {user?.role === "teacher" && <Link to="/teacher">Dashboard</Link>}
        {user?.role === "admin" && <Link to="/admin">Dashboard</Link>}

        {user && <button onClick={logout}>Logout</button>}
      </nav>
    </div>
  );
}