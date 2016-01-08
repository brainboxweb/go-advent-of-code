package main

import (
	"testing"
)

/*
"" is 2 characters of code (the two double quotes), but the string contains zero characters.
"abc" is 5 characters of code, but 3 characters in the string data.
"aaa\"aaa" is 10 characters of code, but the string itself contains six "a" characters and a single, escaped quote character, for a total of 7 characters in the string data.
"\x27" is 6 characters of code, but the string itself contains just one - an apostrophe ('), escaped using hexadecimal notation.
*/

var test = []struct {
	input     string
	expected  int
	expected2 int
}{
	{
		`""`,
		2,
		0,
	},
	{
		`"abc"`,
		5,
		3,
	},
	{
		`"aaa\"aaa"`,
		10,
		7,
	},
	{
		`"\x27"`,
		6,
		1,
	},
}

func TestStrings(t *testing.T) {
	for _, test := range test {

		actual, actual2 := stringInfo(test.input)

		if actual != test.expected {
			t.Errorf("Parse(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}
		if actual2 != test.expected2 {
			t.Errorf("Parse(%s) = %s, expected %s.",
				test.input, actual2, test.expected2)
		}
	}
}

var test2 = []struct {
	input    string
	expected int
}{
	{
		`""
""`,
		4,
	},
	{
		`"abc"
"abc"`,
		4,
	},
	{
		day8data,
		1333,
	},
}

func TestRun(t *testing.T) {
	for _, test := range test2 {
		actual := Run(test.input)
		if actual != test.expected {
			t.Errorf("Parse(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}
	}
}

//"" encodes to "\"\"", an increase from 2 characters to 6.
//"abc" encodes to "\"abc\"", an increase from 5 characters to 9.
//"aaa\"aaa" encodes to "\"aaa\\\"aaa\"", an increase from 10 characters to 16.
//"\x27" encodes to "\"\\x27\"", an increase from 6 characters to 11.

var test3 = []struct {
	input     string
	expected  int
	expected2 int
}{
	{
		`""`,
		2,
		6,
	},
	{
		`"abc"`,
		5,
		9,
	},
	{
		`"aaa\"aaa"`,
		10,
		16,
	},
	{
		`"\x27"`,
		6,
		11,
	},
}

func TestStrings2(t *testing.T) {
	for _, test := range test3 {

		actual, actual2 := stringInfo2(test.input)

		if actual != test.expected {
			t.Errorf("Parse(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}
		if actual2 != test.expected2 {
			t.Errorf("Parse(%s) = %s, expected %s.",
				test.input, actual2, test.expected2)
		}
	}
}

var test4 = []struct {
	input    string
	expected int
}{
	{
		day8data,
		2046,
	},
}

func TestRun2(t *testing.T) {
	for _, test := range test4 {
		actual := run2(test.input)
		if actual != test.expected {
			t.Errorf("Parse(%s) = %s, expected %s.",
				test.input, actual, test.expected)
		}
	}
}

const day8data = `"sjdivfriyaaqa\xd2v\"k\"mpcu\"yyu\"en"
"vcqc"
"zbcwgmbpijcxu\"yins\"sfxn"
"yumngprx"
"bbdj"
"czbggabkzo\"wsnw\"voklp\"s"
"acwt"
"aqttwnsohbzian\"evtllfxwkog\"cunzw"
"ugvsgfv"
"xlnillibxg"
"kexh\"pmi"
"syvugow"
"m\"ktqnw"
"yrbajyndte\\rm"
"f\"kak\x70sn\xc4kjri"
"yxthr"
"alvumfsjni\"kohg"
"trajs\x5brom\xf1yoijaumkem\"\"tahlzs"
"\"oedr\"pwdbnnrc"
"qsmzhnx\""
"\"msoytqimx\\tbklqz"
"mjdfcgwdshrehgs"
"\"rivyxahf\""
"ciagc\x04bp"
"xkfc"
"xrgcripdu\x4c\xc4gszjhrvumvz\"mngbirb"
"gvmae\"yiiujoqvr\"mkxmgbbut\"u"
"ih"
"ncrqlejehs"
"mkno\x43pcfdukmemycp"
"uanzoqxkpsksbvdnkji\"feamp"
"axoufpnbx\\ao\x61pfj\"b"
"dz\\ztawzdjy"
"ihne\"enumvswypgf"
"\"dgazthrphbshdo\\vuqoiy\""
"dlnmptzt\\zahwpylc\\b\"gmslrqysk"
"mhxznyzcp"
"rebr\"amvxw\x5fmbnfpkkeghlntavj"
"lades\x47ncgdof\"\"jmbbk"
"dwxuis\xa5wdkx\\z\"admgnoddpgkt\\zs"
"g\\k\x27qsl\x34hwfglcdxqbeclt\xca\\"
"lhyjky\\m\"pvnm\\xmynpxnlhndmahjl"
"c\"uxabbgorrpprw\"xas\\vefkxioqpt"
"rfrvjxpevcmma\x71gtfipo"
"fgh\"kcwoqwfnjgdlzfclprg\"q"
"onxnwykrba"
"hkkg\x60f\"tjzsanpvarzgkfipl"
"\"aintes\"ofq\"juiaqlqxmvpe\\a"
"wiyczzs\"ciwk"
"mfqeu"
"v\xe1z\x7ftzalmvdmncfivrax\\rjwq"
"k\"vtg"
"exhrtdugeml\xf0"
"behnchkpld"
"mhgxy\"mfcrg\xc5gnp\"\"osqhj"
"rlvjy"
"awe"
"ctwy"
"vt"
"\x54t"
"zugfmmfomz"
"cv\"cvcvfaada\x04fsuqjinbfh\xa9cq\xd2c\"d"
"oj"
"xazanf\"wbmcrn"
"\\\\zkisyjpbzandqikqjqvee"
"dpsnbzdwnxk\\v"
"sj\"tuupr\\oyoh"
"myvkgnw\x81q\xaaokt\\emgejbsyvxcl\\\xee"
"ejeuqvunjcirdkkpt\"nlns"
"twmlvwxyvfyqqzu"
"\"xwtzdp\x98qkcis\"dm\\\"ep\"xyykq"
"vvcq\\expok"
"wgukjfanjgpdjb"
"\"mjcjajnxy\\dcpc"
"wdvgnecw\\ab\x44klceduzgsvu"
"dqtqkukr\"iacngufbqkdpxlwjjt"
"\"xj\"\x66qofsqzkoah"
"nptiwwsqdep"
"gsnlxql\x30mjl"
"yeezwokjwrhelny\""
"bjauamn\\izpmzqqasid"
"tvjdbkn\"tiziw\x82r"
"w"
"xwoakbbnjnypnaa\xa9wft\"slrmoqkl"
"vwxtnlvaaasyruykgygrvpiopzygf\"vq"
"qdancvnvmhlmpj\\isdxs"
"xzc\\elw"
"b\"wxeqvy\"qf\"g\xcaoklsucwicyw\"dovr"
"yomlvvjdbngz\"rly\"afr"
"bfb\"x\"aweuwbwmoa\x13\"t\"zhr"
"\"dmfoxb\"qvpjzzhykt\xd2\"\"ryhxi"
"psqef\"yu\\qiflie\"\x79w"
"arzewkej\"lqmh\\sayyusxxo\\"
"vuvvp"
"hc\"lg\x6bcpupsewzklai\"l"
"cjdfygc\"auorqybnuqghsh\x10"
"j"
"wqjexk\"eyq\\lbroqhk\\dqzsqk"
"dws\"ru\"dvxfiwapif\"oqwzmle"
"agcykg\\jt\\vzklqjvknoe"
"kksd\"jmslja\\z\"y\\b\xaagpyojct"
"nnpipxufvbfpoz\"jno"
"dtw"
"xlolvtahvgqkx\\dgnhj\\spsclpcxv\\"
"mxea\\mbjpi"
"lgbotkk\"zmxh\\\\qji\"jszulnjsxkqf"
"lwckmhwhx\"gmftlb\x91am"
"xxdxqyxth"
"\"lmqhwkjxmvayxy"
"tf"
"qy"
"wdqmwxdztax\"m\"\x09\x11xdxmfwxmtqgwvf"
"\xcbnazlf\"ghziknszmsrahaf"
"e\x6aupmzhxlvwympgjjpdvo\"kylfa"
"\x81vhtlillb\xactgoatva"
"dvnlgr"
"f"
"xg\xfacwizsadgeclm"
"vnnrzbtw\"\\prod\\djbyppngwayy\""
"lrt\xf4jahwvfz"
"aqpnjtom\"ymkak\\dadfybqrso\\fwv"
"gz\"aac\"mrbk\"ktommrojraqh"
"wycamwoecsftepfnlcdkm"
"nrhddblbuzlqsl\x9cben"
"vckxhyqkmqmdseazcykrbysm"
"sil\xbbtevmt\"gvrvybui\"faw\"j"
"cjex\\tp\x45pzf"
"asjobvtxszfodgf\"ibftg"
"gkyjyjdrxdcllnh\"sjcibenrdnxv"
"oswsdpjyxpbwnqbcpl\"yrdvs\\zq"
"\"\"tyowzc\\fycbp\"jbwrbvgui"
"cbpcabqkdgzmpgcwjtrchxp"
"iyrzfh\x45gw\"fdlfpiaap\x31xqq"
"evgksznidz"
"b\\w\\"
"loufizbiy\x57aim\"bgk"
"qjfyk"
"g\"anmloghvgr\x07zwqougqhdz"
"usbbmwcxd\\bdgg"
"htitqcpczml"
"eke\\cqvpexqqk\"to\"tqmljrpn\xe6lji\""
"g\xd2ifdsej"
"h\"sk\"haajajpagtcqnzrfqn\xe6btzo"
"wfkuffdxlvm\\cvlyzlbyunclhmpp"
"myaavh\"spue"
"hqvez\x68d\"eo\"eaioh"
"s\"qd\"oyxxcglcdnuhk"
"ilqvar"
"srh"
"puuifxrfmpc\"bvalwi\x2blu\\"
"yywlbutufzysbncw\\nqsfbhpz\"mngjq"
"zbl\\jfcuop"
"hjdouiragzvxsqkreup\\"
"qi"
"ckx\\funlj\xa7ahi"
"k"
"ufrcnh\"ajteit"
"cqv\"bgjozjj\x60x\xa8yhvmdvutchjotyuz"
"hkuiet\"oku\x8cfhumfpasl"
"\"\\sbe\x4d"
"vhknazqt"
"eyyizvzcahgflvmoowvs\\jhvygci"
"kki\x3ewcefkgtjap\"xtpxh\"lzepoqj"
"wvtk"
"\"ynet"
"zh\\obk\"otagx\x59txfzf"
"ocowhxlx\xe6zqg\x63wx\\tclkhq\\vmaze"
"w\"cf"
"qpniprnrzrnvykghqnalr"
"jctcqra\"\x05dhlydpqamorqjsijt\\xjdgt"
"sig"
"qhlbidbflwxe\"xljbwls\x20vht"
"irmrebfla\xefsg\"j"
"nep"
"hjuvsqlizeqobepf"
"guzbcdp\"obyh"
"\"mjagins\xf9tqykaxy\""
"knvsdnmtr\"zervsb"
"hzuy"
"zza\"k\"buapb\\elm\xfeya"
"lrqar\"dfqwkaaqifig\"uixjsz"
"\"azuo\x40rmnlhhluwsbbdb\x32pk\\yu\"pbcf"
"dplkdyty"
"rfoyciebwlwphcycmguc"
"ivnmmiemhgytmlprq\\eh"
"lhkyzaaothfdhmbpsqd\\yyw"
"tnlzifupcjcaj"
"\\qiyirsdrfpmu\\\x15xusifaag"
"\\lcomf\\s"
"uramjivcirjhqcqcg"
"kkbaklbxfxikffnuhtu\xc6t\"d"
"n\xefai"
"\"toy\"bnbpevuzoc\"muywq\"gz\"grbm"
"\"muu\\wt"
"\\srby\"ee"
"erf\"gvw\"swfppf"
"pbqcgtn\"iuianhcdazfvmidn\\nslhxdf"
"uxbp"
"up\\mgrcyaegiwmjufn"
"nulscgcewj\\dvoyvhetdegzhs\""
"masv\"k\\rzrb"
"qtx\x79d\"xdxmbxrvhj"
"fid\\otpkgjlh\"qgsvexrckqtn\xf4"
"tagzu"
"bvl\\\"noseec"
"\\xgicuuh"
"w\"a\"npemf"
"sxp"
"nsmpktic\x8awxftscdcvijjobnq\"gjd"
"uks\"\"jxvyvfezz\"aynxoev\"cuoav"
"m"
"lkvokj"
"vkfam\"yllr\"q\x92o\x4ebecnvhshhqe\\"
"efdxcjkjverw"
"lmqzadwhfdgmep\x02tzfcbgrbfekhat"
"cpbk\x9azqegbpluczssouop\x36ztpuoxsw"
"cqwoczxdd\"erdjka"
"cwvqnjgbw\\fxdlby"
"mvtm"
"lt\"bbqzpumplkg"
"ntd\xeeuwweucnuuslqfzfq"
"y\xabl\"dbebxjrlbmuoo\\\x1au"
"qjoqx\\a"
"pu\"ekdnfpmly\xbago\""
"fjhhdy"
"arl"
"xcywisim\"bwuwf\"\"raepeawwjub"
"pbe"
"dbnqfpzyaumxtqnd\xc5dcqrkwyop"
"ojv\x40vtkwgkqepm\x8bzft\\vedrry"
"wggqkfbwqumsgajqwphjec\"mstxpwz"
"zjkbem"
"icpfqxbelxazlls"
"pvpqs\\abcmtyielugfgcv\"tjxapxqxnx"
"oqddwlvmtv\"\x39lyybylfb\"jmngnpjrdw"
"gisgbve"
"\"aglg"
"y\"\"ss\xafvhxlrjv"
"qbgqjsra"
"ihshbjgqpdcljpmdwdprwloy"
"djja\\wcdn\"svkrgpqn\"uz\"hc\x43hj"
"cbjm"
"pnn"
"pqvh\"noh"
"\"\\fdktlp"
"ncea"
"pqgzphiyy"
"\xbedovhxuipaohlcvkwtxwmpz\"ckaif\"r"
"arjuzbjowqciunfwgxtph\"vlhy\"n"
"c"
"nrpdxunulgudqzlhtae"
"iefheu\"uru\""
"aqijysxuijud\"np\\opbichhudil\xbesum"
"pfpevmtstl\"lde\"bzr\"vspdxs"
"vparfbdjwvzsocpnzhp"
"g\x4ffxaarafrsjthq\\\xc1rw"
"ng\\rqx\\gwpzucbh\xafl"
"rw\"nf\\dna"
"jkkeahxurxla\\g\xb3czrlsyimmwcwthr"
"twaailoypu\"oas\"kpuuyedlaw\\\xb0vzt"
"hznex\\gdiqvtugi"
"imdibsunjeswhk"
"ta\\icileuzpxro\"cfmv\"mzp"
"coykr\x57luiysucfaflmilhlehmvzeiepo"
"u\x3dfh\xd4yt"
"piw\x1bz\"eowy\"vfk\"wqiekw"
"gan\"y"
"p\"bevidoazcznr\"hddxuuq\""
"bwzucczznutbxe"
"z\"viqgyqjisior\\iecosmjbknol"
"dmlpcglcfkfsctxydjvayhymv\x3c\\gp"
"bfvkqrintbbvgfv"
"xlzntrgdck\"cprc\xadczyarbznqmuhxyuh"
"uqdxnuwioc\"kdytxq\\ig"
"xrafmucpmfi"
"vr\"hltmfrge"
"eonf\"nt\\wtcnsocs"
"j\xb7xoslyjeyjksplkqixncgkylkw"
"njw\"pefgfbez\x9axshdmplxzquqe"
"di\x58bvptfsafirpc"
"l\x1fkco"
"x"
"mprndo\"n"
"psegit"
"svbdnkkuuqs\"sqxu\"oqcyz\"aizashk"
"cwkljukxer\\\"\\nff\"esjwiyaoy"
"ilxrkgbjjxpvhdtq\"cpiuoofdnkpp"
"hlngi\"ulxep\\qohtmqnqjb\"rkgerho"
"gxws\"bcgm\"p"
"bv\"mds\\zhfusiepgrz\\b\x32fscdzz"
"l\xfampwtme\x69qvxnx\"\"\xc4jruuymjxrpsv"
"qqmxhrn"
"xziq\\\x18ybyv\x9am\"neacoqjzytertisysza"
"aqcbvlvcrzceeyx\\j\"\"x"
"yjuhhb"
"\x5em\"squulpy"
"dpbntplgmwb"
"utsgfkm\\vbftjknlktpthoeo"
"ccxjgiocmuhf\"ycnh"
"lltj\"kbbxi"`
