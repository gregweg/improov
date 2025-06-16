import React from 'react';

export default function Stats({ stats }) {
    const [user, setUser] = useState(null);

    useEffect(() => {
        const token = localStorage.getItem("token");

        fetch("/api/me", {
            headers: { Authorization: `Bearer ${token}` },
        })
        .then((res) => res.json())
        .then(setUser);
    }, []);
    return (
      <div>
        <h2>Your Stats</h2>
        <ul>
          {Object.entries(stats).map(([stat, value]) => (
            <li key={stat}>{stat.charAt(0).toUpperCase() + stat.slice(1)}: {value}</li>
          ))}
        </ul>
      </div>
    );
  }