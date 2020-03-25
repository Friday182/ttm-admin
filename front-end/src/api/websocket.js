// const WSS_URL = `wss://wss.xxxx.com/ws?appid=xxx`
const WS_URL = `ws://127.0.0.1:8082/api/ws`
// const WS_URL = `ws://www.transfertestmaster.com/hywt/v1/api/ws`
// const WS_URL = 'wss://www.transfertestmaster.com/hywt/ws/'
let Socket = ''
let setIntervalWesocketPush = null

export function createSocket () {
  console.log('start create websocket')

  if (!Socket) {
    Socket = new WebSocket(WS_URL)
    Socket.onopen = onopenWS
    Socket.onmessage = onmessageWS
    Socket.onerror = onerrorWS
    Socket.onclose = oncloseWS
  } else {
    console.log('websocket已连接')
  }
}
// send heartbeat
export function onopenWS () {
  sessionStorage.setItem('ws-status', 'OK')
  // When connected, client get default data at behind for all pages
  // this will reduce the loading time for all the pages
  sendWSPush({
    'sq': 0,
    'type': 'DEFAULT',
    'data': '',
    'name': '',
    'token': ''
  })
  sendPing()
}
export function onerrorWS () {
  clearInterval(setIntervalWesocketPush)
  sessionStorage.setItem('ws-status', 'NOK')
  Socket.close()
  createSocket()
}
// send data out by window event
export function onmessageWS (e) {
  if (e.data instanceof Blob) {
    console.log('RX binary data: ', e.data)
  } else {
    console.log('Rx Text message: ', e.data)
    let tmp = JSON.parse(e.data)
    if (tmp.type === 'CONNECTED' || tmp.type === 'HEARTBEAT') {
      console.log('Rx a Heartbeat')
    } else {
      window.dispatchEvent(new CustomEvent('onmessageWS', {
        detail: {
          data: tmp
        }
      }))
    }
  }
}

export function sendWSPush (obj) {
  if (Socket !== null && Socket.readyState === 3) {
    Socket.close()
    createSocket()
  } else if (Socket.readyState === 1) {
    Socket.send(JSON.stringify(obj))
  } else if (Socket.readyState === 0) {
    setTimeout(() => {
      Socket.send(JSON.stringify(obj))
    }, 3000)
  }
}
export function oncloseWS () {
  clearInterval(setIntervalWesocketPush)
  sessionStorage.setItem('ws-status', 'NOK')
  console.log('websocket已断开')
}
export function sendPing () {
  setIntervalWesocketPush = setInterval(() => {
    Socket.send(JSON.stringify({
      'sq': 0,
      'type': 'HEARTBEAT',
      'data': 'ping',
      'from': 'ws-id',
      'token': 'ws-id'
    }))
    console.log('sending ping ...')
  }, 30000)
}
