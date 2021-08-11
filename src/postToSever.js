import axios from "axios"
export default async function postToSever (expression) {
    // let output
    console.log(expression);
    return await axios.post('/evaluate', {
        expression
    })
        .then((res) => {return (res.data)})
}