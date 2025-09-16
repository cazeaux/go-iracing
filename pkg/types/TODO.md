
# /data/league/roster

```json
{
  "type": "league_roster",
  "data": {
    "subscribed": false,
    "success": true,
    "roster_count": 144,
    "league_id": 701
  },
  "data_url": "https://scorpio-assets.s3.amazonaws.com/members/messaging-services/short_lived/league-roster/leagueid_701/roster_with_licenses.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T101043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=900&X-Amz-Signature=ad466f0ab079e63a0a86c6865e653b34418d0bee8c5e8db0764f35202f4b1845"
}
```

# /data/league/season_standings

TeamStandings remains interface, lack of data

# /data/lookup/get

The API is not well documented

# /data/member/awards

```json
{
  "type": "member_awards",
  "data": {
    "success": true,
    "cust_id": 394410,
    "award_count": 176
  },
  "data_url": "https://scorpio-assets.s3.amazonaws.com/members/messaging-services/short_lived/awards/394410.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T150131Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=60&X-Amz-Signature=d6e61ae0216f5007e6907c30d6e4a11804fa2758067bdea280ff607c144f06f8"
}
```

# /data/results/event_log

The AWS object gives chunks

```json
{
  "success": true,
  "session_info": {
    "subsession_id": 70845764,
    "session_id": 247522669,
    "simsession_number": 0,
    "simsession_type": 6,
    "simsession_name": "RACE",
    "event_type": 5,
    "event_type_name": "Race",
    "private_session_id": -1,
    "season_name": "IMSA iRacing Series - 2024 Season 3",
    "season_short_name": "2024 Season 3",
    "series_name": "IMSA iRacing Series",
    "series_short_name": "IMSA iRacing Series",
    "start_time": "2024-08-18T08:45:00Z",
    "track": {
      "config_name": "24 Heures du Mans",
      "track_id": 268,
      "track_name": "Circuit des 24 Heures du Mans"
    }
  },
  "chunk_info": {
    "chunk_size": 500,
    "num_chunks": 1,
    "rows": 13,
    "base_download_url": "https://scorpio-assets.s3.amazonaws.com/members/messaging-services/long_lived/lapdata/subsession/70845764/0/eventlog/",
    "chunk_file_names": [
      "ece964167510f8c207f4f023eec9fa60981701335710bea88d5251983e67aa16.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T154428Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=41d0891c0808573ac2f4e2db2f6ee4b1a04cb88bdafe293900da7ae8467364aa"
    ]
  }
}
```

# /data/results/lap_chart_data 

The AWS object gives chunks

