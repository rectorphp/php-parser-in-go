package classresolver_test

import (
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/conf"
	"github.com/rectorphp/php-parser-in-go/pkg/parser"
	"github.com/rectorphp/php-parser-in-go/pkg/version"
	"github.com/rectorphp/php-parser-in-go/pkg/visitor/classresolver"
)

func parse(test *testing.T, source string) (registry *classresolver.Registry) {
	test.Helper()

	rootNode, err := parser.Parse([]byte(source), conf.Config{
		Version: &version.Version{Major: 7, Minor: 2},
	})
	if err != nil {
		test.Fatalf("parse error: %v", err)
	}

	registry = classresolver.NewRegistry()
	registry.Collect(rootNode)
	return registry
}

func TestRegistry_ownMethod(test *testing.T) {
	registry := parse(test, `<?php
class Calculator
{
    private function sum($first, $second) {}
}`)

	parameterCount, ok := registry.LookupMethodParameterCount("Calculator", "sum")
	if !ok || parameterCount != 2 {
		test.Errorf("expected 2 params, got %d (ok=%v)", parameterCount, ok)
	}
}

func TestRegistry_inheritedMethod(test *testing.T) {
	registry := parse(test, `<?php
class Base
{
    protected function change($reservation) {}
}

class Child extends Base {}`)

	parameterCount, ok := registry.LookupMethodParameterCount("Child", "change")
	if !ok || parameterCount != 1 {
		test.Errorf("expected inherited 1 param, got %d (ok=%v)", parameterCount, ok)
	}
}

func TestRegistry_ambiguousName(test *testing.T) {
	registry := parse(test, `<?php
class Ledger
{
    private function record($entity, $data) {}
}

class Ledger
{
    private function record($row) {}
}`)

	if _, ok := registry.LookupMethodParameterCount("Ledger", "record"); ok {
		test.Error("expected lookup to fail for an ambiguous class name")
	}
}

func TestRegistry_unknownMethod(test *testing.T) {
	registry := parse(test, `<?php
class Calculator
{
    private function sum($first) {}
}`)

	if _, ok := registry.LookupMethodParameterCount("Calculator", "missing"); ok {
		test.Error("expected lookup to fail for an unknown method")
	}
	if _, ok := registry.LookupMethodParameterCount("Unknown", "sum"); ok {
		test.Error("expected lookup to fail for an unknown class")
	}
}

func TestRegistry_namespacedInheritance(test *testing.T) {
	registry := classresolver.NewRegistry()

	baseRoot, err := parser.Parse([]byte(`<?php
namespace App\Base;

class Controller
{
    protected function handle($request) {}
}`), conf.Config{Version: &version.Version{Major: 7, Minor: 2}})
	if err != nil {
		test.Fatalf("parse base: %v", err)
	}
	registry.Collect(baseRoot)

	childRoot, err := parser.Parse([]byte(`<?php
namespace App\Web;

use App\Base\Controller;

class RoomController extends Controller {}`), conf.Config{Version: &version.Version{Major: 7, Minor: 2}})
	if err != nil {
		test.Fatalf("parse child: %v", err)
	}
	registry.Collect(childRoot)

	parameterCount, ok := registry.LookupMethodParameterCount("App\\Web\\RoomController", "handle")
	if !ok || parameterCount != 1 {
		test.Errorf("expected inherited 1 param across namespaces, got %d (ok=%v)", parameterCount, ok)
	}
}

func TestRegistry_inheritsVisibleMethodProtected(test *testing.T) {
	registry := parse(test, `<?php
class Base
{
    protected function run() {}
}

class Child extends Base {}`)

	found, resolvable := registry.InheritsVisibleMethod("Child", "run")
	if !found || !resolvable {
		test.Errorf("expected protected ancestor method to be found (found=%v, resolvable=%v)", found, resolvable)
	}
}

func TestRegistry_inheritsVisibleMethodPrivateImposesNoConstraint(test *testing.T) {
	registry := parse(test, `<?php
class Base
{
    private function run() {}
}

class Child extends Base {}`)

	found, resolvable := registry.InheritsVisibleMethod("Child", "run")
	if found || !resolvable {
		test.Errorf("expected private ancestor method to impose no constraint (found=%v, resolvable=%v)", found, resolvable)
	}
}

func TestRegistry_inheritsVisibleMethodUnknownAncestor(test *testing.T) {
	registry := parse(test, `<?php
class Child extends \Vendor\Base {}`)

	if _, resolvable := registry.InheritsVisibleMethod("Vendor\\Base", "run"); resolvable {
		test.Error("expected an unknown ancestor to be unresolvable")
	}
}

