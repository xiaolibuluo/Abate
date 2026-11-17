
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate427_0)(int value);

typedef struct Slot427_0 {
    int slots[16];
    predicate427_0 accept;
} Slot427_0;

int slot_accept427_0(Slot427_0 *store, int value) {
    return store->accept(value) + 27;
}
int gcd427_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale427_1(int value) {
    return gcd427_1(value, 20) * 34;
}
int gcd427_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale427_2(int value) {
    return gcd427_2(value, 31) * 41;
}
int scan427_3(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 48;
}
enum Kind427_4 { KIND_EMPTY427_4 = 0, KIND_FULL427_4 = 55 };

union Cell427_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind427_4(union Cell427_4 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL427_4 : KIND_EMPTY427_4;
}
#define MASK427_5(v) (((v) << 4) | ((v) & 0x0F))

int classify427_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK427_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK427_6(v) (((v) << 4) | ((v) & 0x0F))

int classify427_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK427_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
