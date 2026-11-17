
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd186_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale186_0(int value) {
    return gcd186_0(value, 30) * 95;
}
#if defined(__STDC_VERSION__)
int variant186_1(int value, ...) {
    return value + 5;
}
#else
int variant186_1(int value) {
    return value + 5;
}
#endif
typedef int (*predicate186_2)(int value);

typedef struct Slot186_2 {
    int slots[16];
    predicate186_2 accept;
} Slot186_2;

int slot_accept186_2(Slot186_2 *store, int value) {
    return store->accept(value) + 12;
}
#if defined(__STDC_VERSION__)
int variant186_3(int value, ...) {
    return value + 19;
}
#else
int variant186_3(int value) {
    return value + 19;
}
#endif
enum Kind186_4 { KIND_EMPTY186_4 = 0, KIND_FULL186_4 = 26 };

union Cell186_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind186_4(union Cell186_4 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL186_4 : KIND_EMPTY186_4;
}
struct Node186_5 {
    int value;
    struct Node186_5 *next;
};

int length186_5(const struct Node186_5 *head) {
    int count = 0;
    const struct Node186_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
struct Node186_6 {
    int value;
    struct Node186_6 *next;
};

int length186_6(const struct Node186_6 *head) {
    int count = 0;
    const struct Node186_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
