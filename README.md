# Miner-agent
Miner-Agent 는 희망마이닝 전용 모니터링 툴입니다.

현재는 베타서비스로 제공 되고 있으며 서버 상황에 따라 잠시 중단 될 수 있습니다.

계속 적으로 업그레이드 할 예정입니다.

# Required
- golang (https://golang.org/)

# Getting started

    #go build Miner-agent.go

파일이 생성되면 실행 파라메터와 함께 실행하면 됩니다.

# Usaging

사용하는 파라메터는 총 3가지가 있습니다.

 파라메터        | 내용           | 필수 여부 
 ------------- |:-------------:| -----:
 workername  | 워커 이름을 작성해준다. (ex : test.workername1) | 필수
 currency      | 화폐명을 작성해준다. (ex : zcash)     | 선택
 miner | 현재 사용중인 마이너 이름을 적어준다. {claymore\|ewbfminer\|ccminer\|cgminer\|sgminer} | 선택

    Miner-agent -workername=test.workername1 -currency=zcash -miner=ewbfminer


# Miner Batch file Setting
Miner-agent 는 채굴 프로그램에서 API 환경이 열려 있을 때 사용 가능합니다.

Claymore 를 제외한 나머지 마이너의 경우 특정 세팅이 추가되어야 사용할 수 있습니다.

- Claymore

클레이 모어는 자동으로 3333 포트에 연결됩니다.
- EWBF Zcash Miner

--api 0.0.0.0:42000 를 해주셔야 Miner-agent 가 정보를 수집 할 수 있습니다.

- CCMiner (beta)

--api-bind 0.0.0.0:4068 를 해주셔야 Miner-agent 가 정보를 수집 할 수 있습니다.

- CGMiner (beta)

--api-bind 0.0.0.0:4068 를 해주셔야 Miner-agent 가 정보를 수집 할 수 있습니다.

- SGMiner (beta)

--api-bind 0.0.0.0:4028 를 해주셔야 Miner-agent 가 정보를 수집 할 수 있습니다.

# Simple usaging

이미 빌드된 실행 파일이 포함되어 있습니다.

releases 폴더의 시스템에 맞는 실행 파일의 .bat 파일의 파라메터를 변경 후 실행 하면 됩니다.

# Contact us
Issue : https://github.com/hopemining/Miner-agent/issues

E-mail : info@hopemining.org
