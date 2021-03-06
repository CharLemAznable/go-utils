package str_test

import (
	"testing"

	"github.com/bingoohuang/gou/str"

	"github.com/stretchr/testify/assert"
)

func TestSplitX0(t *testing.T) {
	assert.Equal(t, []string{"a", "b"}, str.SplitX("a;b;", ";"))
	assert.Equal(t, []string{"(a;b;)"}, str.SplitX("(a;b;)", ";"))
	assert.Equal(t, []string{"([a;b;])"}, str.SplitX("([a;b;])", ";"))
	assert.Equal(t, []string{"([a;b;];)", "abc"}, str.SplitX("([a;b;];);abc", ";"))
	assert.Equal(t, []string{`([a;b;];)\;abc`}, str.SplitX(`([a;b;];)\;abc`, ";"))
}

func TestSplitSql(t *testing.T) {
	sql := "create table aaa; drop table aaa;"
	sqls := str.SplitX(sql, ";")

	assert.Equal(t, []string{"create table aaa", "drop table aaa"}, sqls)
}

func TestSplitSql2(t *testing.T) {
	sql := "ADD COLUMN `PREFERENTIAL_WAY` CHAR(3) NULL COMMENT '优\\惠方式:0:现金券;1:减免,2:赠送金额 ;' AFTER `PAY_TYPE`;"
	sqls := str.SplitX(sql, ";")

	assert.Equal(t, []string{"ADD COLUMN `PREFERENTIAL_WAY` CHAR(3) NULL " +
		"COMMENT '优\\惠方式:0:现金券;1:减免,2:赠送金额 ;' AFTER `PAY_TYPE`"}, sqls)
}

func TestSplitSql3(t *testing.T) {
	sql := "ALTER TABLE `tt_l_mbrcard_chg`; \n" +
		"ADD COLUMN `PREFERENTIAL_WAY` CHAR(3) NULL COMMENT '优惠方式:''0:现金券;1:减免,2:赠送金额 ;' AFTER `PAY_TYPE`; "
	sqls := str.SplitX(sql, ";")

	assert.Equal(t, []string{"ALTER TABLE `tt_l_mbrcard_chg`",
		"ADD COLUMN `PREFERENTIAL_WAY` CHAR(3) NULL " +
			"COMMENT '优惠方式:''0:现金券;1:减免,2:赠送金额 ;' AFTER `PAY_TYPE`"}, sqls)
}

func TestSplitSql4(t *testing.T) {
	sql := "hello {aaa,bbb}, hello (ccc,ddd)"
	sqls := str.SplitX(sql, ",")

	assert.Equal(t, []string{"hello {aaa,bbb}", "hello (ccc,ddd)"}, sqls)
}
