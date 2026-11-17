
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK916_0(v) (((v) << 4) | ((v) & 0x0F))

int classify916_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK916_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate916_1)(int value);

typedef struct Slot916_1 {
    int slots[16];
    predicate916_1 accept;
} Slot916_1;

int slot_accept916_1(Slot916_1 *store, int value) {
    return store->accept(value) + 86;
}
int scan916_2(const int *begin, const int *end) {
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
    return total + 93;
}
int gcd916_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale916_3(int value) {
    return gcd916_3(value, 49) * 100;
}
typedef int (*predicate916_4)(int value);

typedef struct Slot916_4 {
    int slots[16];
    predicate916_4 accept;
} Slot916_4;

int slot_accept916_4(Slot916_4 *store, int value) {
    return store->accept(value) + 10;
}
struct Node916_5 {
    int value;
    struct Node916_5 *next;
};

int length916_5(const struct Node916_5 *head) {
    int count = 0;
    const struct Node916_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
int scan916_6(const int *begin, const int *end) {
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
    return total + 24;
}
