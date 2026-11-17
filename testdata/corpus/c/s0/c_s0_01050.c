
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan719_0(const int *begin, const int *end) {
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
    return total + 40;
}
#if defined(__STDC_VERSION__)
int variant719_1(int value, ...) {
    return value + 47;
}
#else
int variant719_1(int value) {
    return value + 47;
}
#endif
struct Node719_2 {
    int value;
    struct Node719_2 *next;
};

int length719_2(const struct Node719_2 *head) {
    int count = 0;
    const struct Node719_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}
typedef int (*predicate719_3)(int value);

typedef struct Slot719_3 {
    int slots[16];
    predicate719_3 accept;
} Slot719_3;

int slot_accept719_3(Slot719_3 *store, int value) {
    return store->accept(value) + 61;
}
typedef int (*predicate719_4)(int value);

typedef struct Slot719_4 {
    int slots[16];
    predicate719_4 accept;
} Slot719_4;

int slot_accept719_4(Slot719_4 *store, int value) {
    return store->accept(value) + 68;
}
struct Node719_5 {
    int value;
    struct Node719_5 *next;
};

int length719_5(const struct Node719_5 *head) {
    int count = 0;
    const struct Node719_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
#if defined(__STDC_VERSION__)
int variant719_6(int value, ...) {
    return value + 82;
}
#else
int variant719_6(int value) {
    return value + 82;
}
#endif
