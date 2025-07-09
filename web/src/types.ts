// generated using github.com/32leaves/bel on 2025-07-09 20:54:15.226311 +0200 CEST m=+0.002321960
// DO NOT MODIFY

export interface PaginatedArtistLookup {
    Rows: Artist[]
    FullLength: number
}
// generated using github.com/32leaves/bel on 2025-07-09 20:54:15.229342 +0200 CEST m=+0.005352793
// DO NOT MODIFY

export interface Artist {
    Name: string
    ArtistId: string
    ArtistPicture: string
}
// generated using github.com/32leaves/bel on 2025-07-09 20:54:15.229499 +0200 CEST m=+0.005509376
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
// generated using github.com/32leaves/bel on 2025-07-09 20:54:15.229644 +0200 CEST m=+0.005655293
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
// generated using github.com/32leaves/bel on 2025-07-09 20:54:15.229812 +0200 CEST m=+0.005822418
// DO NOT MODIFY

export interface MusicInRelease {
    TrackId: string
    Name: string
    Order: number
    ArtistName: string
}
// generated using github.com/32leaves/bel on 2025-07-09 20:54:15.229919 +0200 CEST m=+0.005929585
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
// generated using github.com/32leaves/bel on 2025-07-09 20:54:15.230045 +0200 CEST m=+0.006056001
// DO NOT MODIFY

export interface PaginatedInboxItems {
    Rows: InboxItem[]
    FullLength: number
}
