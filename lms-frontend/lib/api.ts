// const API_URL = "http://localhost:8080";

// export async function apiFetch(
//     endpoint: string,
//     options: RequestInit = {}
// ) {
//     const token = localStorage.getItem("token");

//     return fetch(`${API_URL}/${endpoint}`, {    
//         ...options,
//         headers: {
//             "Content-Type": "application/json",
//             ...(token &&{Authorization: `Bearer ${token}`})
//         },
//     })
// }

const BASE_URL = "http://localhost:8080";

export function apiFetch(path: string, options?: RequestInit) {
  return fetch(`${BASE_URL}${path}`, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...(options?.headers || {}),
    },
  });
}
