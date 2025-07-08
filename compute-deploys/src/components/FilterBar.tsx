import React, { useState } from 'react';
import MultiSelect from './MultiSelect';
import type { Option } from './MultiSelect';
import type { MultiValue } from 'react-select';

interface FilterBarProps {
  onSearch: () => void;
}

const clusterOptions: Option[] = [
  { value: 'cluster-1', label: 'Cluster 1' },
  { value: 'cluster-2', label: 'Cluster 2' },
  { value: 'cluster-3', label: 'Cluster 3' },
];

const FilterBar: React.FC<FilterBarProps> = ({ onSearch }) => {
  const [selectedClusters, setSelectedClusters] = useState<Option[]>([]);

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
