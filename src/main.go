package main

import (
    "encoding/json"
    "log"
    "io/ioutil"
    "html/template"
    "net/http"
    "strings"
    "strconv"
    "github.com/gorilla/mux"
)

type Speller struct {
    Number  int `json:"Number"`
    Text    string `json:"Text"`
}

func digitToString(x int) string{
    switch x{
        case 1:
            return "satu"
        case 2:
            return "dua"
        case 3:
            return "tiga"
        case 4:
            return "empat"
        case 5:
            return "lima"
        case 6:
            return "enam"
        case 7:
            return "tujuh"
        case 8:
            return "delapan"
        case 9:
            return "sembilan"
        default:
            return ""
    }
}

func stringToDigit(s string) int{
    switch s{
        case "satu":
            return 1
        case "dua":
            return 2
        case "tiga":
            return 3
        case "empat":
            return 4
        case "lima":
            return 5
        case "enam":
            return 6
        case "tujuh":
            return 7
        case "delapan":
            return 8
        case "sembilan":
            return 9
        case "sepuluh":
            return 10
        case "sebelas":
            return 11
        case "seratus":
            return 100
        case "seribu":
            return 1000
        default:
            return 0
    }
}

func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

func threeDigitToString(x int) string{
    var ratus int = (x - (x % 100)) / 100
    var puluh int = (x - ratus*100 - (x % 10)) / 10
    var sisa int = (x - ratus*100 - puluh * 10)
    if ratus == 0{
        if puluh == 0{
            return digitToString(sisa)
        } else if puluh == 1{
            if sisa == 0{
                return "sepuluh"
            } else if sisa == 1{
                return "sebelas"
            } else{
                return digitToString(sisa) + " belas"
            }
        } else{
            if sisa == 0{
                return digitToString(puluh) + " puluh"
            }else{
                return digitToString(puluh) + " puluh " + digitToString(sisa)
            }
        }
    } else if ratus == 1{
        if puluh == 0{
            if sisa == 0{
                return "seratus"
            }else{
                return "seratus " + digitToString(sisa)
            }
        } else if puluh == 1{
            if sisa == 0{
                return "seratus sepuluh"
            } else if sisa == 1{
                return "seratus sebelas"
            } else{
                return "seratus " + digitToString(sisa) + " belas"
            }
        } else{
            if sisa == 0{
                return "seratus " + digitToString(puluh) + " puluh"
            } else{
                return "seratus " + digitToString(puluh) + " puluh " + digitToString(sisa)
            }
        }
    } else{
        if puluh == 0{
            if sisa == 0{
                return digitToString(ratus) + " ratus"
            }else{
                return digitToString(ratus) + " ratus " + digitToString(sisa)
            }
        } else if puluh == 1{
            if sisa == 0{
                return digitToString(ratus) + " ratus sepuluh"
            }
            if sisa == 1{
                return digitToString(ratus) + " ratus sebelas"
            } else{
                return digitToString(ratus) + " ratus " + digitToString(sisa) + " belas"
            }
        } else{
            if sisa == 0{
                return digitToString(ratus) + " ratus " + digitToString(puluh) + " puluh"
            }else{
                return digitToString(ratus) + " ratus " + digitToString(puluh) + " puluh " + digitToString(sisa)
            }
        }
    }
}

func intToString(x int) string{
    var milyar int = (x - (x % 1000000000)) / 1000000000
    var juta int = (x - milyar*1000000000 - (x % 1000000)) / 1000000
    var ribu int = (x - milyar*1000000000 - juta*1000000 - (x % 1000)) / 1000
    var sisa int = (x - milyar*1000000000 - juta*1000000 - ribu*1000)
    var res string = ""
    if milyar > 0{
        res += digitToString(milyar) + " milyar "
    }
    if juta > 0{
        res += threeDigitToString(juta) + " juta "
    }
    if ribu > 0{
        if ribu == 1{
            res += "seribu "
        } else{
            res += threeDigitToString(ribu) + " ribu "
        }
    }
    res += threeDigitToString(sisa)
    res = strings.TrimSpace(res)
    return res
}

