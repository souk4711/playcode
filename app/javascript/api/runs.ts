const create = async () => {
  await new Promise((resolve) => setTimeout(resolve, 2000))
}

export default {
  create
}
