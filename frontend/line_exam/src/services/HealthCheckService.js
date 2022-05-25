import axios from 'axios';
import store from '../store'

export default {
  healthCheckUpload(file) {
    console.log('uploading');

    store.dispatch("updateFileNameAction", file.name)

    let formData = new FormData()

    formData.append('file', file);

    var apiURL = "http://localhost:8080/api/v1/health-check"

    return axios.post(apiURL,
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data'
        },
        onUploadProgress: function (progressEvent) {
          let progress = parseInt(Math.round((progressEvent.loaded / progressEvent.total) * 100));
          store.dispatch("updateProgressAction", progress)
        }.bind(this)
      }
    ).then(function (res) {
      return res
    })
      .catch(function (err) {
        console.log(err);
        console.log('FAILURE!!');
      });
  }
}