func stringToInt(s string) int{
    keywords := []string{"satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas", "belas", "puluh", "ratus", "seratus", "ribu", "seribu", "juta", "milyar"}
    stops := []string{"satu", "sepuluh", "sebelas"}
    state1 := []string{"satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}
    state2 := []string{"ratus", "puluh", "belas"}
    state0 := []string{"milyar", "juta", "ribu"}
    words := strings.Split(s, " ")
    var milyar, juta, ribu, ratus, puluh, belas, broken bool = false, false, false, false, false, false, false
    var state int = 1
    var temp int = 0
    var res int = 0
    for _, v := range words{
        if !Include(keywords, v){
            return -1
        }
    }
    for i, v := range words{
        if state == 1{
            if Include(stops, v){
                state = 0
            } else if v == "seribu"{
                ribu = true
            } else if v == "seratus"{
                ratus = true
                var x string
                if i!=len(words)-1{
                    x = words[i+1]
                }
                if Include(state0, x){
                    state = 0
                } else{
                    state = 1
                }
            } else if Include(state1, v){
                var x string
                if i!=len(words)-1{
                    x = words[i+1]
                }
                if Include(state0, x){
                    state = 0
                } else{
                    state = 2
                }
            } else{
                broken = true
            }
            temp += stringToDigit(v)
        } else if state == 2{
            if Include(state2, v){
                if v == "ratus" && !ratus && !puluh && !belas{
                    temp2 := temp % 100
                    temp -= temp2
                    temp2 *= 100
                    temp += temp2
                    ratus = true
                } else if v == "puluh" && !puluh && !belas{
                    temp2 := temp % 10
                    temp -= temp2
                    temp2 *= 10
                    temp += temp2
                    ratus = true
                    puluh = true
                    belas = true
                } else if v == "belas" && !puluh && !belas{
                    temp += 10
                    ratus = true
                    puluh = true
                    belas = true
                } else{
                    broken = true
                }
                var x string
                if i!=len(words)-1{
                    x = words[i+1]
                }
                if Include(state0, x){
                    state = 0
                } else{
                    state = 1
                }
            } else{
                broken = true
            }
        } else if state == 0{
            if Include(state0, v){
                if v == "milyar" && !milyar && !juta && !ribu{
                    temp *= 1000000000
                    milyar = true
                } else if v == "juta" && !juta && !ribu{
                    temp *= 1000000
                    milyar = true
                    juta = true
                } else if v == "ribu" && !ribu{
                    temp *= 1000
                    milyar = true
                    juta = true
                    ribu = true
                } else{
                    broken = true
                }
                ratus = false
                puluh = false
                belas = false
                res += temp
                temp = 0
                state = 1
            } else{
                broken = true
            }
        }
    }
    res += temp
    if broken {
        return -2
    } else{
        return res
    }
}

func jsonIntToString(w http.ResponseWriter, r *http.Request){
    query := r.URL.Query()
    numtext := query.Get("number")
    var spell Speller
    num, _ := strconv.Atoi(numtext)
    txt := intToString(num)
    spell = Speller{Number: num, Text: txt}
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, spell)
}

func jsonStringToInt(w http.ResponseWriter, r *http.Request){
    reqBody, _ := ioutil.ReadAll(r.Body)
    var spell Speller
    json.Unmarshal(reqBody, &spell)
    txt := string(spell.Text)
    num := stringToInt(txt)
    spell = Speller{Number: num, Text: txt}
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, spell)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    s := Speller{Number: 1, Text: "satuqwe"}
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, s)
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/spell", jsonIntToString).Methods("GET")
    myRouter.HandleFunc("/read", jsonStringToInt).Methods("POST")
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    handleRequests()
}