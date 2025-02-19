export const httpRequest = async (
  url: string,
  body?: BodyInit | null,
  method: string = "GET",
  cache: RequestCache = "no-store",
  headers: Record<string, string> = {}
): Promise<{ success: boolean; data?: any; message?: string; error?: any }> => {
  try {
    const response = await fetch(`http://localhost:8080/api/${url}`, {
      method,
      body: body && method !== "GET" ? body : null,
      headers: {
        "Content-Type": "application/json",
        ...headers,
      },
      cache,
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.message || `HTTP Error: ${response.status}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    return {
      success: false,
      message: "An error occurred while making the request.",
      error: error instanceof Error ? error.message : error,
    };
  }
};
