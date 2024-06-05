package nmap

import (
	"encoding/xml"
	"os"

	"github.com/vdjagilev/nmap-formatter/v3/formatter"
)

// GenerateSarif generates a SARIF report from the nmap output xml.
func (n *Client) GenerateSarif() error {
	// err := n.WriteToFile()
	// if err != nil {
	// 	return err
	// }

	// parse reports/nmap-report.xml

	// convert to sarif

	return nil
}

func ConvertNmapXMLToSarif() error {
	// open reports/nmap-report.xml
	// file, err := os.Open("reports/nmap-report.xml")
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()

	// parse reports/nmap-report.xml

	// convert to sarif
	return nil
}

// try parsing nmap-reports/vulscan/nmap-report.xml
// func (n *Client) ReadXML(scan string) error {
// 	outdir := n.Config.OutputDir
// 	var filePath string
// 	if scan == "vulscan" {
// 		filePath = fmt.Sprintf("%s/nmap-reports/vulscan/nmap-report.xml", outdir)
// 	}
// 	if scan == "direct" {
// 		filePath = fmt.Sprintf("%s/vulners/nmap-report.xml", outdir)
// 	}
// 	if scan == "vulners" {
// 		filePath = fmt.Sprintf("%s/vulners/nmap-report.xml", outdir)
// 	}
//
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
//
// 	reader := bufio.NewReader(file)
//
// 	// parse reports/nmap-report.xml using xml.Decoder
// 	decoder := xml.NewDecoder(reader)
//
// 	return nil
// }

func (n *Client) ConverToJSON() {
	fileName := "nmap-reports/direct/nmap-report.xml"
	fileOutput := "nmap-reports/direct/nmap-report.json"
	var nmap formatter.NMAPRun

	var config formatter.Config = formatter.Config{}

	// Read XML file that was produced by nmap (with -oX option)
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// Unmarshal XML and map structure(s) fields accordingly
	if err = xml.Unmarshal(content, &nmap); err != nil {
		panic(err)
	}

	// Output data to console stdout
	// You can use any other io.Writer implementation
	// for example: os.OpenFile("file.json", os.O_CREATE|os.O_EXCL|os.O_WRONLY, os.ModePerm)
	// config.Writer = os.Stdout
	outputFile, err := os.Create(fileOutput)
	if err != nil {
		panic(err)
	}
	config.Writer = outputFile
	// Formatting data to JSON, you can use:
	// CSVOutput, MarkdownOutput, HTMLOutput as well
	config.OutputFormat = formatter.JSONOutput

	// Setting formatter data/options
	templateData := formatter.TemplateData{
		NMAPRun: nmap, // NMAP output data itself
		OutputOptions: formatter.OutputOptions{
			JSONOptions: formatter.JSONOutputOptions{
				PrettyPrint: true, // Additional option to prettify JSON
			},
		},
	}

	// New formatter instance
	formatter := formatter.New(&config)
	if formatter == nil {
		// Not json/markdown/html/csv
		panic("wrong formatter provided")
	}

	// Attempt to format the data
	if err = formatter.Format(&templateData, "" /* no template content for JSON */); err != nil {
		// html template could not be parsed or some other issue occured
		panic(err)
	}
}

// sarif example for eslint
// {
//   "version": "2.1.0",
//   "$schema": "http://json.schemastore.org/sarif-2.1.0-rtm.4",
//   "runs": [
//     {
//       "tool": {
//         "driver": {
//           "name": "ESLint",
//           "informationUri": "https://eslint.org",
//           "rules": [
//             {
//               "id": "no-unused-vars",
//               "shortDescription": {
//                 "text": "disallow unused variables"
//               },
//               "helpUri": "https://eslint.org/docs/rules/no-unused-vars",
//               "properties": {
//                 "category": "Variables"
//               }
//             }
//           ]
//         }
//       },
//       "artifacts": [
//         {
//           "location": {
//             "uri": "file:///C:/dev/sarif/sarif-tutorials/samples/Introduction/simple-example.js"
//           }
//         }
//       ],
//       "results": [
//         {
//           "level": "error",
//           "message": {
//             "text": "'x' is assigned a value but never used."
//           },
//           "locations": [
//             {
//               "physicalLocation": {
//                 "artifactLocation": {
//                   "uri": "file:///C:/dev/sarif/sarif-tutorials/samples/Introduction/simple-example.js",
//                   "index": 0
//                 },
//                 "region": {
//                   "startLine": 1,
//                   "startColumn": 5
//                 }
//               }
//             }
//           ],
//           "ruleId": "no-unused-vars",
//           "ruleIndex": 0
//         }
//       ]
//     }
//   ]
// }
