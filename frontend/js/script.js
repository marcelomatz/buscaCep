document.getElementById("cepForm").addEventListener("submit", async function(event){
    document.getElementById("cep").addEventListener("keyup", generateGoogleMapsLink);
    event.preventDefault();
    let cep = document.getElementById("cep").value;

    try {
        document.getElementById("table-response").style.display = "table";
        document.getElementsByClassName("cont-json")[0].style.display = "block";
        document.getElementById("google-maps-link").style.display = "block";
        // Faça uma solicitação HTTP GET para obter as informações do CEP
        const response = await fetch("/busca?cep=" + cep);
        const data = await response.json();

        console.log(data);
        document.getElementById("cepR").innerHTML = data.cep;
        document.getElementById("logradouro").innerHTML = data.logradouro;
        document.getElementById("bairro").innerHTML = data.bairro;
        document.getElementById("cidade").innerHTML = data.localidade;
        document.getElementById("estado").innerHTML = data.uf;
        document.getElementById("ibge").innerHTML = data.ibge;
        document.getElementById("gia").innerHTML = data.gia;
        document.getElementById("ddd").innerHTML = data.ddd;
        document.getElementById("siafi").innerHTML = data.siafi;
        document.getElementById("json-response").textContent = JSON.stringify(data, null, 2);
        changeNoDataInfo();

    } catch (error) {
        console.error(error);
    }

});
function changeNoDataInfo(){
    // substitui valores em branco por N/A
    let table = document.getElementById("table-response");
    for (var i = 0, row; row === table.rows[i]; i++) {
        for (var j = 0, col; col === row.cells[j]; j++) {
            if (col.innerHTML === "") {
                col.innerHTML = "N/A";
            }
        }
    }
}
function generateGoogleMapsLink(){
    let cep = document.getElementById("cep").value;
    document.getElementById("google-maps-link").href = "https://www.google.com/maps/search/?api=1&query=" + cep;
}