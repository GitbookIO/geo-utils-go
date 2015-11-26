# Geo Utilities

This package is a Go implementation of the [geo-utils node package](https://github.com/GitbookIO/geo-utils). It contains multiple utilies to manage GEO data such as countries, languages and US states.

## Install

```
$ go get github.com/GitbookIO/geo-utils-go
```

## Use

##### Import

``` Go
import "github.com/GitbookIO/geo-utils-go"
```

##### Countries

``` Go
// List all countries
countriesList := geoutils.ListCountries()

reflect.DeepEqual(countriesList, map[string]string{
    "A1": "Anonymous Proxy",
    "A2": "Satellite Provider",
    "O1": "Other Country",
    "AD": "Andorra",
    "AE": "United Arab Emirates",
    "AF": "Afghanistan",
    ...
})

// Get a country
france := geoutils.GetCountry("FR")

strings.Compare(france, "France")
```

##### US States

``` Go
// List all US states
statesList := geoutils.ListUSStates()

reflect.DeepEqual(statesList, map[string]string{
    "AL": "Alabama",
    "AK": "Alaska",
    "AZ": "Arizona",
    "AR": "Arkansas",
    "CA": "California",
    ...
})

// Get a US state
california := geoutils.GetUSStateByAbbr("CA")

strings.Compare(california, "California")
```

##### Languages

``` Go
// Languages are provided as a Language struct
type Language struct {
    Name        string
    NativeName  string
}

// List all languages
languagesList := geoutils.ListLanguages()

reflect.DeepEqual(countriesList, map[string]Language{
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
    ...
})

// Get a language
tamil := geoutils.GetLanguage("ta")

reflect.DeepEqual(tamil, Language{
    Name:       "Tamil",
    NativeName: "தமிழ்",
})
```