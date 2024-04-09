package models

import (
	"fmt"

	common "github.com/freightcms/common/models"
)

// ISO 3166-1 alpha-2 country codes
type CountryCode string

// List of known country codes as of 2024-APR-07
// see https://en.wikipedia.org/wiki/ISO_3166-2
const (
	AD CountryCode = "AD" // Andorra
	AE CountryCode = "AE" // United Arab Emirates
	AF CountryCode = "AF" // Afghanistan
	AG CountryCode = "AG" // Antigua and Barbuda
	AI CountryCode = "AI" // Anguilla
	AL CountryCode = "AL" // Albania
	AM CountryCode = "AM" // Armenia
	AO CountryCode = "AO" // Angola
	AR CountryCode = "AR" // Argentina
	AS CountryCode = "AS" // American Samoa
	AT CountryCode = "AT" // Austria
	AU CountryCode = "AU" // Australia
	AW CountryCode = "AW" // Aruba
	AX CountryCode = "AX" // Åland Islands
	AZ CountryCode = "AZ" // Azerbaijan
	BA CountryCode = "BA" // Bosnia and Herzegovina
	BB CountryCode = "BB" // Barbados
	BD CountryCode = "BD" // Bangladesh
	BE CountryCode = "BE" // Belgium
	BF CountryCode = "BF" // Burkina Faso
	BG CountryCode = "BG" // Bulgaria
	BH CountryCode = "BH" // Bahrain
	BI CountryCode = "BI" // Burundi
	BJ CountryCode = "BJ" // Benin
	BL CountryCode = "BL" // Saint Barthélemy
	BM CountryCode = "BM" // Bermuda
	BN CountryCode = "BN" // Brunei Darussalam
	BO CountryCode = "BO" // Bolivia (Plurinational State of)
	BQ CountryCode = "BQ" // Bonaire, Sint Eustatius and Saba
	BR CountryCode = "BR" // Brazil
	BS CountryCode = "BS" // Bahamas
	BT CountryCode = "BT" // Bhutan
	BV CountryCode = "BV" // Bouvet Island
	BW CountryCode = "BW" // Botswana
	BY CountryCode = "BY" // Belarus
	BZ CountryCode = "BZ" // Belize
	CA CountryCode = "CA" // Canada
	CC CountryCode = "CC" // Cocos (Keeling) Islands
	CD CountryCode = "CD" // Congo, Democratic Republic of the
	CF CountryCode = "CF" // Central African Republic
	CH CountryCode = "CH" // Switzerland
	CI CountryCode = "CI" // Côte d'Ivoire
	CK CountryCode = "CK" // Cook Islands
	CL CountryCode = "CL" // Chile
	CM CountryCode = "CM" // Cameroon
	CN CountryCode = "CN" // China
	CO CountryCode = "CO" // Colombia
	CR CountryCode = "CR" // Costa Rica
	CU CountryCode = "CU" // Cuba
	CV CountryCode = "CV" // Cabo Verde
	CW CountryCode = "CW" // Curaçao
	CX CountryCode = "CX" // Christmas Island
	CY CountryCode = "CY" // Cyprus
	CZ CountryCode = "CZ" // Czechia
	DE CountryCode = "DE" // Germany
	DJ CountryCode = "DJ" // Djibouti
	DK CountryCode = "DK" // Denmark
	DM CountryCode = "DM" // Dominica
	DO CountryCode = "DO" // Dominican Republic
	DZ CountryCode = "DZ" // Algeria
	EC CountryCode = "EC" // Ecuador
	EE CountryCode = "EE" // Estonia
	EG CountryCode = "EG" // Egypt
	EH CountryCode = "EH" // Western Sahara
	ER CountryCode = "ER" // Eritrea
	ES CountryCode = "ES" // Spain
	ET CountryCode = "ET" // Ethiopia
	FI CountryCode = "FI" // Finland
	FJ CountryCode = "FJ" // Fiji
	FK CountryCode = "FK" // Falkland Islands (Malvinas)
	FM CountryCode = "FM" // Micronesia (Federated States of)
	FO CountryCode = "FO" // Faroe Islands
	FR CountryCode = "FR" // France
	GA CountryCode = "GA" // Gabon
	GB CountryCode = "GB" // United Kingdom of Great Britain and Northern Ireland
	GD CountryCode = "GD" // Grenada
	GE CountryCode = "GE" // Georgia
	GF CountryCode = "GF" // French Guiana
	GG CountryCode = "GG" // Guernsey
	GH CountryCode = "GH" // Ghana
	GI CountryCode = "GI" // Gibraltar
	GL CountryCode = "GL" // Greenland
	GM CountryCode = "GM" // Gambia
	GN CountryCode = "GN" // Guinea
	GP CountryCode = "GP" // Guadeloupe
	GQ CountryCode = "GQ" // Equatorial Guinea
	GR CountryCode = "GR" // Greece
	GS CountryCode = "GS" // South Georgia and the South Sandwich Islands
	GT CountryCode = "GT" // Guatemala
	GU CountryCode = "GU" // Guam
	GW CountryCode = "GW" // Guinea-Bissau
	GY CountryCode = "GY" // Guyana
	HK CountryCode = "HK" // Hong Kong
	HM CountryCode = "HM" // Heard Island and McDonald Islands
	HN CountryCode = "HN" // Honduras
	HR CountryCode = "HR" // Croatia
	HT CountryCode = "HT" // Haiti
	HU CountryCode = "HU" // Hungary
	ID CountryCode = "ID" // Indonesia
	IE CountryCode = "IE" // Ireland
	IL CountryCode = "IL" // Israel
	IM CountryCode = "IM" // Isle of Man
	IN CountryCode = "IN" // India
	IO CountryCode = "IO" // British Indian Ocean Territory
	IQ CountryCode = "IQ" // Iraq
	IR CountryCode = "IR" // Iran (Islamic Republic of)
	IS CountryCode = "IS" // Iceland
	IT CountryCode = "IT" // Italy
	JE CountryCode = "JE" // Jersey
	JM CountryCode = "JM" // Jamaica
	JO CountryCode = "JO" // Jordan
	JP CountryCode = "JP" // Japan
	KE CountryCode = "KE" // Kenya
	KG CountryCode = "KG" // Kyrgyzstan
	KH CountryCode = "KH" // Cambodia
	KI CountryCode = "KI" // Kiribati
	KM CountryCode = "KM" // Comoros
	KN CountryCode = "KN" // Saint Kitts and Nevis
	KP CountryCode = "KP" // Korea (Democratic People's Republic of)
	KR CountryCode = "KR" // Korea, Republic of
	KW CountryCode = "KW" // Kuwait
	KY CountryCode = "KY" // Cayman Islands
	KZ CountryCode = "KZ" // Kazakhstan
	LA CountryCode = "LA" // Lao People's Democratic Republic
	LB CountryCode = "LB" // Lebanon
	LC CountryCode = "LC" // Saint Lucia
	LI CountryCode = "LI" // Liechtenstein
	LK CountryCode = "LK" // Sri Lanka
	LR CountryCode = "LR" // Liberia
	LS CountryCode = "LS" // Lesotho
	LT CountryCode = "LT" // Lithuania
	LU CountryCode = "LU" // Luxembourg
	LV CountryCode = "LV" // Latvia
	LY CountryCode = "LY" // Libya
	MA CountryCode = "MA" // Morocco
	MC CountryCode = "MC" // Monaco
	MD CountryCode = "MD" // Moldova, Republic of
	ME CountryCode = "ME" // Montenegro
	MF CountryCode = "MF" // Saint Martin (French part)
	MG CountryCode = "MG" // Madagascar
	MH CountryCode = "MH" // Marshall Islands
	MK CountryCode = "MK" // North Macedonia
	ML CountryCode = "ML" // Mali
	MM CountryCode = "MM" // Myanmar
	MN CountryCode = "MN" // Mongolia
	MO CountryCode = "MO" // Macao
	MP CountryCode = "MP" // Northern Mariana Islands
	MQ CountryCode = "MQ" // Martinique
	MR CountryCode = "MR" // Mauritania
	MS CountryCode = "MS" // Montserrat
	MT CountryCode = "MT" // Malta
	MU CountryCode = "MU" // Mauritius
	MV CountryCode = "MV" // Maldives
	MW CountryCode = "MW" // Malawi
	MX CountryCode = "MX" // Mexico
	MY CountryCode = "MY" // Malaysia
	MZ CountryCode = "MZ" // Mozambique
	NA CountryCode = "NA" // Namibia
	NC CountryCode = "NC" // New Caledonia
	NE CountryCode = "NE" // Niger
	NF CountryCode = "NF" // Norfolk Island
	NG CountryCode = "NG" // Nigeria
	NI CountryCode = "NI" // Nicaragua
	NL CountryCode = "NL" // Netherlands
	NO CountryCode = "NO" // Norway
	NP CountryCode = "NP" // Nepal
	NR CountryCode = "NR" // Nauru
	NU CountryCode = "NU" // Niue
	NZ CountryCode = "NZ" // New Zealand
	OM CountryCode = "OM" // Oman
	PA CountryCode = "PA" // Panama
	PE CountryCode = "PE" // Peru
	PF CountryCode = "PF" // French Polynesia
	PG CountryCode = "PG" // Papua New Guinea
	PH CountryCode = "PH" // Philippines
	PK CountryCode = "PK" // Pakistan
	PL CountryCode = "PL" // Poland
	PM CountryCode = "PM" // Saint Pierre and Miquelon
	PN CountryCode = "PN" // Pitcairn
	PR CountryCode = "PR" // Puerto Rico
	PS CountryCode = "PS" // Palestine, State of
	PT CountryCode = "PT" // Portugal
	PW CountryCode = "PW" // Palau
	PY CountryCode = "PY" // Paraguay
	QA CountryCode = "QA" // Qatar
	RE CountryCode = "RE" // Réunion
	RO CountryCode = "RO" // Romania
	RS CountryCode = "RS" // Serbia
	RU CountryCode = "RU" // Russian Federation
	RW CountryCode = "RW" // Rwanda
	SA CountryCode = "SA" // Saudi Arabia
	SB CountryCode = "SB" // Solomon Islands
	SC CountryCode = "SC" // Seychelles
	SD CountryCode = "SD" // Sudan
	SE CountryCode = "SE" // Sweden
	SG CountryCode = "SG" // Singapore
	SH CountryCode = "SH" // Saint Helena, Ascension and Tristan da Cunha
	SI CountryCode = "SI" // Slovenia
	SJ CountryCode = "SJ" // Svalbard and Jan Mayen
	SK CountryCode = "SK" // Slovakia
	SL CountryCode = "SL" // Sierra Leone
	SM CountryCode = "SM" // San Marino
	SN CountryCode = "SN" // Senegal
	SO CountryCode = "SO" // Somalia
	SR CountryCode = "SR" // Suriname
	SS CountryCode = "SS" // South Sudan
	ST CountryCode = "ST" // Sao Tome and Principe
	SV CountryCode = "SV" // El Salvador
	SX CountryCode = "SX" // Sint Maarten (Dutch part)
	SY CountryCode = "SY" // Syrian Arab Republic
	SZ CountryCode = "SZ" // Eswatini
	TC CountryCode = "TC" // Turks and Caicos Islands
	TD CountryCode = "TD" // Chad
	TF CountryCode = "TF" // French Southern Territories
	TG CountryCode = "TG" // Togo
	TH CountryCode = "TH" // Thailand
	TJ CountryCode = "TJ" // Tajikistan
	TK CountryCode = "TK" // Tokelau
	TL CountryCode = "TL" // Timor-Leste
	TM CountryCode = "TM" // Turkmenistan
	TN CountryCode = "TN" // Tunisia
	TO CountryCode = "TO" // Tonga
	TR CountryCode = "TR" // Turkey
	TT CountryCode = "TT" // Trinidad and Tobago
	TV CountryCode = "TV" // Tuvalu
	TW CountryCode = "TW" // Taiwan, Province of China
	TZ CountryCode = "TZ" // Tanzania, United Republic of
	UA CountryCode = "UA" // Ukraine
	UG CountryCode = "UG" // Uganda
	UM CountryCode = "UM" // United States Minor Outlying Islands
	US CountryCode = "US" // United States of America
	UY CountryCode = "UY" // Uruguay
	UZ CountryCode = "UZ" // Uzbekistan
	VA CountryCode = "VA" // Holy See
	VC CountryCode = "VC" // Saint Vincent and the Grenadines
	VE CountryCode = "VE" // Venezuela (Bolivarian Republic of)
	VG CountryCode = "VG" // Virgin Islands (British)
	VI CountryCode = "VI" // Virgin Islands (U.S.)
	VN CountryCode = "VN" // Viet Nam
	VU CountryCode = "VU" // Vanuatu
	WF CountryCode = "WF" // Wallis and Futuna
	WS CountryCode = "WS" // Samoa
	YE CountryCode = "YE" // Yemen
	YT CountryCode = "YT" // Mayotte
	ZA CountryCode = "ZA" // South Africa
	ZM CountryCode = "ZM" // Zambia
	ZW CountryCode = "ZW" // Zimbabwe
)

