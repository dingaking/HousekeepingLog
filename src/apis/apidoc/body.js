// body.json

/**
 * @api {post} /api/authC authC
 * @apiVersion 0.1.0
 * @apiName AuthCreate
 * @apiDescription 회원가입
 * @apiGroup Auth
 *
 * @apiParam {String} userid userid(email)
 * @apiParam {String} password 비밀번호(8~16자리)
 *
 * @apiSuccess {String} result 요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 *
 * @apiExample {curl} Example usage:
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"userid":"admin@localhost.com", "password":"12345678"}' \
http://localhost:8082/api/authC
 *
 */

/**
 * @api {post} /api/authR authR
 * @apiVersion 0.1.0
 * @apiName AuthRead
 * @apiDescription 회원정보 조회
 * @apiGroup Auth
 *
 * @apiExample {curl} Example : (action=1)
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"action":"1", "userid":"admin@localhost.com", "password":"12345678"}' \
http://localhost:8082/api/authR
 *
 * @apiExample {curl} Example : (action=2)
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"action":"2", "access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/authR
 *
 * @apiParam (AuthRead_Action_1) {String} action 1 : login
 * @apiParam (AuthRead_Action_1) {String} userid userid(email)
 * @apiParam (AuthRead_Action_1) {String} password 비밀번호(8~16자리)
 *
 * @apiParam (AuthRead_Action_2) {String} action 2 : 회원정보 조회
 * @apiParam (AuthRead_Action_2) {String} access_token 인증키
 *
 * @apiSuccess (AuthRead_Success_Action_1) {String} result "success" or "fail"
 * @apiSuccess (AuthRead_Success_Action_1) {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess (AuthRead_Success_Action_1) {String} access_token result가 "success"인 경우
 *
 * @apiSuccess (AuthRead_Success_Action_2) {String} result "success" or "fail"
 * @apiSuccess (AuthRead_Success_Action_2) {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess (AuthRead_Success_Action_2) {String} dispname 대화명
 * @apiSuccess (AuthRead_Success_Action_2) {String} intro 자기소개
 * @apiSuccess (AuthRead_Success_Action_2) {String} profile 프로필 사진
 * @apiSuccess (AuthRead_Success_Action_2) {String} create_datetime 가입일시
 * @apiSuccess (AuthRead_Success_Action_2) {String} phone_number 전화번호
 * @apiSuccess (AuthRead_Success_Action_2) {String} public 그룹에 공개여부
 * @apiSuccess (AuthRead_Success_Action_2) {String} access_token 접속키
 *
 */
/**
 * @apiDefine AuthRead_Action_1
 * Parameter(action=1)
 */
/**
 * @apiDefine AuthRead_Action_2
 * Parameter(action=2)
 */
/**
 * @apiDefine AuthRead_Success_Action_1
 * Response(action=1)
 */
/**
 * @apiDefine AuthRead_Success_Action_2
 * Response(action=2)
 */

/**
 * @api {post} /api/authU authU
 * @apiVersion 0.1.0
 * @apiName AuthUpdate
 * @apiDescription 사용자 정보 갱신
 * @apiGroup Auth
 *
 * @apiExample {curl} Example : (action=1)
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"action":"1", "old_password":"11112222", \
"new_password":"22221111", "access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/authU
 *
 * @apiExample {curl} Example : (action=2)
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"action":"2", "dispname":"my display name", \
"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/authU
 *
 * @apiParam (Param_Action_1) {String} action 1 : 비밀번호 변경
 * @apiParam (Param_Action_1) {String} userid userid(email)
 * @apiParam (Param_Action_1) {String} old_password 이전 비밀번호
 * @apiParam (Param_Action_1) {String} new_password 새 비밀번호
 *
 * @apiParam (Param_Action_2) {String} action 2 : 회원정보 변경
 * @apiParam (Param_Action_2) {String} dispname 대화명
 * @apiParam (Param_Action_2) {String} intro 자기소개
 * @apiParam (Param_Action_2) {String} profile 프로필 사진
 * @apiParam (Param_Action_2) {String} phone_number 전화번호
 * @apiParam (Param_Action_2) {String} public 그룹에 공개여부
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 *
 */
/**
 * @apiDefine Param_Action_1
 * Parameter(action=1)
 */
/**
 * @apiDefine Param_Action_2
 * Parameter(action=2)
 */
/**
 * @apiDefine Response_Action_1
 * Response(action=1)
 */
/**
 * @apiDefine Response_Action_2
 * Response(action=2)
 */

/**
 * @api {post} /api/authD authD
 * @apiVersion 0.1.0
 * @apiName AuthDelete
 * @apiDescription 회원탈퇴
 * @apiGroup Auth
 *
 * @apiParam {String} userid userid
 * @apiParam {String} password 비밀번호
 * @apiParam {String} access_token 인증키
 *
 * @apiSuccess {String} result 요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 *
 * @apiExample {curl} Example usage:
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"userid":"admin@localhost.com", "password":"12345678", \
"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/authD
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
 */


/**
 * @api {post} /api/admin/userU userU
 * @apiVersion 0.1.0
 * @apiName AdminUserU
 * @apiDescription 가계부 관리자 정보 수정
 * @apiGroup Admin_User
 *
 * @apiExample {curl} Example : (action=1)
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"action":"1", "userid":"admin", "old_password":"12345678", "new_password":"12345678"}' \
http://localhost:8082/api/admin/userU
 *
 * @apiParam (Param_Action_1) {String} action 1 : admin init password change
 * @apiParam (Param_Action_1) {String} userid userid
 * @apiParam (Param_Action_1) {String} old_password 이전 비밀번호
 * @apiParam (Param_Action_1) {String} new_password 새 비밀번호
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우
 * @apiSuccess {String} access_token result가 "success"인 경우
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
 * @apiExample {curl} Example usage:
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"action":"1", "access_token":"access_token_value"}' \
http://localhost:8082/api/admin/userL
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
 */
/**
 * @apiDefine AdminUserL_Action_1 Parameter(SubAction 1)
 */
/**
 * @apiDefine AdminUserL_Action_2 Parameter(SubAction 2)
 */
