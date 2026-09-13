type SafeAsync<T, E = Error> = Promise<readonly [T, null] | readonly [null, E]>;

export async function errAsync<T, E = Error>(
  promise: Promise<T>
): SafeAsync<T, E> {
  try {
    const data = await promise;
    return [data, null] as const;
  } catch (error) {
    const actualError = error instanceof Error ? error : new Error(String(error));
    return [null, actualError as unknown as E] as const;
  }
}

type SafeSync<T, E = Error> = readonly [T, null] | readonly [null, E];

export function errSync<T, E = Error>(fn: () => T): SafeSync<T, E> {
  try {
    return [fn(), null] as const;
  } catch (error) {
    const actualError = error instanceof Error ? error : new Error(String(error));
    return [null, actualError as unknown as E] as const;
  }
}