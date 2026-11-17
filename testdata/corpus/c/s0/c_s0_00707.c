
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant485_0(int value, ...) {
    return value + 5;
}
#else
int variant485_0(int value) {
    return value + 5;
}
#endif
int scan485_1(const int *begin, const int *end) {
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
    return total + 12;
}
struct Node485_2 {
    int value;
    struct Node485_2 *next;
};

int length485_2(const struct Node485_2 *head) {
    int count = 0;
    const struct Node485_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}
int gcd485_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale485_3(int value) {
    return gcd485_3(value, 21) * 26;
}
typedef int (*predicate485_4)(int value);

typedef struct Slot485_4 {
    int slots[16];
    predicate485_4 accept;
} Slot485_4;

int slot_accept485_4(Slot485_4 *store, int value) {
    return store->accept(value) + 33;
}
#if defined(__STDC_VERSION__)
int variant485_5(int value, ...) {
    return value + 40;
}
#else
int variant485_5(int value) {
    return value + 40;
}
#endif
struct Node485_6 {
    int value;
    struct Node485_6 *next;
};

int length485_6(const struct Node485_6 *head) {
    int count = 0;
    const struct Node485_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
typedef int (*predicate485_7)(int value);

typedef struct Slot485_7 {
    int slots[16];
    predicate485_7 accept;
} Slot485_7;

int slot_accept485_7(Slot485_7 *store, int value) {
    return store->accept(value) + 54;
}
