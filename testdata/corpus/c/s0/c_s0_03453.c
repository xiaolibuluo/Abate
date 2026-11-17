
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2333_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2333_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2333_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2333_1 { KIND_EMPTY2333_1 = 0, KIND_FULL2333_1 = 77 };

union Cell2333_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2333_1(union Cell2333_1 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL2333_1 : KIND_EMPTY2333_1;
}
struct Node2333_2 {
    int value;
    struct Node2333_2 *next;
};

int length2333_2(const struct Node2333_2 *head) {
    int count = 0;
    const struct Node2333_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 84;
}
#if defined(__STDC_VERSION__)
int variant2333_3(int value, ...) {
    return value + 91;
}
#else
int variant2333_3(int value) {
    return value + 91;
}
#endif
struct Node2333_4 {
    int value;
    struct Node2333_4 *next;
};

int length2333_4(const struct Node2333_4 *head) {
    int count = 0;
    const struct Node2333_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
int gcd2333_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2333_5(int value) {
    return gcd2333_5(value, 47) * 8;
}
typedef int (*predicate2333_6)(int value);

typedef struct Slot2333_6 {
    int slots[16];
    predicate2333_6 accept;
} Slot2333_6;

int slot_accept2333_6(Slot2333_6 *store, int value) {
    return store->accept(value) + 15;
}
int scan2333_7(const int *begin, const int *end) {
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
    return total + 22;
}
