"""textgen: generate random chunks of data

Usage: textgen [--chunk-size=<CHUNK_SIZE>] [--total-size=<TOTAL_SIZE>] [--random-seed=<RANDOM_SEED>]

Options:
    --chunk-size=<CHUNK_SIZE>    Size of a chunk of data [default: 10000]
    --total-size=<TOTAL_SIZE>    Total size of stuff to generate [default: 1000000]
    --random-seed=<RANDOM_SEED>  Seed for the RNG. By default will use current time [default: -1]

"""
import docopt
from io import StringIO
from random import seed, choice
import hashlib
from pathlib import Path
import asyncio
from math import ceil
from time import time
async def main(args):
    chunk_size = int(args['--chunk-size'])
    total_size = int(args['--total-size'])
    random_seed = int(args['--random-seed'])
    if random_seed == -1:
        random_seed = int(time())
    seed(random_seed)
    number = ceil(total_size / chunk_size)
    dictionary = None
    try:
        with open('dictionary.txt', 'r', encoding='utf-8') as fh:
            dictionary = [s.strip() for s in fh.readlines()]
    except Exception as e:
        raise RuntimeError(e)
    t = []
    while number > 0:
        number -= 1
        t.append(generate_chunk(dictionary, chunk_size))
    await asyncio.gather(*t)

async def generate_chunk(dictionary, chunk_size):
    s = StringIO()
    while chunk_size > s.tell():
        i = choice(dictionary)
        s.write(i)
        s.write(' ')
    
    s.write("\n")
    r = s.getvalue()
    fn = hashlib.sha1(r.encode('utf-8')).hexdigest()
    dpath = f'res-py/{fn[0:2]}/{fn[1:3]}'
    fpath = f'{dpath}/{fn}'
    Path(dpath).mkdir(parents=True, exist_ok=True)
    with open(fpath, 'w', encoding='utf-8') as fh:
        fh.write(r)


if __name__ == '__main__':
    args = docopt.docopt(__doc__)
    asyncio.run(main(args))

