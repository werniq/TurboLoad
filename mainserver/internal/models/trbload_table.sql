CREATEA TABLE statistics(
    total_downloads bigint,
    megabits_transfered bigint,
    concurrent_downloads bigint,
    average_download_time bigint
);

CREATE TABLE file_info(
    id bigint primary key not null,
    filename varchar(255) not null,
    downloads_count int,
    size bigint,
    created_at timestamp not null default current_timestamp
);

INSERT INTO file_info(id, filename) VALUES(1, '1GB.bin');
INSERT INTO file_info(id, filename) VALUES(2, '10GB.bin');