var (
	ValidCountryCodes = []CountryCode{
		AD, AE, AF, AG, AI, AL, AM, AO, AR, AS,
		AT, AU, AW, AX, AZ, BA, BB, BD, BE, BF,
		BG, BH, BI, BJ, BL, BM, BN, BO, BQ, BR,
		BS, BT, BV, BW, BY, BZ, CA, CC, CD, CF,
		CH, CI, CK, CL, CM, CN, CO, CR, CU, CV,
		CW, CX, CY, CZ, DE, DJ, DK, DM, DO, DZ,
		EC, EE, EG, EH, ER, ES, ET, FI, FJ, FK,
		FM, FO, FR, GA, GB, GD, GE, GF, GG, GH,
		GI, GL, GM, GN, GP, GQ, GR, GS, GT, GU,
		GW, GY, HK, HM, HN, HR, HT, HU, ID, IE,
		IL, IM, IN, IO, IQ, IR, IS, IT, JE, JM,
		JO, JP, KE, KG, KH, KI, KM, KN, KP, KR,
		KW, KY, KZ, LA, LB, LC, LI, LK, LR, LS,
		LT, LU, LV, LY, MA, MC, MD, ME, MF, MG,
		MH, MK, ML, MM, MN, MO, MP, MQ, MR, MS,
		MT, MU, MV, MW, MX, MY, MZ, NA, NC, NE,
		NF, NG, NI, NL, NO, NP, NR, NU, NZ, OM,
		PA, PE, PF, PG, PH, PK, PL, PM, PN, PR,
		PS, PT, PW, PY, QA, RE, RO, RS, RU, RW,
		SA, SB, SC, SD, SE, SG, SH, SI, SJ, SK,
		SL, SM, SN, SO, SR, SS, ST, SV, SX, SY,
		SZ, TC, TD, TF, TG, TH, TJ, TK, TL, TM,
		TN, TO, TR, TT, TV, TW, TZ, UA, UG, UM,
		US, UY, UZ, VA, VC, VE, VG, VI, VN, VU,
		WF, WS, YE, YT, ZA, ZM, ZW,
	}
	ValidLanguageCodes = []LanguageCode{
		AAV, AFR, ALG, ALV, APA, AQL, ART, ATH, AUF, AUS,
		AWD, AZC, BAD, BAI, BAT, BER, BIH, BNT, BTK, CAL,
		CAU, CBA, CCN, CCS, CDC, CDD, CEL, CMC, CPE, CPF,
		CPP, CRP, CSU, CUS, DAY, DMN, DRA, EGX, EUQ, FIU,
		FOX, GEM, GME, GMQ, GMW, GRK, HMX, HOK, HYX, IIR,
		JPX, KAR, KDO, KHI, KRO, MAP, MKH, MNO, MUN, MYN,
		NAH, NAI, NGF, NIC, NUB, OMQ, OMV, OTO, PAA, PHI,
		PLF, POZ, PQE, PQW, PRA, QWE, ROA, SAI, SAL, SDV,
		SEM, SGN, SIO, SIT, SLA, SMI, SON, SQJ, SSA, SYD,
		TAI, TBQ, TRK, TUP, TUT, TUW, URJ, WAK, WEN, XGN,
		YPK, ZHX, ZLE, ZLS, ZLW, ZND,
	}
)

