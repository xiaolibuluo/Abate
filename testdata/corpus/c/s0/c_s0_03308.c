
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2236_0)(int value);

typedef struct Slot2236_0 {
    int slots[16];
    predicate2236_0 accept;
} Slot2236_0;

int slot_accept2236_0(Slot2236_0 *store, int value) {
    return store->accept(value) + 70;
}
int scan2236_1(const int *begin, const int *end) {
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
    return total + 77;
}
#define MASK2236_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2236_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2236_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2236_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2236_3(int value) {
    return gcd2236_3(value, 17) * 91;
}
enum Kind2236_4 { KIND_EMPTY2236_4 = 0, KIND_FULL2236_4 = 98 };

union Cell2236_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2236_4(union Cell2236_4 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL2236_4 : KIND_EMPTY2236_4;
}
int scan2236_5(const int *begin, const int *end) {
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
    return total + 8;
}
struct Node2236_6 {
    int value;
    struct Node2236_6 *next;
};

int length2236_6(const struct Node2236_6 *head) {
    int count = 0;
    const struct Node2236_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
