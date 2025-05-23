// Code generated by go generate; DO NOT EDIT

package cmds

import "testing"

func hash0(s Builder) {
	s.Hdel().Key("1").Field("1").Field("1").Build()
	s.Hexists().Key("1").Field("1").Build()
	s.Hexists().Key("1").Field("1").Cache()
	s.Hexpire().Key("1").Seconds(1).Nx().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hexpire().Key("1").Seconds(1).Xx().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hexpire().Key("1").Seconds(1).Gt().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hexpire().Key("1").Seconds(1).Lt().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hexpire().Key("1").Seconds(1).Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hexpireat().Key("1").UnixTimeSeconds(1).Nx().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hexpireat().Key("1").UnixTimeSeconds(1).Xx().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hexpireat().Key("1").UnixTimeSeconds(1).Gt().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hexpireat().Key("1").UnixTimeSeconds(1).Lt().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hexpireat().Key("1").UnixTimeSeconds(1).Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hexpiretime().Key("1").Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hget().Key("1").Field("1").Build()
	s.Hget().Key("1").Field("1").Cache()
	s.Hgetall().Key("1").Build()
	s.Hgetall().Key("1").Cache()
	s.Hgetdel().Key("1").Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hgetex().Key("1").Ex(1).Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hgetex().Key("1").Px(1).Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hgetex().Key("1").Exat(1).Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hgetex().Key("1").Pxat(1).Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hgetex().Key("1").Persist().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hgetex().Key("1").Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hincrby().Key("1").Field("1").Increment(1).Build()
	s.Hincrbyfloat().Key("1").Field("1").Increment(1).Build()
	s.Hkeys().Key("1").Build()
	s.Hkeys().Key("1").Cache()
	s.Hlen().Key("1").Build()
	s.Hlen().Key("1").Cache()
	s.Hmget().Key("1").Field("1").Field("1").Build()
	s.Hmget().Key("1").Field("1").Field("1").Cache()
	s.Hmset().Key("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hpersist().Key("1").Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hpexpire().Key("1").Milliseconds(1).Nx().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hpexpire().Key("1").Milliseconds(1).Xx().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hpexpire().Key("1").Milliseconds(1).Gt().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hpexpire().Key("1").Milliseconds(1).Lt().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hpexpire().Key("1").Milliseconds(1).Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hpexpireat().Key("1").UnixTimeMilliseconds(1).Nx().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hpexpireat().Key("1").UnixTimeMilliseconds(1).Xx().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hpexpireat().Key("1").UnixTimeMilliseconds(1).Gt().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hpexpireat().Key("1").UnixTimeMilliseconds(1).Lt().Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hpexpireat().Key("1").UnixTimeMilliseconds(1).Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hpexpiretime().Key("1").Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hpttl().Key("1").Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hrandfield().Key("1").Count(1).Withvalues().Build()
	s.Hrandfield().Key("1").Count(1).Build()
	s.Hrandfield().Key("1").Build()
	s.Hscan().Key("1").Cursor(1).Match("1").Count(1).Novalues().Build()
	s.Hscan().Key("1").Cursor(1).Match("1").Count(1).Build()
	s.Hscan().Key("1").Cursor(1).Match("1").Novalues().Build()
	s.Hscan().Key("1").Cursor(1).Match("1").Build()
	s.Hscan().Key("1").Cursor(1).Count(1).Build()
	s.Hscan().Key("1").Cursor(1).Novalues().Build()
	s.Hscan().Key("1").Cursor(1).Build()
	s.Hset().Key("1").FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Fnx().Ex(1).Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Fnx().Px(1).Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Fnx().Exat(1).Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Fnx().Pxat(1).Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Fnx().Keepttl().Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Fnx().Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Fxx().Ex(1).Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Fxx().Px(1).Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Fxx().Exat(1).Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Fxx().Pxat(1).Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Fxx().Keepttl().Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Fxx().Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Ex(1).Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Px(1).Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Exat(1).Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Pxat(1).Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Keepttl().Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetex().Key("1").Fields().Numfields(1).FieldValue().FieldValue("1", "1").FieldValue("1", "1").Build()
	s.Hsetnx().Key("1").Field("1").Value("1").Build()
	s.Hstrlen().Key("1").Field("1").Build()
	s.Hstrlen().Key("1").Field("1").Cache()
	s.Httl().Key("1").Fields().Numfields(1).Field("1").Field("1").Build()
	s.Hvals().Key("1").Build()
	s.Hvals().Key("1").Cache()
}

func TestCommand_InitSlot_hash(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { hash0(s) })
}

func TestCommand_NoSlot_hash(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { hash0(s) })
}
