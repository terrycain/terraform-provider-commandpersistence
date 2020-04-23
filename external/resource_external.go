package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/ioutil"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceExternal() *schema.Resource {
	return &schema.Resource{
		Create: CreateExternal,
		Read:   schema.Noop,
		Delete: schema.RemoveFromState,
		Importer: &schema.ResourceImporter{
			State: ImportContent,
		},

		Schema: map[string]*schema.Schema{
			"program": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"working_dir": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"query": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"result": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

// CreateExternal Run command and save the output
func CreateExternal(d *schema.ResourceData, meta interface{}) error {

	programI := d.Get("program").([]interface{})
	workingDir := d.Get("working_dir").(string)
	query := d.Get("query").(map[string]interface{})

	if err := validateProgramAttr(programI); err != nil {
		return err
	}

	program := make([]string, len(programI))
	for i, vI := range programI {
		program[i] = vI.(string)
	}

	cmd := exec.Command(program[0], program[1:]...)

	cmd.Dir = workingDir

	queryJSON, err := json.Marshal(query)
	if err != nil {
		return err
	}

	cmd.Stdin = bytes.NewReader(queryJSON)

	resultJSON, err := cmd.Output()
	log.Printf("[TRACE] JSON output: %+v\n", resultJSON)
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.Stderr != nil && len(exitErr.Stderr) > 0 {
				return fmt.Errorf("failed to execute %q: %s", program[0], string(exitErr.Stderr))
			}
			return fmt.Errorf("command %q failed with no error message", program[0])
		} else {
			return fmt.Errorf("failed to execute %q: %s", program[0], err)
		}
	}

	result := map[string]string{}
	err = json.Unmarshal(resultJSON, &result)
	if err != nil {
		return fmt.Errorf("command %q produced invalid JSON: %s", program[0], err)
	}

	d.Set("result", result)
	d.SetId("-")

	return nil
}

func validateProgramAttr(v interface{}) error {
	args := v.([]interface{})
	if len(args) < 1 {
		return fmt.Errorf("'program' list must contain at least one element")
	}

	for i, vI := range args {
		if _, ok := vI.(string); !ok {
			return fmt.Errorf(
				"'program' element %d is %T; a string is required",
				i, vI,
			)
		}
	}

	// first element is assumed to be an executable command, possibly found
	// using the PATH environment variable.
	_, err := exec.LookPath(args[0].(string))
	if err != nil {
		return fmt.Errorf("can't find external program %q", args[0])
	}

	return nil
}

// ImportContent Import content from json file
func ImportContent(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	filepath := d.Id()

	if _, err := os.Stat(filepath); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	result := map[string]string{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return fmt.Errorf("invalid JSON: %s", err)
	}

	d.Set("result", result)
	d.SetId("-")

	return []*schema.ResourceData{d}, nil
}
