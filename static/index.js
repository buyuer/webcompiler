let CompileInfo = {
    Compiler:null,
    Code:null,
    Args:null,
    In:null,
}

let ExecResult = {
    Status:null,
    Err:null,
    Out:null,
}

let textarea_code = document.getElementById("textarea_code")
let textarea_in = document.getElementById("textarea_in")
let textarea_out = document.getElementById("textarea_out")
let request = new XMLHttpRequest()
let result
function button_onclick() {
    textarea_out.value = ""

    let request = new XMLHttpRequest()
    request.onreadystatechange = ()=>{
        if(request.readyState === 4){
            if(request.status === 200){
                result = JSON.parse(request.responseText)
                textarea_out.value = result.Out + result.Err
            }
        }
    }

    CompileInfo.Compiler = "gcc"
    CompileInfo.Args = ""
    CompileInfo.code = textarea_code.value
    CompileInfo.In = textarea_in.value

    body = JSON.stringify(CompileInfo)

    request.open('post', '/compile')
    request.send(body)
}