// LanguageCode ISO 639-5 language codes
// see https://en.wikipedia.org/wiki/List_of_ISO_639-5_codes
type LanguageCode string

const (
	AAV LanguageCode = "AAV" // Austro-Asiatic languages
	AFR LanguageCode = "AFR" // Afro-Asiatic languages
	ALG LanguageCode = "ALG" // Algonguian languages
	ALV LanguageCode = "ALV" // Atlantic-Congo languages
	APA LanguageCode = "APA" // Apachean languages
	AQL LanguageCode = "AQL" // Algic languages
	ART LanguageCode = "ART" // Artificial languages
	ATH LanguageCode = "ATH" // Athapaskan languages
	AUF LanguageCode = "AUF" // Austronesian languages
	AUS LanguageCode = "AUS" // Australian languages
	AWD LanguageCode = "AWD" // Arawakan languages
	AZC LanguageCode = "AZC" // Uto-Aztecan languages
	BAD LanguageCode = "BAD" // Banda languages
	BAI LanguageCode = "BAI" // Bamileke languages
	BAT LanguageCode = "BAT" // Baltic languages
	BER LanguageCode = "BER" // Berber languages
	BIH LanguageCode = "BIH" // Bihari languages
	BNT LanguageCode = "BNT" // Bantu languages
	BTK LanguageCode = "BTK" // Batak languages
	CAL LanguageCode = "CAL" // Cariban languages
	CAU LanguageCode = "CAU" // Caucasian languages
	CBA LanguageCode = "CBA" // Chibchan languages
	CCN LanguageCode = "CCN" // Chinookan languages
	CCS LanguageCode = "CCS" // Kordofanian languages
	CDC LanguageCode = "CDC" // Chadic languages
	CDD LanguageCode = "CDD" // Caddoan languages
	CEL LanguageCode = "CEL" // Celtic languages
	CMC LanguageCode = "CMC" // Chamic languages
	CPE LanguageCode = "CPE" // Creoles and pidgins, English based
	CPF LanguageCode = "CPF" // Creoles and pidgins, French-based
	CPP LanguageCode = "CPP" // Creoles and pidgins, Portuguese-based
	CRP LanguageCode = "CRP" // Creoles and pidgins
	CSU LanguageCode = "CSU" // Central Sudanic languages
	CUS LanguageCode = "CUS" // Cushitic languages
	DAY LanguageCode = "DAY" // Land Dayak languages
	DMN LanguageCode = "DMN" // Mande languages
	DRA LanguageCode = "DRA" // Dravidian languages
	EGX LanguageCode = "EGX" // Egyptian languages
	EUQ LanguageCode = "EUQ" // Basque languages
	FIU LanguageCode = "FIU" // Finno-Ugrian languages
	FOX LanguageCode = "FOX" // Formosan languages
	GEM LanguageCode = "GEM" // Germanic languages
	GME LanguageCode = "GME" // East Germanic languages
	GMQ LanguageCode = "GMQ" // North Germanic languages
	GMW LanguageCode = "GMW" // West Germanic languages
	GRK LanguageCode = "GRK" // Greek languages
	HMX LanguageCode = "HMX" // Hmong-Mien languages
	HOK LanguageCode = "HOK" // Hokan languages
	HYX LanguageCode = "HYX" // Armenian languages
	IIR LanguageCode = "IIR" // Indo-Iranian languages
	JPX LanguageCode = "JPX" // Japanese languages
	KAR LanguageCode = "KAR" // Karen languages
	KDO LanguageCode = "KDO" // Kadu languages
	KHI LanguageCode = "KHI" // Khoisan languages
	KRO LanguageCode = "KRO" // Kru languages
	MAP LanguageCode = "MAP" // Austronesian languages
	MKH LanguageCode = "MKH" // Mon-Khmer languages
	MNO LanguageCode = "MNO" // Manobo languages
	MUN LanguageCode = "MUN" // Munda languages
	MYN LanguageCode = "MYN" // Mayan languages
	NAH LanguageCode = "NAH" // Nahuatl languages
	NAI LanguageCode = "NAI" // North American Indian languages
	NGF LanguageCode = "NGF" // Trans-New Guinea languages
	NIC LanguageCode = "NIC" // Niger-Kordofanian languages
	NUB LanguageCode = "NUB" // Nubian languages
	OMQ LanguageCode = "OMQ" // Oto-Manguean languages
	OMV LanguageCode = "OMV" // Omotic languages
	OTO LanguageCode = "OTO" // Otomian languages
	PAA LanguageCode = "PAA" // Papuan languages
	PHI LanguageCode = "PHI" // Philippine languages
	PLF LanguageCode = "PLF" // Central Malayo-Polynesian languages
	POZ LanguageCode = "POZ" // Malayo-Polynesian languages
	PQE LanguageCode = "PQE" // Eastern Malayo-Polynesian languages
	PQW LanguageCode = "PQW" // Western Malayo-Polynesian languages
	PRA LanguageCode = "PRA" // Prakrit languages
	QWE LanguageCode = "QWE" // Quechuan languages
	ROA LanguageCode = "ROA" // Romance languages
	SAI LanguageCode = "SAI" // South American Indian languages
	SAL LanguageCode = "SAL" // Salishan languages
	SDV LanguageCode = "SDV" // Eastern Sudanic languages
	SEM LanguageCode = "SEM" // Semitic languages
	SGN LanguageCode = "SGN" // Sign languages
	SIO LanguageCode = "SIO" // Siouan languages
	SIT LanguageCode = "SIT" // Sino-Tibetan languages
	SLA LanguageCode = "SLA" // Slavic languages
	SMI LanguageCode = "SMI" // Sami languages
	SON LanguageCode = "SON" // Songhai languages
	SQJ LanguageCode = "SQJ" // Albanian languages
	SSA LanguageCode = "SSA" // Nilo-Saharan languages
	SYD LanguageCode = "SYD" // Samoyedic languages
	TAI LanguageCode = "TAI" // Tai-Kadai languages
	TBQ LanguageCode = "TBQ" // Tibeto-Burman languages
	TRK LanguageCode = "TRK" // Turkic languages
	TUP LanguageCode = "TUP" // Tupian languages
	TUT LanguageCode = "TUT" // Altaic languages
	TUW LanguageCode = "TUW" // Tungusic languages
	URJ LanguageCode = "URJ" // Uralic languages
	WAK LanguageCode = "WAK" // Wakashan languages
	WEN LanguageCode = "WEN" // Sorbian languages
	XGN LanguageCode = "XGN" // Mongolic languages
	YPK LanguageCode = "YPK" // Yupik languages
	ZHX LanguageCode = "ZHX" // Chinese languages
	ZLE LanguageCode = "ZLE" // Zenaga languages
	ZLS LanguageCode = "ZLS" // South Slavic languages
	ZLW LanguageCode = "ZLW" // West Slavic languages
	ZND LanguageCode = "ZND" // Zande languages
)

