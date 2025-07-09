import React, { useState, useEffect } from 'react';
import MultiSelect from './MultiSelect';
import type { Option } from './MultiSelect';
import type { MultiValue } from 'react-select';

interface FilterBarProps {
  onSearch: () => void;
}

interface Service {
  name: string;
  namespace: string;
  kind: string;
}

const initialServices: Service[] = [
  {
    name: 'kube-node-problem-detector',
    namespace: 'kube-system',
    kind: 'daemonsets',
  },
];

const serviceOptions: Option[] = initialServices.map(service => ({
  value: service.name,
  label: service.name,
}));

const FilterBar: React.FC<FilterBarProps> = ({ onSearch }) => {
  const [clusterOptions, setClusterOptions] = useState<Option[]>([]);
  const [selectedClusters, setSelectedClusters] = useState<Option[]>([]);
  const [selectedServices, setSelectedServices] = useState<Option[]>([]);

  useEffect(() => {
    fetch('/api/clusters')
      .then(res => res.json())
      .then(data => setClusterOptions(data));
  }, []);

  const handleClusterChange = (selected: MultiValue<Option>) => {
    setSelectedClusters(selected as Option[]);
  };

  const handleServiceChange = (selected: MultiValue<Option>) => {
    setSelectedServices(selected as Option[]);
  };

  return (
    <div className="filter-bar">
      <MultiSelect
        options={clusterOptions}
        value={selectedClusters}
        onChange={handleClusterChange}
      />
      <MultiSelect
        options={serviceOptions}
        value={selectedServices}
        onChange={handleServiceChange}
      />
      <input type="datetime-local" />
      <input type="datetime-local" />
      <button onClick={onSearch}>Search</button>
    </div>
  );
};

export default FilterBar;
