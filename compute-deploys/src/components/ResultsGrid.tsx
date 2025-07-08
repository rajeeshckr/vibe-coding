
import React from 'react';

interface Result {
  name: string;
  product: string;
  project: string;
  role: string;
  cluster: string;
}

interface ResultsGridProps {
  results: Result[];
}

const ResultsGrid: React.FC<ResultsGridProps> = ({ results }) => {
  return (
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Product</th>
          <th>Project</th>
          <th>Role</th>
          <th>Cluster</th>
        </tr>
      </thead>
      <tbody>
        {results.map((result, index) => (
          <tr key={index}>
            <td>{result.name}</td>
            <td>{result.product}</td>
            <td>{result.project}</td>
            <td>{result.role}</td>
            <td>{result.cluster}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default ResultsGrid;
