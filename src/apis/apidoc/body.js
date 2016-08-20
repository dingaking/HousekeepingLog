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
 * @apiSuccess {String} firstname Firstname of the User.
 * @apiSuccess {String} lastname Lastname of the User.
 *
 * @apiExample {curl} Example usage:
 * curl -H -d '{"userid":"jungtek_whang@hanmail.net","password":"12345678"}' http://localhost:8082/api/authC
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *         "result": "success",
 *         "err_msg":""
 *     }
 *
 * @apiError UserIdNotFound The <code>userid</code> was not found.
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *         "error": "UserNotFound"
 *     }
 *
 */
