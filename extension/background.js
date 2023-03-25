chrome.tabs.onCreated.addListener(function(tab) {
    console.log(tab.url)
    let sub = String(tab.url).split('/')
    fetch('http://localhost:3000'+'/'+String(sub[2]), {
        mode: "no-cors",
        headers: {
           'Accept': 'application/json'
        }
     })
        .then(response => response.text())
        .then(text => console.log(text))
})

chrome.tabs.onUpdated.addListener(function(tabId, changeInfo, tab) {
    if (changeInfo.status === "complete") {
        console.log(tab.url)
        let sub = String(tab.url).split('/')
        fetch('http://localhost:3000'+'/'+String(sub[2]), {
         mode: "no-cors",
            headers: {
               'Accept': 'application/json'
    
            }
         })
            .then(response => response.text())
            .then(text => console.log(text))
    }
})

chrome.tabs.query({}, function(tabs) {
    tabs.forEach(function(tab) {
        console.log(tab.url);
        let sub = String(tab.url).split('/')
        fetch('http://localhost:3000'+'/'+String(sub[2]), {
            mode: "no-cors",
            headers: {
               'Accept': 'application/json'
            }
         })
            .then(response => response.text())
            .then(text => console.log(text))
    })
});