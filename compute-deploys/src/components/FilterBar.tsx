import React from 'react';

interface FilterBarProps {
  onSearch: () => void;
}

const FilterBar: React.FC<FilterBarProps> = ({ onSearch }) => {
  return (
    <div className="filter-bar">
      <select>
        <option>Cluster 1</option>
        <option>Cluster 2</option>
      </select>
      <select>
        <option>Service 1</option>
        <option>Service 2</option>
      </select>
      <input type="datetime-local" />
      <input type="datetime-local" />
      <button onClick={onSearch}>Search</button>
    </div>
  );
};

export default FilterBar;
