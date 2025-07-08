import React from 'react';
import Select from 'react-select';
import type { MultiValue } from 'react-select';

export interface Option {
  value: string;
  label: string;
}

interface MultiSelectProps {
  options: Option[];
  value: Option[];
  onChange: (value: MultiValue<Option>) => void;
}

const MultiSelect: React.FC<MultiSelectProps> = ({ options, value, onChange }) => {
  return (
    <Select
      isMulti
      options={options}
      value={value}
      onChange={onChange}
      className="multi-select"
      classNamePrefix="select"
    />
  );
};

export default MultiSelect;
