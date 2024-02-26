// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package solidserver

import (
	"context"
	"reflect"

	"errors"
	"github.com/AthenaWolfe/pulumi-solidserver/sdk/go/solidserver/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Application Pool resource allows to create and manage a pool that implement a traffic policy.
// Application Pools are groups of nodes serving the same application and monitored by the GSLB(s) DNS servers
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/AthenaWolfe/pulumi-solidserver/sdk/go/solidserver"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := solidserver.NewAppPool(ctx, "myFirstPool", &solidserver.AppPoolArgs{
//				Application:             pulumi.Any(solidserver_app_application.MyFirstApplicaton.Name),
//				Fqdn:                    pulumi.Any(solidserver_app_application.MyFirstApplicaton.Fqdn),
//				LbMode:                  pulumi.Any(latency),
//				Affinity:                pulumi.Bool(true),
//				AffinitySessionDuration: pulumi.Int(300),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type AppPool struct {
	pulumi.CustomResourceState

	// Enable session affinity for the application pool.
	Affinity pulumi.BoolPtrOutput `pulumi:"affinity"`
	// The time each session is maintained in sec (Default: 300).
	AffinitySessionDuration pulumi.IntPtrOutput `pulumi:"affinitySessionDuration"`
	// The name of the application associated to the pool.
	Application pulumi.StringOutput `pulumi:"application"`
	// Number of best active nodes when lbMode is set to latency.
	BestActiveNodes pulumi.IntPtrOutput `pulumi:"bestActiveNodes"`
	// The fqdn of the application associated to the pool.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The IP protocol version used by the application pool to create (Supported: ipv4, ipv6; Default: ipv4).
	IpVersion pulumi.StringPtrOutput `pulumi:"ipVersion"`
	// The load balancing mode of the application pool to create (Supported: weighted,round-robin,latency; Default: round-robin).
	LbMode pulumi.StringPtrOutput `pulumi:"lbMode"`
	// The name of the application pool to create.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAppPool registers a new resource with the given unique name, arguments, and options.
func NewAppPool(ctx *pulumi.Context,
	name string, args *AppPoolArgs, opts ...pulumi.ResourceOption) (*AppPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Application == nil {
		return nil, errors.New("invalid value for required argument 'Application'")
	}
	if args.Fqdn == nil {
		return nil, errors.New("invalid value for required argument 'Fqdn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AppPool
	err := ctx.RegisterResource("solidserver:index/appPool:AppPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppPool gets an existing AppPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppPoolState, opts ...pulumi.ResourceOption) (*AppPool, error) {
	var resource AppPool
	err := ctx.ReadResource("solidserver:index/appPool:AppPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppPool resources.
type appPoolState struct {
	// Enable session affinity for the application pool.
	Affinity *bool `pulumi:"affinity"`
	// The time each session is maintained in sec (Default: 300).
	AffinitySessionDuration *int `pulumi:"affinitySessionDuration"`
	// The name of the application associated to the pool.
	Application *string `pulumi:"application"`
	// Number of best active nodes when lbMode is set to latency.
	BestActiveNodes *int `pulumi:"bestActiveNodes"`
	// The fqdn of the application associated to the pool.
	Fqdn *string `pulumi:"fqdn"`
	// The IP protocol version used by the application pool to create (Supported: ipv4, ipv6; Default: ipv4).
	IpVersion *string `pulumi:"ipVersion"`
	// The load balancing mode of the application pool to create (Supported: weighted,round-robin,latency; Default: round-robin).
	LbMode *string `pulumi:"lbMode"`
	// The name of the application pool to create.
	Name *string `pulumi:"name"`
}

type AppPoolState struct {
	// Enable session affinity for the application pool.
	Affinity pulumi.BoolPtrInput
	// The time each session is maintained in sec (Default: 300).
	AffinitySessionDuration pulumi.IntPtrInput
	// The name of the application associated to the pool.
	Application pulumi.StringPtrInput
	// Number of best active nodes when lbMode is set to latency.
	BestActiveNodes pulumi.IntPtrInput
	// The fqdn of the application associated to the pool.
	Fqdn pulumi.StringPtrInput
	// The IP protocol version used by the application pool to create (Supported: ipv4, ipv6; Default: ipv4).
	IpVersion pulumi.StringPtrInput
	// The load balancing mode of the application pool to create (Supported: weighted,round-robin,latency; Default: round-robin).
	LbMode pulumi.StringPtrInput
	// The name of the application pool to create.
	Name pulumi.StringPtrInput
}

func (AppPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*appPoolState)(nil)).Elem()
}

type appPoolArgs struct {
	// Enable session affinity for the application pool.
	Affinity *bool `pulumi:"affinity"`
	// The time each session is maintained in sec (Default: 300).
	AffinitySessionDuration *int `pulumi:"affinitySessionDuration"`
	// The name of the application associated to the pool.
	Application string `pulumi:"application"`
	// Number of best active nodes when lbMode is set to latency.
	BestActiveNodes *int `pulumi:"bestActiveNodes"`
	// The fqdn of the application associated to the pool.
	Fqdn string `pulumi:"fqdn"`
	// The IP protocol version used by the application pool to create (Supported: ipv4, ipv6; Default: ipv4).
	IpVersion *string `pulumi:"ipVersion"`
	// The load balancing mode of the application pool to create (Supported: weighted,round-robin,latency; Default: round-robin).
	LbMode *string `pulumi:"lbMode"`
	// The name of the application pool to create.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a AppPool resource.
type AppPoolArgs struct {
	// Enable session affinity for the application pool.
	Affinity pulumi.BoolPtrInput
	// The time each session is maintained in sec (Default: 300).
	AffinitySessionDuration pulumi.IntPtrInput
	// The name of the application associated to the pool.
	Application pulumi.StringInput
	// Number of best active nodes when lbMode is set to latency.
	BestActiveNodes pulumi.IntPtrInput
	// The fqdn of the application associated to the pool.
	Fqdn pulumi.StringInput
	// The IP protocol version used by the application pool to create (Supported: ipv4, ipv6; Default: ipv4).
	IpVersion pulumi.StringPtrInput
	// The load balancing mode of the application pool to create (Supported: weighted,round-robin,latency; Default: round-robin).
	LbMode pulumi.StringPtrInput
	// The name of the application pool to create.
	Name pulumi.StringPtrInput
}

func (AppPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appPoolArgs)(nil)).Elem()
}

type AppPoolInput interface {
	pulumi.Input

	ToAppPoolOutput() AppPoolOutput
	ToAppPoolOutputWithContext(ctx context.Context) AppPoolOutput
}

func (*AppPool) ElementType() reflect.Type {
	return reflect.TypeOf((**AppPool)(nil)).Elem()
}

func (i *AppPool) ToAppPoolOutput() AppPoolOutput {
	return i.ToAppPoolOutputWithContext(context.Background())
}

func (i *AppPool) ToAppPoolOutputWithContext(ctx context.Context) AppPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppPoolOutput)
}

// AppPoolArrayInput is an input type that accepts AppPoolArray and AppPoolArrayOutput values.
// You can construct a concrete instance of `AppPoolArrayInput` via:
//
//	AppPoolArray{ AppPoolArgs{...} }
type AppPoolArrayInput interface {
	pulumi.Input

	ToAppPoolArrayOutput() AppPoolArrayOutput
	ToAppPoolArrayOutputWithContext(context.Context) AppPoolArrayOutput
}

type AppPoolArray []AppPoolInput

func (AppPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppPool)(nil)).Elem()
}

