// body.json

/**
 * @api {post} /api/admin/userC userC
 * @apiVersion 0.1.0
 * @apiName AdminUserC
 * @apiDescription 가계부 사용자 추가
 * @apiGroup User
 *
 * @apiExample {curl} Example :
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
 -d '{"access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848",
 "userid":"admin@aaa.com", "password":"admin"}' \
 http://localhost:8082/api/admin/userC
 *
 * @apiParam {String} userid userid
 * @apiParam {String} password 비밀번호
 * @apiParam {String} [dispname] 대화명
 * @apiParam {String} [intro] 자기소개
 * @apiParam {String} [profile] 프로필 사진
 * @apiParam {String} [phone_number] 전화번호
 * @apiParam {String} [public] 그룹에 공개여부
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우
 */

/**
 * @api {post} /api/admin/userR userR
 * @apiVersion 0.1.0
 * @apiName AdminUserR
 * @apiDescription 가계부 사용자 조회
 * @apiGroup User
 *
 * @apiExample {curl} Example : (action=1)
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"action":"1", "userid":"admin@localhost.com", "password":"12345678"}' \
http://localhost:8082/api/admin/userR
 *
 * @apiExample {curl} Example : (action=2)
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"action":"2", "access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848", 
"userno":"57b97e21681b956fd6eb93de"}' \
http://localhost:8082/api/admin/userR
 *
 * @apiParam (Param_Action_1) {String} action 1 : login
 * @apiParam (Param_Action_1) {String} userid userid
 * @apiParam (Param_Action_1) {String} password 비밀번호
 *
 * @apiParam (Param_Action_2) {String} action 2 : 정보 조회
 * @apiParam (Param_Action_2) {String} userno 사용자 no
 * @apiParam (Param_Action_2) {String} access_token 인증키
 *
 * @apiSuccess (Response_Action_1) {String} result "success" or "fail"
 * @apiSuccess (Response_Action_1) {String} err_msg result가 "fail"인 경우
 * @apiSuccess (Response_Action_1) {String} access_token result가 "success"인 경우
 *
 * @apiSuccess (Response_Action_2) {String} result "success" or "fail"
 * @apiSuccess (Response_Action_2) {String} err_msg result가 "fail"인 경우
 * @apiSuccess (Response_Action_2) {Object} data 회원 정보
 * @apiSuccess (Response_Action_2) {String} data.userid login id(email)
 * @apiSuccess (Response_Action_2) {String} data.usertype 5:사용자, 9:관리자
 * @apiSuccess (Response_Action_2) {String} data.dispname 대화명
 * @apiSuccess (Response_Action_2) {String} data.intro 자기소개
 * @apiSuccess (Response_Action_2) {String} data.profile 프로필 사진
 * @apiSuccess (Response_Action_2) {String} data.create_datetime 가입일시
 * @apiSuccess (Response_Action_2) {String} data.phone_number 전화번호
 * @apiSuccess (Response_Action_2) {String} data.state 1:ON(사용), 0:OFF(탈퇴)
 * @apiSuccess (Response_Action_2) {String} data.activated 2:관리자 인증, 1:URL 인증,0:OFF 미인증
 * @apiSuccess (Response_Action_2) {String} data.public 1 = 공개 0 = 비공개
 */

