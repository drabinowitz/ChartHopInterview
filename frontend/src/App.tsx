import React, { useCallback, useEffect } from 'react';
import './App.css';
import { useFetch } from './async/useFetch';
import { CreateTodoForm } from './CreateTodoForm/CreateTodoForm';
import { Table, Todo } from './Table/Table';

function App() {
  const listFetcher = useFetch();
  const createFetcher = useFetch();

  const cTrigger = createFetcher.trigger;
  const create = useCallback((todo: Todo) => {
    cTrigger("http://localhost:8081/todos",{
      method: "POST",
      body: JSON.stringify(todo),
    })
  }, [cTrigger])

  const trigger = listFetcher.trigger;
  const list = useCallback(() => {
    trigger("http://localhost:8081/todos?user=1")
  }, [trigger])

  useEffect(() => {
    list();
  }, [list])

  const listFetcherResponse = listFetcher.result as null | {todos: Todo[]};

  return (
    <div className="App">
      <Table todos={listFetcherResponse?.todos} />
      <CreateTodoForm onSubmit={create} />
    </div>
  );
}

export default App;
