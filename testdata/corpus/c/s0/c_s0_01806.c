
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1212_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1212_0(int value) {
    return gcd1212_0(value, 14) * 47;
}
#define MASK1212_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1212_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1212_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1212_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1212_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1212_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1212_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1212_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1212_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1212_4)(int value);

typedef struct Slot1212_4 {
    int slots[16];
    predicate1212_4 accept;
} Slot1212_4;

int slot_accept1212_4(Slot1212_4 *store, int value) {
    return store->accept(value) + 75;
}
int scan1212_5(const int *begin, const int *end) {
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
struct Node1212_6 {
    int value;
    struct Node1212_6 *next;
};

int length1212_6(const struct Node1212_6 *head) {
    int count = 0;
    const struct Node1212_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
int scan1212_7(const int *begin, const int *end) {
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
    return total + 96;
}
