import { BaseModel } from './base_model';

export class CompanyModel extends BaseModel {
    name: string;
    parent: string;
    parentUUID: string;
    subCompany: CompanyModel;
    type: string;

    constructor(
        ID: bigint, //@meta validate:true
        UUID: string,
        createdAt: string,
        createdBy: string,
        updatedAt: string,
        updatedBy: string,
        deletedAt: string,
        deletedBy: string,
        name: string,
        parent: string,
        parentUUID: string,
        subCompany: CompanyModel,
        type: string
    ) {
        super(
            ID,
            UUID,
            createdAt,
            createdBy,
            updatedAt,
            updatedBy,
            deletedAt,
            deletedBy
        );
        this.name = name;
        this.parent = parent;
        this.parentUUID = parentUUID;
        this.subCompany = subCompany;
        this.type = type;
    }
}
//@Register Company