func (i AppPoolArray) ToAppPoolArrayOutput() AppPoolArrayOutput {
	return i.ToAppPoolArrayOutputWithContext(context.Background())
}

func (i AppPoolArray) ToAppPoolArrayOutputWithContext(ctx context.Context) AppPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppPoolArrayOutput)
}

// AppPoolMapInput is an input type that accepts AppPoolMap and AppPoolMapOutput values.
// You can construct a concrete instance of `AppPoolMapInput` via:
//
//	AppPoolMap{ "key": AppPoolArgs{...} }
type AppPoolMapInput interface {
	pulumi.Input

	ToAppPoolMapOutput() AppPoolMapOutput
	ToAppPoolMapOutputWithContext(context.Context) AppPoolMapOutput
}

type AppPoolMap map[string]AppPoolInput

func (AppPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppPool)(nil)).Elem()
}

func (i AppPoolMap) ToAppPoolMapOutput() AppPoolMapOutput {
	return i.ToAppPoolMapOutputWithContext(context.Background())
}

func (i AppPoolMap) ToAppPoolMapOutputWithContext(ctx context.Context) AppPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppPoolMapOutput)
}

type AppPoolOutput struct{ *pulumi.OutputState }

func (AppPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppPool)(nil)).Elem()
}