/**
 * @api {post} /api/admin/userU userU
 * @apiVersion 0.1.0
 * @apiName AdminUserU
 * @apiDescription 가계부 관리자 정보 수정
 * @apiGroup User
 *
 * @apiExample {curl} Example : (action=1)
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"action":"1", "userid":"admin@aaa.com", "old_password":"admin", "new_password":"12345678"}' \
http://localhost:8082/api/admin/userU
 *
 * @apiExample {curl} Example : (action=2)
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"action":"2", "access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848", 
"userno":"57efaa32aaaf21f8dcb105b8",
"update":{"usertype":"4", "displayname":"대화명"}}' \
http://localhost:8082/api/admin/userU
 *
 * @apiExample {curl} Example : (action=3)
 * curl \
-F "action=3" \
-F "userno=57efaa32aaaf21f8dcb105b8" \
-F "access_token=796f25e350f5b66eb1f8c77ebdeba5db532f3848" \
-F "pimage=@/tmp/go_build_for_window.png" \
http://localhost:8082/api/admin/userU  
 *
 * @apiParam (Param_Action_1) {String} action 1 : admin init password change
 * @apiParam (Param_Action_1) {String} userid userid
 * @apiParam (Param_Action_1) {String} old_password 이전 비밀번호
 * @apiParam (Param_Action_1) {String} new_password 새 비밀번호
 *
 * @apiParam (Param_Action_2) {String} action 2 : update user info
 * @apiParam (Param_Action_2) {String} access_token 인증키
 * @apiParam (Param_Action_2) {String} userno 회원키
 * @apiParam (Param_Action_2) {Object} update 갱신 대상 정보
 * @apiParam (Param_Action_2) {String} [update.usertype] 5:사용자, 9:관리자
 * @apiParam (Param_Action_2) {String} [update.dispname] 대화명
 * @apiParam (Param_Action_2) {String} [update.intro] 자기소개
 * @apiParam (Param_Action_2) {String} [update.phone_number] 전화번호
 * @apiParam (Param_Action_2) {String} [update.state] 1:ON(사용), 0:OFF(탈퇴)
 * @apiParam (Param_Action_2) {String} [update.activated] 2:관리자 인증, 1:URL 인증,0:OFF 미인증
 * @apiParam (Param_Action_2) {String} [update.public] 1 = 공개 0 = 비공개
 *
 * @apiParam (Param_Action_3) {String} action 3 : update user profile
 * @apiParam (Param_Action_3) {String} userno 회원키
 * @apiParam (Param_Action_3) {file} pimage 프로필 이미지
 * @apiParam (Param_Action_3) {String} access_token 인증키
 *
 * @apiSuccess (Response_Action_1) {String} result "success" or "fail"
 * @apiSuccess (Response_Action_1) {String} err_msg result가 "fail"인 경우
 * @apiSuccess (Response_Action_1) {String} access_token result가 "success"인 경우
 *
 * @apiSuccess (Response(action=2,3)) {String} result "success" or "fail"
 * @apiSuccess (Response(action=2,3)) {String} err_msg result가 "fail"인 경우
 */

/**
 * @api {post} /api/admin/userD userD
 * @apiVersion 0.1.0
 * @apiName UserDelete
 * @apiDescription 가계부 관리자 회원 삭제
 * @apiGroup User
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848", "userno":"57f077e5aaaf21f8dcb105ba"}' \
http://localhost:8082/api/admin/userD
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} userno 사용자 no
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 *
 */

/**
 * @api {post} /api/admin/userL userL
 * @apiVersion 0.1.0
 * @apiName AdminUserL
 * @apiDescription 가계부 관리자 회원목록
 * @apiGroup User
 *
 * @apiExample {curl} Example :
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848"}' \
http://localhost:8082/api/admin/userL
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} [page] 페이지
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우
 * @apiSuccess {Object[]} data 멤버 정보 목록
 * @apiSuccess {String} data.userno 사용자 no
 * @apiSuccess {String} data.userid login id(email)
 * @apiSuccess {String} data.usertype 5:사용자, 9:관리자
 * @apiSuccess {String} data.dispname 대화명
 * @apiSuccess {String} data.intro 자기소개
 * @apiSuccess {String} data.profile 프로필 사진
 * @apiSuccess {String} data.create_datetime 가입일시
 * @apiSuccess {String} data.phone_number 전화번호
 * @apiSuccess {String} data.state 1:ON(사용), 0:OFF(탈퇴)
 * @apiSuccess {String} data.activated 2:관리자 인증, 1:URL 인증,0:OFF 미인증
 * @apiSuccess {String} data.public 1 = 공개 0 = 비공개
 */

