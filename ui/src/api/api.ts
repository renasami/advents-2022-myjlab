export const post = async (url:string,opt:{}) => {
    const result =  await fetch(`${import.meta.env.VITE_HOST}${url}`, {
        method: 'POST',
        ...opt,
        // mode:"no-cors"
      } )
      const json = await result?.json()
      window.localStorage.setItem("token",json.token)
      return result
}

export const get = async (url:string, headers?:{}) => {
  console.log(headers)
    const result =  await fetch(`${import.meta.env.VITE_HOST}${url}`, {
        method: "GET",
        headers: { 
          "Content-Type": "application/json" ,
        ...headers
      }
      } )
    const json = await result?.json()
    window.localStorage.setItem("token",json.token)
    return result
}

export const put = async (url:string,opt:{}) => {
    return await fetch(`${import.meta.env.VITE_HOST}${url}`, {
        method: "PUT",
        ...opt
      } )
}
