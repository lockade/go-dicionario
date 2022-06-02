import Home from './views/Home.vue'
import PalavraView from './views/Palavra/View.vue'
import PalavraForm from './views/Palavra/Form.vue'
import NotFound from './views/NotFound.vue'

/** @type {import('vue-router').RouterOptions['routes']} */
export const routes = [
  { path: '/', component: Home, meta: { title: 'Home' } },
  {
    path: '/palavra/:method/:id?',
    meta: { title: 'Palavra' },
    component: PalavraForm,
  },
  { path: '/palavra', meta: { title: 'Palavra' }, component: PalavraView, },
  { path: '/:path(.*)', component: NotFound },
]