func TestRegistry_descendantDeclaresMethodDirectChild(test *testing.T) {
	registry := parse(test, `<?php
class Base
{
    protected function setFilters($name) {}
}

class Child extends Base
{
    protected function setFilters($name) {}
}`)

	declared, resolvable := registry.DescendantDeclaresMethod("Base", "setFilters")
	if !declared || !resolvable {
		test.Errorf("expected the child override to be found, got declared=%v resolvable=%v", declared, resolvable)
	}
}

func TestRegistry_descendantDeclaresMethodGrandchild(test *testing.T) {
	registry := parse(test, `<?php
class Base
{
    protected function setFilters($name) {}
}

class Middle extends Base {}

class Leaf extends Middle
{
    protected function setFilters($name) {}
}`)

	declared, resolvable := registry.DescendantDeclaresMethod("Base", "setFilters")
	if !declared || !resolvable {
		test.Errorf("expected the grandchild override to be found, got declared=%v resolvable=%v", declared, resolvable)
	}
}

func TestRegistry_descendantDeclaresMethodUnrelatedName(test *testing.T) {
	registry := parse(test, `<?php
class Base
{
    protected function setFilters($name) {}
}

class Child extends Base
{
    protected function other($name) {}
}`)

	declared, resolvable := registry.DescendantDeclaresMethod("Base", "setFilters")
	if declared || !resolvable {
		test.Errorf("expected no override, got declared=%v resolvable=%v", declared, resolvable)
	}
}

func TestRegistry_descendantDeclaresMethodUnknownClass(test *testing.T) {
	registry := parse(test, `<?php
class Base
{
    protected function setFilters($name) {}
}`)

	declared, resolvable := registry.DescendantDeclaresMethod("Missing", "setFilters")
	if declared || resolvable {
		test.Errorf("expected an unresolvable answer for an unknown class, got declared=%v resolvable=%v", declared, resolvable)
	}
}

func TestRegistry_descendantDeclaresMethodAnonymousChild(test *testing.T) {
	registry := parse(test, `<?php
class Base
{
    protected function setFilters($name) {}
}

$extension = new class extends Base {};`)

	declared, resolvable := registry.DescendantDeclaresMethod("Base", "setFilters")
	if declared || resolvable {
		test.Errorf("expected an anonymous subclass to block the answer, got declared=%v resolvable=%v", declared, resolvable)
	}
}

func TestRegistry_descendantDeclaresMethodCyclicExtends(test *testing.T) {
	registry := parse(test, `<?php
class First extends Second {}

class Second extends First {}`)

	declared, resolvable := registry.DescendantDeclaresMethod("First", "setFilters")
	if declared || !resolvable {
		test.Errorf("expected a cycle to terminate without a match, got declared=%v resolvable=%v", declared, resolvable)
	}
}

func TestRegistry_descendantOverridesDirectChild(test *testing.T) {
	registry := parse(test, `<?php
class Base
{
    protected function setFilters($name) {}
}

class Child extends Base
{
    protected function setFilters($name) {}
}`)

	overrides, resolvable := registry.DescendantOverrides("Base", "setFilters")
	if !resolvable || len(overrides) != 1 {
		test.Errorf("expected one override, got resolvable=%v count=%d", resolvable, len(overrides))
	}
}

func TestRegistry_descendantOverridesUnrelatedName(test *testing.T) {
	registry := parse(test, `<?php
class Base
{
    protected function setFilters($name) {}
}

class Child extends Base
{
    protected function other($name) {}
}`)

	overrides, resolvable := registry.DescendantOverrides("Base", "setFilters")
	if !resolvable || len(overrides) != 0 {
		test.Errorf("expected no override, got resolvable=%v count=%d", resolvable, len(overrides))
	}
}

func TestRegistry_descendantOverridesVariadic(test *testing.T) {
	// a variadic override is unrecorded in the parameter map, so
	// DescendantDeclaresMethod misses it; DescendantOverrides keeps the node.
	registry := parse(test, `<?php
class Base
{
    protected function setFilters($name) {}
}

class Child extends Base
{
    protected function setFilters(...$names) {}
}`)

	overrides, resolvable := registry.DescendantOverrides("Base", "setFilters")
	if !resolvable || len(overrides) != 1 {
		test.Errorf("expected the variadic override to be found, got resolvable=%v count=%d", resolvable, len(overrides))
	}
}

