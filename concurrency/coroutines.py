import asyncio


async def nested_sleeper():
    await asyncio.sleep(5)
    await asyncio.sleep(2)


if __name__ == "__main__":
    loop = asyncio.get_event_loop()
    loop.run_until_complete(asyncio.wait([
        asyncio.sleep(5),
        asyncio.sleep(2),
    ]))
    # loop.run_until_complete(nested_sleeper())
