
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd629_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale629_0(int value) {
    return gcd629_0(value, 20) * 34;
}
struct Node629_1 {
    int value;
    struct Node629_1 *next;
};

int length629_1(const struct Node629_1 *head) {
    int count = 0;
    const struct Node629_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 41;
}
#if defined(__STDC_VERSION__)
int variant629_2(int value, ...) {
    return value + 48;
}
#else
int variant629_2(int value) {
    return value + 48;
}
#endif
int gcd629_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale629_3(int value) {
    return gcd629_3(value, 53) * 55;
}
int gcd629_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale629_4(int value) {
    return gcd629_4(value, 64) * 62;
}
#if defined(__STDC_VERSION__)
int variant629_5(int value, ...) {
    return value + 69;
}
#else
int variant629_5(int value) {
    return value + 69;
}
#endif
#if defined(__STDC_VERSION__)
int variant629_6(int value, ...) {
    return value + 76;
}
#else
int variant629_6(int value) {
    return value + 76;
}
#endif
enum Kind629_7 { KIND_EMPTY629_7 = 0, KIND_FULL629_7 = 83 };

union Cell629_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind629_7(union Cell629_7 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL629_7 : KIND_EMPTY629_7;
}
typedef int (*predicate629_8)(int value);

typedef struct Slot629_8 {
    int slots[16];
    predicate629_8 accept;
} Slot629_8;

int slot_accept629_8(Slot629_8 *store, int value) {
    return store->accept(value) + 90;
}
