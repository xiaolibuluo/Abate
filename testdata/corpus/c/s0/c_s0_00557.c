
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind381_0 { KIND_EMPTY381_0 = 0, KIND_FULL381_0 = 11 };

union Cell381_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind381_0(union Cell381_0 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL381_0 : KIND_EMPTY381_0;
}
#define MASK381_1(v) (((v) << 4) | ((v) & 0x0F))

int classify381_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK381_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK381_2(v) (((v) << 4) | ((v) & 0x0F))

int classify381_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK381_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd381_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale381_3(int value) {
    return gcd381_3(value, 25) * 32;
}
int scan381_4(const int *begin, const int *end) {
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
    return total + 39;
}
typedef int (*predicate381_5)(int value);

typedef struct Slot381_5 {
    int slots[16];
    predicate381_5 accept;
} Slot381_5;

int slot_accept381_5(Slot381_5 *store, int value) {
    return store->accept(value) + 46;
}
#define MASK381_6(v) (((v) << 4) | ((v) & 0x0F))

int classify381_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK381_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
