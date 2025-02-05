export const httpRequest = (
    url: string,
    body?: BodyInit | null | undefined,
    method?: string,
    cache?: RequestCache,
    headers: any = {}
  ) => {
    return fetch(`api/${url}`, {
      method: method || "GET",
      body: body,
      headers: {
        ...headers,
      },
      cache: cache || "no-store",
    })
      .then((response) => {
        //todo: handle custom error cases
        return response.json();
      })
      .catch((error: any) => {
        return {
          success: false,
          message: "An error occurred:",
          error: error,
        };
      });
  };
  