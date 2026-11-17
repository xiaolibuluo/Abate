
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind673_0 { KIND_EMPTY673_0 = 0, KIND_FULL673_0 = 24 };

union Cell673_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind673_0(union Cell673_0 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL673_0 : KIND_EMPTY673_0;
}
typedef int (*predicate673_1)(int value);

typedef struct Slot673_1 {
    int slots[16];
    predicate673_1 accept;
} Slot673_1;

int slot_accept673_1(Slot673_1 *store, int value) {
    return store->accept(value) + 31;
}
typedef int (*predicate673_2)(int value);

typedef struct Slot673_2 {
    int slots[16];
    predicate673_2 accept;
} Slot673_2;

int slot_accept673_2(Slot673_2 *store, int value) {
    return store->accept(value) + 38;
}
int scan673_3(const int *begin, const int *end) {
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
    return total + 45;
}
#define MASK673_4(v) (((v) << 4) | ((v) & 0x0F))

int classify673_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK673_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd673_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale673_5(int value) {
    return gcd673_5(value, 17) * 59;
}
#define MASK673_6(v) (((v) << 4) | ((v) & 0x0F))

int classify673_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK673_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