func (o AppPoolOutput) ToAppPoolOutput() AppPoolOutput {
	return o
}

func (o AppPoolOutput) ToAppPoolOutputWithContext(ctx context.Context) AppPoolOutput {
	return o
}

// Enable session affinity for the application pool.
func (o AppPoolOutput) Affinity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppPool) pulumi.BoolPtrOutput { return v.Affinity }).(pulumi.BoolPtrOutput)
}

// The time each session is maintained in sec (Default: 300).
func (o AppPoolOutput) AffinitySessionDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppPool) pulumi.IntPtrOutput { return v.AffinitySessionDuration }).(pulumi.IntPtrOutput)
}

// The name of the application associated to the pool.
func (o AppPoolOutput) Application() pulumi.StringOutput {
	return o.ApplyT(func(v *AppPool) pulumi.StringOutput { return v.Application }).(pulumi.StringOutput)
}

// Number of best active nodes when lbMode is set to latency.
func (o AppPoolOutput) BestActiveNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AppPool) pulumi.IntPtrOutput { return v.BestActiveNodes }).(pulumi.IntPtrOutput)
}

// The fqdn of the application associated to the pool.
func (o AppPoolOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *AppPool) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// The IP protocol version used by the application pool to create (Supported: ipv4, ipv6; Default: ipv4).
func (o AppPoolOutput) IpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppPool) pulumi.StringPtrOutput { return v.IpVersion }).(pulumi.StringPtrOutput)
}

// The load balancing mode of the application pool to create (Supported: weighted,round-robin,latency; Default: round-robin).
func (o AppPoolOutput) LbMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppPool) pulumi.StringPtrOutput { return v.LbMode }).(pulumi.StringPtrOutput)
}

// The name of the application pool to create.
func (o AppPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type AppPoolArrayOutput struct{ *pulumi.OutputState }

func (AppPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AppPool)(nil)).Elem()
}

func (o AppPoolArrayOutput) ToAppPoolArrayOutput() AppPoolArrayOutput {
	return o
}

func (o AppPoolArrayOutput) ToAppPoolArrayOutputWithContext(ctx context.Context) AppPoolArrayOutput {
	return o
}

func (o AppPoolArrayOutput) Index(i pulumi.IntInput) AppPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AppPool {
		return vs[0].([]*AppPool)[vs[1].(int)]
	}).(AppPoolOutput)
}

type AppPoolMapOutput struct{ *pulumi.OutputState }

func (AppPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AppPool)(nil)).Elem()
}

func (o AppPoolMapOutput) ToAppPoolMapOutput() AppPoolMapOutput {
	return o
}

func (o AppPoolMapOutput) ToAppPoolMapOutputWithContext(ctx context.Context) AppPoolMapOutput {
	return o
}

func (o AppPoolMapOutput) MapIndex(k pulumi.StringInput) AppPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AppPool {
		return vs[0].(map[string]*AppPool)[vs[1].(string)]
	}).(AppPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppPoolInput)(nil)).Elem(), &AppPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppPoolArrayInput)(nil)).Elem(), AppPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AppPoolMapInput)(nil)).Elem(), AppPoolMap{})
	pulumi.RegisterOutputType(AppPoolOutput{})
	pulumi.RegisterOutputType(AppPoolArrayOutput{})
	pulumi.RegisterOutputType(AppPoolMapOutput{})
}