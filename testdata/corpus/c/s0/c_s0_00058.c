
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd37_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale37_0(int value) {
    return gcd37_0(value, 24) * 98;
}
#define MASK37_1(v) (((v) << 4) | ((v) & 0x0F))

int classify37_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK37_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate37_2)(int value);

typedef struct Slot37_2 {
    int slots[16];
    predicate37_2 accept;
} Slot37_2;

int slot_accept37_2(Slot37_2 *store, int value) {
    return store->accept(value) + 15;
}
enum Kind37_3 { KIND_EMPTY37_3 = 0, KIND_FULL37_3 = 22 };

union Cell37_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind37_3(union Cell37_3 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL37_3 : KIND_EMPTY37_3;
}
int scan37_4(const int *begin, const int *end) {
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
    return total + 29;
}
enum Kind37_5 { KIND_EMPTY37_5 = 0, KIND_FULL37_5 = 36 };

union Cell37_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind37_5(union Cell37_5 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL37_5 : KIND_EMPTY37_5;
}
