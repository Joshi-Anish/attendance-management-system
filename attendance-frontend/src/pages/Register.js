import { useState } from "react";

export default function Register() {
  const [form, setForm] = useState({
    username: "",
    password: "",
    role: "student",
  });

  const register = async () => {
    await fetch("http://localhost:8080/register", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(form),
    });

    alert("User Registered");
  };

  return (
    <div className="auth-box">
      <h2>Register</h2>

      <input placeholder="Username"
        onChange={(e) => setForm({ ...form, username: e.target.value })}
      />

      <input type="password"
        placeholder="Password"
        onChange={(e) => setForm({ ...form, password: e.target.value })}
      />

      <select onChange={(e) => setForm({ ...form, role: e.target.value })}>
        <option value="student">Student</option>
        <option value="teacher">Teacher</option>
        <option value="admin">Admin</option>
      </select>

      <button onClick={register}>Register</button>
    </div>
  );
}