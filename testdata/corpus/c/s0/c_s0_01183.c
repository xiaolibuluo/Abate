
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant810_0(int value, ...) {
    return value + 59;
}
#else
int variant810_0(int value) {
    return value + 59;
}
#endif
int gcd810_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale810_1(int value) {
    return gcd810_1(value, 17) * 66;
}
struct Node810_2 {
    int value;
    struct Node810_2 *next;
};

int length810_2(const struct Node810_2 *head) {
    int count = 0;
    const struct Node810_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
#define MASK810_3(v) (((v) << 4) | ((v) & 0x0F))

int classify810_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK810_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan810_4(const int *begin, const int *end) {
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
    return total + 87;
}
typedef int (*predicate810_5)(int value);

typedef struct Slot810_5 {
    int slots[16];
    predicate810_5 accept;
} Slot810_5;

int slot_accept810_5(Slot810_5 *store, int value) {
    return store->accept(value) + 94;
}