func TestRegistry_descendantOverridesUnknownClass(test *testing.T) {
	registry := parse(test, `<?php
class Base
{
    protected function setFilters($name) {}
}`)

	overrides, resolvable := registry.DescendantOverrides("Missing", "setFilters")
	if resolvable || overrides != nil {
		test.Errorf("expected an unresolvable answer for an unknown class, got resolvable=%v count=%d", resolvable, len(overrides))
	}
}

func TestRegistry_descendantOverridesAnonymousChild(test *testing.T) {
	registry := parse(test, `<?php
class Base
{
    protected function setFilters($name) {}
}

$extension = new class extends Base {};`)

	overrides, resolvable := registry.DescendantOverrides("Base", "setFilters")
	if resolvable || overrides != nil {
		test.Errorf("expected an anonymous subclass to block the answer, got resolvable=%v count=%d", resolvable, len(overrides))
	}
}

// TestRegistry_isSubtypeOf covers the hierarchy question a reporting rule asks
// about a class it can only see from another file: is this a Command, is this a
// form type. Both the extends chain and the implements list carry the answer.
func TestRegistry_isSubtypeOf(test *testing.T) {
	registry := parse(test, `<?php
namespace App;

use Symfony\Component\Console\Command\Command;
use Symfony\Contracts\EventDispatcher\EventSubscriberInterface;

abstract class BaseCommand extends Command {}

class ImportCommand extends BaseCommand {}

interface AuditableInterface extends LoggableInterface {}

class Lead implements AuditableInterface {}

class Listener implements EventSubscriberInterface {}

class Plain {}`)

	tests := []struct {
		classFQN    string
		ancestorFQN string
		expected    bool
	}{
		{"App\\ImportCommand", "Symfony\\Component\\Console\\Command\\Command", true},
		{"App\\BaseCommand", "Symfony\\Component\\Console\\Command\\Command", true},
		{"App\\ImportCommand", "App\\BaseCommand", true},
		{"App\\ImportCommand", "App\\ImportCommand", true},
		{"App\\Plain", "Symfony\\Component\\Console\\Command\\Command", false},
		{"App\\Listener", "Symfony\\Contracts\\EventDispatcher\\EventSubscriberInterface", true},
		{"App\\Lead", "App\\AuditableInterface", true},
		{"App\\Lead", "App\\LoggableInterface", true},
		{"App\\Lead", "App\\ImportCommand", false},
		{"App\\Unknown", "App\\BaseCommand", false},
		{"", "App\\BaseCommand", false},
	}

	for _, testCase := range tests {
		if actual := registry.IsSubtypeOf(testCase.classFQN, testCase.ancestorFQN); actual != testCase.expected {
			test.Errorf("IsSubtypeOf(%q, %q) = %v, want %v", testCase.classFQN, testCase.ancestorFQN, actual, testCase.expected)
		}
	}
}

// TestRegistry_isSubtypeOfStopsAtAmbiguousName pins the conservative side of the
// walk: a name declared twice cannot be climbed through, so what sits above it
// is not claimed.
func TestRegistry_isSubtypeOfStopsAtAmbiguousName(test *testing.T) {
	registry := parse(test, `<?php
class Base extends \Vendor\Root {}
class Base extends \Vendor\Root {}
class Child extends Base {}`)

	if registry.IsSubtypeOf("Child", "Vendor\\Root") {
		test.Error("expected the walk to stop at the ambiguous Base")
	}
	if !registry.IsSubtypeOf("Child", "Base") {
		test.Error("expected the direct parent name itself to still match")
	}
}

// TestRegistry_interfaceIsNoExtendedClass keeps ExtendedClassNames meaning what
// it did: the names a *class* extends, so a rule finalizing leaf classes is not
// swayed by an interface hierarchy.
func TestRegistry_interfaceIsNoExtendedClass(test *testing.T) {
	registry := parse(test, `<?php
interface ChildInterface extends ParentInterface {}
class Service implements ChildInterface {}`)

	if registry.ExtendedClassNames()["ParentInterface"] {
		test.Error("expected an extended interface to stay out of the extended class names")
	}
}

// TestRegistry_usesTraitDeclaringProperty covers the readonly-class guard: a class
// using a trait that declares a property is flagged, since a readonly class would
// force that flattened property readonly too.
func TestRegistry_usesTraitDeclaringProperty(test *testing.T) {
	registry := parse(test, `<?php
trait HasCache
{
    private array $cache = [];
}
class Service
{
    use HasCache;
}`)

	if !registry.UsesTraitDeclaringProperty("Service") {
		test.Error("expected a class using a trait with a property to be flagged")
	}
}

