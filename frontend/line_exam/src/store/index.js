import { createStore } from 'vuex'

export default createStore({
    state: {
        fileName: '',
        uploadProgress: 0,
    },
    getters: {
    },
    mutations: {
        updateProgress(state, value) {
            state.uploadProgress = value
        },
        updateFileName(state, value) {
            state.fileName = value
        }
    },
    actions: {
        updateProgressAction(context, value) {
            context.commit('updateProgress', value)
        },
        updateFileNameAction(context, value) {
            context.commit('updateFileName', value)
        }
    },
    modules: {
    }
})
