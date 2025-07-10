// generated using github.com/32leaves/bel on 2025-07-11 00:13:53.518659 +0200 CEST m=+0.001925542
// DO NOT MODIFY

export interface PaginatedArtistLookup {
    Rows: Artist[]
    FullLength: number
}
// generated using github.com/32leaves/bel on 2025-07-11 00:13:53.519567 +0200 CEST m=+0.002833376
// DO NOT MODIFY

export interface Artist {
    Name: string
    ArtistId: string
    ArtistPicture: string
    Description: string
}
// generated using github.com/32leaves/bel on 2025-07-11 00:13:53.51971 +0200 CEST m=+0.002976667
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
// generated using github.com/32leaves/bel on 2025-07-11 00:13:53.519871 +0200 CEST m=+0.003137417
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
// generated using github.com/32leaves/bel on 2025-07-11 00:13:53.520008 +0200 CEST m=+0.003275334
// DO NOT MODIFY

export interface MusicInRelease {
    TrackId: string
    Name: string
    Position: number
    ArtistName: string
}
// generated using github.com/32leaves/bel on 2025-07-11 00:13:53.520135 +0200 CEST m=+0.003402001
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
// generated using github.com/32leaves/bel on 2025-07-11 00:13:53.520287 +0200 CEST m=+0.003554042
// DO NOT MODIFY

export interface PaginatedInboxItems {
    Rows: InboxItem[]
    PageCount: number
    Count: number
}
// generated using github.com/32leaves/bel on 2025-07-11 00:13:53.520398 +0200 CEST m=+0.003664709
// DO NOT MODIFY

export interface PaginatedMusicLookup {
    Rows: Music[]
    FullLength: number
    Count: number
}
// generated using github.com/32leaves/bel on 2025-07-11 00:13:53.520506 +0200 CEST m=+0.003773334
// DO NOT MODIFY

export interface Analysis {
    TrackId: string
    Tempo: number
    ZCR: number
    RMS: number
    Centroid: number
    Rolloff: number
    Flatness: number
    Duration: number
}
