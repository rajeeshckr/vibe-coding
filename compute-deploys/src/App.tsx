import { useState, useEffect } from 'react';
import './App.css';
import FilterBar from './components/FilterBar';
import ResultsGrid from './components/ResultsGrid';
import type { Option } from './components/MultiSelect';
import type { DatadogLog } from './components/ResultsGrid'; // Import the type

interface Service {
  name: string;
  namespace: string;
  kind: string;
}

function App() {
  const [results, setResults] = useState<DatadogLog[]>([]); // Use DatadogLog type
  const [scale, setScale] = useState(1);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const handleResize = () => {
      const newScale = Math.min(Math.max(window.innerWidth / 1280, 0.5), 2.5);
      setScale(newScale);
    };

    window.addEventListener('resize', handleResize);
    handleResize(); // Initial scale

    return () => window.removeEventListener('resize', handleResize);
  }, []);

  const handleSearch = (filters: { clusters: Option[], services: Option[], from: string, to: string }) => {
    setLoading(true);
    setError(null);
    const { clusters, services, from, to } = filters;
    fetch('/api/search', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ clusters, services, startDate: from, endDate: to }),
    })
      .then(res => {
        if (!res.ok) {
          return res.text().then(text => { throw new Error(text || 'Network response was not ok') });
        }
        return res.json();
      })
      .then(data => {
        // The API returns { data: [...] }, so we access data.data
        if (data && data.data) {
          setResults(data.data);
        } else {
          setResults([]);
        }
      })
      .catch(err => {
        setError(err.message);
        setResults([]);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  return (
    <div className="app-container" style={{ transform: `scale(${scale})` }}>
      <div className="app">
        <FilterBar onSearch={handleSearch} />
        <div className="results-grid-container">
          <ResultsGrid results={results} loading={loading} error={error} />
        </div>
      </div>
    </div>
  );
}

export default App;
