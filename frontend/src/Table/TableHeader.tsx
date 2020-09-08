import React from 'react';
import { Todo } from './Table';

export interface TableHeaderProps {
  todo?: Todo;
}

export function TableHeader({todo}: TableHeaderProps): JSX.Element {

  const keys = todo == null ? [] : Object.keys(todo);
return (
    <thead>
      <tr>
        {keys.map(key => <th id={key}>{key}</th>)}
      </tr>
    </thead>
)
}