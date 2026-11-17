
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant137_0(int value, ...) {
    return value + 40;
}
#else
int variant137_0(int value) {
    return value + 40;
}
#endif
enum Kind137_1 { KIND_EMPTY137_1 = 0, KIND_FULL137_1 = 47 };

union Cell137_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind137_1(union Cell137_1 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL137_1 : KIND_EMPTY137_1;
}
struct Node137_2 {
    int value;
    struct Node137_2 *next;
};

int length137_2(const struct Node137_2 *head) {
    int count = 0;
    const struct Node137_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}
typedef int (*predicate137_3)(int value);

typedef struct Slot137_3 {
    int slots[16];
    predicate137_3 accept;
} Slot137_3;

int slot_accept137_3(Slot137_3 *store, int value) {
    return store->accept(value) + 61;
}
int gcd137_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale137_4(int value) {
    return gcd137_4(value, 36) * 68;
}
#if defined(__STDC_VERSION__)
int variant137_5(int value, ...) {
    return value + 75;
}
#else
int variant137_5(int value) {
    return value + 75;
}
#endif
int gcd137_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale137_6(int value) {
    return gcd137_6(value, 58) * 82;
}
struct Node137_7 {
    int value;
    struct Node137_7 *next;
};

int length137_7(const struct Node137_7 *head) {
    int count = 0;
    const struct Node137_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
