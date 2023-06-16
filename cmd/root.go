/*
Copyright © 2023 Robert Thanulingam

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/robertranjan/wiggle/lib"
	"github.com/robertranjan/wiggle/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var debug bool
var l *logrus.Logger

type Address struct {
	Street  string
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "wiggle",
	Version: version.Version,
	Short:   "A brief description of your application: wiggle",
	Long: `A longer description of application:
	wiggle
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("main command... got a call\n")
		person := Person{
			Name: "John Doe",
			Age:  30,
			Address: Address{
				Street:  "123 Main Street",
				City:    "New York",
				Country: "USA",
			},
		}
		fmt.Printf("flattened struct: %#v\n", lib.FlattenStruct(person, ""))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./wiggle.yaml", "config file")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", true, "run app with this log level")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	InitLogger()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".wiggle" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".wiggle")
	}
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

var (
	mystring string
)

type GlobalHook struct {
}

func (h *GlobalHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *GlobalHook) Fire(e *logrus.Entry) error {
	e.Data["mystring"] = mystring
	return nil
}

func InitLogger() {
	l = logrus.New()
	l.SetReportCaller(true)
	l.Out = os.Stdout
	l.Formatter = &logrus.TextFormatter{
		DisableColors:    true,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 03:04:05",
		PadLevelText:     true,
		QuoteEmptyFields: true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
		},
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			s := strings.Split(f.Function, ".")
			_, l := f.Func.FileLine(f.PC)
			funcname := s[len(s)-1]
			filename := filepath.Join(path.Base(path.Dir(f.File)), path.Base(f.File)) + ":" + strconv.Itoa(l)
			return funcname, filename
		},
	}
	l.AddHook(&GlobalHook{})
	mystring = "abc 1 "

	l.Info("example of custom format caller")
	mystring = "abc 2 "
	l.SetLevel(logrus.ErrorLevel)
	mystring = "abc 3 "
	if debug {
		l.SetLevel(logrus.DebugLevel)
	}
}
