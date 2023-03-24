chrome.tabs.onCreated.addListener(function(tab) {
    console.log(tab.url)
})

chrome.tabs.onUpdated.addListener(function(tabId, changeInfo, tab) {
    if (changeInfo.status === "complete") {
        console.log(tab.url)
    }
})

chrome.tabs.query({}, function(tabs) {
    tabs.forEach(function(tab) {
        console.log(tab.url);
    })
});