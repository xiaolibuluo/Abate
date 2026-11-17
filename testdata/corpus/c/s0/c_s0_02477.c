
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1682_0 {
    int value;
    struct Node1682_0 *next;
};

int length1682_0(const struct Node1682_0 *head) {
    int count = 0;
    const struct Node1682_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
int gcd1682_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1682_1(int value) {
    return gcd1682_1(value, 21) * 53;
}
int scan1682_2(const int *begin, const int *end) {
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
    return total + 60;
}
#define MASK1682_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1682_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1682_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1682_4)(int value);

typedef struct Slot1682_4 {
    int slots[16];
    predicate1682_4 accept;
} Slot1682_4;

int slot_accept1682_4(Slot1682_4 *store, int value) {
    return store->accept(value) + 74;
}
