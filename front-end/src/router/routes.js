
const routes = [
  {
    path: '/',
    component: () => import('layouts/MyLayout.vue'),
    children: [
      { path: '', component: () => import('pages/Index.vue') },
      { path: 'Dashboard', component: () => import('pages/Dashboard.vue') },
      { path: 'StaffAdmin', component: () => import('pages/StaffAdmin.vue') },
      { path: 'OperatorAdmin', component: () => import('pages/OperatorAdmin.vue') },
      { path: 'MentorAdmin', component: () => import('pages/MentorAdmin.vue') },
      { path: 'StudentAdmin', component: () => import('pages/StudentAdmin.vue') },
      { path: 'QuizAdmin', component: () => import('pages/QuizAdmin.vue') },
      { path: '*', component: () => import('pages/Error404.vue') }
    ]
  }
]

// Always leave this as last one
if (process.env.MODE !== 'ssr') {
  routes.push({
    path: '*',
    component: () => import('pages/Error404.vue')
  })
}

export default routes
