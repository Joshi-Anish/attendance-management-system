import { useState } from "react";

export default function Login() {
  const [form, setForm] = useState({ username: "", password: "" });

  const login = async () => {
    const res = await fetch("http://localhost:8080/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(form),
    });

    const data = await res.json();

    if (res.ok) {
      localStorage.setItem("user", JSON.stringify(data));

      if (data.role === "admin") window.location.href = "/admin";
      else if (data.role === "teacher") window.location.href = "/teacher";
      else window.location.href = "/student";
    } else {
      alert(data.message || "Login failed");
    }
  };

  return (
    <div className="auth-box">
      <h2>Login</h2>

      <input
        placeholder="Username"
        onChange={(e) => setForm({ ...form, username: e.target.value })}
      />

      <input
        type="password"
        placeholder="Password"
        onChange={(e) => setForm({ ...form, password: e.target.value })}
      />

      <button onClick={login}>Login</button>
    </div>
  );
}