
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

