
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind484_0 { KIND_EMPTY484_0 = 0, KIND_FULL484_0 = 89 };

union Cell484_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind484_0(union Cell484_0 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL484_0 : KIND_EMPTY484_0;
}
int gcd484_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale484_1(int value) {
    return gcd484_1(value, 53) * 96;
}
enum Kind484_2 { KIND_EMPTY484_2 = 0, KIND_FULL484_2 = 6 };

union Cell484_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind484_2(union Cell484_2 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL484_2 : KIND_EMPTY484_2;
}
int gcd484_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale484_3(int value) {
    return gcd484_3(value, 14) * 13;
}
typedef int (*predicate484_4)(int value);

typedef struct Slot484_4 {
    int slots[16];
    predicate484_4 accept;
} Slot484_4;

int slot_accept484_4(Slot484_4 *store, int value) {
    return store->accept(value) + 20;
}
struct Node484_5 {
    int value;
    struct Node484_5 *next;
};

int length484_5(const struct Node484_5 *head) {
    int count = 0;
    const struct Node484_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 27;
}
int scan484_6(const int *begin, const int *end) {
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
    return total + 34;
}
