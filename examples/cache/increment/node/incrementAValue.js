const { CacheService } = require("m3o/cache");

// Increment a value (if it's a number). If key not found it is equivalent to set.
async function incrementAvalue() {
  let cacheService = new CacheService(process.env.MICRO_API_TOKEN);
  let rsp = await cacheService.increment({
    key: "counter",
    value: 2,
  });
  console.log(rsp);
}

incrementAvalue();
