document.getElementById("cepForm").addEventListener("submit", async function(event){
    document.getElementById("cep").addEventListener("keyup", generateGoogleMapsLink);
    event.preventDefault();
    let cep = document.getElementById("cep").value;

    try {
        const response = await fetch("/busca?cep=" + cep);
        const validations = handleErrors(response);
        if (!validations.ok) throw Error(validations.statusText);
        const data = await response.json();

        displayData(data);
    } catch (error) {
        displayNoData();
    }
});

function handleErrors(response) {
    if (!response.ok) return { statusText: response.statusText, ok: false };
    return response;
}

function displayData(data) {
    // Display the data on the page
    let responseEl = document.getElementById('response');
    // build your HTML string using the data and update responseEl.innerHTML
    responseEl.style.display = "block";
    // similar updates for other elements
}

function displayNoData() {
    let noDataEl = document.getElementById('noData');
    noDataEl.style.display = "block";
}

function generateGoogleMapsLink() {
    let cep = document.getElementById("cep").value;;
    // update Google Maps link with cep
}