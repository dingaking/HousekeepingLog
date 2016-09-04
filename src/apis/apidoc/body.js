// body.json

/**
 * @api {post} /api/authC AuthC
 * @apiVersion 0.1.0
 * @apiName AuthCreate
 * @apiDescription 가계부 서비스 회원가입
 * @apiGroup Auth
 *
 * @apiParam {String} userid userid(email)
 * @apiParam {String} password 비밀번호(8~16자리)
 *
 * @apiSuccess {String} result 가입요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지가 "success"인 경우 ""
 *
 * @apiExample {curl} Example usage:
 * curl -H -d '{"userid":"test1234@hanmail.net","password":"12345678"}' http://localhost:8082/api/authC
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *         "result": "success",
 *         "err_msg":""
 *     }
 *
 */
