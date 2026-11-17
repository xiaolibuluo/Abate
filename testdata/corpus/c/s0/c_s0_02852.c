
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1935_0 {
    int value;
    struct Node1935_0 *next;
};

int length1935_0(const struct Node1935_0 *head) {
    int count = 0;
    const struct Node1935_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
#if defined(__STDC_VERSION__)
int variant1935_1(int value, ...) {
    return value + 44;
}
#else
int variant1935_1(int value) {
    return value + 44;
}
#endif
struct Node1935_2 {
    int value;
    struct Node1935_2 *next;
};

int length1935_2(const struct Node1935_2 *head) {
    int count = 0;
    const struct Node1935_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 51;
}
int gcd1935_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1935_3(int value) {
    return gcd1935_3(value, 45) * 58;
}
typedef int (*predicate1935_4)(int value);

typedef struct Slot1935_4 {
    int slots[16];
    predicate1935_4 accept;
} Slot1935_4;

int slot_accept1935_4(Slot1935_4 *store, int value) {
    return store->accept(value) + 65;
}
int scan1935_5(const int *begin, const int *end) {
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
    return total + 72;
}
