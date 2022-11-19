import axios from 'axios'

// Create a new instance of axios with a custom config
const http = axios.create({
  baseURL: '/api'
})

// Add a request interceptor
axios.interceptors.request.use(
  // Do something before request is sent
  (config) => {
    return config
  },
  // Do something with request error
  async (error) => {
    return await Promise.reject(error)
  }
)

// Add a response interceptor
http.interceptors.response.use(
  // Any status code that lie within the range of 2xx cause this function to trigger
  // Do something with response data
  (response) => {
    return response.data
  },
  // Any status codes that falls outside the range of 2xx cause this function to trigger
  // Do something with response error
  async (error) => {
    return await Promise.reject(error)
  }
)

export default http
