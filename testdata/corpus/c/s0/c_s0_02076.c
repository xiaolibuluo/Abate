
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1398_0 {
    int value;
    struct Node1398_0 *next;
};

int length1398_0(const struct Node1398_0 *head) {
    int count = 0;
    const struct Node1398_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
#define MASK1398_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1398_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1398_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1398_2)(int value);

typedef struct Slot1398_2 {
    int slots[16];
    predicate1398_2 accept;
} Slot1398_2;

int slot_accept1398_2(Slot1398_2 *store, int value) {
    return store->accept(value) + 54;
}
#define MASK1398_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1398_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1398_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1398_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1398_4(int value) {
    return gcd1398_4(value, 18) * 68;
}
int scan1398_5(const int *begin, const int *end) {
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
