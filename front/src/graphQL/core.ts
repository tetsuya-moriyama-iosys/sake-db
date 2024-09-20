export function wrapVariables<T = unknown>(request: T) {
  return {
    variables: request,
  };
}
