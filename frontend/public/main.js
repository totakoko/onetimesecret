if ('serviceWorker' in navigator) {
  window.addEventListener('load', () => {
    navigator.serviceWorker.register('/worker.js', { scope: '/' })
      .catch(function (error) {
        console.log('Service worker registration failed:', error)
      })
  })
}
