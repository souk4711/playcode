const create = async () => {
  return await new Promise((resolve) =>
    setTimeout(() => {
      const result = {
        stdout: 'Hello, World!',
        stderr: ''
      }
      resolve(result)
    }, 2000)
  )
}

export default {
  create
}
