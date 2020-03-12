package validator

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"log"
	"net/http"
	"server/models"
	"server/pkg/file"
	"strconv"
	"strings"
)

var (
	formats      []string
	pokemonNames []string
)

// custom rules
func init() {
	// load formats and pokemon names
	var err error
	formats, err = file.UnmarshalJSONArray("assets/data/Formats.json")
	if err != nil {
		log.Panicln(err)
	}
	pokemonNames, err = file.UnmarshalJSONArray("assets/data/PokemonNames_en.json")
	if err != nil {
		log.Panicln(err)
	}

	// custom rules to take fixed length word.
	// e.g: max_word:5 will throw error if the field contains more than 5 words
	govalidator.AddCustomRule("max_word", func(field string, rule string, message string, value interface{}) error {
		valSlice := strings.Fields(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_word:")) //handle other error
		if len(valSlice) > l {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("The %s field must not be greater than %d words. ", field, l)
		}
		return nil
	})

	// custom rules to check if an element in a slice
	govalidator.AddCustomRule("in_formats", func(field string, rule string, message string, value interface{}) error {
		fname, ok := value.(string)
		if !ok {
			// wrong use case
			return fmt.Errorf("Incorrect Format type. ")
		}
		for _, f := range formats {
			if f == fname {
				return nil
			}
		}
		if message != "" {
			return errors.New(message)
		}
		return fmt.Errorf("The format %s is temporarily not supported. ", fname)
	})

	govalidator.AddCustomRule("in_pokemon", func(field string, rule string, message string, value interface{}) error {
		pname, ok := value.(string)
		if !ok {
			// wrong use case
			return fmt.Errorf("Incorrect pokemon name type. ")
		}
		for _, p := range pokemonNames {
			if p == pname {
				return nil
			}
		}
		if message != "" {
			return errors.New(message)
		}
		return fmt.Errorf("Not a Pokemon: %s. ", pname)
	})
}

func TeamValidator(team *models.Team, r *http.Request) (error string, valid bool) {
	rules := govalidator.MapData{
		"title":        []string{"required", "between:1,50"},
		"author":       []string{"required", "between:1,40"},
		"format":       []string{"in_formats"},
		"pokemon1":     []string{"in_pokemon"},
		"pokemon2":     []string{"in_pokemon"},
		"pokemon3":     []string{"in_pokemon"},
		"pokemon4":     []string{"in_pokemon"},
		"pokemon5":     []string{"in_pokemon"},
		"pokemon6":     []string{"in_pokemon"},
		"rentalImgUrl": []string{"url"},
		"showdown":     []string{"between:200,1600", "max_word:300"}, // 6 * 50 words
		"description":  []string{"max:1100"},
		"state":        []string{"numeric", "bool"},
	}
	// TODO: 验证6只pm名字合法性
	// TODO: 验证showdown语法合法性，考虑post参数添加Koffing解析的JSON
	// TODO: 验证模式合法性
	messages := govalidator.MapData{
		"showdown": []string{"between: Not a valid Showdown paste. "},
	}

	opts := govalidator.Options{
		Request:  r,
		Data:     team,
		Rules:    rules,
		Messages: messages, // custom message map (Optional)
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()

	// valid
	if len(e) == 0 {
		return "", true
	}

	var errMsgs []string
	for _, i := range e {
		for _, j := range i {
			errMsgs = append(errMsgs, j)
		}
	}

	return strings.Join(errMsgs, "\\n"), false
}
