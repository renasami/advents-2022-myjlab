export const post = async (url:string,opt:{}) => {
    return await fetch(url, {
        method: 'POST',
        ...opt,
        // mode:"no-cors"
      } )
}

export const get = async (url:string, headers?:{}) => {
  console.log(headers)
    return await fetch(url, {
        method: "GET",
        headers: { 
          "Content-Type": "application/json" ,
        ...headers
      }
      } )
}

export const put = async (url:string,opt:{}) => {
    return await fetch(url, {
        method: "PUT",
        ...opt
      } )
}
