import React from 'react';

export default function Stats({ stats }) {
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