/* eslint-env serviceworker */
/* global fetch location */

const cacheName = 'cache-v1'

const resourcesToCache = [
  '/_offline',
  '/main.css',
  '/main.js',
  '/manifest.json'
]

self.addEventListener('install', function (event) {
  event.waitUntil(
    caches.open(cacheName)
      .then(cache => {
        return cache.addAll(resourcesToCache)
      })
      .then(() => {
        return self.skipWaiting()
      })
      .catch(err => {
        console.log('could not setup cache', err)
        throw err
      })
  )
})

// TODO expiration date
// https://www.youtube.com/watch?v=ksXwaWHCW6k
// https://github.com/deanhume/Service-Worker-Offline/blob/master/service-worker.js
self.addEventListener('fetch', event => {
  const req = event.request
  const strippedURL = req.url.indexOf(location.origin) === 0 ? req.url.slice(location.origin.length) : req.cacheurl
  let response
  console.log(`fetch to ${strippedURL}`)
  if (resourcesToCache.includes(strippedURL)) {
    response = caches.open(cacheName)
      .then(cache => cache.match(strippedURL))
  } else {
    if (isNavigationRequest(req)) {
      response = networkOrOffline(req)
    } else {
      response = fetch(req)
    }
  }
  event.respondWith(response)
})

function isNavigationRequest (request) {
  return request.method === 'GET' && request.headers.get('accept').includes('text/html')
}

function networkOrOffline (req) {
  return fetch(req)
    .catch(() => caches.open(cacheName).then(cache => cache.match('/_offline')))
}
