async function addTabForm() {
    ts('#addTabsForm').modal('show');
}
async function hideAddTab() {
    ts('#addTabsForm').modal('hide');
}

async function bindval(url = '') {
    if (isValidURLWithProtocol(url)) {
        document.getElementById("label-url-popup").className = "ts positive label";
        document.getElementById("input-url-popup").className = "ts basic labeled success input";
        document.getElementById("error-url-popup").hidden = true;
        document.getElementById("loading-form-popup").hidden = false;
        datas = await postData('/getSiteData', { site: url }).then((datas) => {
            document.getElementById("loading-form-popup").hidden = true;
        });
    } else {
        if (isValidURLWoutProtocol(url)) {
            console.log(url);
            url = "https://" + url;
            console.log(url);
            bindval(url);
            return;
        }
        document.getElementById("label-url-popup").className = "ts negative label";
        document.getElementById("input-url-popup").className = "ts basic labeled error input";
        document.getElementById("error-url-popup").hidden = false;
        document.getElementById("error-url-popup").innerHTML = "Invalid URL";
    };
}

async function postData(url = '', data = { '': `` }) {
    const response = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    });
    console.log(JSON.stringify(data))
    return await response.json();
}

function isValidURLWithProtocol(url) {
    var objRE = /^(?:http(s)?:\/\/)+[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$/;
    return objRE.test(url);
}

function isValidURLWoutProtocol(url) {
    var objRE = /^(?:http(s)+:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$/;
    return objRE.test(url);
}