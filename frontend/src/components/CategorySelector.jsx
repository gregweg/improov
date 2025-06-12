import React, { useState, useEffect } from 'react';
import axios from 'axios';

export default function CategorySelector({ setCategory, setTask, userId }) {
  const [categories, setCategories] = useState([]);

  useEffect(() => {
    axios.get('/api/categories')
      .then(res => setCategories(res.data))
      .catch(console.error);
  }, []);

  const selectCategory = async (cat) => {
    setCategory(cat);
    try {
        const res = await axios.get(`/api/tasks/suggest?category=${cat.name.toLowerCase()}&userId=${userId}`);
        if (res.data) {
            console.log("Suggested task:", res.data);
            setTask(res.data);
        } else {
            console.warn("No task returned");
            setTask(null);
        }
    } catch (err) {
        console.error("Failed to fetch suggested task", err);
        setTask(null);
    }
  };

  return (
    <div>
      <h2>Choose a Category</h2>
      {categories.map(cat => (
        <button key={cat.name} onClick={() => selectCategory(cat)}>{cat.name}</button>
      ))}
    </div>
  );
}