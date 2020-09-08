import React from 'react';
import { Todo } from './Table';

export interface TableRowProps {
  todo: Todo;
}

export function TableRow({todo}: TableRowProps): JSX.Element {
  const keys = Object.keys(todo);
  return <tr>{keys.map(key => <td id={key}>{(todo as any)[key]}</td>)}</tr>
}