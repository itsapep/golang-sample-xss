async function getUsers(){
    let url='http://localhost:8888/api/user';
    try {
        let res = await fetch(url);
        console.log(res)
        return await res.json();
    } catch (error) {
        console.log(error);
    }
}

async function name(params) {
    
}