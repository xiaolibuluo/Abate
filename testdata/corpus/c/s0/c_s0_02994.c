
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2025_0)(int value);

typedef struct Slot2025_0 {
    int slots[16];
    predicate2025_0 accept;
} Slot2025_0;

int slot_accept2025_0(Slot2025_0 *store, int value) {
    return store->accept(value) + 43;
}
#define MASK2025_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2025_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2025_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2025_2(const int *begin, const int *end) {
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
    return total + 57;
}
int gcd2025_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2025_3(int value) {
    return gcd2025_3(value, 65) * 64;
}
typedef int (*predicate2025_4)(int value);

typedef struct Slot2025_4 {
    int slots[16];
    predicate2025_4 accept;
} Slot2025_4;

int slot_accept2025_4(Slot2025_4 *store, int value) {
    return store->accept(value) + 71;
}
enum Kind2025_5 { KIND_EMPTY2025_5 = 0, KIND_FULL2025_5 = 78 };

union Cell2025_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2025_5(union Cell2025_5 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL2025_5 : KIND_EMPTY2025_5;
}
