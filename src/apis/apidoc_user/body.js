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
 * @apiExample {curl} Example :
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
 * @apiParam (Param_Action_1) {String} action 1 : login
 * @apiParam (Param_Action_1) {String} userid userid(email)
 * @apiParam (Param_Action_1) {String} password 비밀번호(8~16자리)
 *
 * @apiParam (Param_Action_2) {String} action 2 : 회원정보 조회
 * @apiParam (Param_Action_2) {String} access_token 인증키
 *
 * @apiSuccess (Response_Action_1) {String} result "success" or "fail"
 * @apiSuccess (Response_Action_1) {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess (Response_Action_1) {String} access_token result가 "success"인 경우
 *
 * @apiSuccess (Response_Action_2) {String} result "success" or "fail"
 * @apiSuccess (Response_Action_2) {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess (Response_Action_2) {String} dispname 대화명
 * @apiSuccess (Response_Action_2) {String} intro 자기소개
 * @apiSuccess (Response_Action_2) {String} profile 프로필 사진
 * @apiSuccess (Response_Action_2) {String} create_datetime 가입일시
 * @apiSuccess (Response_Action_2) {String} phone_number 전화번호
 * @apiSuccess (Response_Action_2) {String} public 그룹에 공개여부
 * @apiSuccess (Response_Action_2) {String} access_token 접속키
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
 * @apiParam (Param_Action_1) {String} access_token 인증키
 *
 * @apiParam (Param_Action_2) {String} action 2 : 회원정보 변경
 * @apiParam (Param_Action_2) {String} dispname 대화명
 * @apiParam (Param_Action_2) {String} intro 자기소개
 * @apiParam (Param_Action_2) {String} profile 프로필 사진
 * @apiParam (Param_Action_2) {String} phone_number 전화번호
 * @apiParam (Param_Action_2) {String} public 그룹에 공개여부
 * @apiParam (Param_Action_2) {String} access_token 인증키
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
 * @apiDefine Param_Action_3
 * Parameter(action=3)
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
 * @apiDefine Response_Action_3
 * Response(action=3)
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
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"userid":"admin@localhost.com", "password":"12345678", \
"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/authD
 *
 */

/**
 * @api {post} /api/deviceC deviceC
 * @apiVersion 0.1.0
 * @apiName DeviceCreate
 * @apiDescription 단말기 추가
 * @apiGroup Device
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"terminal_id":"dk8s88a0ffd8a", "device_type":"i", \
"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/deviceC
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} terminal_id UDID
 * @apiParam {String} device_type i:iOS, a:Android, w:Web
 *
 * @apiSuccess {String} result 요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {String} terminalno 터미널no
 *
 */

/**
 * @api {post} /api/deviceR deviceR
 * @apiVersion 0.1.0
 * @apiName DeviceRead
 * @apiDescription 단말기 정보 조회
 * @apiGroup Device
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"terminalno":"34", \
"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/deviceR
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} terminalno 터미널no
 *
 * @apiSuccess {String} result 요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {String} terminal_id UDID(web null)
 * @apiSuccess {String} device_type i:iOS, a:Android, w:Web
 * @apiSuccess {String} device_token push키
 * @apiSuccess {String} create_datetime 등록일시
 * @apiSuccess {String} notification 1:ON, 0:OFF
 * @apiSuccess {String} state 1:ON(사용),0:OFF(삭제)
 * @apiSuccess {String} terminal_name 단말기 별칭
 * @apiSuccess {String} last_access 최근접속 일시
 */

/**
 * @api {post} /api/deviceU deviceU
 * @apiVersion 0.1.0
 * @apiName DeviceUpdate
 * @apiDescription 단말기 정보 갱신
 * @apiGroup Device
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", \
"terminalno":"36", "terminal_name":"my iphone"}' \
http://localhost:8082/api/deviceU
 *
 * @apiParam (Param_Action_1) {String} access_token 인증키
 * @apiParam (Param_Action_1) {String} terminalno 터미널no
 * @apiParam (Param_Action_1) {String} [device_token] push키
 * @apiParam (Param_Action_1) {String} [terminal_name] 단말기 별칭
 * @apiParam (Param_Action_1) {String} [notification] 1:ON, 0:OFF
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 */

/**
 * @api {post} /api/deviceD deviceD
 * @apiVersion 0.1.0
 * @apiName DeviceDelete
 * @apiDescription 단말기 삭제
 * @apiGroup Device
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", \
"terminalno":"36"}' \
http://localhost:8082/api/deviceD
 *
 * @apiParam (Param_Action_1) {String} access_token 인증키
 * @apiParam (Param_Action_1) {String} terminalno 터미널no
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 *
 */

/**
 * @api {post} /api/deviceL deviceL
 * @apiVersion 0.1.0
 * @apiName DeviceList
 * @apiDescription 단말기 목록
 * @apiGroup Device
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/deviceL
 *
 * @apiParam {String} access_token 인증키
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {Object[]} data 단말기 목록
 * @apiSuccess {String} data.terminal_id UDID(web null)
 * @apiSuccess {String} data.device_type i:iOS, a:Android, w:Web
 * @apiSuccess {String} data.device_token push키
 * @apiSuccess {String} data.create_datetime 등록일시
 * @apiSuccess {String} data.notification 1:ON, 0:OFF
 * @apiSuccess {String} data.state 1:ON(사용),0:OFF(삭제)
 * @apiSuccess {String} data.terminal_name 단말기 별칭
 * @apiSuccess {String} data.last_access 최근접속 일시
 *
 */

/**
 * @api {post} /api/deviceS deviceS
 * @apiVersion 0.1.0
 * @apiName DeviceSearch
 * @apiDescription 단말기 검색
 * @apiGroup Device
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", "terminal_name":"my"}' \
http://localhost:8082/api/deviceS
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} terminal_name 검색어
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {Object[]} data 단말기 목록
 * @apiSuccess {String} data.terminal_id UDID(web null)
 * @apiSuccess {String} data.device_type i:iOS, a:Android, w:Web
 * @apiSuccess {String} data.device_token push키
 * @apiSuccess {String} data.create_datetime 등록일시
 * @apiSuccess {String} data.notification 1:ON, 0:OFF
 * @apiSuccess {String} data.state 1:ON(사용),0:OFF(삭제)
 * @apiSuccess {String} data.terminal_name 단말기 별칭
 * @apiSuccess {String} data.last_access 최근접속 일시
 *
 */

/**
 * @api {post} /api/capitalC capitalC
 * @apiVersion 0.1.0
 * @apiName CapitalCreate
 * @apiDescription 재산항목 추가
 * @apiGroup Capital
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"capital_name":"월급통장", \
"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/capitalC
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} capital_name 재산항목명
 *
 * @apiSuccess {String} result 요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 *
 */

/**
 * @api {post} /api/capitalR capitalR
 * @apiVersion 0.1.0
 * @apiName CapitalRead
 * @apiDescription 재산항목 정보 조회
 * @apiGroup Capital
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"capitalno":"34", \
"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/capitalR
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} capitalno 재산항목 no
 *
 * @apiSuccess {String} result 요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {String} capitalno 재산항목 no
 * @apiSuccess {String} capital_name 재산항목명
 */

/**
 * @api {post} /api/capitalU capitalU
 * @apiVersion 0.1.0
 * @apiName CapitalUpdate
 * @apiDescription 재산항목 정보 갱신
 * @apiGroup Capital
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", \
"capitalno":"36", "capital_name":"my capital"}' \
http://localhost:8082/api/deviceU
 *
 * @apiParam (Param_Action_1) {String} access_token 인증키
 * @apiParam (Param_Action_1) {String} capitalno 재산항목 no
 * @apiParam (Param_Action_1) {String} capital_name 재산항목명
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 */

/**
 * @api {post} /api/capitalD capitalD
 * @apiVersion 0.1.0
 * @apiName CapitalDelete
 * @apiDescription 재산항목 삭제
 * @apiGroup Capital
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", \
"capitalno":"36"}' \
http://localhost:8082/api/capitalD
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} capitalno 재산항목 no
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 *
 */

/**
 * @api {post} /api/capitalL capitalL
 * @apiVersion 0.1.0
 * @apiName CapitalList
 * @apiDescription 단말기 목록
 * @apiGroup Capital
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/capitalL
 *
 * @apiParam {String} access_token 인증키
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {Object[]} data 재산항목 목록
 * @apiSuccess {String} data.capitalno 재산항목 no
 * @apiSuccess {String} data.capital_name 재산항목명
 *
 */

/**
 * @api {post} /api/capitalS capitalS
 * @apiVersion 0.1.0
 * @apiName CapitalSearch
 * @apiDescription 재산항목 검색
 * @apiGroup Capital
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", "capital_name":"my"}' \
http://localhost:8082/api/capitalS
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} capital_name 검색어
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {Object[]} data 재산항목 목록
 * @apiSuccess {String} data.capitalno 재산항목 no
 * @apiSuccess {String} data.capital_name 재산항목명
 *
 */

/**
 * @api {post} /api/categoryC categoryC
 * @apiVersion 0.1.0
 * @apiName CategoryCreate
 * @apiDescription 구분항목 추가
 * @apiGroup Category
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"category_name":"월급", "is_input":"1", \
"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/deviceC
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} category_name 구분항목명
 * @apiParam {String} is_input 1:수입, 0:지출
 *
 * @apiSuccess {String} result 요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 *
 */

/**
 * @api {post} /api/categoryR categoryR
 * @apiVersion 0.1.0
 * @apiName CategoryRead
 * @apiDescription 구분항목 정보 조회
 * @apiGroup Category
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"categoryno":"34", \
"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/categoryR
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} categoryno 카테고리 no
 *
 * @apiSuccess {String} result 요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {String} category_name 구분항목명
 * @apiSuccess {String} is_input 1:수입, 0:지출
 */

/**
 * @api {post} /api/categoryU categoryU
 * @apiVersion 0.1.0
 * @apiName CategoryUpdate
 * @apiDescription 구분항목 정보 갱신
 * @apiGroup Category
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", \
"categoryno":"36", "category_name":"my iphone"}' \
http://localhost:8082/api/categoryU
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} categoryno 카테고리 no
 * @apiParam {String} [category_name] 구분항목명
 * @apiParam {String} [is_input] 1:수입, 0:지출
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 */

/**
 * @api {post} /api/categoryD categoryD
 * @apiVersion 0.1.0
 * @apiName CategoryDelete
 * @apiDescription 구분항목 삭제
 * @apiGroup Category
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", \
"categoryno":"36"}' \
http://localhost:8082/api/categoryD
 *
 * @apiParam (Param_Action_1) {String} access_token 인증키
 * @apiParam (Param_Action_1) {String} categoryno 카테고리 no
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 *
 */

/**
 * @api {post} /api/categoryL categoryL
 * @apiVersion 0.1.0
 * @apiName CategoryList
 * @apiDescription 구분항목 목록
 * @apiGroup Category
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/categoryL
 *
 * @apiParam {String} access_token 인증키
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {Object[]} data 구분항목 목록
 * @apiSuccess {String} data.categoryno 카테고리 no
 * @apiSuccess {String} data.category_name 구분항목명
 * @apiSuccess {String} data.is_input 1:수입, 0:지출
 *
 */

/**
 * @api {post} /api/categoryS categoryS
 * @apiVersion 0.1.0
 * @apiName CategorySearch
 * @apiDescription 구분항목 검색
 * @apiGroup Category
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", "category_name":"my"}' \
http://localhost:8082/api/categoryS
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} category_name 구분항목명
 * @apiParam {String} [is_input] 1:수입, 0:지출
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {Object[]} data 구분항목 목록
 * @apiSuccess {String} data.categoryno 카테고리 no
 * @apiSuccess {String} data.category_name 구분항목명
 * @apiSuccess {String} data.is_input 1:수입, 0:지출
 *
 */

/**
 * @api {post} /api/itemC itemC
 * @apiVersion 0.1.0
 * @apiName ItemCreate
 * @apiDescription 항목 추가
 * @apiGroup Item
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"categoryno":"23", "capitalno":"43", "item_name":"점심", "item_value":"8000", \
"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/itemC
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} categoryno 카테고리 no
 * @apiParam {String} capitalno 재산항목 no
 * @apiParam {String} item_name 항목이름
 * @apiParam {String} item_value 항목금액
 *
 * @apiSuccess {String} result 요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 *
 */

/**
 * @api {post} /api/itemR itemR
 * @apiVersion 0.1.0
 * @apiName ItemRead
 * @apiDescription 항목 조회
 * @apiGroup Item
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"itemno":"34", \
"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/itemR
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} itemno 항목 no
 *
 * @apiSuccess {String} result 요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {String} categoryno 카테고리 no
 * @apiSuccess {String} capitalno 재산항목 no
 * @apiSuccess {String} item_name 항목이름
 * @apiSuccess {String} item_value 항목금액
 * @apiSuccess {String} create_datetime 등록일시
 */

/**
 * @api {post} /api/itemU itemU
 * @apiVersion 0.1.0
 * @apiName ItemUpdate
 * @apiDescription 항목 갱신
 * @apiGroup Item
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", \
"itemno":"36", "item_name":"점심"}' \
http://localhost:8082/api/itemU
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} itemno 항목 no
 * @apiParam {String} [categoryno] 카테고리 no
 * @apiParam {String} [capitalno] 재산항목 no
 * @apiParam {String} [item_name] 항목이름
 * @apiParam {String} [item_value] 항목금액
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 */

/**
 * @api {post} /api/itemD itemD
 * @apiVersion 0.1.0
 * @apiName ItemDelete
 * @apiDescription 항목 삭제
 * @apiGroup Item
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", \
"itemno":"36"}' \
http://localhost:8082/api/itemD
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} itemno 항목 no
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 *
 */

/**
 * @api {post} /api/itemL itemL
 * @apiVersion 0.1.0
 * @apiName ItemList
 * @apiDescription 항목 목록
 * @apiGroup Item
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/deviceL
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} [page] 페이지
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {Object[]} data 데이터 목록
 * @apiSuccess {String} data.categoryno 카테고리 no
 * @apiSuccess {String} data.capitalno 재산항목 no
 * @apiSuccess {String} data.item_name 항목이름
 * @apiSuccess {String} data.create_datetime 등록일시
 *
 */

/**
 * @api {post} /api/itemS itemS
 * @apiVersion 0.1.0
 * @apiName ItemSearch
 * @apiDescription 항목 검색
 * @apiGroup Item
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", "search":{"capitalno":"34"}}' \
http://localhost:8082/api/itemS
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} [page] 페이지
 * @apiParam {Object[]} search 검색 필터
 * @apiParam {String} [search.categoryno] 카테고리 no
 * @apiParam {String} [search.capitalno] 재산항목 no
 * @apiParam {String} [search.item_name] 항목이름
 * @apiParam {String} [search.item_value] 항목금액
 * @apiParam {String} [search.date_begin] 가입 시작 일시
 * @apiParam {String} [search.date_end] 가입 종료 일시
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {Object[]} data 데이터 목록
 * @apiSuccess {String} data.categoryno 카테고리 no
 * @apiSuccess {String} data.capitalno 재산항목 no
 * @apiSuccess {String} data.item_name 항목이름
 * @apiSuccess {String} data.item_value 항목금액
 * @apiSuccess {String} data.create_datetime 등록일시
 *
 */

/**
 * @api {post} /api/systemR systemR
 * @apiVersion 0.1.0
 * @apiName SystemRead
 * @apiDescription 시스템 정보 조회
 * @apiGroup System
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
http://localhost:8082/api/systemR
 *
 * @apiSuccess {String} result 요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {String} api_version
 */

/**
 * @api {post} /api/organizationR organizationR
 * @apiVersion 0.1.0
 * @apiName OrganizationRead
 * @apiDescription 조직 멤버 목록 조회
 * @apiGroup Organization
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", "groupno":"35"}' \
http://localhost:8082/api/organizationR
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} groupno 그룹 no
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {String} group_name 그룹명
 * @apiSuccess {String} create_datetime 등록일시
 * @apiSuccess {Object[]} data 멤버 정보 목록
 * @apiSuccess {String} data.dispname 대화명
 * @apiSuccess {String} data.intro 자기소개
 * @apiSuccess {String} data.profile 프로필 사진
 * @apiSuccess {String} data.create_datetime 가입일시
 * @apiSuccess {String} data.phone_number 전화번호
 */

/**
 * @api {post} /api/organizationL organizationL
 * @apiVersion 0.1.0
 * @apiName OrganizationList
 * @apiDescription 조직 목록 조회
 * @apiGroup Organization
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/organizationL
 *
 * @apiParam {String} access_token 인증키
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {Object[]} data 그룹 목록
 * @apiSuccess {String} data.groupno 그룹 no
 * @apiSuccess {String} data.group_name 그룹명
 * @apiSuccess {String} data.create_datetime 등록일시
 */

/**
 * @api {post} /api/organizationS organizationS
 * @apiVersion 0.1.0
 * @apiName OrganizationSearch
 * @apiDescription 조직 목록 검색
 * @apiGroup Organization
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"ef53163004dd7257c52e9571fff5751f72940bdd", "group_name":"my group"}' \
http://localhost:8082/api/organizationS
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} group_name 그룹명
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {Object[]} data 그룹 목록
 * @apiSuccess {String} data.groupno 그룹 no
 * @apiSuccess {String} data.group_name 그룹명
 * @apiSuccess {String} data.create_datetime 등록일시
 */
