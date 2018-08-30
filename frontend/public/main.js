/* eslint-env browser */

if ('serviceWorker' in navigator) {
  window.addEventListener('load', () => {
    navigator.serviceWorker.register('/worker.js', { scope: '/' })
      .catch(function (error) {
        console.log('Service worker registration failed:', error)
      })
  })
}

class Encryption {
  randomInit () {
    return this.init(crypto.getRandomValues(new Uint8Array(32)), crypto.getRandomValues(new Uint8Array(16)))
  }
  async init (key, iv) {
    if (!key) {
      throw new Error('Encryption.init: missing param key')
    }
    if (!iv) {
      throw new Error('Encryption.init: missing param iv')
    }
    this.key = key
    this.iv = iv
    this.algorithm = {
      name: 'AES-GCM',
      iv: this.iv
    }
    this.cryptoKey = await crypto.subtle.importKey('raw', key, this.algorithm, true, ['encrypt', 'decrypt'])
    return this
  }

  async encrypt (secret) {
    const cipher = await crypto.subtle.encrypt(this.algorithm, this.cryptoKey, new TextEncoder().encode(secret))
    return cipher
  }

  async decrypt (cipĥer) {
    const decodedSecret = await crypto.subtle.decrypt(this.algorithm, this.cryptoKey, cipĥer)
    return new TextDecoder().decode(decodedSecret)
  }

  async exportKey () {
    return `${arrayToBase64url(this.key)}/${arrayToBase64url(this.iv)}`
  }
}

function arrayToBase64url (buffer) {
  return window.btoa(String.fromCharCode.apply(null, new Uint8Array(buffer)))
    .replace(/=/g, '')
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
}

function base64urlToArray (base64urlString) {
  let str = window.atob(base64urlString.replace(/-/g, '+').replace(/_/g, '/'))
  let array = new Uint8Array(str.length)
  for (let i = 0; i < str.length; i++) {
    array[i] = str.charCodeAt(i)
  }
  return array
}
