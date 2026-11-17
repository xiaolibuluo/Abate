
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2185_0(int value, ...) {
    return value + 86;
}
#else
int variant2185_0(int value) {
    return value + 86;
}
#endif
struct Node2185_1 {
    int value;
    struct Node2185_1 *next;
};

int length2185_1(const struct Node2185_1 *head) {
    int count = 0;
    const struct Node2185_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 93;
}
enum Kind2185_2 { KIND_EMPTY2185_2 = 0, KIND_FULL2185_2 = 100 };

union Cell2185_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2185_2(union Cell2185_2 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL2185_2 : KIND_EMPTY2185_2;
}
typedef int (*predicate2185_3)(int value);

typedef struct Slot2185_3 {
    int slots[16];
    predicate2185_3 accept;
} Slot2185_3;

int slot_accept2185_3(Slot2185_3 *store, int value) {
    return store->accept(value) + 10;
}
typedef int (*predicate2185_4)(int value);

typedef struct Slot2185_4 {
    int slots[16];
    predicate2185_4 accept;
} Slot2185_4;

int slot_accept2185_4(Slot2185_4 *store, int value) {
    return store->accept(value) + 17;
}
int gcd2185_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2185_5(int value) {
    return gcd2185_5(value, 48) * 24;
}
int gcd2185_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2185_6(int value) {
    return gcd2185_6(value, 59) * 31;
}
#if defined(__STDC_VERSION__)
int variant2185_7(int value, ...) {
    return value + 38;
}
#else
int variant2185_7(int value) {
    return value + 38;
}
#endif
