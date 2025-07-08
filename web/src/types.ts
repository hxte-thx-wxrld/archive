// generated using github.com/32leaves/bel on 2025-07-08 19:52:00.278425 +0200 CEST m=+0.002076584
// DO NOT MODIFY

export interface Artist {
    Name: string
    ArtistId: string
    ArtistPicture: string
}
// generated using github.com/32leaves/bel on 2025-07-08 19:52:00.278911 +0200 CEST m=+0.002562876
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
// generated using github.com/32leaves/bel on 2025-07-08 19:52:00.279104 +0200 CEST m=+0.002756334
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
// generated using github.com/32leaves/bel on 2025-07-08 19:52:00.279275 +0200 CEST m=+0.002926459
// DO NOT MODIFY

export interface MusicInRelease {
    TrackId: string
    Name: string
    Order: number
    ArtistName: string
}
// generated using github.com/32leaves/bel on 2025-07-08 19:52:00.27939 +0200 CEST m=+0.003042001
// DO NOT MODIFY

export interface TrackEditRequest {
    Tracktitle: string
    ReleaseDate: string
    ArtistId: string
}
// generated using github.com/32leaves/bel on 2025-07-08 19:52:00.279526 +0200 CEST m=+0.003177917
// DO NOT MODIFY

export interface UploadedTrackResponse {
    UploadId: string
}
// generated using github.com/32leaves/bel on 2025-07-08 19:52:00.279618 +0200 CEST m=+0.003270334
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
// generated using github.com/32leaves/bel on 2025-07-08 19:52:00.279734 +0200 CEST m=+0.003386001
// DO NOT MODIFY

export interface PaginatedInboxItems {
    Rows: InboxItem[]
}