/**
 * @api {post} /api/admin/userS userS
 * @apiVersion 0.1.0
 * @apiName AdminUserS
 * @apiDescription 가계부 관리자 회원 검색
 * @apiGroup User
 *
 * @apiExample {curl} Example :
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848","search":{"userid":"admin"}}' \
http://localhost:8082/api/admin/userS
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} [page] 페이지
 * @apiParam {Object} search 검색 필터
 * @apiParam {String} [search.userid] login id(email)
 * @apiParam {String} [search.usertype] 5:사용자, 9:관리자
 * @apiParam {String} [search.dispname] 대화명
 * @apiParam {String} [search.intro] 자기소개
 * @apiParam {String} [search.date_begin] 가입 시작 일시
 * @apiParam {String} [search.date_end] 가입 종료 일시
 * @apiParam {String} [search.phone_number] 전화번호
 * @apiParam {String} [search.state] 1:ON(사용), 0:OFF(탈퇴)
 * @apiParam {String} [search.activated] 2:관리자 인증, 1:URL 인증,0:OFF 미인증
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우
 * @apiSuccess {Object[]} data 멤버 정보 목록
 * @apiSuccess {String} data.userno 사용자 no
 * @apiSuccess {String} data.userid login id(email)
 * @apiSuccess {String} data.usertype 5:사용자, 9:관리자
 * @apiSuccess {String} data.dispname 대화명
 * @apiSuccess {String} data.intro 자기소개
 * @apiSuccess {String} data.profile 프로필 사진
 * @apiSuccess {String} data.create_datetime 가입일시
 * @apiSuccess {String} data.phone_number 전화번호
 * @apiSuccess {String} data.state 1:ON(사용), 0:OFF(탈퇴)
 * @apiSuccess {String} data.activated 2:관리자 인증, 1:URL 인증,0:OFF 미인증
 * @apiSuccess {String} data.public 1 = 공개 0 = 비공개
 */

/**
 * @api {post} /api/admin/systemR systemR
 * @apiVersion 0.1.0
 * @apiName AdminSystemR
 * @apiDescription 가계부 시스템 조회
 * @apiGroup System
 *
 * @apiExample {curl} Example
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848","systemno":"57e78edbaaaf21f8dcb105a7"}' \
http://localhost:8082/api/admin/systemR
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} systemno 관리항목no
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우
 * @apiSuccess {Object} data system 정보
 * @apiSuccess {String} data.smtp_url SMTP 서버
 * @apiSuccess {String} data.smtp_port SMTP 포트
 * @apiSuccess {String} data.smtp_id SMTP 아이디
 * @apiSuccess {String} data.smtp_pw SMTP 패스워드
 * @apiSuccess {String} data.email_success SMTP 인증 완료
 * @apiSuccess {String} data.email_authdate 5분 후 만료
 */

/**
 * @api {post} /api/admin/systemU systemU
 * @apiVersion 0.1.0
 * @apiName AdminSystemU
 * @apiDescription 가계부 시스템 갱신
 * @apiGroup System
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
 -d '{"access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848",
 "systemno":"57e78edbaaaf21f8dcb105a7", "item_value":"admin@test.com"}' \
 http://localhost:8082/api/admin/systemU
 *
 * @apiExample {curl} Example : update desc
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
 -d '{"access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848", "systemno":"57e78edbaaaf21f8dcb105a7", 
 "item_value":"admin222@test.com", "item_desc": "인증 이메일 발송자 계정(email id)"}' \
 http://localhost:8082/api/admin/systemU
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} systemno 관리항목no
 * @apiParam {String} item_value 관리항목 value
 * @apiParam {String} [item_desc] 관리항목 설명
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 */

/**
 * @api {post} /api/admin/systemL systemL
 * @apiVersion 0.1.0
 * @apiName AdminSystemL
 * @apiDescription 가계부 시스템 목록 조회
 * @apiGroup System
 *
 * @apiExample {curl} Example
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848"}' \
http://localhost:8082/api/admin/systemL
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} [page] 페이지
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우
 * @apiSuccess {Object[]} data system 정보
 * @apiSuccess {String} data.systemno 관리항목no
 * @apiSuccess {String} data.item_name 관리항목 이름
 * @apiSuccess {String} data.item_value 관리항목 내용
 */

/**
 * @api {post} /api/admin/systemS systemS
 * @apiVersion 0.1.0
 * @apiName AdminSystemS
 * @apiDescription 가계부 시스템 목록 검색
 * @apiGroup System
 *
 * @apiExample {curl} Example
 * curl -X POST  -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848","search":"system"}' \
http://localhost:8082/api/admin/systemS
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} search 검색어
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우
 * @apiSuccess {Object[]} data system 정보
 * @apiSuccess {String} data.systemno 관리항목no
 * @apiSuccess {String} data.item_name 관리항목 이름
 * @apiSuccess {String} data.item_value 관리항목 내용
 */

