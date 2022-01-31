
CREATE TYPE TASK_TYPE AS ENUM (
    'TERRAFORM_PLAN_TASK',
    'TERRAFORM_APPLY_TASK',
    'ANSIBLE_TASK',
    'PACKER_TASK',
    'CHEF_TASK'
);

-- Status goes from WAITING -> SCHEDULED -> RUNNING -> PASSED or FAILED
CREATE TYPE TASK_STATUS_TYPE AS ENUM (
    'WAITING', -- DEFAULT when the job is created but not run yet.
    'SCHEDULED', -- Scheduled to run
    'RUNNING',
    'FAILED',
    'PASSED'
);

CREATE TYPE USER_ROLE_TYPE AS ENUM (
    'ADMIN', -- User can change permissions for other users and is superuser
    'OPS',     -- User is Operations and can run Jobs
    'READONLY', -- User can only see tasks but can't run anything
    'WRITEONLY' -- User can only create tasks
    );


CREATE TABLE IF NOT EXISTS users (
     id integer GENERATED ALWAYS AS IDENTITY NOT NULL, -- Optimization is to use GUID/UUID instead of integer
     user_name  varchar NOT NULL UNIQUE,
     user_email  varchar NOT NULL UNIQUE,
     user_role USER_ROLE_TYPE NOT NULL,
     is_active boolean DEFAULT FALSE NOT NULL,
     created_at timestamp with time zone DEFAULT NOW()::timestamp NOT NULL,
     updated timestamp with time zone DEFAULT NOW()::timestamp NOT NULL,

     CONSTRAINT user_pk_id PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tasks (
    id integer GENERATED ALWAYS AS IDENTITY NOT NULL, -- Optimization is to use GUID/UUID instead of integer
    task_id BYTEA NOT NULL UNIQUE,
    task_type TASK_TYPE NOT NULL,
    priority SMALLINT NOT NULL CHECK(priority > 0 AND priority < 10),
    task_status TASK_STATUS_TYPE NOT NULL,
    data JSONB,

    git_repo VARCHAR(255) NOT NULL,
    git_commit_id VARCHAR(255) NOT NULL,
    user_created_by integer NOT NULL,
    log_file VARCHAR(255) NOT NULL, -- S3 log file location.
    is_approved boolean DEFAULT FALSE NOT NULL,
    is_completed boolean DEFAULT false NOT NULL,

    scheduled_at timestamp, -- Scheduled field will be used in mvp 2 after we get initial functionality built
    exec_at timestamp with time zone  DEFAULT NOW()::timestamp NOT NULL,
    created_at timestamp with time zone DEFAULT NOW()::timestamp NOT NULL,
    updated timestamp with time zone DEFAULT NOW()::timestamp NOT NULL,


    CONSTRAINT tasks_pk_id PRIMARY KEY (id),
    CONSTRAINT task_user_fk_created_by FOREIGN KEY (user_created_by) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT task_unique_per_user UNIQUE (user_created_by,task_id)
);

