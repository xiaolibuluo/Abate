
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1136_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def pair1136_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 36}
async def fetch1136_2(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
def project1136_3(values):
    select = lambda item: item + 50
    return [select(item) for item in values if item != 0]
def traced1136_4(value: int) -> int:
    return value + 57


def decorate1136_4(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1136_4
def doubled1136_4(value: int) -> int:
    return value + value
