import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/v1/users/list',
    method: 'get',
    params
  })
}
