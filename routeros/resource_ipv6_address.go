package routeros

import (
	"log"
	"strconv"

	roscl "github.com/gnewbury1/terraform-provider-routeros/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPv6Address() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPv6AddressCreate,
		Read:   resourceIPv6AddressRead,
		Update: resourceIPv6AddressUpdate,
		Delete: resourceIPv6AddressDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"actual_interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},
			"invalid": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIPv6AddressCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	ip_addr := new(roscl.IPv6Address)

	ip_addr.Address = d.Get("address").(string)
	ip_addr.Comment = d.Get("comment").(string)
	ip_addr.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	ip_addr.Interface = d.Get("interface").(string)
	ip_addr.ActualInterface = d.Get("actual_interface").(string)
	ip_addr.Invalid = strconv.FormatBool(d.Get("invalid").(bool))

	res, err := c.CreateIPv6Address(ip_addr)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PUT request to the API")
		log.Fatal(err.Error())
		return err
	}

	disabled, _ := strconv.ParseBool(res.Disabled)
	dynamic, _ := strconv.ParseBool(res.Dynamic)
	invalid, _ := strconv.ParseBool(res.Invalid)

	d.SetId(res.ID)
	d.Set("address", res.Address)
	d.Set("comment", res.Comment)
	d.Set("disabled", disabled)
	d.Set("interface", res.Interface)
	d.Set("actual_interface", res.ActualInterface)
	d.Set("dynamic", dynamic)
	d.Set("invalid", invalid)
	return nil
}

func resourceIPv6AddressRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	res, err := c.GetIPv6Address(d.Id())

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a GET request to the API")
		log.Fatal(err.Error())
		return err
	}

	disabled, _ := strconv.ParseBool(res.Disabled)
	dynamic, _ := strconv.ParseBool(res.Dynamic)
	invalid, _ := strconv.ParseBool(res.Invalid)

	d.SetId(res.ID)
	d.Set("address", res.Address)
	d.Set("comment", res.Comment)
	d.Set("disabled", disabled)
	d.Set("interface", res.Interface)
	d.Set("actual_interface", res.ActualInterface)
	d.Set("dynamic", dynamic)
	d.Set("invalid", invalid)

	return nil

}

func resourceIPv6AddressUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	ip_addr := new(roscl.IPv6Address)

	ip_addr.Address = d.Get("address").(string)
	ip_addr.Comment = d.Get("comment").(string)
	ip_addr.Disabled = strconv.FormatBool(d.Get("disabled").(bool))
	ip_addr.Interface = d.Get("interface").(string)
	ip_addr.ActualInterface = d.Get("actual_interface").(string)
	ip_addr.Invalid = strconv.FormatBool(d.Get("invalid").(bool))

	res, err := c.UpdateIPv6Address(d.Id(), ip_addr)

	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a PATCH request to the API")
		log.Fatal(err.Error())
		return err
	}

	disabled, _ := strconv.ParseBool(res.Disabled)
	dynamic, _ := strconv.ParseBool(res.Dynamic)
	invalid, _ := strconv.ParseBool(res.Invalid)

	d.SetId(res.ID)
	d.Set("address", res.Address)
	d.Set("comment", res.Comment)
	d.Set("disabled", disabled)
	d.Set("interface", res.Interface)
	d.Set("actual_interface", res.ActualInterface)
	d.Set("dynamic", dynamic)
	d.Set("invalid", invalid)
	return nil
}

func resourceIPv6AddressDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*roscl.Client)
	ip, _ := c.GetIPv6Address(d.Id())
	err := c.DeleteIPv6Address(ip)
	if err != nil {
		log.Println("[ERROR] An error was encountered while sending a DELETE request to the API")
		log.Fatal(err.Error())
		return err
	}
	d.SetId("")
	return nil
}
