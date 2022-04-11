var request = require('request');
var options = {
  'method': 'POST',
  'url': 'https://maicai.api.ddxq.mobi/order/getMultiReserveTime',
  'headers': {
    'Host': 'maicai.api.ddxq.mobi',
    'Connection': 'keep-alive',
    'Content-Length': '10137',
    'content-type': 'application/x-www-form-urlencoded',
    'ddmc-city-number': '0101',
    'ddmc-build-version': '2.82.0',
    'ddmc-device-id': 'osP8I0exXWR9lPqHwkq0PUnFv0us',
    'ddmc-station-id': '5bc5a86b716de1a94f8b6fb9',
    'ddmc-channel': 'applet',
    'ddmc-os-version': '[object Undefined]',
    'ddmc-app-client-id': '4',
    'Cookie': 'DDXQSESSID=16662b3a106b56be3c326c39ec12d943',
    'ddmc-ip': '',
    'ddmc-longitude': '121.404711',
    'ddmc-latitude': '31.112412',
    'ddmc-api-version': '9.49.2',
    'ddmc-uid': '5e2d814e603cd276d81116cc',
    'Accept-Encoding': 'gzip,compress,br,deflate',
    'User-Agent': 'Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.20(0x18001426) NetType/WIFI Language/zh_CN',
    'Referer': 'https://servicewechat.com/wx1e113254eda17715/422/page-frame.html'
  },
  form: {
    'uid': '5e2d814e603cd276d81116cc',
    'longitude': '121.404711',
    'latitude': '31.112412',
    'station_id': '5bc5a86b716de1a94f8b6fb9',
    'city_number': '0101',
    'api_version': '9.49.2',
    'app_version': '2.82.0',
    'applet_source': '',
    'channel': 'applet',
    'app_client_id': '4',
    'sharer_uid': '',
    's_id': '16662b3a106b56be3c326c39ec12d943',
    'openid': 'osP8I0exXWR9lPqHwkq0PUnFv0us',
    'h5_source': '',
    'device_token': 'WHJMrwNw1k/F0qdLNvE01AQF80HiIUSM/nEjZVI+rpFv2XIAibNjBHpr7A38gs2eXKTj461VexWxEO6AzYueBqw/KwG7ZhPt5dCW1tldyDzmauSxIJm5Txg==1487582755342',
    'address_id': '622e9a929062130001018b5c',
    'group_config_id': '',
    'products': '[[{"type":1,"id":"62073091d66708a57189b081","price":"11.80","count":2,"description":"","sizes":[],"cart_id":"62073091d66708a57189b081","parent_id":"","parent_batch_type":-1,"category_path":"","manage_category_path":"1916,1966,1968","activity_id":"","sku_activity_id":"","conditions_num":"","product_name":"【口口有料】保萝工坊青稞无花果软欧面包 130g/袋","product_type":0,"small_image":"https://imgnew.ddimg.mobi/product/9c355edd90064c4ea094781d99e712c6.jpg?width=800&height=800","total_price":"23.60","origin_price":"11.80","total_origin_price":"23.60","no_supplementary_price":"11.80","no_supplementary_total_price":"23.60","size_price":"0.00","buy_limit":0,"price_type":0,"promotion_num":0,"instant_rebate_money":"0.00","is_invoice":1,"sub_list":[],"is_booking":0,"is_bulk":0,"view_total_weight":"袋","net_weight":"130","net_weight_unit":"g","storage_value_id":0,"temperature_layer":"","sale_batches":{"batch_type":-1},"is_shared_station_product":0,"is_gift":0,"supplementary_list":[],"order_sort":5,"is_presale":0},{"type":1,"id":"61ceb02b750d934830ba5809","price":"12.80","count":2,"description":"","sizes":[],"cart_id":"61ceb02b750d934830ba5809","parent_id":"","parent_batch_type":-1,"category_path":"","manage_category_path":"1916,1943,1944","activity_id":"","sku_activity_id":"","conditions_num":"","product_name":"保萝工坊黄油可颂 100g/袋","product_type":0,"small_image":"https://imgnew.ddimg.mobi/product/995850ba9afd49eba696a5f79ddb6957.jpg?width=800&height=800","total_price":"25.60","origin_price":"12.80","total_origin_price":"25.60","no_supplementary_price":"12.80","no_supplementary_total_price":"25.60","size_price":"0.00","buy_limit":0,"price_type":0,"promotion_num":0,"instant_rebate_money":"0.00","is_invoice":1,"sub_list":[],"is_booking":0,"is_bulk":0,"view_total_weight":"袋","net_weight":"100","net_weight_unit":"g","storage_value_id":0,"temperature_layer":"","sale_batches":{"batch_type":-1},"is_shared_station_product":0,"is_gift":0,"supplementary_list":[],"order_sort":6,"is_presale":0},{"type":1,"id":"61c538bc8e115539dddb80a0","price":"13.80","count":2,"description":"","sizes":[],"cart_id":"61c538bc8e115539dddb80a0","parent_id":"","parent_batch_type":-1,"category_path":"","manage_category_path":"1916,1936,1938","activity_id":"","sku_activity_id":"","conditions_num":"","product_name":"保萝工坊全麦吐司 240g/袋","product_type":0,"small_image":"https://imgnew.ddimg.mobi/product/3f5f734c84b948de90bd24a6d2215c59.jpg?width=800&height=800","total_price":"27.60","origin_price":"13.80","total_origin_price":"27.60","no_supplementary_price":"13.80","no_supplementary_total_price":"27.60","size_price":"0.00","buy_limit":0,"price_type":0,"promotion_num":0,"instant_rebate_money":"0.00","is_invoice":1,"sub_list":[],"is_booking":0,"is_bulk":0,"view_total_weight":"袋","net_weight":"240","net_weight_unit":"g","storage_value_id":0,"temperature_layer":"","sale_batches":{"batch_type":-1},"is_shared_station_product":0,"is_gift":0,"supplementary_list":[],"order_sort":7,"is_presale":0},{"type":1,"id":"6114d8b7d572e7e28fb25ce7","price":"3.80","count":2,"description":"","sizes":[],"cart_id":"6114d8b7d572e7e28fb25ce7","parent_id":"","parent_batch_type":-1,"category_path":"","manage_category_path":"120,128,131","activity_id":"","sku_activity_id":"","conditions_num":"","product_name":"华英老味道鸭血（红膜）300g","product_type":0,"small_image":"https://ddfs-public.ddimg.mobi/img/blind/product-management/202108/edeeefb34ff64a41b29df798787e3774.jpg?width=800&height=800","total_price":"7.60","origin_price":"3.80","total_origin_price":"7.60","no_supplementary_price":"3.80","no_supplementary_total_price":"7.60","size_price":"0.00","buy_limit":0,"price_type":0,"promotion_num":0,"instant_rebate_money":"0.00","is_invoice":1,"sub_list":[],"is_booking":0,"is_bulk":0,"view_total_weight":"份","net_weight":"300","net_weight_unit":"g","storage_value_id":2,"temperature_layer":"4~8℃","sale_batches":{"batch_type":-1},"is_shared_station_product":0,"is_gift":0,"supplementary_list":[],"order_sort":8,"is_presale":0},{"type":1,"id":"5c9471687b37a12d3c79bf10","price":"35.90","count":1,"description":"","sizes":[],"cart_id":"5c9471687b37a12d3c79bf10","parent_id":"","parent_batch_type":-1,"category_path":"58f9e51f936edf88198b5718,58fa2059936edf89778b56e5","manage_category_path":"101,102,991","activity_id":"","sku_activity_id":"","conditions_num":"","product_name":"云南文山冷鲜黄牛牛腩 260g","product_type":0,"small_image":"https://img.ddimg.mobi/product/14893ca05f7211611235007900.jpg?width=800&height=800","total_price":"35.90","origin_price":"35.90","total_origin_price":"35.90","no_supplementary_price":"35.90","no_supplementary_total_price":"35.90","size_price":"0.00","buy_limit":0,"price_type":0,"promotion_num":0,"instant_rebate_money":"0.00","is_invoice":1,"sub_list":[],"is_booking":0,"is_bulk":0,"view_total_weight":"盒","net_weight":"260","net_weight_unit":"g","storage_value_id":2,"temperature_layer":"0~4℃","sale_batches":{"batch_type":-1},"is_shared_station_product":0,"is_gift":0,"supplementary_list":[],"order_sort":9,"is_presale":0},{"type":1,"id":"59112c60916edf46248b5064","price":"49.90","count":1,"description":"","sizes":[],"cart_id":"59112c60916edf46248b5064","parent_id":"","parent_batch_type":-1,"category_path":"58f9e574936edfe4568b5789,5e00bc90b0055a0b5b0b388e","manage_category_path":"201,221,222","activity_id":"","sku_activity_id":"","conditions_num":"","product_name":"光明优+纯牛奶 250ml*12盒/箱","product_type":0,"small_image":"https://img.ddimg.mobi/product/7ab9b81b481af1559738544004.jpg!deliver.product.list","total_price":"49.90","origin_price":"49.90","total_origin_price":"49.90","no_supplementary_price":"49.90","no_supplementary_total_price":"49.90","size_price":"0.00","buy_limit":0,"price_type":0,"promotion_num":0,"instant_rebate_money":"0.00","is_invoice":1,"sub_list":[],"is_booking":0,"is_bulk":0,"view_total_weight":"箱","net_weight":"3000","net_weight_unit":"g","storage_value_id":0,"temperature_layer":"","sale_batches":{"batch_type":-1},"is_shared_station_product":0,"is_gift":0,"supplementary_list":[],"order_sort":13,"is_presale":0}]]',
    'isBridge': 'false',
    'nars': '3b57918e3a4301940015578829b49203',
    'sesi': 'K5ZQC87a249f160fc4626f77e23dac6d65c0504'
  }
};
request(options, function (error, response) {
  if (error) throw new Error(error);
  console.log(response.body);
});
