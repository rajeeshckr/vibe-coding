import React, { useState, useEffect } from 'react';
import MultiSelect from './MultiSelect';
import type { Option } from './MultiSelect';
import type { MultiValue } from 'react-select';

interface Service {
  name: string;
  namespace: string;
  kind: string;
}

interface FilterBarProps {
  onSearch: (filters: { clusters: Option[], services: Service[], from: string, to: string }) => void;
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
  const [selectedServices, setSelectedServices] = useState<Service[]>([]);
  const [from, setFrom] = useState('');
  const [to, setTo] = useState('');

  useEffect(() => {
    fetch('/api/clusters')
      .then(res => {
        if (!res.ok) {
          throw new Error('Failed to fetch clusters');
        }
        return res.json();
      })
      .then(data => setClusterOptions(data))
      .catch(error => {
        console.error("Error fetching clusters:", error);
      });
  }, []);

  const handleClusterChange = (selected: MultiValue<Option>) => {
    setSelectedClusters(selected as Option[]);
  };

  const handleServiceChange = (selectedOptions: MultiValue<Option>) => {
    const services = selectedOptions.map(option => {
        return initialServices.find(service => service.name === option.value);
    }).filter((service): service is Service => service !== undefined);
    setSelectedServices(services);
  };

  const handleSearchClick = () => {
    onSearch({
      clusters: selectedClusters,
      services: selectedServices,
      from,
      to,
    });
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
        value={selectedServices.map(s => ({ value: s.name, label: s.name }))}
        onChange={handleServiceChange}
      />
      <input type="datetime-local" value={from} onChange={e => setFrom(e.target.value)} />
      <input type="datetime-local" value={to} onChange={e => setTo(e.target.value)} />
      <button onClick={handleSearchClick}>Search</button>
    </div>
  );
};

export default FilterBar;
