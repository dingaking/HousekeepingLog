// body.json

/**
 * @api {post} /api/authC authC
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
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"userid":"admin", "password":"12345678"}' \
http://localhost:8082/api/authC
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
 * @apiGroup Admin_User
 *
 * @apiParam {String} userid userid
 * @apiParam {String} password 비밀번호
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우
 * @apiSuccess {String} access_token result가 "success"인 경우
 *
 * @apiExample {curl} Example usage:
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"userid":"admin", "password":"admin"}' \
http://localhost:8082/api/admin/userR
 * 
 * @apiSuccessExample Success-Response:
 * HTTP/1.1 200 OK
 * {
 *     "result":"success",
 *     "access_token":"111111"
 * }
 */


/**
 * @api {post} /api/admin/userU userU
 * @apiVersion 0.1.0
 * @apiName AdminUserU
 * @apiDescription 가계부 관리자 정보 수정
 * @apiGroup Admin_User
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


/**
 * @api {post} /api/admin/userL userL
 * @apiVersion 0.1.0
 * @apiName AdminUserL
 * @apiDescription 가계부 관리자 회원목록
 * @apiGroup Admin_User
 *
 * @apiParam (AdminUserL_Action_1) {String} action 1:회원목록 조회
 * @apiParam (AdminUserL_Action_1) {String} access_token 인증키
 * @apiParam (AdminUserL_Action_1) {String} page 목록 페이지
 *
 * @apiParam (AdminUserL_Action_2) {String} action s:회원목록 검색
 * @apiParam (AdminUserL_Action_2) {String} access_token 인증키
 * @apiParam (AdminUserL_Action_2) {String} search_id 검색어(userid)
 * @apiParam (AdminUserL_Action_2) {String} search_name 검색어(username)
 * @apiParam (AdminUserL_Action_2) {String} page 검색 결과 페이지
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우
 * @apiSuccess {String} data result가 "success"인 경우
 *
 * @apiExample {curl} Example usage:
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"action":"1", "access_token":"access_token_value"}' \
http://localhost:8082/api/admin/userL
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
 * @apiDefine AdminUserL_Action_1 Parameter(SubAction 1)
 */
/**
 * @apiDefine AdminUserL_Action_2 Parameter(SubAction 2)
 */
