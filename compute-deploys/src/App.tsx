import { useState } from 'react';
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

  const handleSearch = () => {
    // Mock API call
    const mockResults: Result[] = [
      { name: 'service-a', product: 'Product A', project: 'Project X', role: 'backend', cluster: 'Cluster 1' },
      { name: 'service-b', product: 'Product B', project: 'Project Y', role: 'frontend', cluster: 'Cluster 1' },
      { name: 'service-c', product: 'Product A', project: 'Project Z', role: 'backend', cluster: 'Cluster 2' },
    ];
    setResults(mockResults);
  };

  return (
    <div className="App">
      <FilterBar onSearch={handleSearch} />
      <ResultsGrid results={results} />
    </div>
  );
}

export default App;
