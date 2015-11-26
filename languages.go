package geoutils

import (
    "strings"
)

type Language struct {
    Name        string
    NativeName  string
}

func ListLanguages() map[string]Language {
    list := map[string]Language{
        "ab": Language{
            Name:       "Abkhaz",
            NativeName: "аҧсуа",
        },
        "aa": Language{
            Name:       "Afar",
            NativeName: "Afaraf",
        },
        "af": Language{
            Name:       "Afrikaans",
            NativeName: "Afrikaans",
        },
        "ak": Language{
            Name:       "Akan",
            NativeName: "Akan",
        },
        "sq": Language{
            Name:       "Albanian",
            NativeName: "Shqip",
        },
        "am": Language{
            Name:       "Amharic",
            NativeName: "አማርኛ",
        },
        "ar": Language{
            Name:       "Arabic",
            NativeName: "العربية",
        },
        "an": Language{
            Name:       "Aragonese",
            NativeName: "Aragonés",
        },
        "hy": Language{
            Name:       "Armenian",
            NativeName: "Հայերեն",
        },
        "as": Language{
            Name:       "Assamese",
            NativeName: "অসমীয়া",
        },
        "av": Language{
            Name:       "Avaric",
            NativeName: "авар мацӀ, магӀарул мацӀ",
        },
        "ae": Language{
            Name:       "Avestan",
            NativeName: "avesta",
        },
        "ay": Language{
            Name:       "Aymara",
            NativeName: "aymar aru",
        },
        "az": Language{
            Name:       "Azerbaijani",
            NativeName: "azərbaycan dili",
        },
        "bm": Language{
            Name:       "Bambara",
            NativeName: "bamanankan",
        },
        "ba": Language{
            Name:       "Bashkir",
            NativeName: "башҡорт теле",
        },
        "eu": Language{
            Name:       "Basque",
            NativeName: "euskara, euskera",
        },
        "be": Language{
            Name:       "Belarusian",
            NativeName: "Беларуская",
        },
        "bn": Language{
            Name:       "Bengali",
            NativeName: "বাংলা",
        },
        "bh": Language{
            Name:       "Bihari",
            NativeName: "भोजपुरी",
        },
        "bi": Language{
            Name:       "Bislama",
            NativeName: "Bislama",
        },
        "bs": Language{
            Name:       "Bosnian",
            NativeName: "bosanski jezik",
        },
        "br": Language{
            Name:       "Breton",
            NativeName: "brezhoneg",
        },
        "bg": Language{
            Name:       "Bulgarian",
            NativeName: "български език",
        },
        "my": Language{
            Name:       "Burmese",
            NativeName: "ဗမာစာ",
        },
        "ca": Language{
            Name:       "Catalan; Valencian",
            NativeName: "Català",
        },
        "ch": Language{
            Name:       "Chamorro",
            NativeName: "Chamoru",
        },
        "ce": Language{
            Name:       "Chechen",
            NativeName: "нохчийн мотт",
        },
        "ny": Language{
            Name:       "Chichewa; Chewa; Nyanja",
            NativeName: "chiCheŵa, chinyanja",
        },
        "zh": Language{
            Name:       "Chinese",
            NativeName: "中文 (Zhōngwén), 汉语, 漢語",
        },
        "cv": Language{
            Name:       "Chuvash",
            NativeName: "чӑваш чӗлхи",
        },
        "kw": Language{
            Name:       "Cornish",
            NativeName: "Kernewek",
        },
        "co": Language{
            Name:       "Corsican",
            NativeName: "corsu, lingua corsa",
        },
        "cr": Language{
            Name:       "Cree",
            NativeName: "ᓀᐦᐃᔭᐍᐏᐣ",
        },
        "hr": Language{
            Name:       "Croatian",
            NativeName: "hrvatski",
        },
        "cs": Language{
            Name:       "Czech",
            NativeName: "česky, čeština",
        },
        "da": Language{
            Name:       "Danish",
            NativeName: "dansk",
        },
        "dv": Language{
            Name:       "Divehi; Dhivehi; Maldivian;",
            NativeName: "ދިވެހި",
        },
        "nl": Language{
            Name:       "Dutch",
            NativeName: "Nederlands, Vlaams",
        },
        "en": Language{
            Name:       "English",
            NativeName: "English",
        },
        "eo": Language{
            Name:       "Esperanto",
            NativeName: "Esperanto",
        },
        "et": Language{
            Name:       "Estonian",
            NativeName: "eesti, eesti keel",
        },
        "ee": Language{
            Name:       "Ewe",
            NativeName: "Eʋegbe",
        },
        "fo": Language{
            Name:       "Faroese",
            NativeName: "føroyskt",
        },
        "fj": Language{
            Name:       "Fijian",
            NativeName: "vosa Vakaviti",
        },
        "fi": Language{
            Name:       "Finnish",
            NativeName: "suomi, suomen kieli",
        },
        "fr": Language{
            Name:       "French",
            NativeName: "français, langue française",
        },
        "ff": Language{
            Name:       "Fula; Fulah; Pulaar; Pular",
            NativeName: "Fulfulde, Pulaar, Pular",
        },
        "gl": Language{
            Name:       "Galician",
            NativeName: "Galego",
        },
        "ka": Language{
            Name:       "Georgian",
            NativeName: "ქართული",
        },
        "de": Language{
            Name:       "German",
            NativeName: "Deutsch",
        },
        "el": Language{
            Name:       "Greek, Modern",
            NativeName: "Ελληνικά",
        },
        "gn": Language{
            Name:       "Guaraní",
            NativeName: "Avañeẽ",
        },
        "gu": Language{
            Name:       "Gujarati",
            NativeName: "ગુજરાતી",
        },
        "ht": Language{
            Name:       "Haitian; Haitian Creole",
            NativeName: "Kreyòl ayisyen",
        },
        "ha": Language{
            Name:       "Hausa",
            NativeName: "Hausa, هَوُسَ",
        },
        "he": Language{
            Name:       "Hebrew (modern)",
            NativeName: "עברית",
        },
        "hz": Language{
            Name:       "Herero",
            NativeName: "Otjiherero",
        },
        "hi": Language{
            Name:       "Hindi",
            NativeName: "हिन्दी, हिंदी",
        },
        "ho": Language{
            Name:       "Hiri Motu",
            NativeName: "Hiri Motu",
        },
        "hu": Language{
            Name:       "Hungarian",
            NativeName: "Magyar",
        },
        "ia": Language{
            Name:       "Interlingua",
            NativeName: "Interlingua",
        },
        "id": Language{
            Name:       "Indonesian",
            NativeName: "Bahasa Indonesia",
        },
        "ie": Language{
            Name:       "Interlingue",
            NativeName: "Originally called Occidental; then Interlingue after WWII",
        },
        "ga": Language{
            Name:       "Irish",
            NativeName: "Gaeilge",
        },
        "ig": Language{
            Name:       "Igbo",
            NativeName: "Asụsụ Igbo",
        },
        "ik": Language{
            Name:       "Inupiaq",
            NativeName: "Iñupiaq, Iñupiatun",
        },
        "io": Language{
            Name:       "Ido",
            NativeName: "Ido",
        },
        "is": Language{
            Name:       "Icelandic",
            NativeName: "Íslenska",
        },
        "it": Language{
            Name:       "Italian",
            NativeName: "Italiano",
        },
        "iu": Language{
            Name:       "Inuktitut",
            NativeName: "ᐃᓄᒃᑎᑐᑦ",
        },
        "ja": Language{
            Name:       "Japanese",
            NativeName: "日本語 (にほんご／にっぽんご)",
        },
        "jv": Language{
            Name:       "Javanese",
            NativeName: "basa Jawa",
        },
        "kl": Language{
            Name:       "Kalaallisut, Greenlandic",
            NativeName: "kalaallisut, kalaallit oqaasii",
        },
        "kn": Language{
            Name:       "Kannada",
            NativeName: "ಕನ್ನಡ",
        },
        "kr": Language{
            Name:       "Kanuri",
            NativeName: "Kanuri",
        },
        "ks": Language{
            Name:       "Kashmiri",
            NativeName: "कश्मीरी, كشميري‎",
        },
        "kk": Language{
            Name:       "Kazakh",
            NativeName: "Қазақ тілі",
        },
        "km": Language{
            Name:       "Khmer",
            NativeName: "ភាសាខ្មែរ",
        },
        "ki": Language{
            Name:       "Kikuyu, Gikuyu",
            NativeName: "Gĩkũyũ",
        },
        "rw": Language{
            Name:       "Kinyarwanda",
            NativeName: "Ikinyarwanda",
        },
        "ky": Language{
            Name:       "Kirghiz, Kyrgyz",
            NativeName: "кыргыз тили",
        },
        "kv": Language{
            Name:       "Komi",
            NativeName: "коми кыв",
        },
        "kg": Language{
            Name:       "Kongo",
            NativeName: "KiKongo",
        },
        "ko": Language{
            Name:       "Korean",
            NativeName: "한국어 (韓國語), 조선말 (朝鮮語)",
        },
        "ku": Language{
            Name:       "Kurdish",
            NativeName: "Kurdî, كوردی‎",
        },
        "kj": Language{
            Name:       "Kwanyama, Kuanyama",
            NativeName: "Kuanyama",
        },
        "la": Language{
            Name:       "Latin",
            NativeName: "latine, lingua latina",
        },
        "lb": Language{
            Name:       "Luxembourgish, Letzeburgesch",
            NativeName: "Lëtzebuergesch",
        },
        "lg": Language{
            Name:       "Luganda",
            NativeName: "Luganda",
        },
        "li": Language{
            Name:       "Limburgish, Limburgan, Limburger",
            NativeName: "Limburgs",
        },
        "ln": Language{
            Name:       "Lingala",
            NativeName: "Lingála",
        },
        "lo": Language{
            Name:       "Lao",
            NativeName: "ພາສາລາວ",
        },
        "lt": Language{
            Name:       "Lithuanian",
            NativeName: "lietuvių kalba",
        },
        "lu": Language{
            Name:       "Luba-Katanga",
            NativeName: "",
        },
        "lv": Language{
            Name:       "Latvian",
            NativeName: "latviešu valoda",
        },
        "gv": Language{
            Name:       "Manx",
            NativeName: "Gaelg, Gailck",
        },
        "mk": Language{
            Name:       "Macedonian",
            NativeName: "македонски јазик",
        },
        "mg": Language{
            Name:       "Malagasy",
            NativeName: "Malagasy fiteny",
        },
        "ms": Language{
            Name:       "Malay",
            NativeName: "bahasa Melayu, بهاس ملايو‎",
        },
        "ml": Language{
            Name:       "Malayalam",
            NativeName: "മലയാളം",
        },
        "mt": Language{
            Name:       "Maltese",
            NativeName: "Malti",
        },
        "mi": Language{
            Name:       "Māori",
            NativeName: "te reo Māori",
        },
        "mr": Language{
            Name:       "Marathi (Marāṭhī)",
            NativeName: "मराठी",
        },
        "mh": Language{
            Name:       "Marshallese",
            NativeName: "Kajin M̧ajeļ",
        },
        "mn": Language{
            Name:       "Mongolian",
            NativeName: "монгол",
        },
        "na": Language{
            Name:       "Nauru",
            NativeName: "Ekakairũ Naoero",
        },
        "nv": Language{
            Name:       "Navajo, Navaho",
            NativeName: "Diné bizaad, Dinékʼehǰí",
        },
        "nb": Language{
            Name:       "Norwegian Bokmål",
            NativeName: "Norsk bokmål",
        },
        "nd": Language{
            Name:       "North Ndebele",
            NativeName: "isiNdebele",
        },
        "ne": Language{
            Name:       "Nepali",
            NativeName: "नेपाली",
        },
        "ng": Language{
            Name:       "Ndonga",
            NativeName: "Owambo",
        },
        "nn": Language{
            Name:       "Norwegian Nynorsk",
            NativeName: "Norsk nynorsk",
        },
        "no": Language{
            Name:       "Norwegian",
            NativeName: "Norsk",
        },
        "ii": Language{
            Name:       "Nuosu",
            NativeName: "ꆈꌠ꒿ Nuosuhxop",
        },
        "nr": Language{
            Name:       "South Ndebele",
            NativeName: "isiNdebele",
        },
        "oc": Language{
            Name:       "Occitan",
            NativeName: "Occitan",
        },
        "oj": Language{
            Name:       "Ojibwe, Ojibwa",
            NativeName: "ᐊᓂᔑᓈᐯᒧᐎᓐ",
        },
        "cu": Language{
            Name:       "Old Church Slavonic, Church Slavic, Church Slavonic, Old Bulgarian, Old Slavonic",
            NativeName: "ѩзыкъ словѣньскъ",
        },
        "om": Language{
            Name:       "Oromo",
            NativeName: "Afaan Oromoo",
        },
        "or": Language{
            Name:       "Oriya",
            NativeName: "ଓଡ଼ିଆ",
        },
        "os": Language{
            Name:       "Ossetian, Ossetic",
            NativeName: "ирон æвзаг",
        },
        "pa": Language{
            Name:       "Panjabi, Punjabi",
            NativeName: "ਪੰਜਾਬੀ, پنجابی‎",
        },
        "pi": Language{
            Name:       "Pāli",
            NativeName: "पाऴि",
        },
        "fa": Language{
            Name:       "Persian",
            NativeName: "فارسی",
        },
        "pl": Language{
            Name:       "Polish",
            NativeName: "polski",
        },
        "ps": Language{
            Name:       "Pashto, Pushto",
            NativeName: "پښتو",
        },
        "pt": Language{
            Name:       "Portuguese",
            NativeName: "Português",
        },
        "qu": Language{
            Name:       "Quechua",
            NativeName: "Runa Simi, Kichwa",
        },
        "rm": Language{
            Name:       "Romansh",
            NativeName: "rumantsch grischun",
        },
        "rn": Language{
            Name:       "Kirundi",
            NativeName: "kiRundi",
        },
        "ro": Language{
            Name:       "Romanian, Moldavian, Moldovan",
            NativeName: "română",
        },
        "ru": Language{
            Name:       "Russian",
            NativeName: "русский язык",
        },
        "sa": Language{
            Name:       "Sanskrit (Saṁskṛta)",
            NativeName: "संस्कृतम्",
        },
        "sc": Language{
            Name:       "Sardinian",
            NativeName: "sardu",
        },
        "sd": Language{
            Name:       "Sindhi",
            NativeName: "सिन्धी, سنڌي، سندھی‎",
        },
        "se": Language{
            Name:       "Northern Sami",
            NativeName: "Davvisámegiella",
        },
        "sm": Language{
            Name:       "Samoan",
            NativeName: "gagana faa Samoa",
        },
        "sg": Language{
            Name:       "Sango",
            NativeName: "yângâ tî sängö",
        },
        "sr": Language{
            Name:       "Serbian",
            NativeName: "српски језик",
        },
        "gd": Language{
            Name:       "Scottish Gaelic; Gaelic",
            NativeName: "Gàidhlig",
        },
        "sn": Language{
            Name:       "Shona",
            NativeName: "chiShona",
        },
        "si": Language{
            Name:       "Sinhala, Sinhalese",
            NativeName: "සිංහල",
        },
        "sk": Language{
            Name:       "Slovak",
            NativeName: "slovenčina",
        },
        "sl": Language{
            Name:       "Slovene",
            NativeName: "slovenščina",
        },
        "so": Language{
            Name:       "Somali",
            NativeName: "Soomaaliga, af Soomaali",
        },
        "st": Language{
            Name:       "Southern Sotho",
            NativeName: "Sesotho",
        },
        "es": Language{
            Name:       "Spanish",
            NativeName: "español, castellano",
        },
        "su": Language{
            Name:       "Sundanese",
            NativeName: "Basa Sunda",
        },
        "sw": Language{
            Name:       "Swahili",
            NativeName: "Kiswahili",
        },
        "ss": Language{
            Name:       "Swati",
            NativeName: "SiSwati",
        },
        "sv": Language{
            Name:       "Swedish",
            NativeName: "svenska",
        },
        "ta": Language{
            Name:       "Tamil",
            NativeName: "தமிழ்",
        },
        "te": Language{
            Name:       "Telugu",
            NativeName: "తెలుగు",
        },
        "tg": Language{
            Name:       "Tajik",
            NativeName: "тоҷикӣ, toğikī, تاجیکی‎",
        },
        "th": Language{
            Name:       "Thai",
            NativeName: "ไทย",
        },
        "ti": Language{
            Name:       "Tigrinya",
            NativeName: "ትግርኛ",
        },
        "bo": Language{
            Name:       "Tibetan Standard, Tibetan, Central",
            NativeName: "བོད་ཡིག",
        },
        "tk": Language{
            Name:       "Turkmen",
            NativeName: "Türkmen, Түркмен",
        },
        "tl": Language{
            Name:       "Tagalog",
            NativeName: "Wikang Tagalog, ᜏᜒᜃᜅ᜔ ᜆᜄᜎᜓᜄ᜔",
        },
        "tn": Language{
            Name:       "Tswana",
            NativeName: "Setswana",
        },
        "to": Language{
            Name:       "Tonga (Tonga Islands)",
            NativeName: "faka Tonga",
        },
        "tr": Language{
            Name:       "Turkish",
            NativeName: "Türkçe",
        },
        "ts": Language{
            Name:       "Tsonga",
            NativeName: "Xitsonga",
        },
        "tt": Language{
            Name:       "Tatar",
            NativeName: "татарча, tatarça, تاتارچا‎",
        },
        "tw": Language{
            Name:       "Twi",
            NativeName: "Twi",
        },
        "ty": Language{
            Name:       "Tahitian",
            NativeName: "Reo Tahiti",
        },
        "ug": Language{
            Name:       "Uighur, Uyghur",
            NativeName: "Uyƣurqə, ئۇيغۇرچە‎",
        },
        "uk": Language{
            Name:       "Ukrainian",
            NativeName: "українська",
        },
        "ur": Language{
            Name:       "Urdu",
            NativeName: "اردو",
        },
        "uz": Language{
            Name:       "Uzbek",
            NativeName: "zbek, Ўзбек, أۇزبېك‎",
        },
        "ve": Language{
            Name:       "Venda",
            NativeName: "Tshivenḓa",
        },
        "vi": Language{
            Name:       "Vietnamese",
            NativeName: "Tiếng Việt",
        },
        "vo": Language{
            Name:       "Volapük",
            NativeName: "Volapük",
        },
        "wa": Language{
            Name:       "Walloon",
            NativeName: "Walon",
        },
        "cy": Language{
            Name:       "Welsh",
            NativeName: "Cymraeg",
        },
        "wo": Language{
            Name:       "Wolof",
            NativeName: "Wollof",
        },
        "fy": Language{
            Name:       "Western Frisian",
            NativeName: "Frysk",
        },
        "xh": Language{
            Name:       "Xhosa",
            NativeName: "isiXhosa",
        },
        "yi": Language{
            Name:       "Yiddish",
            NativeName: "ייִדיש",
        },
        "yo": Language{
            Name:       "Yoruba",
            NativeName: "Yorùbá",
        },
        "za": Language{
            Name:       "Zhuang, Chuang",
            NativeName: "Saɯ cueŋƅ, Saw cuengh",
        },
    }

    return list
}

func GetLanguage(code string) Language {
    code = strings.ToLower(code)
    return ListLanguages()[code]
}