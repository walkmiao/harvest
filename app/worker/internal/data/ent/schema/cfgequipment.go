package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CfgEquipment holds the schema definition for the CfgEquipment entity.
type CfgEquipment struct {
	ent.Schema
}

// Fields of the CfgEquipment.
func (CfgEquipment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("equipid").Unique(),
		field.Int("portid"),
		field.Int("monitorunitid"),
		field.Int("equiptemplateid"),
		field.String("equipname"),
		field.String("exportsetting").Optional().Nillable(),
		field.Int("eventlocked").Default(0),
		field.Int("controllocked").Default(0),
		field.String("extendfield1").Optional(),
		field.String("extendfield2").Optional(),
		field.String("extendfield3").Optional(),
		field.String("extendfield4").Optional(),
		field.String("extendfield5").Optional(),
		field.Int("roomid"),
		field.Int("equiptype"),
		field.Int("sampleinterval"),
		field.String("libname"),
		field.String("description").Optional(),
		field.Int("equipaddress").Default(1),
	}
}

// Edges of the CfgEquipment.
func (CfgEquipment) Edges() []ent.Edge {
	return nil
}
func (CfgEquipment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "cfgequipment"},
	}
}
