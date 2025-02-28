#  Hi,We are free ,please avoid fraud.
#  Generated at 2023-12-31 00:00 by WekiYun
#  Clash.Meta Core
ipv6: true
allow-lan: true
mixed-port: 7890
unified-delay: false
tcp-concurrent: true
external-controller: 0.0.0.0:9090
external-ui: ui
external-ui-url: "https://github.com/MetaCubeX/metacubexd/archive/refs/heads/gh-pages.zip"

geodata-mode: true
geox-url:
  geoip: "https://mirror.ghproxy.com/https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/geoip-lite.dat"
  geosite: "https://mirror.ghproxy.com/https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/geosite.dat"
  mmdb: "https://mirror.ghproxy.com/https://github.com/MetaCubeX/meta-rules-dat/releases/download/latest/country-lite.mmdb"

find-process-mode: strict
global-client-fingerprint: chrome

profile:
  store-selected: true
  store-fake-ip: true

sniffer:
  enable: true
  sniff:
    HTTP:
      ports: [80, 8080-8880]
      override-destination: true
    TLS:
      ports: [443, 8443]
    QUIC:
      ports: [443, 8443]

tun:
  enable: true
  stack: mixed
  dns-hijack:
    - "any:53"
  auto-route: true
  auto-detect-interface: true

dns:
  enable: true
  listen: :1053
  ipv6: true
  enhanced-mode: fake-ip
  fake-ip-filter:
    - "*"
    - "+.lan"
    - "+.local"
  nameserver:
    - https://doh.pub/dns-query
    - https://dns.alidns.com/dns-query
  proxy-server-nameserver:
    - https://doh.pub/dns-query
  nameserver-policy:
    "geosite:private":
      - https://doh.pub/dns-query
      - https://dns.alidns.com/dns-query
    "geosite:geolocation-!cn":
      - "https://dns.cloudflare.com/dns-query#dns"
      - "https://dns.google/dns-query#dns"

proxy-groups:
  - name: 🌐 节点选择
    type: select
    proxies:
      - ♻️ 自动选择
      - 🔯 故障转移
      - 🔮 负载均衡-轮询
      - 🔮 负载均衡-散列
      - DIRECT
    use:
      - 🛩️ Chrome.GO
  - name: ♻️ 自动选择
    type: url-test
    url: http://www.gstatic.com/generate_204
    interval: 300
    tolerance: 50
    proxies:
      - 🔮 负载均衡-轮询
      - 🔮 负载均衡-散列
    use:
      - 🛩️ Chrome.GO
  - name: 🔯 故障转移
    type: fallback
    url: http://www.gstatic.com/generate_204
    interval: 180
    proxies:
    use:
      - 🛩️ Chrome.GO
  - name: 🔮 负载均衡-散列
    type: load-balance
    url: http://www.gstatic.com/generate_204
    interval: 300
    strategy: consistent-hashing
    proxies:
    use:
      - 🛩️ Chrome.GO
  - name: 🔮 负载均衡-轮询
    type: load-balance
    url: http://www.gstatic.com/generate_204
    interval: 300
    strategy: round-robin
    proxies:
    use:
      - 🛩️ Chrome.GO
  - name: 🎯 全球直连
    type: select
    proxies:
      - DIRECT
      - 🌐 节点选择
      - ♻️ 自动选择
  - name: 🛑 全球拦截
    type: select
    proxies:
      - REJECT
      - DIRECT
  - name: 🐟 漏网之鱼
    type: select
    proxies:
      - 🌐 节点选择
      - 🎯 全球直连
      - ♻️ 自动选择
      - 🔯 故障转移
      - 🔮 负载均衡-轮询
      - 🔮 负载均衡-散列
    use:
      - 🛩️ Chrome.GO