// Country Model meant for providing lookup data for countries. Optionally you extend the address model
// and add the country Id and override the Address.Country property with a reference to this model.
type Country struct {
	Id               string          `json:"id"`
	Name             string          `json:"name"`                       // Long display name of Country
	Code             CountryCode     `json:"code"`                       // ISO 3166-5 alpha-3 code
	Currency         common.Currency `json:"currency,omitempty"`         // Currency used in the country
	Language         LanguageCode    `json:"language,omitempty"`         // ISO 639-5 language code
	Region           string          `json:"region,omitempty"`           // Region of the country. E.g. Africa, Americas, Asia, Europe, Oceania
	SubRegion        string          `json:"subRegion,omitempty"`        // Subregion of the country. E.g. Southern Europe
	Capital          string          `json:"capital,omitempty"`          // Capital city of the country
	CallingCode      string          `json:"callingCode,omitempty"`      // International calling code
	Flag             string          `json:"flag,omitempty"`             // URL to flag image
	FlagEmoji        string          `json:"flagEmoji,omitempty"`        // URL to flag emoji
	FlagEmojiUnicode string          `json:"flagEmojiUnicode,omitempty"` // Unicode flag emoji
}

func (c *Country) Validate() []error {
	var (
		errors            []error
		foundCountryCode  = false
		foundLanguageCode = false
	)

	for _, code := range ValidCountryCodes {
		if c.Code == code {
			foundCountryCode = true
			break
		}
	}
	if !foundCountryCode {
		errors = append(errors, fmt.Errorf("invalid country code: %s", c.Code))
	}
	if len(c.Language) > 0 { // language is not required
		for _, code := range ValidLanguageCodes {
			if c.Language == code {
				foundLanguageCode = true
				break
			}
		}
		if !foundLanguageCode {
			errors = append(errors, fmt.Errorf("invalid language code: %s", c.Language))
		}
	}
	if c.Name == "" {
		errors = append(errors, fmt.Errorf("name is required"))
	}

	return errors
}
