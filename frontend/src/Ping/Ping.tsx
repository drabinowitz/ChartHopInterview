import React, { useEffect } from 'react';
import { useFetch } from '../async/useFetch';

export function Ping(): JSX.Element {
  const {pending, result, trigger} = useFetch();

  useEffect(() => {
    trigger('http://localhost:8081/ping');
  }, [trigger])

  return <div>{pending ? 'still loading' : result as string}</div>
}