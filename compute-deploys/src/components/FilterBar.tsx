import React, { useState, useEffect } from 'react';
import MultiSelect from './MultiSelect';
import type { Option } from './MultiSelect';
import type { MultiValue } from 'react-select';

interface FilterBarProps {
  onSearch: () => void;
}

const FilterBar: React.FC<FilterBarProps> = ({ onSearch }) => {
  const [clusterOptions, setClusterOptions] = useState<Option[]>([]);
  const [selectedClusters, setSelectedClusters] = useState<Option[]>([]);

  useEffect(() => {
    fetch('/api/clusters')
      .then(res => res.json())
      .then(data => setClusterOptions(data));
  }, []);

  const handleClusterChange = (selected: MultiValue<Option>) => {
    setSelectedClusters(selected as Option[]);
  };

  return (
    <div className="filter-bar">
      <MultiSelect
        options={clusterOptions}
        value={selectedClusters}
        onChange={handleClusterChange}
      />
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
