# menu-board-app
----
sql
- sql 글자수제한 검사 / front에서 검사
- 사진 용량 제한 / front에서 검사
- restaurant_types, menu_types 전용테이블에서 가져오기 / front에서 검사
- menus table에서 menu_type references restarant.menutype


[ ] middleware logger 만들기
[ ] nginx로 미래에 api-gate 만들기

기본 -> 테스트 -> 기능추가 -> 테스트

기본 (테스트)
- back, db -> front

기능추가
- 요청시 redis cache 1시간씩 추가
- 원산지 업데이트시 업데이트 된부분 저장
- QR코드 공유가능
- email 비밀번호 찾기
- 알러지 물질 클릭시 제외 명도 낮추기 /front 에서 map으로 재구성

next
- 요리 타입(한식, 일식..)에 따라 분석 기능
- 주문 결제 가능
- 공통된 메뉴(주류 etc...) 미리 제공해서 sql 줄이고 편의성 높이기
- webRTC로 고화질 전송 (망사용료 해소) 가계마다 작은 기계(k8s node로 사용)