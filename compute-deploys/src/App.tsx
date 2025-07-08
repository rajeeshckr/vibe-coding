import { useState, useEffect } from 'react';
import './App.css';
import FilterBar from './components/FilterBar';
import ResultsGrid from './components/ResultsGrid';

interface Result {
  name: string;
  product: string;
  project: string;
  role: string;
  cluster: string;
}

function App() {
  const [results, setResults] = useState<Result[]>([]);
  const [scale, setScale] = useState(1);

  useEffect(() => {
    const handleResize = () => {
      const newScale = Math.min(Math.max(window.innerWidth / 1280, 0.5), 2.5);
      setScale(newScale);
    };

    window.addEventListener('resize', handleResize);
    handleResize(); // Initial scale

    return () => window.removeEventListener('resize', handleResize);
  }, []);

  const handleSearch = () => {
    // Mock API call
    const mockResults: Result[] = [
      { name: 'service-a', product: 'Product A', project: 'Project X', role: 'backend', cluster: 'Cluster 1' },
      { name: 'service-b', product: 'Product B', project: 'Project Y', role: 'frontend', cluster: 'Cluster 1' },
      { name: 'service-c', product: 'Product A', project: 'Project Z', role: 'backend', cluster: 'Cluster 2' },
      { name: 'service-d', product: 'Product C', project: 'Project X', role: 'backend', cluster: 'Cluster 1' },
      { name: 'service-e', product: 'Product B', project: 'Project Y', role: 'frontend', cluster: 'Cluster 2' },
      { name: 'service-f', product: 'Product C', project: 'Project Z', role: 'backend', cluster: 'Cluster 2' },
      { name: 'service-g', product: 'Product A', project: 'Project X', role: 'frontend', cluster: 'Cluster 1' },
      { name: 'service-h', product: 'Product B', project: 'Project Z', role: 'backend', cluster: 'Cluster 2' },
      { name: 'service-i', product: 'Product A', project: 'Project Y', role: 'frontend', cluster: 'Cluster 1' },
      { name: 'service-j', product: 'Product C', project: 'Project X', role: 'backend', cluster: 'Cluster 2' },
    ];
    setResults(mockResults);
  };

  return (
    <div className="App" style={{ transform: `scale(${scale})`, transformOrigin: 'top' }}>
      <FilterBar onSearch={handleSearch} />
      <div className="results-container">
        <ResultsGrid results={results} />
      </div>
    </div>
  );
}

export default App;