proxy-providers:
  🛩️ Chrome.GO:
    type: http
    url: "https://api.v1.mk/sub?target=clash&url=https%3A%2F%2Fchromenodes.marcol.top%2F&insert=false&config=https%3A%2F%2Fraw.githubusercontent.com%2FACL4SSR%2FACL4SSR%2Fmaster%2FClash%2Fconfig%2FACL4SSR_Online_Mini.ini&rename=%6051.158.54.46%40%F0%9F%93%A15iClub%60%60167.160.91.115%40%F0%9F%95%895iClub%60%60108.181.22.239%405%EF%B8%8F%E2%83%A35iClub%60%60173.234.25.52%40%F0%9F%8C%805iClub%60%60www.dtku46.xyz%40%F0%9F%8C%905iClub%60%60108.181.22.155%40%F0%9F%90%9F5iClub%60%60167.160.90.252%40%E2%99%BB%EF%B8%8F5iClub%60%60cloudflare.cfgo.cc%40%F0%9F%94%AE5iClub%60%60dongtaiwang3.com%40%F0%9F%8E%AF5iClub%60%60108.181.22.213%40%F0%9F%94%AF5iClub%60%60www.dtku50.xyz%40%F0%9F%9A%805iClub%60%60ip%40%F0%9F%8D%835iClub%60%60quick%40QUICK%60&emoji=true&list=false&xudp=false&udp=false&tfo=false&expand=true&scv=false&fdn=false&new_name=true"
    path: Chrome/GO.yaml
    interval: 3600
    health-check:
      enable: true
      interval: 300
      url: http://www.gstatic.com/generate_204

rule-providers:
  reject:
    type: http
    behavior: domain
    url: "https://mirror.ghproxy.com/https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/reject.txt"
    path: ./ClodQiu/reject.yaml
    interval: 86400
  icloud:
    type: http
    behavior: domain
    url: "https://mirror.ghproxy.com/https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/icloud.txt"
    path: ./ClodQiu/icloud.yaml
    interval: 86400
  apple:
    type: http
    behavior: domain
    url: "https://mirror.ghproxy.com/https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/apple.txt"
    path: ./ClodQiu/apple.yaml
    interval: 86400
  google:
    type: http
    behavior: domain
    url: "https://mirror.ghproxy.com/https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/google.txt"
    path: ./ClodQiu/google.yaml
    interval: 86400
  proxy:
    type: http
    behavior: domain
    url: "https://mirror.ghproxy.com/https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/proxy.txt"
    path: ./ClodQiu/proxy.yaml
    interval: 86400
  direct:
    type: http
    behavior: domain
    url: "https://mirror.ghproxy.com/https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/direct.txt"
    path: ./ClodQiu/direct.yaml
    interval: 86400
  private:
    type: http
    behavior: domain
    url: "https://mirror.ghproxy.com/https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/private.txt"
    path: ./ClodQiu/private.yaml
    interval: 86400
  gfw:
    type: http
    behavior: domain
    url: "https://mirror.ghproxy.com/https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/gfw.txt"
    path: ./ClodQiu/gfw.yaml
    interval: 86400
  tld-not-cn:
    type: http
    behavior: domain
    url: "https://mirror.ghproxy.com/https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/tld-not-cn.txt"
    path: ./ClodQiu/tld-not-cn.yaml
    interval: 86400
  telegramcidr:
    type: http
    behavior: ipcidr
    url: "https://mirror.ghproxy.com/https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/telegramcidr.txt"
    path: ./ClodQiu/telegramcidr.yaml
    interval: 86400
  cncidr:
    type: http
    behavior: ipcidr
    url: "https://mirror.ghproxy.com/https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/cncidr.txt"
    path: ./ClodQiu/cncidr.yaml
    interval: 86400
  lancidr:
    type: http
    behavior: ipcidr
    url: "https://mirror.ghproxy.com/https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/lancidr.txt"
    path: ./ClodQiu/lancidr.yaml
    interval: 86400
  applications:
    type: http
    behavior: classical
    url: "https://mirror.ghproxy.com/https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/applications.txt"
    path: ./ClodQiu/applications.yaml
    interval: 86400

rules:
  - RULE-SET,applications,🎯 全球直连
  - DOMAIN,clash.razord.top,🎯 全球直连
  - DOMAIN,yacd.haishan.me,🎯 全球直连
  - RULE-SET,private,🎯 全球直连
  - RULE-SET,reject,🛑 全球拦截
  - RULE-SET,icloud,🎯 全球直连
  - RULE-SET,apple,🎯 全球直连
  - RULE-SET,google,🎯 全球直连
  - RULE-SET,proxy,🌐 节点选择
  - RULE-SET,direct,🎯 全球直连
  - RULE-SET,lancidr,🎯 全球直连
  - RULE-SET,cncidr,🎯 全球直连
  - RULE-SET,telegramcidr,🌐 节点选择
  - GEOIP,LAN,🎯 全球直连,no-resolve
  - GEOIP,CN,🎯 全球直连,no-resolve
  - MATCH,🌐 节点选择
