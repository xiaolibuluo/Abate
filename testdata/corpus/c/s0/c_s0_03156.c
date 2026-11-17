
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2139_0 {
    int value;
    struct Node2139_0 *next;
};

int length2139_0(const struct Node2139_0 *head) {
    int count = 0;
    const struct Node2139_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 70;
}
enum Kind2139_1 { KIND_EMPTY2139_1 = 0, KIND_FULL2139_1 = 77 };

union Cell2139_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2139_1(union Cell2139_1 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL2139_1 : KIND_EMPTY2139_1;
}
typedef int (*predicate2139_2)(int value);

typedef struct Slot2139_2 {
    int slots[16];
    predicate2139_2 accept;
} Slot2139_2;

int slot_accept2139_2(Slot2139_2 *store, int value) {
    return store->accept(value) + 84;
}
#define MASK2139_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2139_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2139_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2139_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2139_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2139_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2139_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2139_5(int value) {
    return gcd2139_5(value, 31) * 8;
}
int scan2139_6(const int *begin, const int *end) {
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
    return total + 15;
}
#if defined(__STDC_VERSION__)
int variant2139_7(int value, ...) {
    return value + 22;
}
#else
int variant2139_7(int value) {
    return value + 22;
}
#endif
