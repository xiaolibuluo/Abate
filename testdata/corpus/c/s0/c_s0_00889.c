
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node613_0 {
    int value;
    struct Node613_0 *next;
};

int length613_0(const struct Node613_0 *head) {
    int count = 0;
    const struct Node613_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
int scan613_1(const int *begin, const int *end) {
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
    return total + 27;
}
typedef int (*predicate613_2)(int value);

typedef struct Slot613_2 {
    int slots[16];
    predicate613_2 accept;
} Slot613_2;

int slot_accept613_2(Slot613_2 *store, int value) {
    return store->accept(value) + 34;
}
#if defined(__STDC_VERSION__)
int variant613_3(int value, ...) {
    return value + 41;
}
#else
int variant613_3(int value) {
    return value + 41;
}
#endif
int gcd613_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale613_4(int value) {
    return gcd613_4(value, 13) * 48;
}