/**
 * @api {post} /api/admin/groupC groupC
 * @apiVersion 0.1.0
 * @apiName AdminGroupCreate
 * @apiDescription 그룹 추가
 * @apiGroup Group
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"group_name":"group name.", "access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848"}' \
http://localhost:8082/api/admin/groupC
 *
 * @apiExample {curl} Example : permission error
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"group_name":"group name.", "access_token":"ef53163004dd7257c52e9571fff5751f72940bdd"}' \
http://localhost:8082/api/admin/groupC
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} group_name 그룹명
 *
 * @apiSuccess {String} result 요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 *
 */

/**
 * @api {post} /api/admin/groupR groupR
 * @apiVersion 0.1.0
 * @apiName AdminGroupRead
 * @apiDescription 그룹 정보 조회
 * @apiGroup Group
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"groupno":"57ddf6baaaaf21f8dcb0ff69", "access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848"}' \
http://localhost:8082/api/admin/groupR
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} groupno 그룹 no
 *
 * @apiSuccess {String} result 요청 결과 "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {String} group_name 그룹명
 * @apiSuccess {String} create_datetime 등록일시
 * @apiSuccess {String} state 1:ON(사용),0:OFF(삭제)
 */

/**
 * @api {post} /api/admin/groupU groupU
 * @apiVersion 0.1.0
 * @apiName AdminGroupUpdate
 * @apiDescription 그룹 정보 갱신
 * @apiGroup Group
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848", "groupno":"57e643b5aaaf21f8dcb103b2", "state":"1"}' \
http://localhost:8082/api/admin/groupU
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} groupno 그룹no
 * @apiParam {String} [state] 1:ON(사용),0:OFF(삭제)
 * @apiParam {String} [group_name] 그룸명
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 */

/**
 * @api {post} /api/admin/groupD groupD
 * @apiVersion 0.1.0
 * @apiName AdminGroupDelete
 * @apiDescription 그룹 삭제
 * @apiGroup Group
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848", "groupno":"36"}' \
http://localhost:8082/api/admin/groupD
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} groupno 그룹 no
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 *
 */

/**
 * @api {post} /api/admin/groupL groupL
 * @apiVersion 0.1.0
 * @apiName AdminGroupList
 * @apiDescription 그룹 목록 조회
 * @apiGroup Group
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848"}' \
http://localhost:8082/api/admin/groupL
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} [page] 페이지
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {Object[]} data 그룹 목록
 * @apiSuccess {String} data.groupno 그룹 no
 * @apiSuccess {String} data.group_name 그룹명
 * @apiSuccess {String} data.create_datetime 등록일시
 * @apiSuccess {String} data.state 1:ON(사용),0:OFF(삭제)
 */

/**
 * @api {post} /api/admin/groupS groupS
 * @apiVersion 0.1.0
 * @apiName AdminGroupSearch
 * @apiDescription 그룹 목록 검색
 * @apiGroup Group
 *
 * @apiExample {curl} Example :
 * curl -X POST -H "Accept: Application/json" -H "Content-Type: application/json" \
-d '{"access_token":"796f25e350f5b66eb1f8c77ebdeba5db532f3848", "group_name":"my"}' \
http://localhost:8082/api/admin/groupS
 *
 * @apiParam {String} access_token 인증키
 * @apiParam {String} group_name 그룹명
 * @apiParam {String} [page] 페이지
 *
 * @apiSuccess {String} result "success" or "fail"
 * @apiSuccess {String} err_msg result가 "fail"인 경우 에러 메시지
 * @apiSuccess {Object[]} data 그룹 목록
 * @apiSuccess {String} data.groupno 그룹 no
 * @apiSuccess {String} data.group_name 그룹명
 * @apiSuccess {String} data.create_datetime 등록일시
 * @apiSuccess {String} data.state 1:ON(사용),0:OFF(삭제)
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
 * @apiDefine Param_access_token
 * 111122223333
 */
