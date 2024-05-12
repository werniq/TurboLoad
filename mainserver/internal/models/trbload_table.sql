CREATEA TABLE statistics(
    total_downloads bigint,
    megabits_transfered bigint,
    concurrent_downloads bigint,
    average_download_time bigint
);

CREATE TABLE file_info(
    id bigint not null primary key,
    filename varchar(255) not null,
    downloads_count int,
    size bigint,
    created_at timestamp not null default current_timestamp
);

