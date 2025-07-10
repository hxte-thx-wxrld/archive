// generated using github.com/32leaves/bel on 2025-07-10 15:29:47.657006 +0200 CEST m=+0.002898251
// DO NOT MODIFY

export interface PaginatedArtistLookup {
    Rows: Artist[]
    FullLength: number
}
// generated using github.com/32leaves/bel on 2025-07-10 15:29:47.65771 +0200 CEST m=+0.003602626
// DO NOT MODIFY

export interface Artist {
    Name: string
    ArtistId: string
    ArtistPicture: string
    Description: string
}
// generated using github.com/32leaves/bel on 2025-07-10 15:29:47.657816 +0200 CEST m=+0.003708001
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
// generated using github.com/32leaves/bel on 2025-07-10 15:29:47.657946 +0200 CEST m=+0.003838292
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
// generated using github.com/32leaves/bel on 2025-07-10 15:29:47.658054 +0200 CEST m=+0.003946126
// DO NOT MODIFY

export interface MusicInRelease {
    TrackId: string
    Name: string
    Position: number
    ArtistName: string
}
// generated using github.com/32leaves/bel on 2025-07-10 15:29:47.658199 +0200 CEST m=+0.004091626
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
// generated using github.com/32leaves/bel on 2025-07-10 15:29:47.658314 +0200 CEST m=+0.004206376
// DO NOT MODIFY

export interface PaginatedInboxItems {
    Rows: InboxItem[]
    PageCount: number
    Count: number
}
