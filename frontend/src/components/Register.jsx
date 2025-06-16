import React from 'react';
import { useState } from "react"

export default function Register({ onRegister }) {
    const [username, setUsername] = useState("");
    const [name, setName] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");

    const handleRegister = async (e) => {
        e.preventDefault();
        setError("");

        try {
            const res = await fetch("/api/register", {
               method: "POST",
               headers: {
                "Content-Type": "application/json"
               },
               body: JSON.stringify({
                username,
                password,
                name,
               }),
            });

            if (!res.ok) {
                const errText = await res.text();
                throw new Error(errText || "Registration failed");
            }

            const { token } = await res.json();
            localStorage.setItem("token", token);

            onRegister(token); // Pass token to parent (App)
        } catch (err) {
            setError(err.message || "Something went wrong");
        }
    };

    return (
        <form onSubmit={handleRegister} className="register-form">
            <h2>Register</h2>
            {error && <p style={{ color: "red" }}>{error}</p>}

            <input
                type="text"
                placeholder="Username"
                value={username}
                required
                onChange={(e) => setUsername(e.target.value)}
            />

            <input
                type="text"
                placeholder="Full Name"
                value={name}
                required
                onChange={(e) => setName(e.target.value)}
            />

            <input
                type="password"
                placeholder="Password"
                value={password}
                required
                onChange={(e) => setPassword(e.target.value)}
            />

            <button type="submit">Register</button>
        </form>
    );
}