// -------------------------------------------- PRASER ------------------------------------------------ \\
//                                   official part of LunaCompiler                                      \\
//                                        Authors: LunaTeam                                             \\
// ---------------------------------------------------------------------------------------------------- \\

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func mainPrase(filename string) map[string]lunafunc {
	file, err := os.Open(filename)
	if err != nil {
		ShowErr("Failed to open file: ", err.Error())
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	inComment := false

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimLeft(line, " \t")

		// comment start
		if strings.HasPrefix(line, "/*") {
			inComment = true
			comndts -= 1
		}
		// ignore coments
		if inComment {
			comndts += 1
			// comment end
			if strings.Contains(line, "*/") {
				inComment = false
			}
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		ShowErr("Error reading file: ", err.Error())
		return nil
	}

	functions := make(map[string]lunafunc)

	for i, line := range lines {
		i += comndts + 1
		line = strings.TrimSpace(line)
		if matches := funcRegex.FindStringSubmatch(line); matches != nil {
			if curFunc != nil {
				ShowErri(i, fif, funcName) // FunctionNesting == false :emoji_ok:
				return nil
			}

			funcName = matches[1]
			if _, exists := functions[funcName]; exists {
				ShowErri(i, fmt.Sprintf("Duplicate function name: `%s`", funcName))
				return nil
			}
			argss, err := pArgs(matches[2])
			if err != nil {
				ShowErri(i, err)
			}
			curFunc = &lunafunc{
				Args:    argss,
				Returns: pReturns(i, matches[3]),
				Lines:   []string{},
			}
			functions[funcName] = *curFunc
			idx := strings.Index(line, "{")
			if idx != -1 {
				afterB := strings.TrimSpace(line[idx+1:])
				if afterB != "" {
					curFunc.Lines = append(curFunc.Lines, afterB)
				}
				dep = 1
				nextbody = false
				continue
			}

			nextbody = true
			continue
		}

		if nextbody {
			if line == "{" {
				dep = 1
				nextbody = false
				continue
			} else {
				ShowErri(i, "Expected '{' after function declaration", funcName)
				return nil
			}
		}

		// if int function body
		if curFunc != nil {
			openc := strings.Count(line, "{")  // ADD
			closec := strings.Count(line, "}") // DEDUCT
			dep += openc
			dep -= closec

			if dep == 0 {
				functions[funcName] = *curFunc // save func to map
				curFunc = nil
				comndts += 1
				continue
			}
			curFunc.Lines = append(curFunc.Lines, line)
		}
	}
	return functions
}

func pReturns(i int, returns string) []string {
	if returns == "" {
		return nil
	}

	var returnTypes []string
	for _, r := range strings.Split(returns, ",") {
		type_ := strings.TrimSpace(r)

		if !aTypes[type_] {
			ShowErri(i, fmt.Sprintf("invalid return type '%s'", type_))
			return nil
		}

		returnTypes = append(returnTypes, type_)
	}
	return returnTypes
}

func pArgs(args string) (map[string]string, error) {
	argMap := make(map[string]string)
	if args == "" { // brak argumentów
		return argMap, nil
	}

	for _, arg := range strings.Split(args, ",") {
		parts := strings.Fields(strings.TrimSpace(arg))
		if len(parts) != 2 {
			return nil, errors.New("invalid argument format: expected 'name type'")
		}

		name, typ := parts[0], parts[1]

		if !aTypes[typ] {
			return nil, errors.New(fmt.Sprintf("invalid type '%s' for argument '%s'", typ, name))
		}

		argMap[name] = typ
	}

	return argMap, nil
}
