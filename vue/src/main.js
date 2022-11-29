import starportLibrary from '@starport/vue'
import { createApp } from 'vue'

import App from './App.vue'
import router from './router'
import store from './store'
import 'mdb-vue-ui-kit/css/mdb.min.css'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'

const app = createApp(App)
app.use(store).use(router).use(starportLibrary).mount('#app')
