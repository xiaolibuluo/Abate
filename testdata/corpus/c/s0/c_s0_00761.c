
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node527_0 {
    int value;
    struct Node527_0 *next;
};

int length527_0(const struct Node527_0 *head) {
    int count = 0;
    const struct Node527_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
typedef int (*predicate527_1)(int value);

typedef struct Slot527_1 {
    int slots[16];
    predicate527_1 accept;
} Slot527_1;

int slot_accept527_1(Slot527_1 *store, int value) {
    return store->accept(value) + 73;
}
enum Kind527_2 { KIND_EMPTY527_2 = 0, KIND_FULL527_2 = 80 };

union Cell527_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind527_2(union Cell527_2 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL527_2 : KIND_EMPTY527_2;
}
int gcd527_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale527_3(int value) {
    return gcd527_3(value, 10) * 87;
}
enum Kind527_4 { KIND_EMPTY527_4 = 0, KIND_FULL527_4 = 94 };

union Cell527_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind527_4(union Cell527_4 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL527_4 : KIND_EMPTY527_4;
}
struct Node527_5 {
    int value;
    struct Node527_5 *next;
};

int length527_5(const struct Node527_5 *head) {
    int count = 0;
    const struct Node527_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
#if defined(__STDC_VERSION__)
int variant527_6(int value, ...) {
    return value + 11;
}
#else
int variant527_6(int value) {
    return value + 11;
}
#endif
