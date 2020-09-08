import React, { FormEvent, useCallback } from 'react';
import { Todo } from '../Table/Table';

export interface CreateTodoFormProps {
  onSubmit: (todo: Todo) => void;
}

export function CreateTodoForm({onSubmit}: CreateTodoFormProps): JSX.Element {
  const handleSubmit = useCallback((e: FormEvent) => {
    e.preventDefault();

    onSubmit({
      orgId: '1',
      id: 'ignoreme',
      entityId: 'form1',
      userId: '1',
      status: 'PENDING',
      todoType: 'CHANGE_APPROVE',
    })

  }, [onSubmit])
  return <form onSubmit={handleSubmit} id="create-todo-form">
    <button type="submit" form="create-todo-form">
      Create Todo!
    </button>
  </form>
}