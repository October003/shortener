package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shortener/pkg/base62"
	"shortener/pkg/md5"
	"shortener/pkg/urltool"

	"shortener/internal/svc"
	"shortener/internal/types"
	"shortener/pkg/connect"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Convert 转链业务逻辑: 输入一个长链接 --> 转为短链接
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 1. 校验输入的数据 长链接
	// 1.1 数据不能为null
	// if len(req.LongUrl) == 0 {}
	// 使用validator包用来做参数校验
	// 1.2 输入的长链接是有效的、能请求通的
	if ok := connect.Get(req.LongUrl); !ok {
		return nil, errors.New("无效的链接")
	}
	// 1.3 判断之前是否已经转过 (数据库中是否已存在该长链接)
	// 1.3.1 给长链接生成md5
	md5Value := md5.Sum([]byte(req.LongUrl)) // 这里使用的是项目中封装的md5包
	// 1.3.2 拿md5去数据库中查找是否存在
	u, err := l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, sql.NullString{String: md5Value, Valid: true})
	if err != sqlx.ErrNotFound {
		if err == nil {
			return nil, fmt.Errorf("该链接已被转为%s", u.Surl.String)
		}
		logx.Errorw("ShortUrlModel.FindOneByMd5 failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	// 1.4 输入的不能是一个短链接 (避免循环转链)
	// 输入的是一个完整的url october.cn/1d12a?name=october
	basePath, err := urltool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("urltool.GetBasePath failed", logx.LogField{Key: "lurl", Value: req.LongUrl}, logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	_, err = l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: basePath, Valid: true})
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, errors.New("该链接已经是短链了")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.LogField{Key: "err", Value: err.Error()})
	}
	// 2. 取号 转链  基于MySQL实现的发号器
	// 每来一个转链请求，我们就使用 REPLACE INTO 语句往sequence 表中插入一条数据，并且取出主键id作为号码
	seq, err := l.svcCtx.Sequence.Next()
	if err != nil {
		logx.Errorw("Sequence.Next() failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	fmt.Println(seq)
	// 3. 号码转短链
	// 3.1 安全性   1En = 6347
	// 3.2 短域名避免某些特殊词
	short := base62.Int2String(seq)

	// 4. 存储长链接映射关系

	// 5. 返回响应
	return
}
