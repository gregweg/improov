import axios from 'axios';
import React from 'react';

export default function TaskDisplay({ userId, category, task, setTask, setStats }) {
  if (!task) return null;

  const completeTask = async () => {
    const categoryKey = category?.name?.toLowerCase();
    const payload = {
        user_id: userId,
        task_id: task?.id,
        category: categoryKey,
      };
    const token = localStorage.getItem("token");
    console.log("POST payload", payload); // ✅ debug log
    try {
        /*
        fetch("/api/suggest-task", {
  headers: {
    "Authorization": "Bearer " + token,
  },
});
        */
        await axios.post("/api/tasks/complete", payload);
        /*await axios.post('/api/tasks/complete', {   
                user_id: userId,
                task_id: task.id, 
                category: category.toLowerCase 
        });*/

        setTask(null);
        setStats(prev => ({
            ...prev,
            [categoryKey]: (prev[categoryKey] || 0) + 10
        }));
    } catch (err) {
        console.error("Error completing task:", err);
    }
  };

  return (
    <div>
      <h3>Your Task</h3>
      <p>{task.description}</p>
      <button onClick={completeTask}>Complete Task</button>
    </div>
  );
}