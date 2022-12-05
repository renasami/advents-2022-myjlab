export const post = async (url:string,opt:{}) => {
    const result =  await fetch(`${import.meta.env.VITE_HOST}${url}`, {
        method: 'POST',
        ...opt,
      } )
      return await result.json()
}

export const get = async (url:string, headers?:{}) => {
    const result =  await fetch(`${import.meta.env.VITE_HOST}${url}`, {
        method: "GET",
        headers: { 
          "Content-Type": "application/json" ,
        ...headers
      }
      } )
    const json = await result?.json()
   
    return json
}

export const put = async (url:string,opt:{}) => {
    const result =  await fetch(`${import.meta.env.VITE_HOST}${url}`, {
        method: "PUT",
        ...opt
      } )

      return await result.json()
}

export const del = async (url:string,headers?:{}) => {
  const result =  await fetch(`${import.meta.env.VITE_HOST}${url}`, {
      method: "DELETE",
      headers: { 
        "Content-Type": "application/json" ,
      ...headers
    }
    })

    return await result.json()
}