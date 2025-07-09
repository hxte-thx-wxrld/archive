// generated using github.com/32leaves/bel on 2025-07-09 23:45:05.163661 +0200 CEST m=+0.001975917
// DO NOT MODIFY

export interface PaginatedArtistLookup {
    Rows: Artist[]
    FullLength: number
}
// generated using github.com/32leaves/bel on 2025-07-09 23:45:05.16445 +0200 CEST m=+0.002765334
// DO NOT MODIFY

export interface Artist {
    Name: string
    ArtistId: string
    ArtistPicture: string
}
// generated using github.com/32leaves/bel on 2025-07-09 23:45:05.164561 +0200 CEST m=+0.002876459
// DO NOT MODIFY

export interface Music {
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
// generated using github.com/32leaves/bel on 2025-07-09 23:45:05.164696 +0200 CEST m=+0.003011417
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
// generated using github.com/32leaves/bel on 2025-07-09 23:45:05.164811 +0200 CEST m=+0.003126126
// DO NOT MODIFY

export interface MusicInRelease {
    TrackId: string
    Name: string
    Order: number
    ArtistName: string
}
// generated using github.com/32leaves/bel on 2025-07-09 23:45:05.16492 +0200 CEST m=+0.003235751
// DO NOT MODIFY

export interface InboxItem {
    UploadId: string
    Uri: string
    Trackname: string
    ArtistId: string
    CreatedBy: string
    CreatedAt: string
    Status: string
    ReleaseDate: string
}
// generated using github.com/32leaves/bel on 2025-07-09 23:45:05.165046 +0200 CEST m=+0.003361792
// DO NOT MODIFY

export interface PaginatedInboxItems {
    Rows: InboxItem[]
    PageCount: number
    Count: number
}
