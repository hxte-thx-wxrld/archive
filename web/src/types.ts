// generated using github.com/32leaves/bel on 2025-07-09 14:36:06.878561 +0200 CEST m=+0.002265043
// DO NOT MODIFY

export interface Artist {
    Name: string
    ArtistId: string
    ArtistPicture: string
}
// generated using github.com/32leaves/bel on 2025-07-09 14:36:06.879358 +0200 CEST m=+0.003061543
// DO NOT MODIFY

export interface Release {
    Name: string
    ReleaseDate: string
    Isrc: string
    ReleaseId: string
    CatalogId: string
    CoverUrl: string
    RelatedMusic: MusicInRelease[]
}
// generated using github.com/32leaves/bel on 2025-07-09 14:36:06.879536 +0200 CEST m=+0.003239126
// DO NOT MODIFY

export interface MusicRow {
    TrackId: string
    Tracktitle: string
    Artist: string
    ArtistId: string
    CatalogNo: string
    ReleaseDate: string
    PublicUrl: string
    ReleaseId: string
    CoverUrl: string
    Length: string
}
// generated using github.com/32leaves/bel on 2025-07-09 14:36:06.879655 +0200 CEST m=+0.003358668
// DO NOT MODIFY

export interface MusicInRelease {
    TrackId: string
    Name: string
    Order: number
    ArtistName: string
}
// generated using github.com/32leaves/bel on 2025-07-09 14:36:06.879769 +0200 CEST m=+0.003473043
// DO NOT MODIFY

export interface TrackEditRequest {
    Tracktitle: string
    ReleaseDate: string
    ArtistId: string
}
// generated using github.com/32leaves/bel on 2025-07-09 14:36:06.879895 +0200 CEST m=+0.003598835
// DO NOT MODIFY

export interface UploadedTrackResponse {
    UploadId: string
}
// generated using github.com/32leaves/bel on 2025-07-09 14:36:06.879987 +0200 CEST m=+0.003690376
// DO NOT MODIFY

export interface InboxItem {
    UploadId: string
    Uri: string
    Trackname: string
    ArtistId: string
    CreatedBy: string
    CreatedAt: string
    Status: string
}
// generated using github.com/32leaves/bel on 2025-07-09 14:36:06.880088 +0200 CEST m=+0.003791960
// DO NOT MODIFY

export interface PaginatedInboxItems {
    Rows: InboxItem[]
    FullLength: number
}
