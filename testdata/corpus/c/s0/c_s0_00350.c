
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK238_0(v) (((v) << 4) | ((v) & 0x0F))

int classify238_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK238_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate238_1)(int value);

typedef struct Slot238_1 {
    int slots[16];
    predicate238_1 accept;
} Slot238_1;

int slot_accept238_1(Slot238_1 *store, int value) {
    return store->accept(value) + 99;
}
int scan238_2(const int *begin, const int *end) {
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
    return total + 9;
}
#define MASK238_3(v) (((v) << 4) | ((v) & 0x0F))

int classify238_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK238_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd238_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale238_4(int value) {
    return gcd238_4(value, 11) * 23;
}
int scan238_5(const int *begin, const int *end) {
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
    return total + 30;
}
struct Node238_6 {
    int value;
    struct Node238_6 *next;
};

int length238_6(const struct Node238_6 *head) {
    int count = 0;
    const struct Node238_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
