// generated using github.com/32leaves/bel on 2025-07-07 11:49:39.894861 +0200 CEST m=+0.004108501
// DO NOT MODIFY

export interface Artist {
    Name: string
    ArtistId: string
    ArtistPicture: string
}
// generated using github.com/32leaves/bel on 2025-07-07 11:49:39.89533 +0200 CEST m=+0.004577709
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
// generated using github.com/32leaves/bel on 2025-07-07 11:49:39.895472 +0200 CEST m=+0.004719126
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
// generated using github.com/32leaves/bel on 2025-07-07 11:49:39.895607 +0200 CEST m=+0.004854584
// DO NOT MODIFY

export interface MusicInRelease {
    TrackId: string
    Name: string
    Order: number
    ArtistName: string
}
// generated using github.com/32leaves/bel on 2025-07-07 11:49:39.895711 +0200 CEST m=+0.004958459
// DO NOT MODIFY

export interface TrackEditRequest {
    Tracktitle: string
    ReleaseDate: string
    ArtistId: string
}
