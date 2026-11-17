
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant960_0(int value, ...) {
    return value + 69;
}
#else
int variant960_0(int value) {
    return value + 69;
}
#endif
int gcd960_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale960_1(int value) {
    return gcd960_1(value, 30) * 76;
}
struct Node960_2 {
    int value;
    struct Node960_2 *next;
};

int length960_2(const struct Node960_2 *head) {
    int count = 0;
    const struct Node960_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 83;
}
int scan960_3(const int *begin, const int *end) {
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
    return total + 90;
}
#if defined(__STDC_VERSION__)
int variant960_4(int value, ...) {
    return value + 97;
}
#else
int variant960_4(int value) {
    return value + 97;
}
#endif
typedef int (*predicate960_5)(int value);

typedef struct Slot960_5 {
    int slots[16];
    predicate960_5 accept;
} Slot960_5;

int slot_accept960_5(Slot960_5 *store, int value) {
    return store->accept(value) + 7;
}
typedef int (*predicate960_6)(int value);

typedef struct Slot960_6 {
    int slots[16];
    predicate960_6 accept;
} Slot960_6;

int slot_accept960_6(Slot960_6 *store, int value) {
    return store->accept(value) + 14;
}
