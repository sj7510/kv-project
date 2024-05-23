use std::result;
use thiserror::Error;

#[derive(Error, Debug)]
pub enum Errors {
    #[error("failed to read form data file")]
    FailedReadFromDAtaFile,

    #[error("failed to write to data file")]
    FailedWriteToDAtaFile,

    #[error("failed to sync data file")]
    FailedSyncDAtaFile,

    #[error("failed to open data file")]
    FailedToOpenDataFile,
}

pub type Result<T> = result::Result<T, Errors>;
