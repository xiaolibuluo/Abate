
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate825_0)(int value);

typedef struct Slot825_0 {
    int slots[16];
    predicate825_0 accept;
} Slot825_0;

int slot_accept825_0(Slot825_0 *store, int value) {
    return store->accept(value) + 60;
}
struct Node825_1 {
    int value;
    struct Node825_1 *next;
};

int length825_1(const struct Node825_1 *head) {
    int count = 0;
    const struct Node825_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
typedef int (*predicate825_2)(int value);

typedef struct Slot825_2 {
    int slots[16];
    predicate825_2 accept;
} Slot825_2;

int slot_accept825_2(Slot825_2 *store, int value) {
    return store->accept(value) + 74;
}
int scan825_3(const int *begin, const int *end) {
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
    return total + 81;
}
#if defined(__STDC_VERSION__)
int variant825_4(int value, ...) {
    return value + 88;
}
#else
int variant825_4(int value) {
    return value + 88;
}
#endif
int gcd825_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale825_5(int value) {
    return gcd825_5(value, 44) * 95;
}
int gcd825_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale825_6(int value) {
    return gcd825_6(value, 55) * 5;
}
int scan825_7(const int *begin, const int *end) {
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
