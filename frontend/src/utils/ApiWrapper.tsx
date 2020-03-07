class API {
  masterUrl: string;

  constructor(url: string) {
    this.masterUrl = url + '/v1'

    this.openSolenoid = this.openSolenoid.bind(this)
  }

  async getClusterInfo() {
    const response = await fetch(this.masterUrl + '/cluster_info', {
      method: 'GET',
    })

    return await response.json()
  }

  openSolenoid(uid: string) {
    const queryUrl = 'http://' + this.masterUrl + '/component/' + uid + '/open'
    fetch(queryUrl, {
      method: 'GET'
    }).then(data => {
      console.log('open_solenoid_response', data)
    })
  }

  closeSolenoid(uid: string) {
    const queryUrl = 'http://' + this.masterUrl + '/component/' + uid + '/close'
    fetch(queryUrl, {
      method: 'GET'
    }).then(data => {
      console.log('close_solenoid_response', data)
    })
  }
}

export default API