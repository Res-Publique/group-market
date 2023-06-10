data = JSON.stringify({
    'initData': Telegram.WebApp.initData,
    'initDataUnsafe': Telegram.WebApp.initDataUnsafe
})

document.getElementById('data').textContent = data