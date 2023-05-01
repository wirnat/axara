export class BaseModel {
    ID: bigint;
    UUID: string;
    createdAt: string;
    createdBy: string;
    updatedAt: string;
    updatedBy: string;
    deletedAt: string;
    deletedBy: string;

    constructor(
        ID: bigint,
        UUID: string,
        createdAt: string,
        createdBy: string,
        updatedAt: string,
        updatedBy: string,
        deletedAt: string,
        deletedBy: string
    ) {
        this.ID = ID;
        this.UUID = UUID;
        this.createdAt = createdAt;
        this.createdBy = createdBy;
        this.updatedAt = updatedAt;
        this.updatedBy = updatedBy;
        this.deletedAt = deletedAt;
        this.deletedBy = deletedBy;
    }
}
