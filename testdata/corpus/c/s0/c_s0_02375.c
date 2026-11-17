
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1605_0 { KIND_EMPTY1605_0 = 0, KIND_FULL1605_0 = 15 };

union Cell1605_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1605_0(union Cell1605_0 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL1605_0 : KIND_EMPTY1605_0;
}
typedef int (*predicate1605_1)(int value);

typedef struct Slot1605_1 {
    int slots[16];
    predicate1605_1 accept;
} Slot1605_1;

int slot_accept1605_1(Slot1605_1 *store, int value) {
    return store->accept(value) + 22;
}
struct Node1605_2 {
    int value;
    struct Node1605_2 *next;
};

int length1605_2(const struct Node1605_2 *head) {
    int count = 0;
    const struct Node1605_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 29;
}
#define MASK1605_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1605_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1605_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1605_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1605_4(int value) {
    return gcd1605_4(value, 64) * 43;
}
int gcd1605_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1605_5(int value) {
    return gcd1605_5(value, 14) * 50;
}
#define MASK1605_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1605_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1605_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1605_7(const int *begin, const int *end) {
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
    return total + 64;
}