// TestRegistry_usesTraitDeclaringPropertyNested reaches a property declared by a
// trait a used trait itself uses, so the guard is not fooled by one level of
// indirection.
func TestRegistry_usesTraitDeclaringPropertyNested(test *testing.T) {
	registry := parse(test, `<?php
trait HasCache
{
    private array $cache = [];
}
trait Cacheable
{
    use HasCache;
}
class Service
{
    use Cacheable;
}`)

	if !registry.UsesTraitDeclaringProperty("Service") {
		test.Error("expected a nested trait property to be flagged")
	}
}

// TestRegistry_usesTraitDeclaringPropertyNoProperty keeps the guard off a trait
// that declares only methods, so such a class stays free to collapse to readonly.
func TestRegistry_usesTraitDeclaringPropertyNoProperty(test *testing.T) {
	registry := parse(test, `<?php
trait Greets
{
    public function greet(): string { return 'hi'; }
}
class Service
{
    use Greets;
}`)

	if registry.UsesTraitDeclaringProperty("Service") {
		test.Error("expected a trait with no property to leave the class unflagged")
	}
}

// TestRegistry_hasUnresolvedTraitReference flags a class that uses a trait whose
// declaration never entered the registry — the vendored-trait case, where the
// trait's file was skipped and its members are out of view.
func TestRegistry_hasUnresolvedTraitReference(test *testing.T) {
	registry := parse(test, `<?php
class Service
{
    use HasCache;
}`)

	if !registry.HasUnresolvedTraitReference() {
		test.Error("expected a class using an uncollected trait to be flagged")
	}
}

// TestRegistry_hasUnresolvedTraitReferenceResolved keeps the flag off once every
// used trait is in the registry, so no extra vendor scan is triggered for a
// project whose traits are all in view.
func TestRegistry_hasUnresolvedTraitReferenceResolved(test *testing.T) {
	registry := parse(test, `<?php
trait HasCache
{
    private array $cache = [];
}
class Service
{
    use HasCache;
}`)

	if registry.HasUnresolvedTraitReference() {
		test.Error("expected an in-view trait to leave the class unflagged")
	}
}

// TestRegistry_collectTraitsResolvesReference mirrors folding a vendored trait in:
// the class is collected first with the trait out of view, then the trait alone is
// added from a second file, resolving the reference and exposing its property.
func TestRegistry_collectTraitsResolvesReference(test *testing.T) {
	registry := classresolver.NewRegistry()

	classRoot, err := parser.Parse([]byte(`<?php
class Service
{
    use HasCache;
}`), conf.Config{Version: &version.Version{Major: 8, Minor: 2}})
	if err != nil {
		test.Fatalf("parse class: %v", err)
	}
	registry.Collect(classRoot)

	if !registry.HasUnresolvedTraitReference() {
		test.Fatal("expected the trait unresolved before it is collected")
	}

	traitRoot, err := parser.Parse([]byte(`<?php
trait HasCache
{
    private array $cache = [];
}`), conf.Config{Version: &version.Version{Major: 8, Minor: 2}})
	if err != nil {
		test.Fatalf("parse trait: %v", err)
	}
	registry.CollectTraitsWithResolvedNames(traitRoot, classresolver.ResolveNames(traitRoot))

	if registry.HasUnresolvedTraitReference() {
		test.Error("expected the trait resolved once collected")
	}
	if !registry.UsesTraitDeclaringProperty("Service") {
		test.Error("expected the collected trait's property to be visible to the class")
	}
}

// TestRegistry_collectTraitsSkipsClasses proves the trait-only collect leaves the
// file's classes out of the registry, so vendor classes never affect class-level
// lookups such as method resolution.
func TestRegistry_collectTraitsSkipsClasses(test *testing.T) {
	registry := classresolver.NewRegistry()

	root, err := parser.Parse([]byte(`<?php
trait HasCache
{
    private array $cache = [];
}
class VendorClass
{
    public function handle($request) {}
}`), conf.Config{Version: &version.Version{Major: 8, Minor: 2}})
	if err != nil {
		test.Fatalf("parse: %v", err)
	}
	registry.CollectTraitsWithResolvedNames(root, classresolver.ResolveNames(root))

	if _, ok := registry.LookupMethodParameterNames("VendorClass", "handle"); ok {
		test.Error("expected the trait-only collect to leave classes out of the registry")
	}
}
