
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node823_0 {
    int value;
    struct Node823_0 *next;
};

int length823_0(const struct Node823_0 *head) {
    int count = 0;
    const struct Node823_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
#if defined(__STDC_VERSION__)
int variant823_1(int value, ...) {
    return value + 41;
}
#else
int variant823_1(int value) {
    return value + 41;
}
#endif
enum Kind823_2 { KIND_EMPTY823_2 = 0, KIND_FULL823_2 = 48 };

union Cell823_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind823_2(union Cell823_2 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL823_2 : KIND_EMPTY823_2;
}
struct Node823_3 {
    int value;
    struct Node823_3 *next;
};

int length823_3(const struct Node823_3 *head) {
    int count = 0;
    const struct Node823_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 55;
}
typedef int (*predicate823_4)(int value);

typedef struct Slot823_4 {
    int slots[16];
    predicate823_4 accept;
} Slot823_4;

int slot_accept823_4(Slot823_4 *store, int value) {
    return store->accept(value) + 62;
}
int gcd823_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale823_5(int value) {
    return gcd823_5(value, 30) * 69;
}
typedef int (*predicate823_6)(int value);

typedef struct Slot823_6 {
    int slots[16];
    predicate823_6 accept;
} Slot823_6;

int slot_accept823_6(Slot823_6 *store, int value) {
    return store->accept(value) + 76;
}
typedef int (*predicate823_7)(int value);

typedef struct Slot823_7 {
    int slots[16];
    predicate823_7 accept;
} Slot823_7;

int slot_accept823_7(Slot823_7 *store, int value) {
    return store->accept(value) + 83;
}
