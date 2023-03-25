async function checkWebsite(website) {
  try {
    const response = await fetch('https://secureex.azurewebsites.net/'+website, {
      method: 'GET',
      mode: 'no-cors',
      headers: {
        accept: 'application/json',
      },
    });

    if (!response.ok) {
      throw new Error(`Error! status: ${response.status}`);
    }

    const result = await response.json();
    console.log(result);
    return result;
  } catch (err) {
    console.log(err);
  }
}

chrome.tabs.onCreated.addListener(function(tab) {
    console.log(tab.url)
    let sub = String(tab.url).split('/')
    checkWebsite(String(sub[2]));
})

chrome.tabs.onUpdated.addListener(function(tabId, changeInfo, tab) {
    if (changeInfo.status === "complete") {
        console.log(tab.url)
        let sub = String(tab.url).split('/')
        checkWebsite(String(sub[2]));
        
    }
})

chrome.tabs.query({}, function(tabs) {
    tabs.forEach(function(tab) {
        console.log(tab.url);
        let sub = String(tab.url).split('/')
        checkWebsite(String(sub[2]));
    })
});

chrome.tabs.query({active: true, currentWindow: true}, function(tabs) {
  tabs.forEach(function(tab) {
    console.log("this is the current tab "+tab.url)
    let sub = String(tab.url).split('/')
    document.getElementById("website").innerHTML = String(sub[2])

  })
});