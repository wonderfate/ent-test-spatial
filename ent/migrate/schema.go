// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LocationsColumns holds the columns for the "locations" table.
	LocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "coords", Type: field.TypeOther, SchemaType: map[string]string{"mysql": "POINT"}},
		{Name: "location_children", Type: field.TypeInt, Nullable: true},
	}
	// LocationsTable holds the schema information for the "locations" table.
	LocationsTable = &schema.Table{
		Name:       "locations",
		Columns:    LocationsColumns,
		PrimaryKey: []*schema.Column{LocationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "locations_locations_children",
				Columns:    []*schema.Column{LocationsColumns[3]},
				RefColumns: []*schema.Column{LocationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LocationsTable,
	}
)

func init() {
	LocationsTable.ForeignKeys[0].RefTable = LocationsTable
}
