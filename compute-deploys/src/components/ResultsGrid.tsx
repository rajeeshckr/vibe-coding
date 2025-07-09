import React from 'react';

// Define a type for the nested attributes within a log entry
interface LogAttributes {
  stageTimestamp: string;
  objectRef?: {
    name?: string;
  };
  requestObject?: {
    metadata?: {
      labels?: {
        product?: string;
        project?: string;
        role?: string;
      };
    };
  };
  // Add other potential properties to avoid TypeScript errors
  [key: string]: unknown; // Use unknown instead of any for better type safety
}

// Define the structure of a single log from the Datadog API response
export interface DatadogLog {
  id: string;
  type: string;
  attributes: {
    attributes: LogAttributes;
    // Add other potential properties
    [key: string]: unknown; // Use unknown instead of any
  };
}

interface ResultsGridProps {
  results: DatadogLog[];
  loading: boolean;
  error: string | null;
}

// Function to format UTC date string to AEST (UTC+10)
const formatToAEST = (utcDateString: string) => {
  const date = new Date(utcDateString);
  // AEST is UTC+10. We can create a new date by adding the offset.
  const aestDate = new Date(date.getTime() + 10 * 60 * 60 * 1000);
  // Format it to a more readable string, indicating the timezone.
  // Using toISOString and replacing Z with the offset.
  return aestDate.toISOString().replace('T', ' ').replace('Z', ' AEST');
};


const ResultsGrid: React.FC<ResultsGridProps> = ({ results, loading, error }) => {
  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  if (!results || results.length === 0) {
    return <div>No results found.</div>;
  }

  return (
    <table>
      <thead>
        <tr>
          <th>Date</th>
          <th>Name</th>
          <th>Product</th>
          <th>Project</th>
          <th>Role</th>
        </tr>
      </thead>
      <tbody>
        {results.map((log) => {
          const attrs = log.attributes?.attributes || {};
          return (
            <tr key={log.id}>
              <td>{attrs.stageTimestamp ? formatToAEST(attrs.stageTimestamp) : 'N/A'}</td>
              <td>{attrs.objectRef?.name ?? 'N/A'}</td>
              <td>{attrs.requestObject?.metadata?.labels?.product ?? 'N/A'}</td>
              <td>{attrs.requestObject?.metadata?.labels?.project ?? 'N/A'}</td>
              <td>{attrs.requestObject?.metadata?.labels?.role ?? 'N/A'}</td>
            </tr>
          );
        })}
      </tbody>
    </table>
  );
};

export default ResultsGrid;
