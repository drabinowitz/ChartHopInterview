import React from 'react';
import { TableHeader } from './TableHeader';
import { TableRow } from './TableRow';

export interface Todo {
  id: string;
  entityId: string;
  orgId: string;
  status: string;
  todoType: string;
  userId: string;
}

export interface TableProps {
  todos?: Todo[];
}

export function Table({todos}: TableProps): JSX.Element {
  return <table>
<TableHeader todo={todos?.[0]} />
    <tbody>
      {todos?.map(todo => <TableRow key={todo.id} todo={todo} />)}
    </tbody>
  </table>
}