package main

import (
	"flag"
	"fmt"

	"github.com/museop/design-pattern-in-go/behavioral/chainofresponsibility"
	"github.com/museop/design-pattern-in-go/behavioral/command"
	"github.com/museop/design-pattern-in-go/behavioral/iterator"
	"github.com/museop/design-pattern-in-go/behavioral/mediator"
	"github.com/museop/design-pattern-in-go/behavioral/observer"
	"github.com/museop/design-pattern-in-go/behavioral/state"
	"github.com/museop/design-pattern-in-go/behavioral/strategy"
	"github.com/museop/design-pattern-in-go/creational/abstractfactory"
	"github.com/museop/design-pattern-in-go/creational/builder"
	"github.com/museop/design-pattern-in-go/creational/factorymethod"
	"github.com/museop/design-pattern-in-go/creational/prototype"
	"github.com/museop/design-pattern-in-go/creational/singleton"
	"github.com/museop/design-pattern-in-go/structural/adapter"
	"github.com/museop/design-pattern-in-go/structural/bridge"
	"github.com/museop/design-pattern-in-go/structural/composite"
	"github.com/museop/design-pattern-in-go/structural/decorator"
	"github.com/museop/design-pattern-in-go/structural/facade"
	"github.com/museop/design-pattern-in-go/structural/proxy"
)

func main() {
	patternName := flag.String("pattern", "none", "The design pattern name")
	flag.Parse()

	fmt.Printf("Pattern name: %s\n", *patternName)
	fmt.Println("=========================================")

	switch *patternName {
	/* Creational Patterns */
	case "abstract factory":
		abstractfactory.TestAbstractFactory()
	case "factory method":
		factorymethod.TestFactoryMethod()
	case "builder":
		builder.TestBuilder()
	case "prototype":
		prototype.TestPrototype()
	case "singleton":
		singleton.TestSingleton()

	/* Structural Patterns */
	case "adapter":
		adapter.TestAdapter()
	case "bridge":
		bridge.TestBridge()
	case "composite":
		composite.TestComposite()
	case "decorator":
		decorator.TestDecorator()
	case "facade":
		facade.TestFacade()
	case "proxy":
		proxy.TestProxy()

	/* Behavioral Patterns */
	case "chain of responsibility":
		chainofresponsibility.TestChainOfResponsibility()
	case "observer":
		observer.TestObserver()
	case "command":
		command.TestCommand()
	case "iterator":
		iterator.TestIterator()
	case "mediator":
		mediator.TestMediator()
	case "state":
		state.TestState()
	case "strategy":
		strategy.TestStrategy()

	default:
		fmt.Println("Undefined pattern")
	}

	fmt.Println("=========================================")
}
