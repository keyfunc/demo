-- 用户表
CREATE TABLE users (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
name VARCHAR(64),
avatar TEXT,
status SMALLINT NOT NULL DEFAULT 1,
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMENT ON TABLE users IS '用户表';
COMMENT ON COLUMN users.id IS '用户唯一ID';
COMMENT ON COLUMN users.name IS '昵称';
COMMENT ON COLUMN users.avatar IS '头像URL';
COMMENT ON COLUMN users.status IS '状态：1=正常，2=禁用';
COMMENT ON COLUMN users.created_at IS '创建时间';
COMMENT ON COLUMN users.updated_at IS '更新时间';

-- 用户认证信息表
CREATE TABLE user_auths (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
user_id UUID REFERENCES users(id) ON DELETE CASCADE,
provider VARCHAR(32) NOT NULL,
identifier VARCHAR(255) NOT NULL,
credential TEXT,
last_login_at TIMESTAMPTZ,
created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE user_auths IS '用户认证信息表';
COMMENT ON COLUMN user_auths.user_id IS '所属用户';
COMMENT ON COLUMN user_auths.provider IS '登录方式';
COMMENT ON COLUMN user_auths.identifier IS '登录账号';
COMMENT ON COLUMN user_auths.credential IS '认证凭证（密码Hash等）';
COMMENT ON COLUMN user_auths.last_login_at IS '最后登录时间';
COMMENT ON COLUMN users.created_at IS '创建时间';
COMMENT ON COLUMN users.updated_at IS '更新时间';