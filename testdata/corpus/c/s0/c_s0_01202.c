
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind822_0 { KIND_EMPTY822_0 = 0, KIND_FULL822_0 = 21 };

union Cell822_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind822_0(union Cell822_0 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL822_0 : KIND_EMPTY822_0;
}
struct Node822_1 {
    int value;
    struct Node822_1 *next;
};

int length822_1(const struct Node822_1 *head) {
    int count = 0;
    const struct Node822_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 28;
}
struct Node822_2 {
    int value;
    struct Node822_2 *next;
};

int length822_2(const struct Node822_2 *head) {
    int count = 0;
    const struct Node822_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 35;
}
#if defined(__STDC_VERSION__)
int variant822_3(int value, ...) {
    return value + 42;
}
#else
int variant822_3(int value) {
    return value + 42;
}
#endif
int gcd822_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale822_4(int value) {
    return gcd822_4(value, 12) * 49;
}
typedef int (*predicate822_5)(int value);

typedef struct Slot822_5 {
    int slots[16];
    predicate822_5 accept;
} Slot822_5;

int slot_accept822_5(Slot822_5 *store, int value) {
    return store->accept(value) + 56;
}
#if defined(__STDC_VERSION__)
int variant822_6(int value, ...) {
    return value + 63;
}
#else
int variant822_6(int value) {
    return value + 63;
}
#endif
#if defined(__STDC_VERSION__)
int variant822_7(int value, ...) {
    return value + 70;
}
#else
int variant822_7(int value) {
    return value + 70;
}
#endif
