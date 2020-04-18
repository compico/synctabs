async function addTabForm() {
    ts('#addTabsForm').modal('show');
}
async function hideAddTab() {
    ts('#addTabsForm').modal('hide');
}

async function test(url) {
    x = await bindval(url)
}

async function bindval(url = '') {

    datas = await postData('/getSiteData', { site: url }).then((datas) => {
        console.log(datas)
    });

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

function isValidUrl(url) {
    var objRE = /(^https+:\/\/)?[a-z0-9~_\-\.]+\.[a-z]{2,9}(\/|:|\?[!-~]*)?$/i;
    return objRE.test(url);
}
function haveProtocol(url) {

}