import React from 'react';
import { useState, useEffect } from 'react';
import jwt_decode from "jwt-decode";
import CategorySelector from './components/CategorySelector';
import TaskDisplay from './components/TaskDisplay';
import Stats from './components/Stats';
import Login from "./components/Login";
import Register from "./components/Register";

function App() {
  const [category, setCategory] = useState(null);
  const [task, setTask] = useState(null);
  const [stats, setStats] = useState({ fitness: 0, learning: 0, mindfulness: 0 });
  const [loggedIn, setLoggedIn] = useState(false);
  const [userId, setUserId] = useState(null);
  const [mode, setMode] = useState("login");

  useEffect(() => {
    console.log("Current task:", task);
  }, [task]);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      try {
        const decoded = jwt_decode(token);
        setUserId(decoded.username || decoded.userId || decoded.sub);
        setLoggedIn(true);
      } catch (err) {
        console.error("Invalid token", err);
        localStorage.removeItem("token");
      }
    }
  }, []);


  const handleLoginOrRegister = (token) => {
    localStorage.setItem("token", token);
    const decoded = jwt_decode(token);
    setUserId(decoded.username || decoded.userId || decoded.sub);
    setLoggedIn(true);
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    setLoggedIn(false);
    setUserId(null);
  };

  return (
    <div className="App">
      {loggedIn ? (
        <>
          <button onClick={handleLogout}>Logout</button>
          <h1>Improov</h1>
          <CategorySelector 
            setCategory={setCategory}
            userId={userId} // or a real user from state/auth
            setTask={setTask} 
          />
          <TaskDisplay 
            userId={userId}
            category={category} 
            task={task} 
            setTask={setTask} 
            setStats={setStats} />
          <Stats stats={stats} />
        </>
      ) : (
        <>
          {mode == "login" ? (
            <>
              <Login onLogin={handleLoginOrRegister} />
              <p>
                Don't have an account?{" "}
                <button onClick={() => setMode("register")}>Register</button>
              </p>
            </>
          ) : (
            <>
              <Register onRegister={handleLoginOrRegister} />
              <p>
                Already have an account?{" "}
                <button onClick={() => setMode("login")}>Login</button>
              </p>
            </>
          )}
        </>
      )}
    </div>
  );
}

export default App;