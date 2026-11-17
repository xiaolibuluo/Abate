
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node242_0 {
    int value;
    struct Node242_0 *next;
};

int length242_0(const struct Node242_0 *head) {
    int count = 0;
    const struct Node242_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
enum Kind242_1 { KIND_EMPTY242_1 = 0, KIND_FULL242_1 = 54 };

union Cell242_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind242_1(union Cell242_1 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL242_1 : KIND_EMPTY242_1;
}
typedef int (*predicate242_2)(int value);

typedef struct Slot242_2 {
    int slots[16];
    predicate242_2 accept;
} Slot242_2;

int slot_accept242_2(Slot242_2 *store, int value) {
    return store->accept(value) + 61;
}
#define MASK242_3(v) (((v) << 4) | ((v) & 0x0F))

int classify242_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK242_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan242_4(const int *begin, const int *end) {
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
    return total + 75;
}
int scan242_5(const int *begin, const int *end) {
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
    return total + 82;
}
int gcd242_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale242_6(int value) {
    return gcd242_6(value, 61) * 89;
}
#if defined(__STDC_VERSION__)
int variant242_7(int value, ...) {
    return value + 96;
}
#else
int variant242_7(int value) {
    return value + 96;
}
#endif
