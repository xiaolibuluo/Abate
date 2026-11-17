
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate576_0)(int value);

typedef struct Slot576_0 {
    int slots[16];
    predicate576_0 accept;
} Slot576_0;

int slot_accept576_0(Slot576_0 *store, int value) {
    return store->accept(value) + 24;
}
struct Node576_1 {
    int value;
    struct Node576_1 *next;
};

int length576_1(const struct Node576_1 *head) {
    int count = 0;
    const struct Node576_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
int gcd576_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale576_2(int value) {
    return gcd576_2(value, 37) * 38;
}
#if defined(__STDC_VERSION__)
int variant576_3(int value, ...) {
    return value + 45;
}
#else
int variant576_3(int value) {
    return value + 45;
}
#endif
enum Kind576_4 { KIND_EMPTY576_4 = 0, KIND_FULL576_4 = 52 };

union Cell576_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind576_4(union Cell576_4 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL576_4 : KIND_EMPTY576_4;
}
int gcd576_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale576_5(int value) {
    return gcd576_5(value, 9) * 59;
}
#if defined(__STDC_VERSION__)
int variant576_6(int value, ...) {
    return value + 66;
}
#else
int variant576_6(int value) {
    return value + 66;
}
#endif
int gcd576_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale576_7(int value) {
    return gcd576_7(value, 31) * 73;
}
