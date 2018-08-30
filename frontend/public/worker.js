/* eslint-env browser, serviceworker */

const cacheName = 'cache-v1'

const offlinePageURL = '/_offline'

self.addEventListener('install', async function (event) {
  event.waitUntil(
    (async () => {
      const cache = await caches.open(cacheName)
      await cache.add(offlinePageURL)
      await self.skipWaiting()
    })()
  )
})

self.addEventListener('fetch', event => {
  const request = event.request
  event.respondWith(isNavigationRequest(request) ? networkOrOffline(request) : fetch(request))
})

function isNavigationRequest (request) {
  return request.method === 'GET' && request.headers.get('accept').includes('text/html')
}

async function networkOrOffline (request) {
  try {
    return await fetch(request)
  } catch (err) {
    const cache = await caches.open(cacheName)
    return cache.match('/_offline')
  }
}
