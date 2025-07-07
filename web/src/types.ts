// generated using github.com/32leaves/bel on 2025-07-07 22:15:30.451844 +0200 CEST m=+0.006063917
// DO NOT MODIFY

export interface Artist {
    Name: string
    ArtistId: string
    ArtistPicture: string
}
// generated using github.com/32leaves/bel on 2025-07-07 22:15:30.453155 +0200 CEST m=+0.007374876
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
// generated using github.com/32leaves/bel on 2025-07-07 22:15:30.453564 +0200 CEST m=+0.007784209
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
// generated using github.com/32leaves/bel on 2025-07-07 22:15:30.45385 +0200 CEST m=+0.008070417
// DO NOT MODIFY

export interface MusicInRelease {
    TrackId: string
    Name: string
    Order: number
    ArtistName: string
}
// generated using github.com/32leaves/bel on 2025-07-07 22:15:30.454073 +0200 CEST m=+0.008292709
// DO NOT MODIFY

export interface TrackEditRequest {
    Tracktitle: string
    ReleaseDate: string
    ArtistId: string
}
// generated using github.com/32leaves/bel on 2025-07-07 22:15:30.454328 +0200 CEST m=+0.008547834
// DO NOT MODIFY

export interface UploadedTrackResponse {
    TrackId: string
}
