package main

import (
	"fmt"
	"strings"

	"github.com/yosssi/gohtml"
)

func main() {
	h := `<p class=\"union-media-editor__paragraph\">亲爱的合作伙伴，</p><p class=\"union-media-editor__paragraph\">pangle 致力于帮助应用开发者实现更有效的流量变现，推动用户增长，并释放创造潜力，在开发者、广告主和最终用户之间建立一个良好的生态系统。 pangle 邀请所有应用程序开发人员与 pangle 合作，以确保所有广告活动都是真实有效的。</p><p class=\"union-media-editor__paragraph\">我们观察到您的 pangle 账户存在违规行为。</p><p class=\"union-media-editor__paragraph\">以下是我们审核的 [广告位/应用/账号] 的详细信息：</p><br><ul><li>应用：</li></ul><p class=\"union-media-editor__paragraph\"><table border = \"1\" cellspacing=\"0\" cellpadding=\"10px\"><tr align=\"center\"><th colspan=\"1\">应用id</th><th colspan=\"1\">应用名称</th><th colspan=\"1\">包名</th><th colspan=\"1\">国家</th><th colspan=\"1\">处置方式</th><th colspan=\"1\">是否扣款</th><th colspan=\"1\">违规原因/建议</th></tr><tr align=\"center\"><td>50051732</td><td>zombie tsunami</td><td>net.mobigame.zombietsunami</td><td>ma</td><td>停止填充</td><td>否</td><td>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作系统、系统版本等信息的准确性和集中度等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>50053106</td><td>hair color match</td><td>hair.colormatch</td><td>om</td><td>停止填充</td><td>否</td><td>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作系统、系统版本等信息的准确性和集中<br>度等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>50052834</td><td>watt the bulb</td><td>watt.thebulb</td><td>ar</td><td>限制填充</td><td>否</td><td>[您的应用被识别为vpn应用，目前pangle不支持在vpn应用上变现。/请检查并确保您账户下的应用不是vpn应用。]</td></tr><tr align=\"center\"><td>50052547</td><td>rare pets - merge game mystery</td><td>com.drihq.pets</td><td>jp</td><td>限制填充</td><td>是</td><td>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作系统、系统版本等信息的准确性和集中度等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>50053369</td><td>idle casino manager - tycoon</td><td>com.coldfiregames.idle.casino.tycoon</td><td>ar,ca,th</<br>td><td>停止填充</td><td>是</td><td>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>50053370</td><td>idle hustle kingdom</td><td>com.tiamatgames.idlehustlekingdom</td><td>ae,ar</td><td>停止填充</td><td>否</td><td>[您的应用被识别为vpn应用，目前pangle不支持在vpn应用上变现。/请检查并确保您账户下的应用不是vpn应用。]</td></tr><tr align=\"center\"><td>50053106</td><td>hair color match</td><td>hair.colormatch</td><td>ae,bh,br</td><td>停止填充</td><td>否</td><td>[您的应用被识别为vpn应用，目前pangle不支持在vpn应用上变现。/请检查并确保您账户下的应用不是vpn应用。]</td></tr><tr align=\"center\"><td>50051733</td><td>compare them</td><td>com.rsamy.comparethem</td><td>global</td><td>停止填充</td><td>否</td><td>[无效流量由流量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指<br>标。]</td></tr><tr align=\"center\"><td>50052834</td><td>watt the bulb</td><td>watt.thebulb</td><td>au,bh</td><td>停止填充</td><td>否</td><td>[您的应用被识别为vpn应用，目前pangle不支持在vpn应用上变现。/请检查并确保您账户下的应用不是vpn应用。]</td></tr><tr align=\"center\"><td>50052835</td><td>dino factory</td><td>com.ohbibi.motorworlddinofactory</td><td>global</td><td>限制填充</td><td>否</td><td>[您的应用被识别为vpn应用，目前pangle不支持在vpn应用上变现。/请检查并确保您账户下的应用不是vpn应用。]</td></tr><tr align=\"center\"><td>50052833</td><td>digs & balls</td><td>com.quok.digsandballs</td><td>global</td><td>限制填充</td><td>否</td><td>[您的应用被识别为vpn应用，目前pangle不支持在vpn应用上变现。/请检查并确保您账户下的应用不是vpn应用。]</td></tr><tr align=\"center\"><td>50052494</td><td>stickman heroes: battle of god</td><td>com.stickman.heroes.battle.of.warriors.stickmanwarriors</td><t<br>d>global</td><td>停止填充</td><td>否</td><td>[您的应用被识别为vpn应用，目前pangle不支持在vpn应用上变现。/请检查并确保您账户下的应用不是vpn应用。]</td></tr><tr align=\"center\"><td>50052493</td><td>blade bouncer 2: revolution</td><td>com.threeswords.bladebouncerrevolution</td><td>global</td><td>停止填充</td><td>否</td><td>[您的应用被识别为vpn应用，目前pangle不支持在vpn应用上变现。/请检查并确保您账户下的应用不是vpn应用。]</td></tr><tr align=\"center\"><td>50051733</td><td>compare them</td><td>com.rsamy.comparethem</td><td>kr</td><td>停止填充</td><td>否</td><td>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]</td></tr><tr align=\"center\"><td>50053106</td><td>hair color match</td><td>hair.colormatch</td><td>il</td><td>停止填充</td><td>是</td><td>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作系统、系统版本等信息的准确性和集中度等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发<br>生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>50053111</td><td>geometry dash lite</td><td>com.robtopx.geometryjumplite</td><td>global</td><td>停止填充</td><td>是</td><td>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>50053370</td><td>idle hustle kingdom</td><td>com.tiamatgames.idlehustlekingdom</td><td>au</td><td>限制填充</td><td>否</td><td>[您的应用被识别为vpn应用，目前pangle不支持在vpn应用上变现。/请检查并确保您账户下的应用不是vpn应用。]</td></tr><tr align=\"center\"><td>50053144</td><td>metal madness pvp: car shooter</td><td>com.gdcompany.metalmadness</td><td>global</td><td>停止填充</td><td>否</td><td>[您的应用被识别为vpn应用，目前pangle不支持在vpn应用上变现。/请检查并确保您账户下的应用不是vpn应用。]</td></tr><tr align=\"center\"><td>50053140</td><td<br>>zombie hunter king</td><td>com.mobirix.zombieking</td><td>global</td><td>停止填充</td><td>否</td><td>[您的应用被识别为vpn应用，目前pangle不支持在vpn应用上变现。/请检查并确保您账户下的应用不是vpn应用。]</td></tr></table></p><br><ul><li>广告位：</li></ul><p class=\"union-media-editor__paragraph\"><table border = \"1\" cellspacing=\"0\" cellpadding=\"10px\"><tr align=\"center\"><th colspan=\"1\">广告位id</th><th colspan=\"1\">应用id</th><th colspan=\"1\">应用名称</th><th colspan=\"1\">包名</th><th colspan=\"1\">国家</th><th colspan=\"1\">处置方式</th><th colspan=\"1\">是否扣款</th><th colspan=\"1\">违规原因/建议</th></tr><tr align=\"center\"><td>980145488</td><td>50051874</td><td>stickman warriors legend fight</td><td>com.stickman.dragon.pvp.online</td><td>jp</td><td>停止填充</td><td>是</td><td>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]</td></tr><tr align=\"center\"><td>897345621</td><td>50051874</td><td>stickman warriors legend fight</td><td>com.stickman.dragon.pvp.online</td><td>kr</td><td>停止填充</td><td>否</td><td>[无效流量由流<br>量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980146467</td><td>50053369</td><td>idle casino manager - tycoon</td><td>com.coldfiregames.idle.casino.tycoon</td><td>jp</td><td>停止填充</td><td>是</td><td>[无效流量由流量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指标。]<br>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作系统、系统版本等信息的准确性和集中度等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980146470</td><td>50053369</td><td>idle casino manager - tycoon</td><td>com.coldfiregames.idle.casino.tycoon</td><td>kr,th</td><td>停止填充<br></td><td>是</td><td>[无效流量由流量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980146471</td><td>50053369</td><td>idle casino manager - tycoon</td><td>com.coldfiregames.idle.casino.tycoon</td><td>il,th</td><td>停止填充</td><td>是</td><td>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980145483</td><td>50051732</td><td>zombie tsunami</td><td>net.mobigame.zombietsunami</td><td>om</td><td>停止填充</td><td>是</td><td>[无效流量由流量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时<br>间、地域等指标。]</td></tr><tr align=\"center\"><td>980145482</td><td>50051733</td><td>compare them</td><td>com.rsamy.comparethem</td><td>ar</td><td>停止填充</td><td>是</td><td>[无效流量由流量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指标。]<br>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作系统、系统版本等信息的准确性和集中度等指标。]</td></tr><tr align=\"center\"><td>980145825</td><td>50052547</td><td>rare pets - merge game mystery</td><td>com.drihq.pets</td><td>th</td><td>停止填充</td><td>是</td><td>[无效流量由流量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指标。]<br>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作系统、系统版本等信息的准确性和集中度等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]</td></tr><tr align=\"center\"><td>980146467</td><td>50053369</td><td>idle casino manager - tycoon</td><td>com.col<br>dfiregames.idle.casino.tycoon</td><td>om,th</td><td>停止填充</td><td>是</td><td>[无效流量由流量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980146468</td><td>50053369</td><td>idle casino manager - tycoon</td><td>com.coldfiregames.idle.casino.tycoon</td><td>il</td><td>停止填充</td><td>是</td><td>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980146469</td><td>50053369</td><td>idle casino manager - tycoon</td><td>com.coldfiregames.idle.casino.tycoon</td><td>kr</td><td>停止填充</td><td>是</td><td>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，<br>如机型、操作系统、系统版本等信息的准确性和集中度等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980145820</td><td>50052547</td><td>rare pets - merge game mystery</td><td>com.drihq.pets</td><td>il</td><td>限制填充</td><td>否</td><td>[无效流量由流量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指标。]<br>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作系统、系统版本等信息的准确性和集中度等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]</td></tr><tr align=\"center\"><td>980145821</td><td>50052547</td><td>rare pets - merge game mystery</td><td>com.drihq.pets</td><td>ca</td><td>停止填充</td><td>是</td><td>[无效流量由流量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指标。]<br>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作<br>系统、系统版本等信息的准确性和集中度等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980145824</td><td>50052547</td><td>rare pets - merge game mystery</td><td>com.drihq.pets</td><td>om,th</td><td>停止填充</td><td>是</td><td>[无效流量由流量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980146468</td><td>50053369</td><td>idle casino manager - tycoon</td><td>com.coldfiregames.idle.casino.tycoon</td><td>kr</td><td>停止填充</td><td>是</td><td>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作系统、系统版本等信息的准确性和集中度等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量<br>中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980146469</td><td>50053369</td><td>idle casino manager - tycoon</td><td>com.coldfiregames.idle.casino.tycoon</td><td>il,jp,th</td><td>停止填充</td><td>是</td><td>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980145826</td><td>50052547</td><td>rare pets - merge game mystery</td><td>com.drihq.pets</td><td>ca</td><td>限制填充</td><td>是</td><td>[无效流量由流量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指标。]<br>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作系统、系统版本等信息的准确性和集中度等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980145825</td><td>50052547</td><td>rare pets - merge game mystery<br></td><td>com.drihq.pets</td><td>il,om</td><td>停止填充</td><td>否</td><td>[无效流量由流量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980146471</td><td>50053369</td><td>idle casino manager - tycoon</td><td>com.coldfiregames.idle.casino.tycoon</td><td>kr</td><td>停止填充</td><td>是</td><td>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作系统、系统版本等信息的准确性和集中度等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980145819</td><td>50052547</td><td>rare pets - merge game mystery</td><td>com.drihq.pets</td><td>om</td><td>停止填充</td><td>是</td><br><td>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作系统、系统版本等信息的准确性和集中度等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr><tr align=\"center\"><td>980146470</td><td>50053369</td><td>idle casino manager - tycoon</td><td>com.coldfiregames.idle.casino.tycoon</td><td>jp,om</td><td>停止填充</td><td>是</td><td>[无效流量由流量来源异常导致，如，ip地址集中或ip出现跨区域/地区等情况。/请仔细检查您流量中的地区分布集中度和跨地区情况等指标。]<br>[无效流量由设备参数异常导致。/请仔细检查您流量中的设备相关信息，如机型、操作系统、系统版本等信息的准确性和集中度等指标。]<br>[用户行为：由用户行为异常导致/请仔细检查您流量中的用户观看、点击、转化动作发生的时间、地域等指标。]<br>[转化效果/转化行为：由转化效果异常导致，留存等指标表现弱/请仔细检查您流量中的广告场景、跳转链路等]</td></tr></table></p><p class=\"union-media-editor__paragraph\">为了保护您和广告商的利益，您的处置将从即日起生效。如果有任何异议，请尽快提出申诉反馈。请参阅 pangle 帮助中心了解整个申诉流程：<a href=\"https://www.pangleglobal.com/zh/knowledge/appeal-procedure\" target=\"_blank\" rel=\"nofollow noreferrer\" class=\"union-media-editor__link\">https://www.pangleglobal.com/zh/knowledge/appeal-procedure</a></p><p class=\"union-media-editor__paragraph\">如果涉及扣款，则表明相应金额的款项将在结算时扣减，如后续申诉成功，扣减的金额将在下次结算时返还，若申诉失败，扣减金额将不会返还。</p><p class=\"union-media-editor__paragraph\">如不涉及扣款，您的款项将正常结算支付，但我们依然建议您对流量进行整改，防止广告填充持续被限制。</p><p class=\"union-media-editor__paragraph\"><br></p><p class=\"union-media-editor__paragraph\">pangle反作弊团队</p>`
	htmlWithLine := gohtml.Format(h)
	lines := strings.Split(htmlWithLine, "\n")
	width := 100
	result := ""
	newLine := ""
	for _, line := range lines {
		if len(newLine)+len(line) < width {
			newLine += line
		} else {
			result += newLine + "\n"
			newLine = line
		}
	}
	// fmt.Println(result)
	fmt.Println(gohtml.Format(h))
	fmt.Printf("%s", "\\n")
}
