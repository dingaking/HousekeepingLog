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
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" -d '{"userid":"admin", "password":"12345678"}' http://localhost:8082/api/admin/userU
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *         "result": "success",
 *         "err_msg":""
 *     }
 *
 */

/**
 * @api {post} /api/admin/userR userR
 * @apiVersion 0.1.0
 * @apiName AdminUserR
 * @apiDescription 가계부 관리자 로그인
 * @apiGroup AdminUser
 *
 * @apiParam {String} userid userid
 * @apiParam {String} password 비밀번호
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우
 * @apiSuccess {String} access_token result가 "success"인 경우
 *
 * @apiExample {curl} Example usage:
 * curl -H -d '{"userid":"admin", "password":"12345678"}' http://localhost:8082/api/admin/userR
 * 
 * @apiSuccessExample Success-Response:
 * HTTP/1.1 200 OK
 * {
 *     "result":"success",
 *     "access_token":"111111"
 * }
 */


/**
 * @api {post} /api/admin/UserU UserU
 * @apiVersion 0.1.0
 * @apiName AdminUserU
 * @apiDescription 가계부 관리자 정보 수정
 * @apiGroup AdminUser
 *
 * @apiParam (AdminUserU_Action_1) {String} action 1:admin init password change
 * @apiParam (AdminUserU_Action_1) {String} userid userid
 * @apiParam (AdminUserU_Action_1) {String} old_password 이전 비밀번호
 * @apiParam (AdminUserU_Action_1) {String} new_password 새 비밀번호
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우
 * @apiSuccess {String} access_token result가 "success"인 경우
 *
 * @apiExample {curl} Example usage:
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
 -d '{"action":"1", "userid":"admin", "old_password":"12345678", "new_password":"12345678"}' \
http://localhost:8082/api/admin/userU
 * 
 * @apiSuccessExample Success-Response[SubAction:1]:
 * HTTP/1.1 200 OK
 * {
 *     "result":"success",
 *     "access_token":"111111"
 * }
 * 
 * @apiSuccessExample Success-Response[SubAction:2]:
 * HTTP/1.1 200 OK
 * {
 *     "result":"success",
 *     "access_token":"111111"
 * }
 * 
 * 
 */
/**
 * @apiDefine AdminUserU_Action_1
 * Parameter(SubAction 1)
 */