```json
{
  "success": true,
  "session_info": {
    "subsession_id": 78354169,
    "session_id": 279907628,
    "simsession_number": 0,
    "simsession_type": 6,
    "simsession_name": "RACE",
    "num_laps_for_qual_average": 2,
    "num_laps_for_solo_average": 4,
    "event_type": 5,
    "event_type_name": "Race",
    "private_session_id": -1,
    "season_name": "2025 24 hours of Spa Presented by Falken Tyre",
    "season_short_name": "2025 Spa 24",
    "series_name": "24 Hours of Spa",
    "series_short_name": "24 Hours of Spa",
    "start_time": "2025-07-12T12:00:00Z",
    "track": {
      "config_name": "Endurance",
      "track_id": 525,
      "track_name": "Circuit de Spa-Francorchamps"
    }
  },
  "best_lap_num": 310,
  "best_lap_time": 1354196,
  "best_nlaps_num": -1,
  "best_nlaps_time": -1,
  "best_qual_lap_num": -1,
  "best_qual_lap_time": -1,
  "best_qual_lap_at": null,
  "chunk_info": {
    "chunk_size": 500,
    "num_chunks": 43,
    "rows": 21206,
    "base_download_url": "https://scorpio-assets.s3.amazonaws.com/members/messaging-services/long_lived/lapdata/subsession/78354169/0/chart/",
    "chunk_file_names": [
      "050b4b424143c63b8f3e79b0292619e850ee03c567737080518e53983e3a463e.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=b66e0dbb9e231ca92011add8b90fab2691ae28d00c3bd2380d8c194e609eb537",
      "c5be737c38d487ab1d5243d1e4b03d8d3a565dafb245be3de95db5695d69e153.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=04316ec8d14035fa2f8277fbdfab213c2fb946e2df34a80bad545ddb3f57ae0f",
      "fa96e2c04ab75d67843e9bb74deb7324f8138d2dfd119133429696e01b3a1d4d.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=603feaac9607d6beec3955549206628f9cba54652456a2d339cf7401c501141d",
      "ae376394b2492d1b5e679be9df8ae9e701ec6a69567944903d4090dadde47062.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=72ba89e4cae59977103f7bd4bae36bd561787b9939806bbd9ca6bdd9ad7e2f8a",
      "2218329a76c9a2edf1406f9010b9e4ab2b7d31828b31768893061edc75dc98b4.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=90090ffc4ef7607875dba7e4e2e3783f7befcb1385c43303d3297f018915944e",
      "094868d185826a1e0b7a57682df11f249ff5d31cc6dc7ee6c72529d681d3ca43.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=6f47c487efc11ed309aa95ee2cf13b9b1871a8bbcc19a470bb0dcbcbf17823e7",
      "63cf7673fcbd309d43f946fb88269b512348794bed7315faa5702a95a0ad8456.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=56bc05f2dd9f61620b9c8f755f891eade5c396d85c2297b1c281f919ecf4e619",
      "9109ff313d1fbe273cbd9e4effca648a510a0ebb6d929f1c723785ccbc066645.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=c3f27be6f85cb68abb6d47570ee2e0c646bc8de82183cb7388c62c3d0155ae66",
      "2ba6cba210ba8de06dee6b0660a8e930130ed1f8857de4b97a162bc77b6a48bd.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=50675d2567fc158cdd40db31f97a8de01cd64da79f61c87a57c0718f918bb12f",
      "3622f43724f44fc9b5f03017626d3ac881ba86d1ec0ba2f8014cba40bd482fb6.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=775972887054258c5f19494fcc82248b6402feffa5a06fca235affdc7ab70a01",
      "5134750e3494d2d22a18d961e736aabcf8044fa60b6edcab909de1c8ac91f294.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=654c017076a442fa38c272b587bc79b49ebd0713efd62764195bf09d5c3ea18d",
      "b90fd9cafb24f17384132bb546400dd1855d73225cde5403b2d6d2c30f678ce3.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=3619108c1479b9f6ee91ba1e258d6d5a9e63bf50e35f5c4e1a653672932b53aa",
      "b3deaaf4b866e3ab5907185460378900adbee44c5818af524c06848147fdbe38.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=babbe060c4d0d9d5bd3ebed15ecda044f76f775b5b8e965ada71baa5bdf99987",
      "ab2d8c032790113336103c08ca8f72ad4b3714e63a671ee74fbfb94c05df2146.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=918c6b29d4fda59e3fe0f990574beb9c944ad490cf4aad0c9ee5a4ccaa29b2a8",
      "541a0a713ea8e7e41118b86187edef03e8fad99bbadcc27f336518d38682ce18.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=961847458679293ad7333024069f5812cd7f9b47684335fd6a2d39dff024fe68",
      "5eb9dd7a04d0c0876ed9dd491cd56fe9b7cb8335e4f68f03b74915132a887d37.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=baf09ea1804da5f61ef56c580fb7fad61dd1af9c4152f3434396778cd8c54220",
      "a2ff9af3372f1736163be1ccd3c2dd3171407ffed6f0183a2056ea9430069cde.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=231cfbd25a79e5db6162c8eb6090da6a58b26229426f95f2075f06798fb83e83",
      "356ce9ffc325068e6609bc8a5a9163982ad99d58682753e7c90a074654379a27.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=025acda193195d2808fff56de7290eca62ee929278b7220d86652f943d3d40ec",
      "db237df185abe6782123b40777de0d3161c6892d2a13333bc23f500e6d19b250.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=0f94ca6b87eb4cf408e024f9232402197ee7129d951e2d49efbf18538d3a8206",
      "16b837a5ead79165c36cffe846bf8a398dc247ecd531f9a923dddfe75b6045d3.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=1b10a10e3b8d7d5e146862daae5fe1880d713be8000aff67eb0cdc6f10dd3ec2",
      "8cb524fb25cc8e21c8d33bd7a6aede09fd3a652fae25772f2bc1ecf1d60f727a.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=b57a1fdddf1c579859e68a902d8c4a19ee926b88ed4d8434ab0efef0b9af5143",
      "bd9281fb0e1965b8b7f90b8d6b2d69886e35bf67859240127e3002c353bb8dc2.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=347a82fbe8918fae6db87620183c079d28dbf6e30fec35375ea5761b726036ba",
      "dfd572485dd23f56b7e8928e86ef8ce2fd1d33c41ec001ef80b19f1ab0843869.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=3060761d127f5e2d099b8b66c1b9f4b66b7cbfa227bf7684cc08c576c4e37b39",
      "813854c517fd9fa943f9fbdce54c1f6447915c2f3f7c9ac1b4e39866e391c5ce.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=f7792faba58e03a30cd7cae705715b39ec2a95d668153d859b006035f81fa5a2",
      "94d023c043c1c536855cfcbce42feb73d4de13257bb5446b9ff7df772b5fefff.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=c3c33d54981cb394d6a3149435005f0d9d13f608ccc43514249091243e4119fc",
      "b74217dadaf081843e1e04dc88badd5a93cddd11572d9bb1de6202c5e5cefa9a.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=4c67eb4249da17a4c2c1937abe71c144a5127fe7a8f947e70cb0a285c5a0468e",
      "eed75c796d42a3f35aa052620c63d415a777bfa5664ae0c774e1401a25d8cdb7.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=00ecb4133422639baf3e54a3b317bc4f5847826565736dcac4c79e53e1d5b21a",
      "cf0bf0181395f17fa202b1227e096deea15fc82f48ef2173f8b30e8e7b87c74d.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=653323fb76ba6035239d30f2b48d4ab7615e4495b079e914d0a4786ea40c8442",
      "06cf4c16c9477c51ea235b8806a000fe4bf35027df20ee34a0c3625b91d57698.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=e9a0257165560f35417d3c6116c6d1219510facc8d7c291fbdfbbafbbc3f576d",
      "a125816331dab8c145bdb40e1dfb16a957b2a7d49f47e04a027361c8e562c249.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=691a8223cdfd9ab29dd66757040fe6dd7d51d9ad71faedfe55407cf9b20fa02a",
      "a69459f520bd8ea0dcb67e82ed18628c220a4c1046cf47cd3d9f1422af0a0a65.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=904f1da5012ba786db2280b86bf806d232e92ce0baa6d275b3a95afdaf016809",
      "2e09032f3ed27cae65a89a44539981c642815e26f678c081204f44c681ff3bec.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=32583e3bd46902e569ed43528c10f9eee6426526b6fa061da51c7394378ea69f",
      "6b47f260a1bc9099a83a09a655fac584abd43809d4682b4f251d8c93f02a993e.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=236e93d27ae81b87e66900af36524d34a500d283ad0c39b8de6c15dbda557c40",
      "68ec43d3be7f440ab8a0c29326557d30bd3aefc75e8dc44056d576224aded208.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=9d77fac18ea33db8235ff148ba0a6e1a83fc565cb6fd7530edaacaa8420a0f80",
      "fecbbda0307898c3d5ff61771b14601db2f0b7f6a88ee84b4d0b3c8ad11893a3.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=2ac8bcd362fdd48b0286152af43d6258f9a1e15e381472298f854c6883705131",
      "e47b9e7d9f8c31084baddde182058b7250fc10a6fb81405587ef930f0fc79d5a.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=d43878de7d3a72beb03089d7d778f1dfd7b7c674af3d407c933a4cdcd8b3774f",
      "d613dc02ebe3a669068f6cdc3fe70f50370f56f11ba68f8c6ee826771e68673f.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=178703cd15e5db3c93e4c9f52332f84b27b8baccecd81c05021ef5ceddede303",
      "ea2f1e60a034a7ae67670a46f147c9c3aa5d0874013a95f0418efb80c8c7d43e.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=959ea5515141e82b1efd20510266ca7f0a189c6f8d17e2d14a4e1812fae59b8c",
      "c052ee43fe6895fa580d0948c42012e08d93f4a5afd4f08aed1465de4ab46bdb.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=e78f5470d3267d8eef6a156278a74d5b8030ab2ad9fd53aa87d4137331b5cd8e",
      "01257f00a2f176b761b1cad066d29dc2dfebaebdad562dab96be7cfe6a2f4997.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=8812e41356814cd6d6ddf9b80d691539b7c7b3f63e4ad010ad956fe0509fcd32",
      "bcf5db97039e9ea9b2f22126a63ee78bc1237d5eb1c11662f40ee273c78d4fb1.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=c4531e7fe1c29cde298312416a9e1d646f380a5497211ec6763388dddd712b7e",
      "55522e4acf51d1a3f015dd7328f7922f6a9f09b3db9202bde0e730c9e1959f7b.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=bc7f1dd13f650f62bce081304ab4e882f8ac733d0eedc1aaaa3492c7edaaad07",
      "e0a2f5269e6999cc007eec744e99d1689755bdc27ab474660f7040dffbde07d6.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250914T155043Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250914%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=1e44d1ffd81705c802f9c13b8245a2b428d9fd8c7cbf7f42811910e28bd3b552"
    ]
  },
  "last_updated": "2025-09-13T07:07:52.668467488Z"
}
```

