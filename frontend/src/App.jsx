import React from 'react';
import { useState, useEffect } from 'react';
import CategorySelector from './components/CategorySelector';
import TaskDisplay from './components/TaskDisplay';
import Stats from './components/Stats';
import Login from "./components/Login";

export default function App() {
  const [category, setCategory] = useState(null);
  const [task, setTask] = useState(null);
  const [stats, setStats] = useState({ fitness: 0, learning: 0, mindfulness: 0 });
  const [loggedIn, setLoggedIn] = useState(false);

  useEffect(() => {
    console.log("Current task:", task);
  }, [task]);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) setLoggedIn(true);
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("token");
    setLoggedIn(false);
  };

  return (
    <div className="App">
      {loggedIn ? (
        <>
          <h1>Improov</h1>
          <CategorySelector 
            setCategory={setCategory}
            userId="tester" // or a real user from state/auth
            setTask={setTask} 
          />
          <TaskDisplay 
            userId="tester" 
            category={category} 
            task={task} 
            setTask={setTask} 
            setStats={setStats} />
          <Stats stats={stats} />
        </>
      ) : (
        <Login onLogin={() => setLoggedIn(true)} />
      )}
    </div>
  );
}