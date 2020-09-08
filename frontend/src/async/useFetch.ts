import { useCallback, useState } from 'react';

export interface FetchOptions {
  json?: boolean;
  method?: 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE';
  headers?: Headers | Record<string, string>;
  body?: BodyInit;
}

interface Fetch {
  (uriOrRequest: string | Request, options?: FetchOptions): Promise<void>;
}

export interface UseFetchReturn {
  succeeded: boolean;
  pending: boolean;
  failed: boolean;
  error: null | Error;
  rawResponse: null | Response;
  result: null | unknown;
  reset: () => void;
  trigger: (uriOrRequest: string | Request, options?: FetchOptions) => void;
}

export function useFetch(): UseFetchReturn {
  const [succeeded, setSucceeded] = useState(false);
  const [pending, setPending] = useState(false);
  const [error, setError] = useState<Error | null>(null);
  const [failed, setFailed] = useState(false);
  const [rawResponse, setRawResponse] = useState<null | Response>(null);
  const [result, setResult] = useState<null | unknown>(null);

  const reset = useCallback(() => {
    setSucceeded(false);
    setPending(false);
    setError(null);
    setFailed(false);
    setRawResponse(null);
    setResult(null);
  }, []);

  const handleError = useCallback((e: Error, rawResponse: null | Response) => {
    setError(e);
    setFailed(true);
    setPending(false);
    setRawResponse(rawResponse);
  }, []);

  const handleSuccess = useCallback((result: unknown, rawResponse: Response) => {
    setResult(result);
    setSucceeded(true);
    setPending(false);
    setRawResponse(rawResponse);
  }, [])

  const trigger = useCallback(
    async (
      uriOrRequest: string | Request,
      options?: FetchOptions,
    ): Promise<void> => {
      let response: Response;
      let result: unknown

      try {
        if (uriOrRequest instanceof Request) {
          response = await fetch(uriOrRequest);
        } else {
          response = await fetch(uriOrRequest, {
            method: options?.method,
            body: options?.body,
            headers: options?.json
              ? {
                  Accept: 'application/json',
                  'Content-Type': 'application/json',
                  ...options?.headers,
                }
              : options?.headers,
          });
        }

        result = await (options?.json ? response.json() : response.text());
      } catch (e) {
        return handleError(e, null);
      }

      if (!response.ok) {
        return handleError(new NetworkError(await response.text(), response.status), response);
      }

      return handleSuccess(result, response);
    },
    [handleError, handleSuccess],
  );

  return {
    succeeded,
    pending,
    failed,
    error,
    result,
    rawResponse,
    reset,
    trigger,
  };
}

export class NetworkError extends Error {
  public constructor(
    message: string,
    public readonly statusCode: number | null = null,
  ) {
    super(message);
  }
}