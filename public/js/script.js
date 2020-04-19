; let tooltipElem;

async function addTabForm() {
    ts('#addTabsForm').modal('show');
}
async function hideAddTab() {
    ts('#addTabsForm').modal('hide');
}

async function bindval(url = '') {
    let label_url_addtab = document.getElementById("label-url-addtab")
        , input_url_addtab = document.getElementById("input-url-addtab")
        , loading_form_addtab = document.getElementById("loading-form-addtab")
        , pushbtn_url_addtab = document.getElementById("pushbtn-url-addtab")
        , error_url_addtab = document.getElementById("error-url-addtab")
        , x = isValidURLWithProtocol(url);
    if (x) {
        loading_form_addtab.className = "ts loading segment";
        label_url_addtab.className = "ts positive label";
        input_url_addtab.className = "ts basic labeled fluid success input";
        loading_form_addtab.hidden = false;
        pushbtn_url_addtab.disabled = false;
        error_url_addtab.innerHTML = "";
        datas = await postData('/getSiteData', { site: url }).then((datas) => {
            console.log(datas);
            if (datas.error != "") {
                error_url_addtab_error(datas.error);
            };
            loading_form_addtab.className = "ts segment";
            template = generate_inner(datas);
            loading_form_addtab.innerHTML = template;
        });

    } else if (!x || error != '') {
        if (isValidURLWoutProtocol(url)) {
            console.log(url);
            url = "https://" + url;
            console.log(url);
            bindval(url);
            return;
        }
        loading_form_addtab.hidden = true;
        label_url_addtab.className = "ts negative label";
        input_url_addtab.className = "ts basic labeled fluid error input";
        error_url_addtab.innerHTML = "Invalid URL";
    };
}

function error_url_addtab_error(error = '') {
    let label_url_addtab = document.getElementById("label-url-addtab")
        , input_url_addtab = document.getElementById("input-url-addtab")
        , loading_form_addtab = document.getElementById("loading-form-addtab")
        , error_url_addtab = document.getElementById("error-url-addtab");

    loading_form_addtab.hidden = true;
    label_url_addtab.className = "ts negative label";
    input_url_addtab.className = "ts basic labeled fluid error input";
    error_url_addtab.innerHTML = error;
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
//google.com
function generate_inner(datas) {
    let template;
    newline = `<br><br>`
    div_title = `<div class="ts labeled fluid input"><div class="ts basic label">Title:</div> <input type="text" value="` + datas.title + `"></div>`;
    if (datas.url == "") datas.url = "Add URL..."
    div_url = `<div class="ts labeled fluid input"><div class="ts basic label">URL:</div> <input type="text" value="` + datas.url + `"></div>`;
    if (datas.description == "") datas.description = "Add description...";
    div_description = `<div class="ts basic label">Description:</div><div class="ts fluid input"><textarea type="text">` + datas.description + `</textarea></div>`;
    if (datas.icon != "") {
        div_description = `<div class="ts basic image label"><img src="` + datas.icon + `">Description:</div><div class="ts fluid input"><textarea type="text">` + datas.description + `</textarea></div>`;
    }
    template = div_title + newline + div_url + newline + div_description + newline;
    return template;
}
;