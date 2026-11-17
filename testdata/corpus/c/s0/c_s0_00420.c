
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate285_0)(int value);

typedef struct Slot285_0 {
    int slots[16];
    predicate285_0 accept;
} Slot285_0;

int slot_accept285_0(Slot285_0 *store, int value) {
    return store->accept(value) + 24;
}
int scan285_1(const int *begin, const int *end) {
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
    return total + 31;
}
#if defined(__STDC_VERSION__)
int variant285_2(int value, ...) {
    return value + 38;
}
#else
int variant285_2(int value) {
    return value + 38;
}
#endif
#define MASK285_3(v) (((v) << 4) | ((v) & 0x0F))

int classify285_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK285_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd285_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale285_4(int value) {
    return gcd285_4(value, 35) * 52;
}
#define MASK285_5(v) (((v) << 4) | ((v) & 0x0F))

int classify285_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK285_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind285_6 { KIND_EMPTY285_6 = 0, KIND_FULL285_6 = 66 };

union Cell285_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind285_6(union Cell285_6 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL285_6 : KIND_EMPTY285_6;
}
struct Node285_7 {
    int value;
    struct Node285_7 *next;
};

int length285_7(const struct Node285_7 *head) {
    int count = 0;
    const struct Node285_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
