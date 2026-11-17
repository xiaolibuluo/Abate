
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd240_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale240_0(int value) {
    return gcd240_0(value, 42) * 21;
}
int scan240_1(const int *begin, const int *end) {
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
    return total + 28;
}
enum Kind240_2 { KIND_EMPTY240_2 = 0, KIND_FULL240_2 = 35 };

union Cell240_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind240_2(union Cell240_2 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL240_2 : KIND_EMPTY240_2;
}
typedef int (*predicate240_3)(int value);

typedef struct Slot240_3 {
    int slots[16];
    predicate240_3 accept;
} Slot240_3;

int slot_accept240_3(Slot240_3 *store, int value) {
    return store->accept(value) + 42;
}
#if defined(__STDC_VERSION__)
int variant240_4(int value, ...) {
    return value + 49;
}
#else
int variant240_4(int value) {
    return value + 49;
}
#endif
#define MASK240_5(v) (((v) << 4) | ((v) & 0x0F))

int classify240_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK240_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan240_6(const int *begin, const int *end) {
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
    return total + 63;
}
struct Node240_7 {
    int value;
    struct Node240_7 *next;
};

int length240_7(const struct Node240_7 *head) {
    int count = 0;
    const struct Node240_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 70;
}
