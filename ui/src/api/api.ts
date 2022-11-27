export const post = async (url:string,body:{}) => {
    return await fetch(url, {
        method: "POST",
        body: JSON.stringify(body),
        headers: { "Content-Type": "Application-JSON" }
      } )
}

export const get = async (url:string) => {
    return await fetch(url, {
        method: "GET",
        headers: { "Content-Type": "Application-JSON" }
      } )
}

export const put = async (url:string,body:{}) => {
    return await fetch(url, {
        method: "PUT",
        body: JSON.stringify(body),
        headers: { "Content-Type": "Application-JSON" }
      } )
}
