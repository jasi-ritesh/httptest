package test

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"src/server/api"
	"testing"
	"time"
)

func TestArithmeticEngine(t *testing.T) {
	engineName := "Simple_Arithmetic"

	//Stage-1 Create a new Engine
	AddEngine(t, engineName)
	//Stage-7 Delete the Engine
	defer DeleteEngine(t, engineName)

	//Stage-2 Add Expressions to Engine
	AddExpression(t, engineName, &api.Expression{Name: "First", Expr: "2+3"})
	AddExpression(t, engineName, &api.Expression{Name: "Second", Expr: "10%3"})

	//Stage-3 Engine is asked to evaluate all Expressions
	Evaluate(t, engineName)

	//Stage-4 Fetch Result From Engine & Validate
	resultMap := FetchResult(t, engineName)
	assert.Equal(t, 2, len(resultMap))
	//Map Should have a key called First
	assert.Contains(t, resultMap, "First")
	//Map should have the correct value for the key called First
	assert.Equal(t, "5", resultMap["First"])
	//Map Should have a key called Second
	assert.Contains(t, resultMap, "Second")
	//Map should have the correct value for the key called Second
	assert.Equal(t, "1", resultMap["Second"])

	//Stage-5  Delete the Expression and Validate
	DeleteExpression(t, engineName, "First")
	//Fetch Result From Engine
	resultMap2 := FetchResult(t, engineName)
	assert.Equal(t, 1, len(resultMap2))

	//Stage-6 Clear the Engine
	ClearEngine(t, engineName)
}

func TestComparisonEngine(t *testing.T) {
	engineName := "Simple_Comparison"

	//Stage-1 Create a new Engine
	AddEngine(t, engineName)
	//Stage-7 Delete the Engine
	defer DeleteEngine(t, engineName)

	//Stage-2 Add Expressions to Engine
	AddExpression(t, engineName, &api.Expression{Name: "First", Expr: "3>2"})
	AddExpression(t, engineName, &api.Expression{Name: "Second", Expr: "4<2"})
	AddExpression(t, engineName, &api.Expression{Name: "Third", Expr: "100<=(50+60)"})

	//Stage-3 Engine is asked to evaluate all Expressions
	Evaluate(t, engineName)

	//Stage-4 Fetch Result From Engine & Validate
	resultMap := FetchResult(t, engineName)
	assert.Equal(t, 3, len(resultMap))
	//Map Should have a key called First
	assert.Contains(t, resultMap, "First")
	//Map should have the correct value for the key called First
	assert.Equal(t, "true", resultMap["First"])
	//Map Should have a key called Second
	assert.Contains(t, resultMap, "Second")
	//Map should have the correct value for the key called Second
	assert.Equal(t, "false", resultMap["Second"])
	//Map Should have a key called Third
	assert.Contains(t, resultMap, "Third")
	//Map should have the correct value for the key called Second
	assert.Equal(t, "true", resultMap["Third"])

	//Stage-5  Delete the Expression and Validate
	DeleteExpression(t, engineName, "First")
	resultMap2 := FetchResult(t, engineName)
	assert.Equal(t, 2, len(resultMap2))

	//Stage-6 Clear the Engine
	ClearEngine(t, engineName)
}

func TestMain(m *testing.M) {
	// Setup logger
	lumberjackLogrotate := &lumberjack.Logger{
		Filename:   "testlog.log",
		MaxSize:    5,  // Max megabytes before log is rotated
		MaxBackups: 90, // Max number of old log files to keep
		MaxAge:     60, // Max number of days to retain log files
		Compress:   true,
	}

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC1123Z})

	logMultiWriter := io.MultiWriter(os.Stdout, lumberjackLogrotate)
	log.SetOutput(logMultiWriter)

	m.Run()
}
