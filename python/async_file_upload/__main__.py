import asyncio
from time import time

import aiohttp


def write_file(binary_file: bytes) -> None:
    filename = f"{time()}.mp3"
    with open(filename, "wb") as file:
        file.write(binary_file)


async def fetch_content(url: str, session: aiohttp.ClientSession):
    async with session.get(url, allow_redirects=True) as response:
        binary_file = await response.read()
        write_file(binary_file)


async def main() -> None:
    urls = [
        """ Add download links """
    ]
    tasks = []

    async with aiohttp.ClientSession() as session:
        for url in urls:
            print(url)
            task = asyncio.create_task(fetch_content(url, session))
            tasks.append(task)

        await asyncio.gather(*tasks)


if __name__ == '__main__':
    asyncio.run(main())
