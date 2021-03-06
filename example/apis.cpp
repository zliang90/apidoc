// Copyright 2018 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// @api GET /users/{id}/logs/{lid} 获取用户日志信息
// @apitags t1,t2
// @apiservers admin
// @apideprecated 新版本不再使用
// @apiquery page int default desc
// @apiquery size int default desc
// @apiquery state array.string [normal,lock] 状态码
// @apiparam id int required desc
// @apiparam lid int required desc
//
// @apirequest object * 通用的请求主体
// @apiheader name optional desc
// @apiheader name1 optional desc markdown
// @apiparam count int optional desc
// @apiparam list array.string optional desc
// @apiparam list.id int optional desc
// @apiparam list.name int reqiured desc
// @apiparam list.groups array.string optional.xxxx desc markdown enum:
//  - xx1 xxxxx
//  - xx2 xxxxx
// @apiexample application/json summary
// {
//  count: 5,
//  list: [
//    {id:1, name: 'name1', 'groups': [1,2]},
//    {id:2, name: 'name2', 'groups': [1,2]}
//  ]
// }
//
// @apirequest object application/xml 特定的请求主体
//
// @apiresponse 200 array.object * 通用的返回内容定义
// @apiheader string required desc
// @apiparam id int reqiured desc
// @apiparam name string reqiured desc
// @apiparam group object reqiured desc
// @apiparam group.id int reqiured desc
//
// @apiresponse 404 object application/json 错误的返回内容
// @apiheader string required desc38G
// @apiparam code int reqiured desc
// @apiparam message string reqiured desc
// @apiparam detail array.object reqiured desc
// @apiparam detail.id string reqiured desc
// @apiparam detail.message string reqiured desc
//
// @apiCallback GET 回调内容
// @apirequest object application/xml 特定的请求主体
// @apiresponse 404 object application/json 错误的返回内容
void logs() {}
