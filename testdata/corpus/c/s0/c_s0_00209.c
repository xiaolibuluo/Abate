
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate144_0)(int value);

typedef struct Slot144_0 {
    int slots[16];
    predicate144_0 accept;
} Slot144_0;

int slot_accept144_0(Slot144_0 *store, int value) {
    return store->accept(value) + 34;
}
typedef int (*predicate144_1)(int value);

typedef struct Slot144_1 {
    int slots[16];
    predicate144_1 accept;
} Slot144_1;

int slot_accept144_1(Slot144_1 *store, int value) {
    return store->accept(value) + 41;
}
int gcd144_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale144_2(int value) {
    return gcd144_2(value, 63) * 48;
}
int gcd144_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale144_3(int value) {
    return gcd144_3(value, 13) * 55;
}
enum Kind144_4 { KIND_EMPTY144_4 = 0, KIND_FULL144_4 = 62 };

union Cell144_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind144_4(union Cell144_4 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL144_4 : KIND_EMPTY144_4;
}
#if defined(__STDC_VERSION__)
int variant144_5(int value, ...) {
    return value + 69;
}
#else
int variant144_5(int value) {
    return value + 69;
}
#endif
#if defined(__STDC_VERSION__)
int variant144_6(int value, ...) {
    return value + 76;
}
#else
int variant144_6(int value) {
    return value + 76;
}
#endif
enum Kind144_7 { KIND_EMPTY144_7 = 0, KIND_FULL144_7 = 83 };

union Cell144_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind144_7(union Cell144_7 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL144_7 : KIND_EMPTY144_7;
}
struct Node144_8 {
    int value;
    struct Node144_8 *next;
};

int length144_8(const struct Node144_8 *head) {
    int count = 0;
    const struct Node144_8 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 90;
}