# /data/results/lap_data

Test with team id

# /data/stats/season_driver_standings

Simple CSV or full data ?

```json
{
  "success": true,
  "season_id": 5420,
  "season_name": "GT Sprint Series by Simucube - 2025 Season 2",
  "season_short_name": "2025 Season 2",
  "series_id": 228,
  "series_name": "GT Sprint Series by Simucube",
  "car_class_id": 2708,
  "race_week_num": -1,
  "division": -1,
  "customer_rank": 1321,
  "chunk_info": {
    "chunk_size": 500,
    "num_chunks": 44,
    "rows": 21710,
    "base_download_url": "https://scorpio-assets.s3.amazonaws.com/members/messaging-services/long_lived/standings2/season-driver/season/5420/2708/all/",
    "chunk_file_names": [
      "60c0573da6a6e45e624aa19773cac37c25052e8dbcf0aa1e12c6bb5ae2b0c022.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=1f3397fe18a9d27ef17e31605307aba7823f2baadc00b204e1d50da1b81e959c",
      "7e15572c36a6c024f37a636c5daf181538d87bd53a178304f7c9995d2afacc77.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=616b8bcfb420e93e986993d4dd2dc681e397c87466d322823956da36021719a5",
      "01cc0a15cb80639f23ee97882a4b558317347a2cfcc5dec5eb87feb67dfde36f.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=9454e3b5d58c32362487dd0bd2dd6ca8ae0f51c0441bdc273d53537da99502f3",
      "6d139990aa3416673024b02c65d7e59ab7db343c060a943a1277a0141b0e168a.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=68e785850b80a388d9d4133171fbd00da7fa6946f61bf7b49d847a2b27201c99",
      "b9a823b7adbd0ef2719249b35bd17f967e2e41c1beaabea82181da60143360a5.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=1331516c137722f8ce3c54f73ddecbe32627e6466fbd0b00261360ddcbebd11f",
      "a4249464eb6b8fa608bda83c0396ccf63734c992df6579c753f338c1bd7cc9a2.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=a8799b18ac17cdf3bebb9f709ad7ac790898e505c63b464856b0540b33b51658",
      "ae56a20be713da87f713cd5a3f317dc4e2cf10c76942c70b4f3d20340bea6278.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=6939e7462e66bbe8b5bda017ed928e6a3a243ea6b3da98bc2676aecbd167fcd6",
      "cae029460a9f6aea24e4c3b7032883078dd9674d95c838175c8444f17aa4695c.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=b5091a313162cc3f19d09d99220dbf6d44047346c9697ff52ea695598beb9a04",
      "276f95c5d0742cedc7427ea15c81f451803256839dccbf5790aac74d0361d30b.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=5c34101823f5d739fd1d953994d3f17c6fa744a87106c8802d3ad0f901b2a8c3",
      "0e37a5d18a30ec12607326b0f0e5d00fcf1898857a940a19bc97e9475df0664d.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=150f9b80a69c36d574ce237a7f970de8a5cd4a7f4709994f7301d7fbed3035c8",
      "c40d431975eb4e7a7a577b8b867960c0267e7648fac4a56b2ef4fa7a1c62788c.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=01eb4f0da55106e13dc36a0cab2a54144ce7d5b63c1d52225f1e743cd48e7b42",
      "981ce087317439c9b97d376dfdd975348daff61a7afc6e898d33d869ccd0f659.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=868865d6a2efca86bbb8f5c896b1ca6967be5ad71fb25206fee1d588ff450d7c",
      "696c37f12ff1961c2fa5f96303d1d3039e1e5e82d1e3c13f6b9a21d2548224c1.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=8aae718bc06b363bdd43115d30d1856cb771c64a3e0e6ecc0abbe47b72691700",
      "f1c23f0610dafe96add44da18bcfb4bc8e4b3df813b7734f61cbb12e75187e05.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=565811c935a4d4ba4dcf33620b14a49aabfee8a3dd77b68fa5b92786cf2f5954",
      "744a8253c504aacb21afeb4179b8469ee0586b6259c7e4bb64fc68f27d5b4489.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=5615f56db14189b3860c59fb9629042e47f00cfe2e403300fda52902b4090c4e",
      "43500d3e7dc223caf135dec4b58b0d1d7dc238d67de602cc01a2cdee2dfd3681.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=cd9bfbcc217e5c2dce7b145732e94167c59b9bd033aa162b857e7ce5acf65319",
      "fe7981b36cbbc57b5a320a423ee200864a7996ede2f7d77ef8f653afc2970f72.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=dff3b03f5acedd7fa1b6ff3217d52182b25f359230c159d0bddd8700887c2dd3",
      "fa3868cf641d292ad6379408f89d047c101b75e23914c1e79c2bef85b7b61aaf.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=34ce9bf2d5234a0daca7079c284da84cd7bd8fdb91a529bb00fd810455079cbf",
      "3793073ad7850c910820dd196e5ab2d95d1d4a0d1ae1d5a8ddc525b5a27a81bd.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=047f3d8a095f3f6ed404157f5aa2067b4457abadb2e3022c81a82c275db405bb",
      "e99bc8cfff699ba01d619a3e1babc2352cf0f54473741a5f45020e1adeac5bec.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=89bed90c7d406a49e2b99bc90e2928661651ac0a3cdf7873e94b7fca8b615b54",
      "17a4a2f6335cda99a428836e209132092fc0871437a3cbdaad1f369de592a380.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=f2b192d5da139231cbf2e9329155a02e9d7ceab9a949917f31792174f49a5c31",
      "09d51dade8c76bdebbbd0f8177aa6c9f041eb9dd16815873a9eecf086eebd69a.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=1041b6a9be23d960cb0710109118456cf6ff9a10ceae58ab1d0bc06311dc28a4",
      "17ebb1dc1de5ce4e6200fe7e76d4d51143c46c2521dfbdc3aa325477c28abc77.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=a9ad7f0293ad3aaba4bbffb8aa84cd690a4a9687939a5439d20271d5734becaa",
      "f9be19db6d1a9eec207e343243f2489d8ee47f98ad6ef5e2659d91268fac3bca.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=d7f7a96c822784910c7376c7bd20476735158db6f73739796ec1bc512a7015c5",
      "e21fa5ad736f2b4d66abc1f0300d6d4141c123f07ad8e43cf60d63db4acdbe88.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=e632ec094214d4b26bd5b7b761738cc0de3dcff273d824ffa02f935ddc411538",
      "94c3668f0f1e49ed824e4c2258e6d21d1fcf45d3bf8a9def1a8e58001593e6b6.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=b931aca8b94ee64331e90b009c997421c66c6ff754554cd4cc38e33f4c7a9d78",
      "26ee25953d367fb2365301704a723479bfc58de79dd91d12ddc1cb3ca22cf254.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=e9311c58bb9e158bdd11fbd553a36a00dd927bf739d3d44e70af94bfe732d73e",
      "dfcadb57133b4a741f9635e1c2ad45e63cb7a6b7dbc3ef669fc345c48b065762.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=cc2fa18587b23a2522a705aaf9219ae381db1adc216dcb7ff961f52a2555ba44",
      "707413b9c8fc31ee8d216eb74e99806942f39bc275d2f7042bf74900c71f1e2c.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=b592d9c97069ff499c01acbd4d9f79e8e7557a8e9fc5b602dca2f406bbe84f09",
      "d5e59737877bac7009f5fcc75b226f4eb048e252b6934f37581be1a4bb7d7717.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=93adac4ac1da23ea247194a74968534249b8f4f44ce1dcf3f97a43329b4f8829",
      "eed849c1b84401f5dfe8e6f705c496a31a2e0c52f75c89008e2d478ef7ba2ac2.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=3ded2f3e4e254c11f31ec046292b49c5c0a387bd47c833ce30388809f685d57a",
      "a7173c1abd925d569e264915baeb6280e888937ceab17f62ff67e9ce09b76856.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=78f4539556a3baa1711e9a2e45c5f12733afeb070d30f590757ce505fd86e455",
      "83a3e934e9f07adb6d5358b2bb843769e30a58142da4546971668b45d0d41a65.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=b1e2196d5255c38b4a30a588c6ddb929d3ba86ea49531500bc27c7ead2fda95f",
      "1741f8995887297c8e694f53d308a53f57d09547d60d77915af5c7ed9a785c4e.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=aee7bb105708b6c0eb46dc4e6118a74b3dfd53584de56dae18a3d9a2a4500a21",
      "549439b49d16c4ad35c7317515b9081b9624a764bb2883763482d374b03c82d6.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=bb4c46a4104f6e45c43c417049295500159ffe3d54c766088c9a574e70fa96f7",
      "250c364cafa40306b29bda38f430d14c4bb8030f2e2be67aeb235a7f43a346f2.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=7faf020664a7ebf5219f09e036299015611053eaae6ef7498391f6f7053d9fbd",
      "067284504ffd2c6236203bd9a09278588b23dc4df07cf93deca85f051ec3a91f.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=03e166c0014cb91e319eaf9321ac7ceaf5dad517a0c8e95e0517d880c1eb7028",
      "386186611bc3a609b4d0503a780cbea29a3e367bad7a3ebdb322c9e1dc6b7dc0.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=001d6d3bad288617040091ad2b7faed345146446dc1cab5e3c5b99752c7ac1c2",
      "24d6778146f810a91d6dfae3475bf7ab4975322f51eb4ba783df4c3a48ff5651.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=78299acd4952bea1ee2f95b06e6313ee18bc1b6b41e185da67befb637117ab3b",
      "49d3e54827ab195479c21fdd00f68013c74d0d9027a01d3c03c3acf15f1de699.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=88e1d24d701fc9803762eb6daf0d7a11c4196f303ad54aefa9013c50e4d4e077",
      "37fa3b87b9ade926497b1035d19f6a22a70e2066c6a84a31256cf2c6371b8f23.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=58389a9ce58690b0e604db3c99cd7076de4a9a2eb7382aa3e9eb9b48c02f796d",
      "e12126e7af3ac8105f31453755307e2e7209a6faa3f6fbdfd7969f65a7e4e9fc.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=7625249672e6916bdf5d0d586da67706e6684f330178cd4add0d5fc5787d9dfb",
      "0a39cc4572babf1d4062445f93265824dfb0eb92829fa7ccf236f8edfae9cf08.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=b86bef98d606a97eed0303e07a7cc22b21113ff51dc877ea69a6dc39380530d5",
      "9331e9ed3f18501d083360c863c73307f8e3dca4895293ccaaa8e12b066e5377.json?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=b9650d53f098d6d2c7ee171ec65cfd80e9888e8c11c9fb7b43f041a771f2fafd"
    ]
  },
  "last_updated": "2025-09-16T17:45:29.176036893Z",
  "csv_url": "https://scorpio-assets.s3.amazonaws.com/members/messaging-services/long_lived/standings2/season-driver/season/5420/2708/all/export.csv?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20250916T174529Z&X-Amz-SignedHeaders=host&X-Amz-Credential=AKIAUO6OO4A3WX3RTXUZ%2F20250916%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Expires=3600&X-Amz-Signature=f663a459351408d36debfb60c41bc9b50f3f72b5811e6b64a5185ae73ec7848a"
}
```